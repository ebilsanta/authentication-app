import uuid
import os
import jwt
import time
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.primitives import serialization
from unittest.mock import patch
from app.pkce import generate_pkce_code_verifier, generate_pkce_code_challenge
from app.database import AuthCodeRecord
from app.setup_utils import build_url
from app.dpop_service import create_dpop_jwt

# Variables
allowed_client = 'testing_client'
issuer = 'authn.testing'
subject = 'testing@test.com'
audience = 'authz.testing'
redirect_url = 'http://localhost:8000'
authz_url = 'http://localhost:8000'

ac = 'abcdef0123456789abcdef0123456789'
state = uuid.uuid4().hex
expiry_delta = 300

# Setup

private_key = rsa.generate_private_key(
    public_exponent=65537,
    key_size=2048,
)

pvk = private_key.private_bytes(
        encoding=serialization.Encoding.PEM,
        format=serialization.PrivateFormat.TraditionalOpenSSL,
        encryption_algorithm=serialization.NoEncryption()
    )

pbk = private_key.public_key().public_bytes(
        encoding=serialization.Encoding.PEM,
        format=serialization.PublicFormat.SubjectPublicKeyInfo
    )

iat = time.time()
exp = iat + 3600
payload = {'iss': issuer,
           'sub': subject,
           'aud': audience,
           'iat': iat,
           'exp': exp}
additional_headers = {'kid': uuid.uuid4().hex}
public_key = private_key.public_key()
testing_jwt = jwt.encode(payload, private_key, headers=additional_headers, algorithm="RS256")

code_verifier = generate_pkce_code_verifier()
code_challenge = generate_pkce_code_challenge(code_verifier)

os.environ['ALLOWED_CLIENT'] = allowed_client
os.environ['ALLOWED_ISSUER'] = issuer
os.environ['ALLOWED_REDIRECT'] = redirect_url

os.environ['PUB_KEY'] = str(public_key.public_bytes(
    encoding=serialization.Encoding.PEM,
    format=serialization.PublicFormat.SubjectPublicKeyInfo
))[2:-1]

os.environ['DPOP_HTM'] = 'POST'
os.environ['DPOP_HTU'] = authz_url

os.environ['AUDIENCE'] = audience
os.environ['DB_URL'] = 'mongodb://localhost:27017'
os.environ['DB_NAME'] = 'authz'
os.environ['DB_COLLECTION_AUTHCODES'] = 'authcodes'
os.environ['DB_COLLECTION_USERS'] = 'credentials'



from fastapi.testclient import TestClient
from main import app

client = TestClient(app)

# Token Tests
@patch('main.db.get_authcode_record')
def test_token_ok(m1):
    m1.return_value = AuthCodeRecord(ac, state, code_challenge, subject, expiry_delta)

    dpop = create_dpop_jwt(pvk, pbk, authz_url, 'POST')[1:-1]
    ca = 'client_assertion'
    params = {
        'authcode': ac,
        'dpop': dpop,
        'client_assertion': ca,
        'redirect_url': redirect_url,
        'code_verifier': code_verifier
    }

    response = client.post(build_url('/token', params), follow_redirects=False)
    print(response)
    print(response.headers)