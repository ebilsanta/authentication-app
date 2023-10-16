import uuid
from dotenv import load_dotenv
import os
import jwt
import time
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.primitives import serialization
from unittest.mock import patch
from app.pkce import generate_pkce_code_verifier, generate_pkce_code_challenge

# Variables
allowed_client = 'testing_client'
issuer = 'authn.testing'
subject = 'testing@test.com'
audience = 'authz.testing'
redirect_url = 'http://localhost:8000'

# Setup

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

os.environ['AUDIENCE'] = audience
os.environ['DB_URL'] = 'mongodb://localhost:27017'
os.environ['DB_NAME'] = 'authz'
os.environ['DB_COLLECTION_AUTHCODES'] = 'authcodes'
os.environ['DB_COLLECTION_ALLOWED'] = 'allowed'

from fastapi.testclient import TestClient
from main import app

client = TestClient(app)

def build_url(url, params):
    ret = url + '?'
    for k in params:
        ret += k + '=' + params[k] + '&'
    return ret[:-1]

def is_uuid(test):
    try:
        uuid.UUID(str(test))
        return True
    except ValueError:
        return False

# Tests start here

def test_hello_world():
    response = client.get("/")
    assert response.status_code == 200
    assert response.json() == {"Hello": "World"}

@patch('main.db.insert_authcode_record')
def test_authcode_ok(db_mock):
    db_mock.return_value = 0
    params = {'response_type': 'code',
              'client_id': allowed_client,
              'redirect_url': redirect_url,
              'state': uuid.uuid4().hex,
              'code_challenge': code_challenge,
              'id_jwt': testing_jwt,
              'code_challenge_method': 'S256'}

    response = client.post(build_url('/authcode', params), follow_redirects=False)

    assert response.status_code == 302
    assert response.headers['location'].split('=')[0] == redirect_url + '?code'
    assert is_uuid(response.headers['location'].split('=')[1])

@patch('main.db.insert_authcode_record')
def test_authcode_no_redirect_ok(db_mock):
    db_mock.return_value = 0
    params = {'response_type': 'code',
              'client_id': allowed_client,
              'state': uuid.uuid4().hex,
              'code_challenge': code_challenge,
              'id_jwt': testing_jwt,
              'code_challenge_method': 'S256'}

    response = client.post(build_url('/authcode', params), follow_redirects=False)
    assert response.status_code == 200
    assert is_uuid(response.json()['code'])
    
