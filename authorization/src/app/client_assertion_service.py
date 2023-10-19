from cryptography.hazmat.primitives.asymmetric import rsa
from functools import lru_cache
import time
from fastapi import Depends
import jwt
from config import Settings

@lru_cache()
def get_settings():
    return Settings()

class ClientAssertionService:
    def __init__(self) -> None:
        sets = get_settings()
        self.client_id = sets.allowed_client
        self.client_pub = sets.allowed_client_pub_key.replace('\\n', '\n').replace('\\t', '\t')
        self.client_pvt = sets.allowed_client_pvt_key.replace('\\n', '\n').replace('\\t', '\t')
        self.audience = sets.audience

    def generate_client_assertion(self):
        current_time = int(time.time())
        return do_generate_client_assertion(self.client_id, self.audience, \
                                                 current_time + 300, \
                                          current_time, self.client_pvt)
    
    # RFC 7521
    def verify_client_assertion(self, assertion):
        try:
            decoded_payload = jwt.decode(assertion, self.client_pub, algorithms=["RS256"], audience=self.audience)

            if decoded_payload['iss'] != self.client_id:
                return 'invalid_client', 'invalid_issuer'
            if decoded_payload['sub'] != self.client_id:
                return 'invalid_client', 'invalid_subject'
            if decoded_payload['aud'] != self.audience:
                return 'invalid_client', 'invalid_audience'
            
            return None, None

        except jwt.ExpiredSignatureError:
            return 'invalid_client', 'assertion has expired'
        except jwt.InvalidTokenError as e:
            return 'invalid_client', 'Error verifying Client Assertion: ' + str(e)
        

def do_generate_client_assertion(client_id, audience, expiry, iat, client_pvt):
    payload = {
        "iss": client_id,
        "sub": client_id,
        "aud": audience,
        "exp": expiry,
        "iat": iat
    }

    return jwt.encode(payload, client_pvt, algorithm="RS256")