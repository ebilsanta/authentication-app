import jwt
import time
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.primitives import serialization

private_key = rsa.generate_private_key(
    public_exponent=65537,
    key_size=2048,
)

iat = time.time()
exp = iat + 36000
payload = {'iss': 'authn.gandalf',
           'sub': 'user1@gmail.com',
           'aud': 'authz.gandalf',
           'iat': iat,
           'exp': exp}

additional_headers = {'kid': 'kid'}

print()

public_key = private_key.public_key()

encoded = jwt.encode(payload, private_key,
                     headers=additional_headers, algorithm="RS256")
print(encoded)

decoded = jwt.decode(encoded, public_key, algorithms=[
                     "RS256"], audience='authz.gandalf')
{'some': 'payload'}

pbk = public_key.public_bytes(
    encoding=serialization.Encoding.PEM,
    format=serialization.PublicFormat.SubjectPublicKeyInfo
)

pvk = private_key.private_bytes(
    encoding=serialization.Encoding.PEM,
    format=serialization.PrivateFormat.TraditionalOpenSSL,
    encryption_algorithm=serialization.NoEncryption()
)

f = open("pvk.pem", "a")
f.write(str(pvk))
f.close()

f = open("pbk.pem", "a")
f.write(str(pbk))
f.close()
