
from fastapi import FastAPI, Response

from app.AuthCodeService import AuthCodeService


app = FastAPI()
ac = AuthCodeService()

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
                        state: str, code_challenge: str, id_jwt: str, response: Response):
    if response_type != "code":
        return ac.make_error(redirect_url, 'unsupported_response_type')
    
    if not ac.is_client_allowed(client_id):
        return ac.make_error(redirect_url, 'unauthorized_client')

    if ac.verify_jwt(id_jwt):
        return ac.make_error(redirect_url, 'access_denied')
    
    # Generate code
    # Persist [Code, State, Challenge, expiry]

    return "ok"

