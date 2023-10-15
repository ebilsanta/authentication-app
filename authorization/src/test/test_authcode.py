from dotenv import load_dotenv
from fastapi.testclient import TestClient
load_dotenv(".env")

from main import app

client = TestClient(app)

def test_hello_world():
    response = client.get("/")
    assert response.status_code == 200
    assert response.json() == {"Hello": "World"}

# def test_response_type():
#     response = client.post("/authcode?code=none")
#     assert response.status_code == 400
#     print(response.json)
    