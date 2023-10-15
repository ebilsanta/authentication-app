import os
from motor.motor_asyncio import AsyncIOMotorClient
from pydantic import BaseModel

class AuthCodeRecord():
    authcode: str
    state: str
    code_challenge: str

    def __init__(self, authcode, state, code_challenge):
        self.authcode = authcode
        self.state = state
        self.code_challenge = code_challenge

class Database:
    def __init__(self):
        self.client = AsyncIOMotorClient(os.getenv('DB_URL'))
        self.db = self.client[os.getenv('DB_NAME')]
        self.collection = self.db[os.getenv('DB_COLLECTION')]

    async def insert_authcode_record(self, acr: AuthCodeRecord):
        return await self.collection.insert_one(acr.__dict__)



    