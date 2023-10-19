import os
import time
import jwt

class ClientAssertionService:
    def __init__(self) -> None:
        self.client_id = os.getenv('ALLOWED_CLIENT')

        self.client_pvt = os.getenv('ALLOWED_CLIENT_PVT_KEY')
        self.audience = os.getenv('AUDIENCE')

    def generate_client_assertion(self):
        current_time = int(time.time())
        return self.do_generate_client_assertion(self.client_id, self.audience, \
                                                 current_time + 300, \
                                          current_time, self.client_pvt)
    
    # RFC 7521
    def verify_client_assertion(self, assertion):
        print(self.client_id, self.client_pvt, self.audience)
        try:
            decoded_payload = jwt.decode(assertion, os.getenv('ALLOWED_CLIENT_PUB_KEY').replace('\\n', '\n').replace('\\t', '\t'), algorithms=["RS256"])

            if decoded_payload['iss'] != self.client_id:
                return 'invalid_client', 'invalid_issuer'
            if decoded_payload['sub'] != self.client_id:
                return 'invalid_client', 'invalid_subject'
            if decoded_payload['aud'] != self.audience:
                return 'invalid_client', 'invalid_audience'

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