import time
import uuid
from unittest.mock import patch

import jwt
from app.pkce import generate_pkce_code_challenge, generate_pkce_code_verifier
from app.setup_utils import build_url, is_uuid
from config import get_settings
from cryptography.hazmat.primitives.asymmetric import rsa
from fastapi.testclient import TestClient
from main import app

# Setup
sets = get_settings()

allowed_client = sets.allowed_client
issuer = sets.allowed_issuer
subject = "testing@test.com"
audience = sets.audience
redirect_url = sets.allowed_redirect

iat = time.time()
exp = iat + 3600
payload = {"iss": issuer, "sub": subject, "aud": audience, "iat": iat, "exp": exp}
additional_headers = {"kid": uuid.uuid4().hex}
private_key = sets.pvt_key.replace("\\n", "\n").replace("\\t", "\t")
testing_jwt = jwt.encode(
    payload, private_key, headers=additional_headers, algorithm="RS256"
)

code_verifier = generate_pkce_code_verifier()
code_challenge = generate_pkce_code_challenge(code_verifier)


client = TestClient(app)

# Tests start here


def test_hello_world():
    response = client.get("/")
    assert response.status_code == 200
    assert response.json() == {"Hello": "World"}


@patch("main.db.insert_authcode_record")
@patch("main.db.exists_valid_user")
def test_authcode_ok(method2, method1):
    method1.return_value = 0
    method2.return_value = True
    params = {
        "response_type": "code",
        "client_id": allowed_client,
        "redirect_url": redirect_url,
        "state": uuid.uuid4().hex,
        "code_challenge": code_challenge,
        "id_jwt": testing_jwt,
        "code_challenge_method": "S256",
    }

    response = client.post(build_url("/authcode", params), follow_redirects=False)

    assert response.status_code == 302
    assert response.headers["location"].split("=")[0] == redirect_url + "?code"
    assert is_uuid(response.headers["location"].split("=")[1])


@patch("main.db.insert_authcode_record")
@patch("main.db.exists_valid_user")
def test_authcode_no_redirect_ok(method2, method1):
    method1.return_value = 0
    method2.return_value = True
    params = {
        "response_type": "code",
        "client_id": allowed_client,
        "state": uuid.uuid4().hex,
        "code_challenge": code_challenge,
        "id_jwt": testing_jwt,
        "code_challenge_method": "S256",
    }

    response = client.post(build_url("/authcode", params), follow_redirects=False)
    assert response.status_code == 200
    assert is_uuid(response.json()["code"])


