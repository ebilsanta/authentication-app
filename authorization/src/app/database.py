import time
from functools import lru_cache

import boto3
from boto3.session import Session
from botocore.credentials import RefreshableCredentials
from botocore.session import get_session
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

class TokenRecord():
    token: str
    removeAt: str
    active: bool

    def __init__(self, token, removeAt, active):
        self.token = token
        self.removeAt = removeAt
        self.active = active

class Database:
    def __init__(self):
        sets = get_settings()
        self.session_name = "RoleSession1"

        if sets.db_access_key_id != "" and sets.db_secret_access_key != "":
            print("Using Credentials")
            ddb = boto3.resource(
                "dynamodb",
                region_name=sets.db_region_name,
                aws_access_key_id=sets.db_access_key_id,
                aws_secret_access_key=sets.db_secret_access_key,
            )
        elif sets.role_arn != "":
            print("Using Role")
            self.sts_client = boto3.client("sts", region_name=sets.db_region_name)

            session_credentials = RefreshableCredentials.create_from_metadata(
                metadata=self._refresh(),
                refresh_using=self._refresh,
                method="sts-assume-role",
            )

            session = get_session()
            session._credentials = session_credentials
            session.set_config_variable("region", sets.db_region_name)
            autorefresh_session = Session(botocore_session=session)

            ddb = autorefresh_session.resource("dynamodb")
        else:
            print("No DynamoDB login method")

        self.ac_table = ddb.Table(sets.db_collection_authcodes)
        self.users_table = ddb.Table(sets.db_collection_users)
        self.token_table = ddb.Table(sets.db_collection_tokens)

    def _refresh(self):
        params = {
            "RoleArn": get_settings().role_arn,
            "RoleSessionName": self.session_name,
            "DurationSeconds": 3600,
        }

        response = self.sts_client.assume_role(**params).get("Credentials")
        credentials = {
            "access_key": response.get("AccessKeyId"),
            "secret_key": response.get("SecretAccessKey"),
            "token": response.get("SessionToken"),
            "expiry_time": response.get("Expiration").isoformat(),
        }
        return credentials

    async def insert_authcode_record(self, acr: AuthCodeRecord):
        return self.ac_table.put_item(Item=acr.__dict__)
    
    async def get_token_record(self, token: str):
        response = self.token_table.get_item(Key={"token": token})
        return response.get("Item")
    
    async def insert_token_record(self, tr: TokenRecord):
        return self.token_table.put_item(Item=tr.__dict__)

    async def get_authcode_record(self, ac: str):
        response = self.ac_table.get_item(Key={"authcode": ac})
        return response.get("Item")

    async def exists_valid_user(self, email):
        key = {"company": "ascenda", "email": email}  # TODO: Find way to change company

        response = self.users_table.get_item(Key=key)

        return "Item" in response
