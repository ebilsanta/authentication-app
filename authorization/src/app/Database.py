import os
from motor.motor_asyncio import AsyncIOMotorClient
import time

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
        user = user
        self.created = time.time()
        self.expiry = self.created + expiry_delta

class Database:
    def __init__(self):
        self.client = AsyncIOMotorClient(os.getenv('DB_URL'))
        self.db = self.client[os.getenv('DB_NAME')]
        self.collection = self.db[os.getenv('DB_COLLECTION_AUTHCODES')]

    async def insert_authcode_record(self, acr: AuthCodeRecord):
        return await self.collection.insert_one(acr.__dict__)



    