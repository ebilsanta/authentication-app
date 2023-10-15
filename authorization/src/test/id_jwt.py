import jwt
import time
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.primitives import serialization

private_key = rsa.generate_private_key(
    public_exponent=65537,
    key_size=2048,
)

iat = time.time()
exp = iat + 3600
payload = {'iss': 'authn.gandalf',
           'sub': 'user1@gmail.com',
           'aud': 'authz.gandalf',
           'iat': iat,
           'exp': exp}

additional_headers = {'kid': 'kid'}
signed_jwt = jwt.encode(payload, private_key, headers=additional_headers,
                       algorithm='RS256')

print(signed_jwt)