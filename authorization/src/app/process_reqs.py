import base64
import hashlib
import time
import uuid
from functools import lru_cache
from typing import Union

import jwt
from app.authcode_service import AuthCodeService
from app.client_assertion_service import ClientAssertionService
from app.database import AuthCodeRecord, Database
from app.dpop_service import DpopService
from app.pkce import generate_pkce_code_challenge
from config import Settings
from fastapi import status
from fastapi.encoders import jsonable_encoder
from fastapi.responses import JSONResponse, RedirectResponse


@lru_cache()
def get_settings():
    return Settings()


ac = AuthCodeService()
dps = DpopService()
db = Database()
cas = ClientAssertionService()


async def process_authcode(
    response_type: str,
    client_id: str,
    state: str,
    id_jwt: str,
    code_challenge: Union[str, None] = None,
    code_challenge_method: Union[str, None] = None,
    redirect_url: Union[str, None] = None,
):
    def respond(redirect_url, message, desc=None):
        if redirect_url and desc:
            return ac.make_error_desc(redirect_url, message, desc)
        elif redirect_url:
            return ac.make_error(redirect_url, message)
        elif desc:
            return JSONResponse(
                status_code=status.HTTP_400_BAD_REQUEST,
                content=jsonable_encoder({"error": message, "error_description": desc}),
            )
        else:
            return JSONResponse(
                status_code=status.HTTP_400_BAD_REQUEST,
                content=jsonable_encoder({"error": message}),
            )

    # Verify redirection url is registered
    # RFC 6749 10.6 check redirection url is same as get_auth_code
    if redirect_url and not ac.is_redirect_valid(redirect_url):
        return respond(redirect_url, "access_denied", "invalid_redirect")

    if response_type != "code":
        return respond(redirect_url, "unsupported_response_type")

    if not ac.is_client_allowed(client_id):
        return respond(redirect_url, "unauthorized_client")

    if not code_challenge:
        return respond(redirect_url, "invalid_request", desc="code challenge required")
    elif len(code_challenge) != 44:
        return respond(redirect_url, "invalid_request", desc="invalid code challenge")

    if code_challenge_method != "S256":
        return respond(
            redirect_url, "invalid_request", desc="transform algorithm not supported"
        )

    if len(state) < 16 or len(state) > 64:
        return respond(redirect_url, "invalid_request", desc="invalid state")

    err_message, decoded = ac.verify_jwt(id_jwt)
    if err_message:
        return respond(redirect_url, "access_denied", err_message)

    # Check if user is actually in our system first
    if not await db.exists_valid_user(decoded["sub"]):
        return respond(redirect_url, "access_denied", "unknown user")

    code = uuid.uuid4().hex

    await db.insert_authcode_record(
        AuthCodeRecord(code, state, code_challenge, decoded["sub"], 600)
    )
    if redirect_url:
        return RedirectResponse(url=redirect_url + "?code=" + code, status_code=302)
    else:
        return JSONResponse(
            status_code=status.HTTP_200_OK,
            content=jsonable_encoder({"code": code}),
        )


async def process_token(grant_type, authcode, dpop, client_assertion, redirect_url, code_verifier):
    def respond(message, desc=None):
        if desc:
            return JSONResponse(
                status_code=status.HTTP_400_BAD_REQUEST,
                content=jsonable_encoder({"error": message, "error_description": desc}),
            )
        else:
            return JSONResponse(
                status_code=status.HTTP_400_BAD_REQUEST,
                content=jsonable_encoder({"error": message}),
            )

    if grant_type != "authorization_code":
        return respond("unsupported grant type")

    err, err_desc = cas.verify_client_assertion(client_assertion)
    if err:
        return respond(err, err_desc)

    jwk, err = dps.verify_dpop(dpop)
    if err:
        return respond(err)

    authc = await db.get_authcode_record(authcode)

    print(generate_pkce_code_challenge(code_verifier))
    if generate_pkce_code_challenge(code_verifier) != authc['code_challenge']:
        return respond("Invalid PKCE Code Verifier")

    sets = get_settings()

    now = int(time.time())
    payload = {
        "sub": authc['user'],
        "iss": sets.audience,
        "exp": now + 3600,
        "iat": now,
        "cnf.jkt": base64.b64encode(hashlib.sha256(jwk).digest()).decode(),
    }
    access_token = jwt.encode(
        payload,
        get_settings().authz_pvt_key.replace("\\n", "\n").replace("\\t", "\t"),
        algorithm="RS256",
        headers={"typ": "dpop"},
    )

    payload = {
        "sub": authc['user'],
        "iss": sets.audience,
        "exp": now + 86400,
        "iat": now,
        "cnf.jkt": base64.b64encode(hashlib.sha256(jwk).digest()).decode(),
    }

    refresh_token = jwt.encode(
        payload,
        get_settings().authz_pvt_key.replace("\\n", "\n").replace("\\t", "\t"),
        algorithm="RS256",
        headers={"typ": "dpop+refresh"},  # Enforce a check
    )
    if redirect_url:
        return JSONResponse(content=jsonable_encoder(
                {
                    "access_token": access_token,
                    "token_type": "DPoP",
                    "expires_in": 3600,
                    "refresh_token": refresh_token,
                }
            ), 
            headers={'Location': redirect_url}, 
            status_code=status.HTTP_302_FOUND)
    else:
        return JSONResponse(
            status_code=status.HTTP_200_OK,
            content=jsonable_encoder(
                {
                    "access_token": access_token,
                    "token_type": "DPoP",
                    "expires_in": 3600,
                    "refresh_token": refresh_token,
                }
            ),
        )


async def process_refresh(grant_type, dpop, refresh_token):
    def respond(message, desc=None):
        if desc:
            return JSONResponse(
                status_code=status.HTTP_400_BAD_REQUEST,
                content=jsonable_encoder({"error": message, "error_description": desc}),
            )
        else:
            return JSONResponse(
                status_code=status.HTTP_400_BAD_REQUEST,
                content=jsonable_encoder({"error": message}),
            )

    if grant_type != "authorization_code":
        return respond("unsupported grant type")

    err, jwk = dps.verify_dpop(dpop, at=refresh_token)

    if err:
        print(err)
        return respond(err)

    decoded_ref = jwt.decode(
        refresh_token,
        get_settings().authz_pub_key.replace("\\n", "\n").replace("\\t", "\t"),
        algorithms=["RS256"],
    )

    sets = get_settings()

    now = int(time.time())
    payload = {
        "sub": decoded_ref["sub"],
        "iss": sets.audience,
        "exp": now + 3600,
        "iat": now,
        "cnf.jkt": base64.b64encode(
            hashlib.sha256(
                jwk.encode("ascii")
            ).digest()  # Confirm this is ok, compared to access token's impl
        ).decode(),
    }
    access_token = jwt.encode(
        payload,
        bytes(
            get_settings().authz_pvt_key.replace("\\n", "\n").replace("\\t", "\t"),
            "ascii",
        ),
        algorithm="RS256",
        headers={"typ": "dpop"},
    )

    return JSONResponse(
        status_code=status.HTTP_200_OK,
        content=jsonable_encoder(
            {"access_token": access_token, "token_type": "DPoP", "expires_in": 3600}
        ),
    )
