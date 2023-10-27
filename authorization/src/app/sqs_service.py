from functools import lru_cache
import json
import boto3
from config import Settings
from app.process_reqs import process_authcode, process_token, process_refresh


@lru_cache()
def get_settings():
    return Settings()


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
                    # You can process the message here

                    payload = json.loads(message_body)
                    self.process_message(
                        payload["operation"], payload["callback"], payload["body"]
                    )

                    # Delete the message from the queue once processed
                    receipt_handle = message["ReceiptHandle"]
                    self.sqs.delete_message(
                        QueueUrl=self.queue_url, ReceiptHandle=receipt_handle
                    )

                except Exception as e:
                    print(e)
        except Exception as e:
            print(e)

    async def process_message(self, op: str, callback: str, req_body):
        print(op, callback, req_body)
        if op == "authcode":
            return await self.process_authcode(req_body)
        elif op == "token":
            return await self.process_token(req_body)
        elif op == "refresh":
            return await self.process_refresh(req_body)

    async def process_authcode(self, rq):
        try:
            return await post_authcode(
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

    async def process_token(self):
        pass

    async def process_refresh(self):
        pass
