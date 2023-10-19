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

from fastapi.testclient import TestClient
from main import app

# Token Tests
@patch('main.db.get_authcode_record')
def test_token_ok(m1):
    sets = get_settings()
    pvk = bytes(sets.allowed_client_pvt_key.replace('\\n', '\n').replace('\\t', '\t'), 'utf-8')
    pbk = bytes(sets.allowed_client_pub_key.replace('\\n', '\n').replace('\\t', '\t'), 'utf-8')

    client = TestClient(app)
    code_verifier = generate_pkce_code_verifier()
    ac = '0123456789abcdef0123456789abcdef'
    m1.return_value = AuthCodeRecord(ac, uuid.uuid4().hex, \
                                     generate_pkce_code_challenge(code_verifier), subject, 300)


    dpop = create_dpop_jwt(pvk, pbk, sets.authz_url, 'POST')[1:-1]

    now = int(time.time())
    ca = do_generate_client_assertion(sets.allowed_client, sets.audience, now+300, now, pvk)
    params = {
        'grant_type': 'authorization_code',
        'authcode': ac,
        'dpop': dpop,
        'client_assertion': ca,
        'redirect_url': sets.allowed_redirect,
        'code_verifier': code_verifier
    }

    response = client.post('/token', json=params, follow_redirects=False)
    print(response)
    print(response.headers)
    print(response.content)