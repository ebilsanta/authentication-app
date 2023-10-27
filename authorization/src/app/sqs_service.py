import asyncio
from functools import lru_cache
import json
import boto3
from fastapi import status
from fastapi.encoders import jsonable_encoder
from fastapi.responses import JSONResponse
from config import Settings
from app.models import TokenRequest, RefreshRequest
from app.process_reqs import process_authcode, process_token, process_refresh


@lru_cache()
def get_settings():
    return Settings()


invalid_response = JSONResponse(
    status_code=status.HTTP_400_BAD_REQUEST,
    content=jsonable_encoder(
        {
            "error": "invalid_request",
            "error_description": "The request is missing a required parameter, includes an invalid parameter value, includes a parameter more than once, or is otherwise malformed.",
        }
    ),
)


class SQS_Service:
    def __init__(self):
        self.sqs = boto3.client("sqs", region_name=get_settings().db_region_name)
        self.queue_url = get_settings().sqs_url

    async def poll_sqs(self):
        print("POLLING!")
        try:
            response = self.sqs.receive_message(
                QueueUrl=self.queue_url,
                AttributeNames=["All"],
                MaxNumberOfMessages=10,  # Maximum number of messages to receive
                MessageAttributeNames=["All"],
            )

            messages = response.get("Messages", [])

            print(len(messages))

            if not messages:
                return {"message": "No messages in the queue"}

            # Process the received messages
            for message in messages:
                try:
                    message_body = message["Body"]

                    payload = json.loads(message_body)
                    asyncio.ensure_future(
                        self.handle_message(
                            payload["operation"], payload["callback"], payload["body"]
                        )
                    )

                    receipt_handle = message["ReceiptHandle"]
                    self.sqs.delete_message(
                        QueueUrl=self.queue_url, ReceiptHandle=receipt_handle
                    )

                except Exception as e:
                    print(e)
        except Exception as e:
            print(e)

    async def handle_message(self, op: str, callback: str, req_body):
        if op == "authcode":
            response = await self.handle_authcode(req_body)
        elif op == "token":
            response = await self.handle_token(req_body)
        elif op == "refresh":
            response = await self.handle_refresh(req_body)

        print(response.__dict__)

    async def handle_authcode(self, rq):
        try:
            return await process_authcode(
                rq["response_type"],
                rq["client_id"],
                rq["state"],
                rq["id_jwt"],
                rq["code_challenge"],
                rq["code_challenge_method"],
                rq["redirect_url"],
            )
        except Exception as e:
            print(e)
            return invalid_response

    async def handle_token(self, rq):
        try:
            return await process_authcode(
                TokenRequest(
                    rq["grant_type"],
                    rq["authcode"],
                    rq["dpop"],
                    rq["client_assertion"],
                    rq["redirect_url"],
                    rq["code_verifier"],
                )
            )
        except Exception as e:
            print(e)
            return invalid_response

    async def handle_refresh(self, rq):
        try:
            return await process_authcode(
                RefreshRequest(rq["grant_type"], rq["dpop"], rq["refresh_token"])
            )
        except Exception as e:
            print(e)
            return invalid_response
