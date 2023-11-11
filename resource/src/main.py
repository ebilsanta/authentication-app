from fastapi import FastAPI, Depends, HTTPException, status
from jose import jwt, jwk, JWTError
import requests
import os
from cachetools import TTLCache
from dotenv import load_dotenv

app = FastAPI()

load_dotenv()
JWKS_URL = os.getenv("JWKS_URL")
if not JWKS_URL:
    raise Exception("JWKS_URL environment variable not set")

# Cache configuration
cache = TTLCache(maxsize=100, ttl=3600)  # Adjust maxsize and ttl as needed


def get_jwk(jwks_url: str, kid: str):
    # Check if JWKS is in the cache
    if jwks_url in cache:
        jwks_resp = cache[jwks_url]
    else:
        response = requests.get(jwks_url)
        response.raise_for_status()  # Raise an exception for HTTP errors
        jwks_resp = response.json()
        cache[jwks_url] = jwks_resp  # Store JWKS in the cache

    keys = jwks_resp.get("keys")
    if keys:
        for key in keys:
            if key.get("kid") == kid:
                return key
    return None


def verify_jwt(token: str):
    try:
        unverified_header = jwt.get_unverified_header(token)
        jwk_data = get_jwk(JWKS_URL, unverified_header["kid"])
        if not jwk_data:
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED, detail="JWK not found"
            )

        # Use the alg from the JWKS server
        algorithm = jwk_data.get("alg", "ES256")  # Default to ES256 if not specified
        key = jwk.construct(jwk_data, algorithm)
        decoded_token = jwt.decode(token, key=key, algorithms=[algorithm])
        return decoded_token
    except JWTError:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED, detail="Invalid token"
        )


@app.get("/user")
async def read_user(decoded_token: dict = Depends(verify_jwt)):
    # Extract user details from the decoded token. Adjust the keys according to your token's payload structure.
    user_details = {
        "user_id": decoded_token.get("sub"),  # 'sub' is typically used for the user ID
        "email": decoded_token.get("email"),
    }
    return user_details
