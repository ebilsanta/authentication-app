import base64
import hashlib
import time
import uuid
from unittest.mock import patch

import jwt
from app.client_assertion_service import do_generate_client_assertion
from app.database import AuthCodeRecord
from app.dpop_service import create_dpop_jwt, verify_dpop_jwt
from app.pkce import generate_pkce_code_challenge, generate_pkce_code_verifier
from config import get_settings
from fastapi.testclient import TestClient
from main import app

# Variables
subject = "testing@test.com"
sets = get_settings()
pvk = bytes(
    sets.allowed_client_pvt_key.replace("\\n", "\n").replace("\\t", "\t"), "utf-8"
)
pbk = bytes(
    sets.allowed_client_pub_key.replace("\\n", "\n").replace("\\t", "\t"), "utf-8"
)

client = TestClient(app)
code_verifier = generate_pkce_code_verifier()
ac = "0123456789abcdef0123456789abcdef"

# Token Tests


@patch("main.db.get_authcode_record")
def test_token_ok(m1):
    m1.return_value = AuthCodeRecord(
        ac, uuid.uuid4().hex, generate_pkce_code_challenge(code_verifier), subject, 300
    )

    dpop = create_dpop_jwt(pvk, pbk, sets.authz_url, "POST")[1:-1]

    now = int(time.time())
    ca = do_generate_client_assertion(
        sets.allowed_client, sets.audience, now + 300, now, pvk
    )
    params = {
        "grant_type": "authorization_code",
        "authcode": ac,
        "dpop": dpop,
        "client_assertion": ca,
        "redirect_url": sets.allowed_redirect,
        "code_verifier": code_verifier,
    }

    response = client.post("/token", json=params, follow_redirects=False)
    assert response.status_code == 200
    assert response.json()["token_type"] == "DPoP"
    assert response.json()["access_token"]
    assert response.json()["refresh_token"]


