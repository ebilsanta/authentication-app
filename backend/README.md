## Hosted Endpoints
### Register
Request to register user with AuthN service
- **Method**: `POST`
- **Endpoint**: `/api/hosted/register`

#### Request

##### Body (JSON)

- **email** (string): User's email address.
- **firstName** (string): User's first name.
- **lastName** (string): User's last name.
- **birthdate** (string): User's birthdate in YYYY-MM-DD format.
- **password** (string): User's password.
- **company** (string): User's company eg. 'ascenda'.

#### Responses
- **200 OK**
  - Response Body:
    - "Successful Registration"

- **400 Bad Request**
  - Response Body (JSON):
    - **errors** (array):
      - (objects)
        - **type** (string): `field`
        - **value** (string): `200-01-27`
        - **msg** (string): `Invalid birthdate`
        - **path** (string): `birthdate`
        - **location** (string): `body`

- **500 Internal Server Error**
  - Response Body (JSON):
    - **error** (string): Error message
  - Example: `Timeout waiting for verification key`, `You are already a registered user, please proceed to login instead!`

#### Sample Request
```
curl --location --request POST 'http://localhost:3000/api/hosted/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "thaddeusleezx@gmail.com",
    "firstName": "Thaddeus",
    "lastName": "Lee",
    "birthdate": "2000-01-27",
    "password": "1234567",
    "company": "ascenda"
}'
```

### Verify-email

- **Method**: `POST`
- **Endpoint**: `/api/hosted/verify-email`

#### Request

##### Body (JSON)

- **otp** (string): OTP user received in email

#### Responses
- **200 OK**
  - Response Body:
    - "Your email is verified, please proceed to login!"

- **400 Bad Request**
  - Response Body (JSON):
    - **errors** (array):
      - (objects)
        - **type** (string): `field`
        - **value** (string): `abcd`
        - **msg** (string): `OTP must be a 6-digit numeric`
        - **path** (string): `otp`
        - **location** (string): `body`

- **500 Internal Server Error**
  - Response Body (JSON):
    - **error** (string): Error message
  - Example: `"Some details did not match"`

#### Sample Request
```
curl --location --request POST 'http://localhost:3000/api/hosted/verify-email' \
--header 'Content-Type: application/json' \
--data-raw '{
    "otp": "086742"
}'
```

### Login
Request to login user
- **Method**: `POST`
- **Endpoint**: `/api/hosted/login`

#### Request

##### Body (JSON)

- **email** (string): User's email address.
- **password** (string): User's password.
- **company** (string): User's company eg. 'ascenda'.

#### Responses
- **200 OK**
  - Response Body (JSON):
    - **email** (integer): User's email. (temporary now, will return full user data later)

- **400 Bad Request**
  - Response Body (JSON):
    - **errors** (array):
      - (objects)
        - **type** (string): `field`
        - **value** (string): `thaddeusleezxgmail.com`
        - **msg** (string): `Invalid email`
        - **path** (string): `email`
        - **location** (string): `body`

- **500 Internal Server Error**
  - Response Body (JSON):
    - **error** (string): Error message
  - Example: `"Password did not match"`

#### Sample Request
```
curl --location --request POST 'http://localhost:3000/api/hosted/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "company": "ascenda",
    "email": "thaddeusleezx@gmail.com",
    "password": "1234567"
}'
```

### User
Request to get user data
- **Method**: `GET`
- **Endpoint**: `/api/hosted/user`

#### Responses
- **200 OK**
  - Response Body (JSON):
    - **email** (integer): User's email. (temporary now, will return full user data later)

- **500 Internal Server Error**
  - Response Body (JSON):
    - **error** (string): Error message
  - Example: `"Refresh token expired."`, `"No tokens in session"`

#### Sample Request
```
curl --location --request GET 'http://localhost:3000/api/hosted/user'
```



