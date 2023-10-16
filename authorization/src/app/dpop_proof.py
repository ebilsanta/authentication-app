import base64
import uuid
import jwt
import time
import json

def create_dpop_jwt(private_key, public_key, htu, ath=None, htm='POST'):
    header = {
        "alg": "RS256",
        "typ": "dpop+jwt",
        "jwk": base64.b64encode(public_key).decode("ascii")
    }

    payload = {
        "iat": int(time.time()),    # Creation time
        "jti": str(uuid.uuid4()),   # Unique identifier
        "htm": htm,                 # HTTP Method
        "htu": htu,                 # HTTP Target site w/o ? and fragments
    }

    if ath:
        # Base64 encoded SHA256 of associated access token's ASCII
        # Needed if access token also presented
        payload.update("ath", ath)  # Access Token's Hash

    token = jwt.encode(payload, private_key, algorithm='RS256', headers=header)

    return json.dumps(token)
