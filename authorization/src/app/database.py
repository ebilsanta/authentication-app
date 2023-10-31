import time
from functools import lru_cache

import boto3
from boto3.session import Session
from config import Settings


@lru_cache()
def get_settings():
    return Settings()


class AuthCodeRecord:
    authcode: str
    state: str
    code_challenge: str
    user: str
    created: int
    expiry: int

    def __init__(self, authcode, state, code_challenge, user, expiry_delta):
        self.authcode = authcode
        self.state = state
        self.code_challenge = code_challenge
        self.user = user
        self.created = round(time.time())
        self.expiry = self.created + expiry_delta


class Database:
    def __init__(self):
        sets = get_settings()

        if sets.db_access_key_id != "" and sets.db_secret_access_key != "":
            print("Using Credentials")
            ddb = boto3.resource(
                "dynamodb",
                region_name=sets.db_region_name,
                aws_access_key_id=sets.db_access_key_id,
                aws_secret_access_key=sets.db_secret_access_key,
            )
        elif sets.role_arn != "":
            print('Using Role')
            sts_client = boto3.client("sts")
            assumed_role_object = sts_client.assume_role(
                RoleArn=sets.role_arn, RoleSessionName="RoleSession1"
            )

            credentials = assumed_role_object["Credentials"]
            session = Session(
                aws_access_key_id=credentials["AccessKeyId"],
                aws_secret_access_key=credentials["SecretAccessKey"],
                aws_session_token=credentials["SessionToken"],
            )
            ddb = session.resource("dynamodb")
        else:
            print("No DynamoDB login method")

        self.ac_table = ddb.Table(sets.db_collection_authcodes)
        self.users_table = ddb.Table(sets.db_collection_users)

    async def insert_authcode_record(self, acr: AuthCodeRecord):
        return self.ac_table.put_item(Item=acr.__dict__)

    async def get_authcode_record(self, ac: str):
        response = self.ac_table.get_item(Key={"authcode": ac})
        return response.get("Item")

    async def exists_valid_user(self, email):
        key = {"company": "ascenda", "email": email}  # TODO: Find way to change company

        response = self.users_table.get_item(Key=key)

        return "Item" in response
