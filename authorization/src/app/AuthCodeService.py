import os
from requests.models import PreparedRequest
from fastapi.responses import RedirectResponse
import jwt

class AuthCodeService:
    def __init__(self):
        self.allowed_client = os.getenv('ALLOWED_CLIENT')
        self.allowed_issuer = os.getenv('ALLOWED_ISSUER')
        self.pub_key = os.getenv('PUB_KEY').replace('\\n', '\n').replace('\\t', '\t')
        self.audience = os.getenv('AUDIENCE')

    def is_client_allowed(self, client):
        return client == self.allowed_client
    
    def is_issuer_allowed(self, issuer):
        return issuer == self.allowed_issuer
    
    def make_error(self, url, err_msg):
        req = PreparedRequest()
        req.prepare_url(url, {'error':err_msg})
        return RedirectResponse(url=req.url, status_code=302) 
    
    def verify_jwt(self, id_jwt):
        try:
            decoded = jwt.decode(id_jwt, self.pub_key ,algorithms=['RS256'], audience=self.audience)
            if decoded.iss != self.allowed_issuer:
                return "Unknown Issuer"
        except jwt.ExpiredSignatureError:
            return "JWT has expired"
        except jwt.InvalidTokenError:
            return "Invalid JWT."
        except Exception as e:
            return "An error occurred during JWT decoding:" + str(e)
    
