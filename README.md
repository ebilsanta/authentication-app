# CS301 G2T1 Authentication App

## Description

This project is an Authentication (AuthN) service that allows a company's partners to leverage its managed user services. We have implemented a system that provides identity and Authorization (AuthZ) for secure data exchange between partners and the company's existing services, such as a rewards platform.

## Live Website

We have deployed a frontend to demonstrate the end-to-end flow of our features.  
It was deployed on https://www.itsag2t1.com/ 

## Solution Architecture

![CS301-G2T1-Architecture](https://github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/assets/101983505/37b04a71-dc13-4161-b1f9-3895d87a1fd6)

## Microservices  

We have designed our solution with the following microservices for modularity.

- `authentication` - A `gRPC`/`REST` API that manages user credentials and handles registration, sign in, password change requests. Built with `Golang`.
- `authorization` - A `REST` API that issues and manages the tokens used to access protected resources. Built with `Python`/`FastAPI`.
- `otp` - A `gRPC`/`REST` API that generates and validates One-Time Passwords (OTP) for sensitive actions. Built with `Golang`.
- `resource` - A `REST` API that returns user data to authorized users. Built with `Python`/`FastAPI`.
- `backend` - A `REST` API that models how partners can integrate with our solution to build a customized authorization flow. Built with `Express.JS`.
- `frontend` - A web application to demonstrate the end-to-end authorization flow with our solution. Built with `Next.js`.

## API Endpoints
**API Base URL**: `https://api.itsag2t1.com/V0/`  

**Resource Server Base URL**: `https://resource.itsag2t1.com/`

### User Registration

Registers a new user with the provided information. Sends an OTP for verification to their email. 

- **Method**: `POST`
- **Endpoint**: `/register`

#### Request

##### Body (JSON)

- **company** (string): Company name.
- **email** (string): User's email address.
- **firstName** (string): User's first name.
- **lastName** (string): User's last name.
- **birthdate** (string): User's birthdate in YYYY-mm-dd.
- **password** (string): User's password.

#### Example

```json
{
  "company": "Example Corp",
  "email": "user@example.com",
  "firstName": "John",
  "lastName": "Doe",
  "birthdate": "1990-01-01",
  "password": "securePassword123"
}
```

## Email Verification

Verifies the user's email address using the provided OTP (One-Time Password) and verification key.

- **Method**: `POST`
- **Endpoint**: `/verify-email`

### Request

#### Body (JSON)

- **otp** (string): One-Time Password received by the user.
- **email** (string): User's email address.
- **verificationKey** (string): Key for verifying the email. Obtained from the previous step.
- **callback** (string): Callback URL for handling the verification result.

#### Example

```json
{
  "otp": "123456",
  "email": "user@example.com",
  "verificationKey": "abcdef",
  "callback": "https://example.com/verify-email/session123"
}
```

## User Login

Logs in a user with the provided credentials.

- **Method**: `POST`
- **Endpoint**: `/login`

### Request

#### Body (JSON)

- **company** (string): Company name.
- **email** (string): User's email address.
- **password** (string): User's password.
- **callback** (string): Callback URL for handling the login result.

#### Example

```json
{
  "company": "Example Corp",
  "email": "user@example.com",
  "password": "securePassword123",
  "callback": "https://example.com/login/session123"
}
```

## Authorization Code

Initiates the request for an authorization code from the authentication server.

- **Method**: `POST`
- **Endpoint**: `/hosted/authcode`

### Request

#### Parameters

- **response_type** (string): The type of response (e.g., "code").
- **state** (string): A nonce, a random hex string between 16 to 64 characters. Helps to prevent MiTM.
- **id_jwt** (string): JWT with the following fields, signed with AuthN's private key
  - Header:
  - alg=RS256
  - kid=<A UUID4 in Hexadecimal, doesn't really matter>
  - Payload:
    - iss=<authN server's name>
    - sub=<user's email>
    - aud=<authZ server's name>
    - iat=<current time in seconds>
    - exp=<expiry time in seconds>
- **client_id** (string): Client ID registered with our service.
- **code_challenge** (string): Code challenge with the following requirements:
  - 44 characters long
  - Base64 of the SHA256 hash of the code verifier in ASCII
  - Refer to [Sample Code](#sample-code).
- **code_challenge_method** (string): Method used for code challenge generation (must be "S256").
- **redirect_url** (string): Redirect URL for the authentication response.
  - Any absolute URL registered with our service.
  - Example: https://localhost:8000
- **callback_url** (string): Callback URL for handling the authorization code result.

#### Example
Note: The following go into the query params according to RFC
Eg. `?response_type=code&state=...`
```json
{
  "response_type": "code",
  "state": "generatedState123",
  "id_jwt": "exampleIdentityJwt",
  "client_id": "exampleClientID",
  "code_challenge": "generatedCodeChallenge123",
  "code_challenge_method": "S256",
  "redirect_url": "http://localhost:8000",
  "callback_url": "https://example.com/authcode/session123"
}
```

## Access Token

Requests an access token using the authorization code obtained from the authentication server.

- **Method**: `POST`
- **Endpoint**: `/hosted/token`

### Request

#### Body (JSON)

- **grant_type** (string): The type of grant being requested (must be "authorization_code").
- **authcode** (string): Authorization code obtained from the /authcode endpoint.
- **dpop** (string): DPoP (Distributed Proof-of-Possession) proof.
  - A JWT with the following fields, signed with an ephemeral private key. Refer to [Sample Code](#sample-code).
  - Header
    - alg=RS256
    - typ=dpop+jwt
    - jwk=<ephemeral public key>
  - Payload
    - iat=<created time in seconds>
    - jti=<random UUID4 in hex>
    - htm="POST"
    - htu=<Token endpoint URL, without anything towards the end like `?q1=abc` etc. i.e `http://localhost:8080/token`
    - exp=<iat+120>

- **client_assertion** (string): JWT with the following fields, signed with client's private key. Refer to [Sample Code](#sample-code).
  - Header
    - alg=RS256
  - Payload
    - iss=<your client_id>
    - sub=<your client_id>
    - aud=<the authZ server name>
    - iat=<created time in seconds>
    - exp=<iat+300>
 

- **redirect_url** (string):  Redirect URL for the response.
  - Any absolute URL registered with our service.
  - Example: https://localhost:8000
- **code_verifier** (string): Full 43-128 character Code Verifier used in the initial auth code request.
- **callback_url** (string): Callback URL for handling the access token result.

#### Example

```json
{
  "grant_type": "authorization_code",
  "authcode": "exampleAuthorizationCode",
  "dpop": "exampleDPoPProof",
  "client_assertion": "exampleClientAssertion",
  "redirect_url": "http://localhost:8000",
  "code_verifier": "exampleCodeVerifier",
  "callback_url": "https://example.com/token/session123"
}
```

## Refresh Token

Refreshes user's access token.

- **Method**: `POST`
- **Endpoint**: `/hosted/refresh`

### Request

#### Body (JSON)

- **grant_type** (string): The type of grant being requested (must be "authorization_code").
- **refresh_token** (string): Refresh token obtained from the authentication server.
- **dpop** (string): DPoP (Distributed Proof-of-Possession) proof similar to access token request, but with additional `ath` field in payload. Signed with the same ephemeral private key in access token request. Refer to [Sample Code](#sample-code).
  - Header
    - alg=RS256
    - typ=dpop+jwt
    - jwk=<ephemeral public key>
  - Payload
    - iat=<created time in seconds>
    - jti=<random UUID4 in hex>
    - ath=<base64 encoded SHA256 hash of the access token in ASCII>
    - htm="POST"
    - htu=<Token endpoint URL, without anything towards the end like `?q1=abc` etc. i.e `http://localhost:8080/token`
    - exp=<iat+120>
- **callback_url** (string): Callback URL for handling the refresh token result.

#### Example

```json
{
  "grant_type": "authorization_code",
  "refresh_token": "exampleRefreshToken",
  "dpop": "exampleDPoPProof",
  "callback_url": "https://example.com/refresh/session123"
}
```

## Introspect Token

Check if an access token is active.

- **Method**: `POST`
- **Endpoint**: `/hosted/introspect`

### Request

#### Body (JSON)

- **token** (string): The JWT token

#### Example

```json
{
  "token": "accessToken",
}
```

## Request OTP 

Requests a One-Time Password (OTP) for sensitive operations such as changing password.

- **Method**: `POST`
- **Endpoint**: `/otp`

### Request

#### Body (JSON)

- **company** (string): Company name.
- **email** (string): User's email address.
- **callback** (string): Callback URL for handling the OTP result.

#### Example

```json
{
  "company": "Example Corp",
  "email": "user@example.com",
  "callback": "https://example.com/otp/session123"
}
```

## Verify OTP 

Verifies the provided One-Time Password for sensitive actions.

- **Method**: `POST`
- **Endpoint**: `/valid-token`

### Request

#### Body (JSON)

- **otp** (string): One-Time Password to be verified.
- **email** (string): User's email address.
- **verificationKey** (string): Key for verifying the OTP from previous step.
- **callback** (string): Callback URL for handling the OTP verification result.

#### Example

```json
{
  "otp": "123456",
  "email": "user@example.com",
  "verificationKey": "abcdef",
  "callback": "https://example.com/valid-token/session123"
}
```

## Request to Change Password

Initiates a request to change the user's password.

- **Method**: `POST`
- **Endpoint**: `/change-password`

### Request

#### Body (JSON)

- **company** (string): Company name.
- **email** (string): User's email address.
- **token** (string): Valid token obtained from verify OTP step.
- **password** (string): New password for the user's account.
- **callback** (string): Callback URL for handling the change password result.

#### Example

```json
{
  "company": "Example Corp",
  "email": "user@example.com",
  "token": "exampleValidToken",
  "password": "newSecurePassword123",
  "callback": "https://example.com/change-password/session123"
}
```

## User Data

Requests user data from the **resource server** using the access token issued.

- **Method**: `GET`
- **Endpoint**: `/user`

### Request

#### Headers

- **Authorization** (string): Bearer token containing the access token.
- **DPoP** (string): DPoP proof similar to refresh token request. Signed with the same ephemeral private key in access token request.
  - Header
    - alg=RS256
    - typ=dpop+jwt
    - jwk=<ephemeral public key>
  - Payload
    - iat=<created time in seconds>
    - jti=<random UUID4 in hex>
    - ath=<base64 encoded SHA256 hash of the access token in ASCII>
    - htm="POST"
    - htu=<Resource endpoint URL, without anything towards the end like `?q1=abc` etc. i.e `http://localhost:8080/token`
    - exp=<iat+120>

#### Example Headers

```json
{
  "Authorization": "Bearer exampleAccessToken",
  "DPoP": "exampleDPoPProof"
}
```

## Sample Code

### Generate Code Verifier and Code Challenge

#### Node.js
```
const crypto = require('node:crypto');

function base64URLEncode(str) {
  return str.toString('base64')
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=/g, '');
}

function generateCodeVerifier() {
  const length = Math.floor(Math.random() * (128 - 43 + 1)) + 43;
  const randomBytes = crypto.randomBytes(length / 2);
  const codeVerifier = base64URLEncode(randomBytes);
  return codeVerifier;
}

function generateCodeChallenge(code_verifier) {
  const encoder = crypto.createHash('sha256');
  encoder.update(Buffer.from(code_verifier, 'ascii'));
  const hash = encoder.digest();

  const codeChallenge = base64URLEncode(hash);

  return codeChallenge;
}
```

#### Python 
```
import base64
import hashlib
import os
import random


def generate_pkce_code_verifier():
    length = random.randint(43, 128)
    random_bytes = os.urandom(length // 2)
    code_verifier = base64.urlsafe_b64encode(random_bytes).decode("ascii")[:length]
    return code_verifier


def generate_pkce_code_challenge(code_verifier):
    encoder = hashlib.sha256()
    encoder.update(bytes(code_verifier, "ascii"))
    hash = encoder.digest()

    code_challenge = base64.urlsafe_b64encode(hash).decode("ascii").replace("=", "")
    return code_challenge
```

### Generate DPoP

#### Node.js
```
const crypto = require('node:crypto');
const jose = require('node-jose');

function base64URLEncode(str) {
  return str.toString('base64')
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=/g, '');
}

function generateRandomString(length) {
  let result = '';
  const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';

  for (let i = 0; i < length; i++) {
    const randomIndex = Math.floor(Math.random() * characters.length);
    result += characters.charAt(randomIndex);
  }
  return result;
}

// ath: Base64urlencoded Hash of the access_token
async function generateDpop(url, ath, method, ephemeralKeyPair) {
  let now = Math.floor(Date.now() / 1000);
  let payload = {
    htu: url,
    htm: method,
    jti: generateRandomString(40),
    iat: now,
    exp: now + 120,
  };

  if (ath) {
    const encoder = crypto.createHash('sha256');
    encoder.update(Buffer.from(ath, 'utf-8'));
    const hash = encoder.digest();
    const encodedHash = base64URLEncode(hash);
    payload.ath = encodedHash;
  }

  let privateKey = await jose.JWK.asKey(ephemeralKeyPair.privateKey, "pem");
  const asciiEncodedPublicKey = Buffer.from(ephemeralKeyPair.publicKey).toString('base64').toString('ascii');
  let DPoP = await jose.JWS.createSign(
    { format: "compact", fields: { alg: 'RS256', typ: "dpop+jwt", jwk: asciiEncodedPublicKey } },
    privateKey
  )
    .update(JSON.stringify(payload))
    .final();

  return DPoP;
}
```

#### Python
```
def create_dpop_jwt(private_key, public_key, htu, htm, ath=None):
    header = {
        "alg": "RS256",
        "typ": "dpop+jwt",
        "jwk": base64.b64encode(public_key).decode("ascii"),
    }

    now = int(time.time())
    payload = {
        "iat": now,  # Creation time
        "jti": str(uuid.uuid4()),  # Unique identifier
        "htm": htm,  # HTTP Method
        "htu": htu,  # HTTP Target site w/o ? and fragments
        "exp": now + 120,
    }

    if ath:
        # Base64 encoded SHA256 of associated access token's ASCII
        # Needed if access token also presented
        payload.update({"ath": ath})  # Access Token's Hash

    token = jwt.encode(payload, private_key, algorithm="RS256", headers=header)

    return json.dumps(token)
```







