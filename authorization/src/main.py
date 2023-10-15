
from typing import Union
from fastapi import FastAPI
import uuid
from app.AuthCodeService import AuthCodeService
from app.Database import Database, AuthCodeRecord

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
async def post_authcode(response_type: str, client_id: str, redirect_url: str, 
                        state: str, id_jwt: str,
                        code_challenge: Union[str, None] = None, 
                        code_challenge_method:Union[str, None] = None):
    if response_type != "code":
        return ac.make_error(redirect_url, 'unsupported_response_type')
    
    if not ac.is_client_allowed(client_id):
        return ac.make_error(redirect_url, 'unauthorized_client')
    
    if not code_challenge:
        return ac.make_error_desc(redirect_url, 'invalid_request', 'code challenge required')
    
    if code_challenge_method != "S256":
        return ac.make_error_desc(redirect_url, 'invalid_request', 'transform algorithm not supported')

    err_message = ac.verify_jwt(id_jwt)
    if err_message:
        return ac.make_error(redirect_url, 'access_denied', err_message)
    
    code = uuid.uuid4().hex

    await db.insert_authcode_record(AuthCodeRecord(code, state, code_challenge))
    return code

