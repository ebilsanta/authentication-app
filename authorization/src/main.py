from typing import Union

from app.jwks import update_authN_key
from app.models import TokenRequest, TokenIntrospectionRequest, RefreshRequest
from app.process_reqs import (
    introspect,
    process_authcode,
    process_token,
    process_refresh,
)
from app.sqs_service import SQS_Service
from fastapi import FastAPI, Request, status
from fastapi.encoders import jsonable_encoder
from fastapi.exceptions import RequestValidationError
from fastapi.responses import JSONResponse
from fastapi_restful.tasks import repeat_every

sqss = SQS_Service()
app = FastAPI()


@app.get("/health")
def read_root():
    return {"Hello": "World"}


# response_type: str              RFC 6749 4.1.1. Value MUST be set to "code"
# client_id: str                  RFC 6749 4.1.1
# redirect_url: str               RFC 6749 4.1.1
# state: str                      RFC 6749 4.1.1 & 10.12
# code_challenge: str             RFC 7636 4.3


@app.post("/authcode")
async def post_authcode(
    response_type: str,
    client_id: str,
    state: str,
    id_jwt: str,
    code_challenge: Union[str, None] = None,
    code_challenge_method: Union[str, None] = None,
    redirect_url: Union[str, None] = None,
):
    return await process_authcode(
        response_type,
        client_id,
        state,
        id_jwt,
        code_challenge,
        code_challenge_method,
        redirect_url,
    )


@app.exception_handler(RequestValidationError)
async def standard_validation_exception_handler(
    request: Request, exc: RequestValidationError
):
    print(request, exc)
    return JSONResponse(
        status_code=status.HTTP_400_BAD_REQUEST,
        content=jsonable_encoder(
            {
                "error": "invalid_request",
                "error_description": "The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed.",
            }
        ),
    )


@app.post("/token")
async def post_token(token_req: TokenRequest):
    return await process_token(
        token_req.grant_type,
        token_req.authcode,
        token_req.dpop,
        token_req.client_assertion,
        token_req.redirect_url,
        token_req.code_verifier,
    )


@app.post("/refresh")
async def post_refresh(refresh_req: RefreshRequest):
    return await process_refresh(
        refresh_req.grant_type, refresh_req.dpop, refresh_req.refresh_token
    )


@app.post("/introspect")
async def post_introspect(token: TokenIntrospectionRequest):
    print(token.token)
    return await introspect(token.token)


@app.on_event("startup")
@repeat_every(seconds=0.2)
async def pull_from_sqs():
    await sqss.poll_sqs()


@app.on_event("startup")
@repeat_every(seconds=10)
async def update_authN_pub_key():
    update_authN_key()