def test_response_type_fail():
    params = {
        "response_type": "nocode",
        "client_id": allowed_client,
        "redirect_url": redirect_url,
        "state": uuid.uuid4().hex,
        "code_challenge": code_challenge,
        "id_jwt": testing_jwt,
        "code_challenge_method": "S256",
    }

    response = client.post(build_url("/authcode", params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers["location"].split("?")[0] == redirect_url + "/"
    assert (
        response.headers["location"].split("?")[1] == "error=unsupported_response_type"
    )


def test_client_fail():
    params = {
        "response_type": "code",
        "client_id": "badclient",
        "redirect_url": redirect_url,
        "state": uuid.uuid4().hex,
        "code_challenge": code_challenge,
        "id_jwt": testing_jwt,
        "code_challenge_method": "S256",
    }

    response = client.post(build_url("/authcode", params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers["location"].split("?")[0] == redirect_url + "/"
    assert response.headers["location"].split("?")[1] == "error=unauthorized_client"


def test_code_challenge_empty_fail():
    params = {
        "response_type": "code",
        "client_id": allowed_client,
        "redirect_url": redirect_url,
        "state": uuid.uuid4().hex,
        "id_jwt": testing_jwt,
        "code_challenge_method": "S256",
    }

    response = client.post(build_url("/authcode", params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers["location"].split("?")[0] == redirect_url + "/"

    queries = response.headers["location"].split("?")[1].split("&")

    assert "error=invalid_request" in queries
    assert "error_description=code+challenge+required" in queries


def test_code_challenge_invalid_fail():
    params = {
        "response_type": "code",
        "client_id": allowed_client,
        "redirect_url": redirect_url,
        "state": uuid.uuid4().hex,
        "code_challenge": "aaa",
        "id_jwt": testing_jwt,
        "code_challenge_method": "S256",
    }

    response = client.post(build_url("/authcode", params), follow_redirects=False)

    assert response.status_code == 302
    assert response.headers["location"].split("?")[0] == redirect_url + "/"

    queries = response.headers["location"].split("?")[1].split("&")

    assert "error=invalid_request" in queries
    assert "error_description=invalid+code+challenge" in queries


def test_code_challenge_method_empty_fail():
    params = {
        "response_type": "code",
        "client_id": allowed_client,
        "redirect_url": redirect_url,
        "state": uuid.uuid4().hex,
        "code_challenge": code_challenge,
        "id_jwt": testing_jwt,
    }

    response = client.post(build_url("/authcode", params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers["location"].split("?")[0] == redirect_url + "/"

    queries = response.headers["location"].split("?")[1].split("&")

    assert "error=invalid_request" in queries
    assert "error_description=transform+algorithm+not+supported" in queries


def test_code_challenge_method_invalid_fail():
    params = {
        "response_type": "code",
        "client_id": allowed_client,
        "redirect_url": redirect_url,
        "state": uuid.uuid4().hex,
        "code_challenge": code_challenge,
        "id_jwt": testing_jwt,
        "code_challenge_method": "S128",
    }

    response = client.post(build_url("/authcode", params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers["location"].split("?")[0] == redirect_url + "/"

    queries = response.headers["location"].split("?")[1].split("&")

    assert "error=invalid_request" in queries
    assert "error_description=transform+algorithm+not+supported" in queries


def test_invalid_jwt_fail():
    params = {
        "response_type": "code",
        "client_id": allowed_client,
        "redirect_url": redirect_url,
        "state": uuid.uuid4().hex,
        "code_challenge": code_challenge,
        "id_jwt": "abc",
        "code_challenge_method": "S256",
    }

    response = client.post(build_url("/authcode", params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers["location"].split("?")[0] == redirect_url + "/"

    queries = response.headers["location"].split("?")[1].split("&")

    assert "error=access_denied" in queries
    assert "error_description=Invalid+JWT" in queries


def test_expired_jwt_fail():
    iat = 0
    exp = 1
    payload = {"iss": issuer, "sub": subject, "aud": audience, "iat": iat, "exp": exp}
    additional_headers = {"kid": uuid.uuid4().hex}
    expired_jwt = jwt.encode(
        payload, private_key, headers=additional_headers, algorithm="RS256"
    )

    params = {
        "response_type": "code",
        "client_id": allowed_client,
        "redirect_url": redirect_url,
        "state": uuid.uuid4().hex,
        "code_challenge": code_challenge,
        "id_jwt": expired_jwt,
        "code_challenge_method": "S256",
    }

    response = client.post(build_url("/authcode", params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers["location"].split("?")[0] == redirect_url + "/"

    queries = response.headers["location"].split("?")[1].split("&")

    assert "error=access_denied" in queries
    assert "error_description=JWT+has+expired" in queries


def test_jwt_future_fail():
    iat = time.time() + 100000
    exp = iat + 3600
    payload = {"iss": issuer, "sub": subject, "aud": audience, "iat": iat, "exp": exp}
    additional_headers = {"kid": uuid.uuid4().hex}
    future_jwt = jwt.encode(
        payload, private_key, headers=additional_headers, algorithm="RS256"
    )

    params = {
        "response_type": "code",
        "client_id": allowed_client,
        "redirect_url": redirect_url,
        "state": uuid.uuid4().hex,
        "code_challenge": code_challenge,
        "id_jwt": future_jwt,
        "code_challenge_method": "S256",
    }

    response = client.post(build_url("/authcode", params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers["location"].split("?")[0] == redirect_url + "/"

    queries = response.headers["location"].split("?")[1].split("&")

    assert "error=access_denied" in queries
    assert "error_description=Invalid+JWT" in queries


def test_jwt_invalid_issuer_fail():
    iat = time.time()
    exp = iat + 3600
    payload = {
        "iss": "newperson",
        "sub": subject,
        "aud": audience,
        "iat": iat,
        "exp": exp,
    }
    additional_headers = {"kid": uuid.uuid4().hex}
    invalid_jwt = jwt.encode(
        payload, private_key, headers=additional_headers, algorithm="RS256"
    )

    params = {
        "response_type": "code",
        "client_id": allowed_client,
        "redirect_url": redirect_url,
        "state": uuid.uuid4().hex,
        "code_challenge": code_challenge,
        "id_jwt": invalid_jwt,
        "code_challenge_method": "S256",
    }

    response = client.post(build_url("/authcode", params), follow_redirects=False)
    assert response.status_code == 302
    assert response.headers["location"].split("?")[0] == redirect_url + "/"

    queries = response.headers["location"].split("?")[1].split("&")

    assert "error=access_denied" in queries
    assert "error_description=Unknown+Issuer" in queries


def test_state_invalid_fail():
    params = {
        "response_type": "code",
        "client_id": allowed_client,
        "redirect_url": redirect_url,
        "state": "NewYork",
        "code_challenge": code_challenge,
        "id_jwt": testing_jwt,
        "code_challenge_method": "S256",
    }

    response = client.post(build_url("/authcode", params), follow_redirects=False)

    assert response.status_code == 302
    assert response.headers["location"].split("?")[0] == redirect_url + "/"

    queries = response.headers["location"].split("?")[1].split("&")

    assert "error=invalid_request" in queries
    assert "error_description=invalid+state" in queries


def test_missing_params_fail():
    params = {"response_type": "code", "code_challenge_method": "S256"}

    response = client.post(build_url("/authcode", params), follow_redirects=False)

    assert response.status_code == 400
    assert response.json()["error"] == "invalid_request"
    assert (
        response.json()["error_description"]
        == "The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed."
    )
