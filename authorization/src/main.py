
from typing import Union
from fastapi import FastAPI, Request, status
from fastapi.encoders import jsonable_encoder
from fastapi.exceptions import RequestValidationError
from fastapi.responses import JSONResponse, RedirectResponse
import uuid
from app.authcode_service import AuthCodeService
from app.database import Database, AuthCodeRecord

app = FastAPI()
ac = AuthCodeService()
db = Database()

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
                        code_challenge_method:Union[str, None] = None,
                        redirect_url: Union[str, None] = None):
    
    def respond(redirect_url, message, desc = None):
        if redirect_url and desc:
            return ac.make_error_desc(redirect_url, message, desc)
        elif redirect_url:
            return ac.make_error(redirect_url, message)
        elif desc:
            return JSONResponse(
                status_code=status.HTTP_400_BAD_REQUEST,
                content=jsonable_encoder({'error': message, 'error_description': desc}),
            )
        else:
            return JSONResponse(
                status_code=status.HTTP_400_BAD_REQUEST,
                content=jsonable_encoder({'error': message}),
            )
    
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
        content=jsonable_encoder({'error': 'invalid_request', 'error_description': 'The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed.'}),
    )

