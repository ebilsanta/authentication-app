import jwt
from main import app
from fastapi.testclient import TestClient
import uuid
from dotenv import load_dotenv
import time
from unittest.mock import patch
from app.pkce import generate_pkce_code_verifier, generate_pkce_code_challenge
from app.database import AuthCodeRecord
from app.dpop_service import create_dpop_jwt
from app.keypair_gen import gen_keypair
from app.client_assertion_service import do_generate_client_assertion
from config import get_settings

# Variables
subject = 'testing@test.com'
sets = get_settings()
pvk = bytes(sets.allowed_client_pvt_key.replace(
    '\\n', '\n').replace('\\t', '\t'), 'utf-8')
pbk = bytes(sets.allowed_client_pub_key.replace(
    '\\n', '\n').replace('\\t', '\t'), 'utf-8')

client = TestClient(app)
code_verifier = generate_pkce_code_verifier()
ac = '0123456789abcdef0123456789abcdef'

# Token Tests

@patch('main.db.get_authcode_record')
def test_token_ok(m1):
    m1.return_value = AuthCodeRecord(ac, uuid.uuid4().hex,
                                     generate_pkce_code_challenge(code_verifier), subject, 300)

    dpop = create_dpop_jwt(pvk, pbk, sets.authz_url, 'POST')[1:-1]

    now = int(time.time())
    ca = do_generate_client_assertion(
        sets.allowed_client, sets.audience, now+300, now, pvk)
    params = {
        'grant_type': 'authorization_code',
        'authcode': ac,
        'dpop': dpop,
        'client_assertion': ca,
        'redirect_url': sets.allowed_redirect,
        'code_verifier': code_verifier
    }

    response = client.post('/token', json=params, follow_redirects=False)
    assert response.status_code == 200
    assert response.json()['token_type']=='DPoP'
    assert response.json()['access_token']
    assert response.json()['refresh_token']

def test_token_invalid_grant_failure():
    dpop = create_dpop_jwt(pvk, pbk, sets.authz_url, 'POST')[1:-1]

    now = int(time.time())
    ca = do_generate_client_assertion(
        sets.allowed_client, sets.audience, now+300, now, pvk)
    params = {
        'grant_type': 'funny_method',
        'authcode': ac,
        'dpop': dpop,
        'client_assertion': ca,
        'redirect_url': sets.allowed_redirect,
        'code_verifier': code_verifier
    }

    response = client.post('/token', json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()['error']=='unsupported grant type'

def test_token_invalid_ca_failure():
    dpop = create_dpop_jwt(pvk, pbk, sets.authz_url, 'POST')[1:-1]

    now = int(time.time())
    params = {
        'grant_type': 'authorization_code',
        'authcode': ac,
        'dpop': dpop,
        'client_assertion': 'beans',
        'redirect_url': sets.allowed_redirect,
        'code_verifier': code_verifier
    }

    response = client.post('/token', json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()['error']=='invalid_client'

def test_token_expired_ca_failure():
    dpop = create_dpop_jwt(pvk, pbk, sets.authz_url, 'POST')[1:-1]

    now = int(time.time())
    ca = do_generate_client_assertion(
        sets.allowed_client, sets.audience, now-100, now-10, pvk)
    params = {
        'grant_type': 'authorization_code',
        'authcode': ac,
        'dpop': dpop,
        'client_assertion': ca,
        'redirect_url': sets.allowed_redirect,
        'code_verifier': code_verifier
    }

    response = client.post('/token', json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()['error']=='invalid_client'
    assert response.json()['error_description']=='assertion has expired'

def test_token_invalid_issuer_ca_failure():
    dpop = create_dpop_jwt(pvk, pbk, sets.authz_url, 'POST')[1:-1]

    now = int(time.time())
    ca = do_generate_client_assertion(
        'badguy', sets.audience, now+300, now, pvk)
    params = {
        'grant_type': 'authorization_code',
        'authcode': ac,
        'dpop': dpop,
        'client_assertion': ca,
        'redirect_url': sets.allowed_redirect,
        'code_verifier': code_verifier
    }

    response = client.post('/token', json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()['error']=='invalid_client'
    assert response.json()['error_description']=='invalid_issuer'

def test_token_invalid_subject_ca_failure():
    dpop = create_dpop_jwt(pvk, pbk, sets.authz_url, 'POST')[1:-1]

    now = int(time.time())
    ca = jwt.encode({
        "iss": sets.allowed_client,
        "sub": 'who',
        "aud": sets.audience,
        "exp": now+300,
        "iat": now
    }, pvk, algorithm="RS256")
    params = {
        'grant_type': 'authorization_code',
        'authcode': ac,
        'dpop': dpop,
        'client_assertion': ca,
        'redirect_url': sets.allowed_redirect,
        'code_verifier': code_verifier
    }

    response = client.post('/token', json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()['error']=='invalid_client'
    assert response.json()['error_description']=='invalid_subject'

def test_token_invalid_audience_ca_failure():
    dpop = create_dpop_jwt(pvk, pbk, sets.authz_url, 'POST')[1:-1]

    now = int(time.time())
    ca = do_generate_client_assertion(
        sets.allowed_client, 'notaudience', now+300, now, pvk)
    params = {
        'grant_type': 'authorization_code',
        'authcode': ac,
        'dpop': dpop,
        'client_assertion': ca,
        'redirect_url': sets.allowed_redirect,
        'code_verifier': code_verifier
    }

    response = client.post('/token', json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()['error']=='invalid_client'
    assert response.json()['error_description']=='invalid_audience'