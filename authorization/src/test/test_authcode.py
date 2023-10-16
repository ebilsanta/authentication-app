import uuid
from dotenv import load_dotenv
import os
import jwt
import time
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.primitives import serialization
from requests.models import PreparedRequest

# Variables
allowed_client = 'testing_client'
issuer = 'authn.testing'
subject = 'testing@test.com'
audience = 'authz.testing'
redirect_url = 'http://localhost:8000'

private_key = rsa.generate_private_key(
    public_exponent=65537,
    key_size=2048,
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

os.environ['ALLOWED_CLIENT'] = allowed_client
os.environ['ALLOWED_ISSUER'] = issuer
os.environ['PUB_KEY'] = str(public_key.public_bytes(
    encoding=serialization.Encoding.PEM,
    format=serialization.PublicFormat.SubjectPublicKeyInfo
))

os.environ['AUDIENCE'] = audience
os.environ['DB_URL'] = 'mongodb://localhost:27017'
os.environ['DB_NAME'] = 'authz'
os.environ['DB_COLLECTION_AUTHCODES'] = 'authcodes'
os.environ['DB_COLLECTION_ALLOWED'] = 'allowed'

testing_jwt = encoded = jwt.encode(payload, private_key, headers=additional_headers, algorithm="RS256")

from fastapi.testclient import TestClient
from main import app

client = TestClient(app)

def build_url(url, params):
    ret = url + '?'
    for k in params:
        ret += k + '=' + params[k] + '&'
    return ret[:-1]

def test_hello_world():
    response = client.get("/")
    assert response.status_code == 200
    assert response.json() == {"Hello": "World"}

def test_authcode_ok():
    params = {'response_type': 'code',
              'client_id': allowed_client,
              'redirect_url': redirect_url,
              'state': uuid.uuid4().hex,
              'code_challenge': uuid.uuid4().hex,
              'id_jwt': testing_jwt,
              'code_challenge_method': 'S256'}

    response = client.post(build_url('/authcode', params))
    assert response.status_code == 200
    print(response.json)
    