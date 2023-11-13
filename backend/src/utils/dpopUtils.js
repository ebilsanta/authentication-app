const crypto = require('node:crypto');
const jose = require('node-jose');
const { base64URLEncode } = require('./pkceUtils');

async function generateEphemeralKeys() {
  let options = {
    modulusLength: 2048, // Change the modulus length as needed
    publicKeyEncoding: {
      type: "spki",
      format: "pem",
    },
    privateKeyEncoding: {
      type: "pkcs8",
      format: "pem",
    },
  };

  let rsaKeyPair = crypto.generateKeyPairSync("rsa", options);

  return rsaKeyPair;
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

function generateRandomString(length) {
  let result = '';
  const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';

  for (let i = 0; i < length; i++) {
    const randomIndex = Math.floor(Math.random() * characters.length);
    result += characters.charAt(randomIndex);
  }
  return result;
}

async function generateClientAssertion(audience, clientId, privateKey) {
  let now = Math.floor((Date.now() / 1000));
  const payload = {
    iss: clientId,
    sub: clientId,
    aud: audience,
    exp: now + 120,
    iat: now,
  };

  const privateKeyPEM = await jose.JWK.asKey(privateKey, "pem");
  const clientAssertion = await jose.JWS.createSign(
    { format: "compact", fields: { typ: 'JWT', alg: 'RS256' } },
    privateKeyPEM, 
  )
    .update(JSON.stringify(payload))
    .final();

  return clientAssertion;
}

module.exports = {
  generateEphemeralKeys,
  generateDpop,
  generateClientAssertion
}