import base64
import uuid
import jwt
import time
import json
from functools import lru_cache
from config import Settings


@lru_cache()
def get_settings():
    return Settings()


class DpopService:
    def __init__(self):
        sets = get_settings()
        self.htu = sets.dpop_htu
        self.htm = sets.dpop_htm

    def create_dpop(self, private_key, public_key, ath=None):
        create_dpop_jwt(private_key, public_key, self.htu, self.htm, ath=ath)

    def verify_dpop(self, dpop_jwt):
        return verify_dpop_jwt(dpop_jwt, self.htu, self.htm)


def create_dpop_jwt(private_key, public_key, htu, htm, ath=None):
    header = {
        "alg": "RS256",
        "typ": "dpop+jwt",
        "jwk": base64.b64encode(public_key).decode("ascii")
    }

    now = int(time.time())
    payload = {
        "iat": now,                     # Creation time
        "jti": str(uuid.uuid4()),       # Unique identifier
        "htm": htm,                # HTTP Method
        "htu": htu,                # HTTP Target site w/o ? and fragments
        "exp": now + 120
    }

    if ath:
        # Base64 encoded SHA256 of associated access token's ASCII
        # Needed if access token also presented
        payload.update("ath", ath)  # Access Token's Hash

    token = jwt.encode(payload, private_key, algorithm='RS256', headers=header)

    return json.dumps(token)

# cnf.jkt holds the hash of the public key in the access token


def verify_dpop_jwt(dpop_jwt, htu, htm):
    try:
        uvh = jwt.get_unverified_header(dpop_jwt)
        if uvh['alg'] != 'RS256':
            return None, "Invalid algorithm"
        jwk = base64.b64decode(uvh['jwk'])

        decoded = jwt.decode(dpop_jwt, jwk, algorithms=['RS256'])

        if uvh['typ'] != 'dpop+jwt':
            return None, 'Invalid token type'
        if decoded['htm'] != htm:
            return None, 'Invalid dPoP HTTP Method'

        if decoded['htu'] != htu:
            return None, 'Invalid dPoP URL'

        return jwk, None

    except jwt.ExpiredSignatureError:
        return None, "JWT has expired"
    except jwt.InvalidTokenError:
        return None, "Invalid JWT"
    except Exception as e:
        return None, "An error occurred during JWT decoding:" + str(e)
