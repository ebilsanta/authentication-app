import asyncio
from functools import lru_cache
import json
import boto3
from fastapi import status
from fastapi.encoders import jsonable_encoder
from fastapi.responses import JSONResponse
import httpx
from config import Settings
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
        messages = await self.receive_sqs_msg()

        if not messages:
            return {"message": "No messages in the queue"}

        # Process the polled messages
        print('Polled ', str(len(messages)), ' messages')
        for message in messages:
            try:
                message_body = message["Body"]

                payload = json.loads(message_body)
                asyncio.ensure_future(
                    self.handle_message(
                        payload["operation"], payload["callback"], payload["body"]
                    )
                )

                self.sqs.delete_message(
                    QueueUrl=self.queue_url, ReceiptHandle=message["ReceiptHandle"]
                )

            except Exception as e:
                print(e)

    async def receive_sqs_msg(self):
        try:
            return self.sqs.receive_message(
                QueueUrl=self.queue_url,
                AttributeNames=["All"],
                MaxNumberOfMessages=10,
                MessageAttributeNames=["All"],
            ).get("Messages", [])
        except Exception as e:
            print(e)
            return None

    async def handle_message(self, op: str, callback: str, req_body):
        if op == "authcode":
            print("SQS Authcode Request " + callback)
            response = await self.handle_authcode(req_body)
        elif op == "token":
            print("SQS Token Request " + callback)
            response = await self.handle_token(req_body)
        elif op == "refresh":
            print("SQS Refresh Request " + callback)
            response = await self.handle_refresh(req_body)

        await self.use_callback(callback, response)

    async def handle_authcode(self, rq):
        try:
            return await process_authcode(
                rq["response_type"],
                rq["client_id"],
                rq["state"],
                rq["id_jwt"],
                rq["code_challenge"] if "code_challenge" in rq else None,
                rq["code_challenge_method"] if "code_challenge_method" in rq else None,
                rq["redirect_url"] if "redirect_url" in rq else None,
            )
        except Exception as e:
            print(e)
            return invalid_response

    async def handle_token(self, rq):
        try:
            return await process_token(
                rq["grant_type"],
                rq["authcode"],
                rq["dpop"],
                rq["client_assertion"],
                rq["redirect_url"] if "redirect_url" in rq else None,
                rq["code_verifier"],
            )
        except Exception as e:
            print(e)
            return invalid_response

    async def handle_refresh(self, rq):
        try:
            return await process_refresh(
                rq["grant_type"], rq["dpop"], rq["refresh_token"]
            )
        except Exception as e:
            print(e)
            return invalid_response

    async def use_callback(self, callback: str, body):
        async with httpx.AsyncClient() as client:
            response = {
                "status_code": body.__dict__["status_code"],
                "headers": dict(
                    (x.decode("utf-8"), y.decode("utf-8"))
                    for x, y in body.__dict__["raw_headers"]
                ),
            }
            if "body" in body.__dict__ and body.__dict__["body"] != b"":
                response.update({"body": json.loads(body.__dict__["body"])})

            # print(response)
            # print(json.dumps(response))

            try:
                await client.post(callback, json={"response": json.dumps(response)})
            except Exception as e:
                print(e)
