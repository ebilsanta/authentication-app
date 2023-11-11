from typing import Optional
from pydantic import BaseModel


class TokenRequest(BaseModel):
    grant_type: str
    authcode: str
    dpop: str
    client_assertion: str
    code_verifier: str
    redirect_url: Optional[str] = None


class RefreshRequest(BaseModel):
    grant_type: str
    dpop: str
    refresh_token: str

class TokenIntrospectionRequest(BaseModel):
    token: str
