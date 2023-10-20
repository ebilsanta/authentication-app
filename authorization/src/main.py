
import base64
from functools import lru_cache
import hashlib
from typing import Union
from fastapi import FastAPI, Request, status
from fastapi.encoders import jsonable_encoder
from fastapi.exceptions import RequestValidationError
from fastapi.responses import JSONResponse, RedirectResponse
import uuid
import time
import jwt

from pydantic import BaseModel
from app.authcode_service import AuthCodeService
from app.database import Database, AuthCodeRecord
from app.dpop_service import DpopService
from app.pkce import generate_pkce_code_challenge
from app.client_assertion_service import ClientAssertionService
from config import Settings


@lru_cache()
def get_settings():
    return Settings()


ac = AuthCodeService()
dps = DpopService()
db = Database()
cas = ClientAssertionService()

app = FastAPI()


def respond(redirect_url, message, desc=None):
    if redirect_url and desc:
        return ac.make_error_desc(redirect_url, message, desc)
    elif redirect_url:
        return ac.make_error(redirect_url, message)
    elif desc:
        return JSONResponse(
            status_code=status.HTTP_400_BAD_REQUEST,
            content=jsonable_encoder(
                {'error': message, 'error_description': desc}),
        )
    else:
        return JSONResponse(
            status_code=status.HTTP_400_BAD_REQUEST,
            content=jsonable_encoder({'error': message}),
        )


@app.get("/")
def read_root():
    return {"Hello": "World"}

# response_type: str              RFC 6749 4.1.1. Value MUST be set to "code"
# client_id: str                  RFC 6749 4.1.1
# redirect_url: str               RFC 6749 4.1.1
# state: str                      RFC 6749 4.1.1 & 10.12
# code_challenge: str             RFC 7636 4.3


@app.post("/authcode")
async def post_authcode(response_type: str, client_id: str,
                        state: str, id_jwt: str,
                        code_challenge: Union[str, None] = None,
                        code_challenge_method: Union[str, None] = None,
                        redirect_url: Union[str, None] = None):

    # Verify redirection url is registered
    # RFC 6749 10.6 check redirection url is same as get_auth_code
    if redirect_url and not ac.is_redirect_valid(redirect_url):
        return respond(redirect_url, 'access_denied', 'invalid_redirect')

    if response_type != "code":
        return respond(redirect_url, 'unsupported_response_type')

    if not ac.is_client_allowed(client_id):
        return respond(redirect_url, 'unauthorized_client')

    if not code_challenge:
        return respond(redirect_url, 'invalid_request', desc='code challenge required')
    elif len(code_challenge) != 44:
        return respond(redirect_url, 'invalid_request', desc='invalid code challenge')

    if code_challenge_method != "S256":
        return respond(redirect_url, 'invalid_request', desc='transform algorithm not supported')

    if len(state) < 16 or len(state) > 64:
        return respond(redirect_url, 'invalid_request', desc='invalid state')

    err_message, decoded = ac.verify_jwt(id_jwt)
    if err_message:
        return respond(redirect_url, 'access_denied', err_message)

    # Check if user is actually in our system first
    if not await db.exists_valid_user(decoded['sub']):
        return respond(redirect_url, 'access_denied', 'unknown user')

    code = uuid.uuid4().hex

    await db.insert_authcode_record(AuthCodeRecord(code, state, code_challenge, decoded['sub'], 600))
    if redirect_url:
        return RedirectResponse(url=redirect_url + '?code=' + code, status_code=302)
    else:
        return JSONResponse(
            status_code=status.HTTP_200_OK,
            content=jsonable_encoder({'code': code}),
        )


@app.exception_handler(RequestValidationError)
async def standard_validation_exception_handler(request: Request, exc: RequestValidationError):
    return JSONResponse(
        status_code=status.HTTP_400_BAD_REQUEST,
        content=jsonable_encoder(
            {'error': 'invalid_request', 'error_description': 'The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed.'}),
    )


class TokenRequest(BaseModel):
    grant_type: str
    authcode: str
    dpop: str
    client_assertion: str
    redirect_url: str
    code_verifier: str

class RefreshRequest(BaseModel):
    grant_type: str
    dpop: str
    refresh_token: str


@app.post("/token")
async def post_token(token_req: TokenRequest):
    if token_req.grant_type != "authorization_code":
        return respond(token_req.redirect_url, 'unsupported grant type')

    err, err_desc = cas.verify_client_assertion(token_req.client_assertion)
    if err:
        return respond(token_req.redirect_url, err, err_desc)

    jwk, err = dps.verify_dpop(token_req.dpop)
    if err:
        return respond(token_req.redirect_url, err)

    ac = await db.get_authcode_record(token_req.authcode)
    print(ac.__dict__)

    print(generate_pkce_code_challenge(token_req.code_verifier))
    if generate_pkce_code_challenge(token_req.code_verifier) != ac.code_challenge:
        return respond(token_req.redirect_url, 'Invalid PKCE Code Verifier')

    sets = get_settings()

    now = int(time.time())
    payload = {
        "sub": ac.user,                                                  
        "iss": sets.audience,                                            
        "exp": now + 3600,                                                 
        "iat": now,
        "cnf.jkt": base64.b64encode(hashlib.sha256(jwk).digest()).decode(),
        "typ": "dpop"
    }
    access_token = jwt.encode(payload, get_settings().authz_pvt_key.replace(
        '\\n', '\n').replace('\\t', '\t'), algorithm="RS256")
    
    payload = {
        "sub": ac.user,                                                  
        "iss": sets.audience,                                            
        "exp": now + 86400,                                                 
        "iat": now,
        "cnf.jkt": base64.b64encode(hashlib.sha256(jwk).digest()).decode(),
        "typ": "dpop+refresh"
    }

    refresh_token = jwt.encode(payload, get_settings().authz_pvt_key.replace(
        '\\n', '\n').replace('\\t', '\t'), algorithm="RS256")

    return JSONResponse(
        status_code=status.HTTP_200_OK,
        content=jsonable_encoder({"access_token": access_token,
                                  "token_type": "DPoP",
                                  "expires_in": 3600,
                                  "refresh_token": refresh_token}) 
    )

@app.post("/refresh")
async def post_refresh(refresh_req: RefreshRequest):
    if refresh_req.grant_type != "authorization_code":
        return respond(None, 'unsupported grant type')
    
    jwk, err = dps.verify_dpop(refresh_req.dpop, at=refresh_req.refresh_token)

    sets = get_settings()

    now = int(time.time())
    payload = {
        "sub": ac.user,                                                  
        "iss": sets.audience,                                            
        "exp": now + 3600,                                                 
        "iat": now,
        "cnf.jkt": base64.b64encode(hashlib.sha256(jwk).digest()).decode(),
        "typ": "dpop"
    }
    access_token = jwt.encode(payload, get_settings().authz_pvt_key.replace(
        '\\n', '\n').replace('\\t', '\t'), algorithm="RS256")

    return JSONResponse(
        status_code=status.HTTP_200_OK,
        content=jsonable_encoder({"access_token": access_token,
                                  "token_type": "DPoP",
                                  "expires_in": 3600}) 
    )