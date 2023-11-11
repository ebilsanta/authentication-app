from fastapi import FastAPI, Depends, HTTPException, status, Request
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


# Thanks matt!
def verify_dpop_jwt(dpop_jwt, htu, htm, at=None):
    try:
        uvh = jwt.get_unverified_header(dpop_jwt)
        if uvh["alg"] != "RS256":
            return None, "Invalid algorithm"
        jwk = uvh["jwk"]
        decoded = jwt.decode(dpop_jwt, base64.b64decode(jwk), algorithms=["RS256"])

        if uvh["typ"] != "dpop+jwt":
            return None, "Invalid token type"
        if decoded["htm"] != htm:
            return None, "Invalid dPoP HTTP Method"

        if decoded["htu"] != htu:
            return None, "Invalid dPoP URL"

        if (
            "ath" in decoded
            and decoded["ath"]
            != base64.b64encode(
                hashlib.sha256(at.encode("ascii")).digest()
            ).decode()  # Same, verify output is same
        ):
            return None, "Invalid access token hash"

        if at:
            decoded_at = jwt.decode(
                at,
                (
                    get_settings().authz_pub_key
                    if get_settings().authz_pub_key
                    else update_authZ_key()
                )
                .replace("\\n", "\n")
                .replace("\\t", "\t"),
                algorithms=["RS256"],
            )
            if (
                decoded_at["cnf.jkt"]
                != base64.b64encode(
                    hashlib.sha256(base64.b64decode(jwk)).digest()
                ).decode()
            ):
                return None, "dPoP and JWT mismatch"

        return jwk, None

    except jwt.ExpiredSignatureError:
        return None, "JWT has expired"
    except jwt.InvalidTokenError:
        return None, "Invalid JWT"
    except Exception as e:
        return None, "An error occurred during dpop JWT decoding:" + str(e) + str(
            dpop_jwt
        )


@app.get("/user")
async def read_user(request: Request, decoded_token: dict = Depends(verify_jwt)):
    await verify_dpop_jwt(request, request.url.path, request.method)
    user_details = {
        "user_id": decoded_token.get("sub"),  # 'sub' is typically used for the user ID
        "email": decoded_token.get("email"),
    }
    return user_details


@app.get("/health")
def get_health():
    return {"status": "ok"}
