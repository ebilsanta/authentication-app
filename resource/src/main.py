from typing import Union
from fastapi import FastAPI
from database.users import get_user
app = FastAPI()

# TODO: Find out whats the endpoint gonna be like 
@app.get("/users/{company}/{email}")
def get_user_handler(company: str, email: str):
    user = get_user(company, email)

    return user