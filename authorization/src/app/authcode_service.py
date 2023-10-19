from functools import lru_cache
import os
from requests.models import PreparedRequest
from fastapi.responses import RedirectResponse
import jwt
from config import Settings

@lru_cache()
def get_settings():
    return Settings()

class AuthCodeService:
    def __init__(self):
        sets = get_settings()

        self.allowed_client = sets.allowed_client
        self.allowed_issuer = sets.allowed_issuer
        self.pub_key = sets.pub_key.replace('\\n', '\n').replace('\\t', '\t')
        self.audience = sets.audience
        self.redirect = sets.allowed_redirect

    def is_client_allowed(self, client):
        return client == self.allowed_client
    
    def is_issuer_allowed(self, issuer):
        return issuer == self.allowed_issuer
    
    def is_redirect_valid(self, redir):
        return redir == self.redirect
    
    def make_error(self, url, err_msg):
        req = PreparedRequest()
        req.prepare_url(url, {'error':err_msg})
        return RedirectResponse(url=req.url, status_code=302) 
    
    def make_error_desc(self, url, err_msg, err_desc):
        req = PreparedRequest()
        req.prepare_url(url, {'error':err_msg, 'error_description':err_desc})
        return RedirectResponse(url=req.url, status_code=302) 
    
    def verify_jwt(self, id_jwt):
        try:
            decoded = jwt.decode(id_jwt, self.pub_key ,algorithms=['RS256'], audience=self.audience)
            if decoded['iss'] != self.allowed_issuer:
                return "Unknown Issuer", None
            return None, decoded
        except jwt.ExpiredSignatureError:
            return "JWT has expired", None
        except jwt.InvalidTokenError:
            return "Invalid JWT", None
        except Exception as e:
            return "An error occurred during JWT decoding:" + str(e), None
    