def test_token_invalid_grant_failure():
    dpop = create_dpop_jwt(pvk, pbk, sets.authz_url, "POST")[1:-1]

    now = int(time.time())
    ca = do_generate_client_assertion(
        sets.allowed_client, sets.audience, now + 300, now, pvk
    )
    params = {
        "grant_type": "funny_method",
        "authcode": ac,
        "dpop": dpop,
        "client_assertion": ca,
        "redirect_url": sets.allowed_redirect,
        "code_verifier": code_verifier,
    }

    response = client.post("/token", json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()["error"] == "unsupported grant type"


def test_token_invalid_ca_failure():
    dpop = create_dpop_jwt(pvk, pbk, sets.authz_url, "POST")[1:-1]

    now = int(time.time())
    params = {
        "grant_type": "authorization_code",
        "authcode": ac,
        "dpop": dpop,
        "client_assertion": "beans",
        "redirect_url": sets.allowed_redirect,
        "code_verifier": code_verifier,
    }

    response = client.post("/token", json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()["error"] == "invalid_client"


def test_token_expired_ca_failure():
    dpop = create_dpop_jwt(pvk, pbk, sets.authz_url, "POST")[1:-1]

    now = int(time.time())
    ca = do_generate_client_assertion(
        sets.allowed_client, sets.audience, now - 100, now - 10, pvk
    )
    params = {
        "grant_type": "authorization_code",
        "authcode": ac,
        "dpop": dpop,
        "client_assertion": ca,
        "redirect_url": sets.allowed_redirect,
        "code_verifier": code_verifier,
    }

    response = client.post("/token", json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()["error"] == "invalid_client"
    assert response.json()["error_description"] == "assertion has expired"


def test_token_invalid_issuer_ca_failure():
    dpop = create_dpop_jwt(pvk, pbk, sets.authz_url, "POST")[1:-1]

    now = int(time.time())
    ca = do_generate_client_assertion("badguy", sets.audience, now + 300, now, pvk)
    params = {
        "grant_type": "authorization_code",
        "authcode": ac,
        "dpop": dpop,
        "client_assertion": ca,
        "redirect_url": sets.allowed_redirect,
        "code_verifier": code_verifier,
    }

    response = client.post("/token", json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()["error"] == "invalid_client"
    assert response.json()["error_description"] == "invalid_issuer"


def test_token_invalid_subject_ca_failure():
    dpop = create_dpop_jwt(pvk, pbk, sets.authz_url, "POST")[1:-1]

    now = int(time.time())
    ca = jwt.encode(
        {
            "iss": sets.allowed_client,
            "sub": "who",
            "aud": sets.audience,
            "exp": now + 300,
            "iat": now,
        },
        pvk,
        algorithm="RS256",
    )
    params = {
        "grant_type": "authorization_code",
        "authcode": ac,
        "dpop": dpop,
        "client_assertion": ca,
        "redirect_url": sets.allowed_redirect,
        "code_verifier": code_verifier,
    }

    response = client.post("/token", json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()["error"] == "invalid_client"
    assert response.json()["error_description"] == "invalid_subject"


def test_token_invalid_audience_ca_failure():
    dpop = create_dpop_jwt(pvk, pbk, sets.authz_url, "POST")[1:-1]

    now = int(time.time())
    ca = do_generate_client_assertion(
        sets.allowed_client, "notaudience", now + 300, now, pvk
    )
    params = {
        "grant_type": "authorization_code",
        "authcode": ac,
        "dpop": dpop,
        "client_assertion": ca,
        "redirect_url": sets.allowed_redirect,
        "code_verifier": code_verifier,
    }

    response = client.post("/token", json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()["error"] == "invalid_client"
    assert response.json()["error_description"] == "invalid_audience"


def test_token_invalid_dpop_failure():
    now = int(time.time())
    ca = do_generate_client_assertion(
        sets.allowed_client, sets.audience, now + 300, now, pvk
    )
    params = {
        "grant_type": "authorization_code",
        "authcode": ac,
        "dpop": "not a dpop",
        "client_assertion": ca,
        "redirect_url": sets.allowed_redirect,
        "code_verifier": code_verifier,
    }

    response = client.post("/token", json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()["error"] == "Invalid JWT"


def test_token_invalid_dpop_typ_failure():
    header = {
        "alg": "RS256",
        "typ": "jwt",
        "jwk": base64.b64encode(pbk).decode("ascii"),
    }

    now = int(time.time())
    payload = {
        "iat": now,
        "jti": str(uuid.uuid4()),
        "htm": sets.dpop_htm,
        "htu": sets.dpop_htu,
        "exp": now + 120,
    }
    dpop = jwt.encode(payload, pvk, algorithm="RS256", headers=header)

    now = int(time.time())
    ca = do_generate_client_assertion(
        sets.allowed_client, sets.audience, now + 300, now, pvk
    )

    params = {
        "grant_type": "authorization_code",
        "authcode": ac,
        "dpop": dpop,
        "client_assertion": ca,
        "redirect_url": sets.allowed_redirect,
        "code_verifier": code_verifier,
    }

    response = client.post("/token", json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()["error"] == "Invalid token type"


def test_token_invalid_dpop_htm_failure():
    header = {
        "alg": "RS256",
        "typ": "dpop+jwt",
        "jwk": base64.b64encode(pbk).decode("ascii"),
    }

    now = int(time.time())
    payload = {
        "iat": now,
        "jti": str(uuid.uuid4()),
        "htm": "PATCH",
        "htu": sets.dpop_htu,
        "exp": now + 120,
    }
    dpop = jwt.encode(payload, pvk, algorithm="RS256", headers=header)

    now = int(time.time())
    ca = do_generate_client_assertion(
        sets.allowed_client, sets.audience, now + 300, now, pvk
    )

    params = {
        "grant_type": "authorization_code",
        "authcode": ac,
        "dpop": dpop,
        "client_assertion": ca,
        "redirect_url": sets.allowed_redirect,
        "code_verifier": code_verifier,
    }

    response = client.post("/token", json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()["error"] == "Invalid dPoP HTTP Method"


def test_token_invalid_dpop_htu_failure():
    header = {
        "alg": "RS256",
        "typ": "dpop+jwt",
        "jwk": base64.b64encode(pbk).decode("ascii"),
    }

    now = int(time.time())
    payload = {
        "iat": now,
        "jti": str(uuid.uuid4()),
        "htm": sets.dpop_htm,
        "htu": "www.wrongwebsite.com",
        "exp": now + 120,
    }
    dpop = jwt.encode(payload, pvk, algorithm="RS256", headers=header)

    now = int(time.time())
    ca = do_generate_client_assertion(
        sets.allowed_client, sets.audience, now + 300, now, pvk
    )

    params = {
        "grant_type": "authorization_code",
        "authcode": ac,
        "dpop": dpop,
        "client_assertion": ca,
        "redirect_url": sets.allowed_redirect,
        "code_verifier": code_verifier,
    }

    response = client.post("/token", json=params, follow_redirects=False)
    assert response.status_code == 400
    assert response.json()["error"] == "Invalid dPoP URL"


def test_token_with_ath_ok():
    sets = get_settings()

    now = int(time.time())
    payload = {
        "sub": "tester@gmail.com",
        "iss": sets.audience,
        "exp": now + 3600,
        "iat": now,
        "cnf.jkt": base64.b64encode(hashlib.sha256(pbk).digest()).decode(),
        "typ": "dpop",
    }
    access_token = jwt.encode(
        payload,
        get_settings().authz_pvt_key.replace("\\n", "\n").replace("\\t", "\t"),
        algorithm="RS256",
    )

    dpop = create_dpop_jwt(
        pvk,
        pbk,
        sets.authz_url,
        "POST",
        base64.b64encode(
            hashlib.sha256(access_token.encode("utf-8")).digest()
        ).decode(),
    )[1:-1]

    res, err = verify_dpop_jwt(dpop, sets.dpop_htu, sets.dpop_htm, at=access_token)

    dpop_decoded = jwt.decode(dpop, pbk, algorithms=["RS256"])

    assert not err
    assert res
    assert (
        dpop_decoded["ath"]
        == base64.b64encode(
            hashlib.sha256(access_token.encode("utf-8")).digest()
        ).decode()
    )


def test_token_invalid_at_dpop_failure():
    sets = get_settings()

    now = int(time.time())
    payload = {
        "sub": "tester@gmail.com",
        "iss": sets.audience,
        "exp": now + 3600,
        "iat": now,
        "cnf.jkt": "unmatchingcnfjkt",
        "typ": "dpop",
    }
    access_token = jwt.encode(
        payload,
        get_settings().authz_pvt_key.replace("\\n", "\n").replace("\\t", "\t"),
        algorithm="RS256",
    )

    dpop = create_dpop_jwt(
        pvk,
        pbk,
        sets.authz_url,
        "POST",
        base64.b64encode(
            hashlib.sha256(access_token.encode("utf-8")).digest()
        ).decode(),
    )[1:-1]

    res, err = verify_dpop_jwt(dpop, sets.dpop_htu, sets.dpop_htm, at=access_token)

    assert err == "dPoP and JWT mismatch"
    assert res == None


def test_refresh_ok():
    now = time.time()
    payload = {
        "sub": "tester@gmail.com",
        "iss": sets.audience,
        "exp": now + 86400,
        "iat": now,
        "cnf.jkt": base64.b64encode(
            hashlib.sha256(sets.authz_pub_key.encode("utf-8")).digest()
        ).decode(),
        "typ": "dpop+refresh",
    }
    refresh_token = jwt.encode(
        payload,
        get_settings().authz_pvt_key.replace("\\n", "\n").replace("\\t", "\t"),
        algorithm="RS256",
    )

    dpop = create_dpop_jwt(
        pvk,
        pbk,
        sets.authz_url,
        "POST",
        base64.b64encode(
            hashlib.sha256(refresh_token.encode("utf-8")).digest()
        ).decode(),
    )[1:-1]

    refresh_req = {
        "grant_type": "authorization_code",
        "dpop": dpop,
        "refresh_token": refresh_token,
    }

    response = client.post("/refresh", json=refresh_req, follow_redirects=False)
    assert response.status_code == 200
    assert response.json()["token_type"] == "DPoP"
    assert response.json()["access_token"]