def test_response_type_fail():
    params = {'response_type': 'nocode',
            'client_id': allowed_client,
            'redirect_url': redirect_url,
            'state': uuid.uuid4().hex,
            'code_challenge': code_challenge,
            'id_jwt': testing_jwt,
            'code_challenge_method': 'S256'}

    response = client.post(build_url('/authcode', params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers['location'].split('?')[0] == redirect_url + '/'
    assert response.headers['location'].split('?')[1] == 'error=unsupported_response_type'

def test_client_fail():
    params = {'response_type': 'code',
            'client_id': 'badclient',
            'redirect_url': redirect_url,
            'state': uuid.uuid4().hex,
            'code_challenge': code_challenge,
            'id_jwt': testing_jwt,
            'code_challenge_method': 'S256'}

    response = client.post(build_url('/authcode', params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers['location'].split('?')[0] == redirect_url + '/'
    assert response.headers['location'].split('?')[1] == 'error=unauthorized_client'

def test_code_challenge_empty_fail():
    params = {'response_type': 'code',
            'client_id': allowed_client,
            'redirect_url': redirect_url,
            'state': uuid.uuid4().hex,
            'id_jwt': testing_jwt,
            'code_challenge_method': 'S256'}

    response = client.post(build_url('/authcode', params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers['location'].split('?')[0] == redirect_url + '/'

    queries = response.headers['location'].split('?')[1].split('&')

    assert 'error=invalid_request' in queries
    assert 'error_description=code+challenge+required' in queries

def test_code_challenge_invalid_fail():
    params = {'response_type': 'code',
              'client_id': allowed_client,
              'redirect_url': redirect_url,
              'state': uuid.uuid4().hex,
              'code_challenge': 'aaa',
              'id_jwt': testing_jwt,
              'code_challenge_method': 'S256'}

    response = client.post(build_url('/authcode', params), follow_redirects=False)

    assert response.status_code == 302
    assert response.headers['location'].split('?')[0] == redirect_url + '/'

    queries = response.headers['location'].split('?')[1].split('&')

    assert 'error=invalid_request' in queries
    assert 'error_description=invalid+code+challenge' in queries

def test_code_challenge_method_empty_fail():
    params = {'response_type': 'code',
            'client_id': allowed_client,
            'redirect_url': redirect_url,
            'state': uuid.uuid4().hex,
            'code_challenge': code_challenge,
            'id_jwt': testing_jwt}

    response = client.post(build_url('/authcode', params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers['location'].split('?')[0] == redirect_url + '/'

    queries = response.headers['location'].split('?')[1].split('&')

    assert 'error=invalid_request' in queries
    assert 'error_description=transform+algorithm+not+supported' in queries

def test_code_challenge_method_invalid_fail():
    params = {'response_type': 'code',
            'client_id': allowed_client,
            'redirect_url': redirect_url,
            'state': uuid.uuid4().hex,
            'code_challenge': code_challenge,
            'id_jwt': testing_jwt,
            'code_challenge_method': 'S128'}

    response = client.post(build_url('/authcode', params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers['location'].split('?')[0] == redirect_url + '/'

    queries = response.headers['location'].split('?')[1].split('&')

    assert 'error=invalid_request' in queries
    assert 'error_description=transform+algorithm+not+supported' in queries

def test_invalid_jwt_fail():
    params = {'response_type': 'code',
            'client_id': allowed_client,
            'redirect_url': redirect_url,
            'state': uuid.uuid4().hex,
            'code_challenge': code_challenge,
            'id_jwt': 'abc',
            'code_challenge_method': 'S256'}

    response = client.post(build_url('/authcode', params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers['location'].split('?')[0] == redirect_url + '/'

    queries = response.headers['location'].split('?')[1].split('&')

    assert 'error=access_denied' in queries
    assert 'error_description=Invalid+JWT' in queries

def test_expired_jwt_fail():    
    iat = 0
    exp = 1
    payload = {'iss': issuer,
            'sub': subject,
            'aud': audience,
            'iat': iat,
            'exp': exp}
    additional_headers = {'kid': uuid.uuid4().hex}
    expired_jwt = jwt.encode(payload, private_key, headers=additional_headers, algorithm="RS256")

    params = {'response_type': 'code',
    'client_id': allowed_client,
    'redirect_url': redirect_url,
    'state': uuid.uuid4().hex,
    'code_challenge': code_challenge,
    'id_jwt': expired_jwt,
    'code_challenge_method': 'S256'}

    response = client.post(build_url('/authcode', params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers['location'].split('?')[0] == redirect_url + '/'

    queries = response.headers['location'].split('?')[1].split('&')

    assert 'error=access_denied' in queries
    assert 'error_description=JWT+has+expired' in queries

def test_jwt_future_fail():    
    iat = time.time()+100000
    exp = iat+3600
    payload = {'iss': issuer,
            'sub': subject,
            'aud': audience,
            'iat': iat,
            'exp': exp}
    additional_headers = {'kid': uuid.uuid4().hex}
    future_jwt = jwt.encode(payload, private_key, headers=additional_headers, algorithm="RS256")

    params = {'response_type': 'code',
    'client_id': allowed_client,
    'redirect_url': redirect_url,
    'state': uuid.uuid4().hex,
    'code_challenge': code_challenge,
    'id_jwt': future_jwt,
    'code_challenge_method': 'S256'}

    response = client.post(build_url('/authcode', params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers['location'].split('?')[0] == redirect_url + '/'

    queries = response.headers['location'].split('?')[1].split('&')

    assert 'error=access_denied' in queries
    assert 'error_description=Invalid+JWT' in queries

def test_jwt_invalid_issuer_fail():    
    iat = time.time()
    exp = iat + 3600
    payload = {'iss': 'newperson',
            'sub': subject,
            'aud': audience,
            'iat': iat,
            'exp': exp}
    additional_headers = {'kid': uuid.uuid4().hex}
    invalid_jwt = jwt.encode(payload, private_key, headers=additional_headers, algorithm="RS256")

    params = {'response_type': 'code',
    'client_id': allowed_client,
    'redirect_url': redirect_url,
    'state': uuid.uuid4().hex,
    'code_challenge': code_challenge,
    'id_jwt': invalid_jwt,
    'code_challenge_method': 'S256'}

    response = client.post(build_url('/authcode', params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers['location'].split('?')[0] == redirect_url + '/'

    queries = response.headers['location'].split('?')[1].split('&')

    assert 'error=access_denied' in queries
    assert 'error_description=Unknown+Issuer' in queries

def test_state_invalid_fail():
    params = {'response_type': 'code',
              'client_id': allowed_client,
              'redirect_url': redirect_url,
              'state': 'NewYork',
              'code_challenge': code_challenge,
              'id_jwt': testing_jwt,
              'code_challenge_method': 'S256'}

    response = client.post(build_url('/authcode', params), follow_redirects=False)

    assert response.status_code == 302
    assert response.headers['location'].split('?')[0] == redirect_url + '/'

    queries = response.headers['location'].split('?')[1].split('&')

    assert 'error=invalid_request' in queries
    assert 'error_description=invalid+state' in queries

def test_missing_params_fail():
    params = {'response_type': 'code',
              'code_challenge_method': 'S256'}

    response = client.post(build_url('/authcode', params), follow_redirects=False)

    assert response.status_code == 400
    assert response.json()['error'] == 'invalid_request'
    assert response.json()['error_description'] == 'The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed.'
