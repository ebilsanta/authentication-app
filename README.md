# CS301 G2T1 Authentication App

## Description

This project is an Authentication (AuthN) service that allows a company's partners to leverage its managed user services. We have implemented a system that provides identity and Authorization (AuthZ) for secure data exchange between partners and the company's existing services, such as a rewards platform.

## Solution Architecture

![CS301-G2T1-Architecture](https://github.com/cs301-itsa/project-2023-24t1-project-2023-24t1-g2-t1/assets/101983505/37b04a71-dc13-4161-b1f9-3895d87a1fd6)

## Live Website

We have deployed a frontend to test the end-to-end flow of our features.  
You can check it out at https://www.itsag2t1.com/

## API Endpoints
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
  - Refer to Appendix for sample code
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
  - A JWT with the following fields, signed with an ephemeral private key
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

- **client_assertion** (string): JWT with the following fields, signed with client's private key
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
- **dpop** (string): DPoP (Distributed Proof-of-Possession) proof similar to access token request, but with additional `ath` field in payload. Signed with the same ephemeral private key in access token request.
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





