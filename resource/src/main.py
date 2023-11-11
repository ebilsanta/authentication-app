from fastapi import FastAPI
from database.users import get_user

app = FastAPI()

# TODO: Verify the token from the request
@app.get("/users/{company}/{email}")
def get_user_handler(company: str, email: str):
    user = get_user(company, email)

    return user