# import time

# import jwt
# from cryptography.hazmat.primitives import serialization
# from cryptography.hazmat.primitives.asymmetric import rsa

# private_key = rsa.generate_private_key(
#     public_exponent=65537,
#     key_size=2048,
# )

# iat = time.time()
# exp = iat + 36000
# payload = {
#     "iss": "authn.gandalf",
#     "sub": "user1@gmail.com",
#     "aud": "authz.gandalf",
#     "iat": iat,
#     "exp": exp,
# }

# additional_headers = {"kid": "kid"}

# print()

# public_key = private_key.public_key()

# encoded = jwt.encode(
#     payload, private_key, headers=additional_headers, algorithm="RS256"
# )
# print(encoded)

# decoded = jwt.decode(
#     encoded, public_key, algorithms=["RS256"], audience="authz.gandalf"
# )
# {"some": "payload"}

# pbk = public_key.public_bytes(
#     encoding=serialization.Encoding.PEM,
#     format=serialization.PublicFormat.SubjectPublicKeyInfo,
# )

# pvk = private_key.private_bytes(
#     encoding=serialization.Encoding.PEM,
#     format=serialization.PrivateFormat.TraditionalOpenSSL,
#     encryption_algorithm=serialization.NoEncryption(),
# )

# f = open("pvk.pem", "a")
# f.write(str(pvk))
# f.close()

# f = open("pbk.pem", "a")
# f.write(str(pbk))
# f.close()
import base64
import hashlib
import os
import random


def generate_pkce_code_verifier():
    length = random.randint(43, 128)
    random_bytes = os.urandom(length // 2)
    code_verifier = base64.urlsafe_b64encode(random_bytes).decode("ascii")[:length]
    return code_verifier


def generate_pkce_code_challenge(code_verifier):
    encoder = hashlib.sha256()
    encoder.update(bytes(code_verifier, "ascii"))
    hash = encoder.digest()

    code_challenge = base64.urlsafe_b64encode(hash).decode("ascii").replace("=", "")
    return code_challenge
