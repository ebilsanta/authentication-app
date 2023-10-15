from typing import Union

from fastapi import FastAPI, Response, status
from fastapi.responses import RedirectResponse
from pydantic import BaseModel
from requests.models import PreparedRequest

app = FastAPI()

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
        req = PreparedRequest()
        req.prepare_url(redirect_url, {'error':'unsupported_response_type'})
        print(req.url)
        return RedirectResponse(url=req.url, status_code=302) 

    return "ok"

