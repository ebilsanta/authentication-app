const jwt = require('jsonwebtoken');
const uuid = require('uuid');
require("dotenv").config();

const allowedClient = process.env.ALLOWED_CLIENT;
const issuer = process.env.ALLOWED_ISSUER;
const subject = "testing@test.com";
const audience = process.env.AUDIENCE;
const redirectUrl = process.env.ALLOWED_REDIRECT;

const iat = Math.floor(Date.now() / 1000);
const exp = iat + 3600;

const payload = {
  iss: issuer,
  sub: subject,
  aud: audience,
  iat: iat,
  exp: exp,
};

const additionalHeaders = {
  kid: uuid.v4(),
};

const privateKey = process.env.PVT_KEY.replace(/\\n/g, '\n').replace(/\\t/g, '\t');

const getIdentityJwt = () => {
  const testingJwt = jwt.sign(payload, privateKey, {
    algorithm: 'RS256',
    header: additionalHeaders,
  });
  return testingJwt;
}



module.exports = {
  getIdentityJwt
}