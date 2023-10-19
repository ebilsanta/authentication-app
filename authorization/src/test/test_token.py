import uuid
import os
from dotenv import load_dotenv
import jwt
import time
from unittest.mock import patch
import pytest
from app.pkce import generate_pkce_code_verifier, generate_pkce_code_challenge
from app.database import AuthCodeRecord
from app.setup_utils import build_url
from app.dpop_service import create_dpop_jwt
from app.keypair_gen import gen_keypair
from app.client_assertion_service import do_generate_client_assertion

# Variables
allowed_client = 'testing_client'
issuer = 'authn.testing'
subject = 'testing@test.com'
audience = 'authz.testing'
redirect_url = 'http://localhost:8000'
authz_url = 'http://localhost:8000'

from fastapi.testclient import TestClient
from main import app

_, _, ca_pvk, ca_pbk = gen_keypair()
_, _, pvk, pbk = gen_keypair()

@pytest.fixture(scope='session', autouse=True)
def env_fixture() -> pytest.fixture():
    load_dotenv('test.env')

# Token Tests
@patch('main.db.get_authcode_record')
def test_token_ok(m1, env_fixture):
    client = TestClient(app)
    code_verifier = generate_pkce_code_verifier()
    ac = '0123456789abcdef0123456789abcdef'
    m1.return_value = AuthCodeRecord(ac, uuid.uuid4().hex, \
                                     generate_pkce_code_challenge(code_verifier), subject, 300)


    dpop = create_dpop_jwt(pvk, pbk, authz_url, 'POST')[1:-1]

    now = int(time.time())
    ca = do_generate_client_assertion(allowed_client, audience, now+300, now, pvk)
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