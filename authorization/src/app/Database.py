import os
from motor.motor_asyncio import AsyncIOMotorClient
import time
import boto3
from boto3.resources.base import ServiceResource
from boto3.dynamodb.conditions import Key

class AuthCodeRecord():
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
        ddb = boto3.resource('dynamodb',
                            region_name=os.getenv('DB_REGION_NAME'),
                            aws_access_key_id=os.getenv('DB_ACCESS_KEY_ID'),
                            aws_secret_access_key=os.getenv('DB_SECRET_ACCESS_KEY'))

        self.ac_table = ddb.Table(os.getenv('DB_COLLECTION_AUTHCODES'))
        self.users_table = ddb.Table(os.getenv('DB_COLLECTION_USERS'))

    async def insert_authcode_record(self, acr: AuthCodeRecord):
        return self.ac_table.put_item(Item=acr.__dict__)
    
    async def exists_valid_user(self, email):
        key = {
            'company': 'ascenda', # TODO: Find way to change company
            'email': email
        }
        
        response = self.users_table.get_item(Key=key)

        return 'Item' in response
