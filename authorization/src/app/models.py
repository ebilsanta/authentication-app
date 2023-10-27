from pydantic import BaseModel

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