import json
import time
import uuid
from unittest.mock import patch, ANY
from unittest import mock

import jwt
import pytest
from app.pkce import generate_pkce_code_challenge, generate_pkce_code_verifier
from app.sqs_service import SQS_Service
from app.dpop_service import create_dpop_jwt
from app.client_assertion_service import do_generate_client_assertion
from config import get_settings
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


@pytest.mark.asyncio
async def test_poll_sqs_authcode_302_ok():
    sqss = SQS_Service()
    with patch.object(
        sqss, "receive_sqs_msg", new=mock.AsyncMock()
    ) as m1, patch.object(
        sqss, "handle_message", new=mock.AsyncMock()
    ) as m2, patch.object(
        sqss.sqs, "delete_message"
    ) as m3:
        body = {
            "response_type": "code",
            "client_id": allowed_client,
            "redirect_url": redirect_url,
            "state": uuid.uuid4().hex,
            "code_challenge": code_challenge,
            "id_jwt": testing_jwt,
            "code_challenge_method": "S256",
        }

        to_return = [
            {
                "Body": json.dumps(
                    {
                        "operation": "authcode",
                        "callback": "http://api-gateway/callback",
                        "body": body,
                    }
                ),
                "ReceiptHandle": "handle",
            }
        ]

        print(json.dumps(to_return))

        m1.return_value = to_return

        await sqss.poll_sqs()
        m1.assert_called_once()
        m2.assert_called_once_with("authcode", "http://api-gateway/callback", body)
        m3.assert_called_once_with(QueueUrl=ANY, ReceiptHandle="handle")

@pytest.mark.asyncio
async def test_poll_sqs_authcode_200_ok():
    sqss = SQS_Service()
    with patch.object(
        sqss, "receive_sqs_msg", new=mock.AsyncMock()
    ) as m1, patch.object(
        sqss, "handle_message", new=mock.AsyncMock()
    ) as m2, patch.object(
        sqss.sqs, "delete_message"
    ) as m3:
        body = {
            "response_type": "code",
            "client_id": allowed_client,
            "state": uuid.uuid4().hex,
            "code_challenge": code_challenge,
            "id_jwt": testing_jwt,
            "code_challenge_method": "S256",
        }

        to_return = [
            {
                "Body": json.dumps(
                    {
                        "operation": "authcode",
                        "callback": "http://api-gateway/callback",
                        "body": body,
                    }
                ),
                "ReceiptHandle": "handle",
            }
        ]

        print(json.dumps(to_return))

        m1.return_value = to_return

        await sqss.poll_sqs()
        m1.assert_called_once()
        m2.assert_called_once_with("authcode", "http://api-gateway/callback", body)
        m3.assert_called_once_with(QueueUrl=ANY, ReceiptHandle="handle")

@pytest.mark.asyncio
async def test_poll_sqs_authcode_400_failure():
    sqss = SQS_Service()
    with patch.object(
        sqss, "receive_sqs_msg", new=mock.AsyncMock()
    ) as m1, patch.object(
        sqss, "handle_message", new=mock.AsyncMock()
    ) as m2, patch.object(
        sqss.sqs, "delete_message"
    ) as m3:
        body = {
            "client_id": allowed_client,
            "state": uuid.uuid4().hex,
            "code_challenge": code_challenge,
            "id_jwt": testing_jwt,
            "code_challenge_method": "S256",
        }

        to_return = [
            {
                "Body": json.dumps(
                    {
                        "operation": "authcode",
                        "callback": "http://api-gateway/callback",
                        "body": body,
                    }
                ),
                "ReceiptHandle": "handle",
            }
        ]

        print(json.dumps(to_return))

        m1.return_value = to_return

        await sqss.poll_sqs()
        m1.assert_called_once()
        m2.assert_called_once_with("authcode", "http://api-gateway/callback", body)
        m3.assert_called_once_with(QueueUrl=ANY, ReceiptHandle="handle")



@pytest.mark.asyncio
@patch("app.sqs_service.process_authcode")
async def test_handle_authcode_message_ok(m1):
    sqss = SQS_Service()
    with patch.object(sqss, "use_callback", new=mock.AsyncMock()) as m2:
        callback = "http://callback.url"

        m1.return_value = callback

        body = {
            "response_type": "code",
            "client_id": allowed_client,
            "redirect_url": redirect_url,
            "state": uuid.uuid4().hex,
            "code_challenge": code_challenge,
            "id_jwt": testing_jwt,
            "code_challenge_method": "S256",
        }

        await sqss.handle_message("authcode", callback, body)

        m1.assert_called_once()
        m2.assert_called_once()


@pytest.mark.asyncio
@patch("app.sqs_service.process_token")
async def test_handle_token_message_ok(m1):
    sqss = SQS_Service()
    with patch.object(sqss, "use_callback", new=mock.AsyncMock()) as m2:
        callback = "http://callback.url"

        m1.return_value = callback

        now = int(time.time())
        sets = get_settings()
        pvk = bytes(
            sets.allowed_client_pvt_key.replace("\\n", "\n").replace("\\t", "\t"),
            "utf-8",
        )
        pbk = bytes(
            sets.allowed_client_pub_key.replace("\\n", "\n").replace("\\t", "\t"),
            "utf-8",
        )
        ac = "0123456789abcdef0123456789abcdef"

        body = {
            "grant_type": "authorization_code",
            "authcode": ac,
            "dpop": create_dpop_jwt(pvk, pbk, sets.authz_url, "POST")[1:-1],
            "client_assertion": do_generate_client_assertion(
                sets.allowed_client, sets.audience, now + 300, now, pvk
            ),
            "redirect_url": sets.allowed_redirect,
            "code_verifier": code_verifier,
        }

        await sqss.handle_message("token", callback, body)

        m1.assert_called_once()
        m2.assert_called_once()


@pytest.mark.asyncio
@patch("app.sqs_service.process_refresh")
async def test_handle_refresh_message_ok(m1):
    sqss = SQS_Service()
    with patch.object(sqss, "use_callback", new=mock.AsyncMock()) as m2:
        callback = "http://callback.url"

        m1.return_value = callback

        body = {
            "grant_type": "authorization_code",
            "dpop": "dpop",
            "refresh_token": "refresh_token",
        }

        await sqss.handle_message("refresh", callback, body)

        m1.assert_called_once()
        m2.assert_called_once()
