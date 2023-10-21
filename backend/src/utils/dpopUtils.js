const crypto = require('crypto');
const jose = require('jose');

async function generateEphemeralKeys() {
  let options = {
    namedCurve: "P-256",
    publicKeyEncoding: {
      type: "spki",
      format: "pem",
    },
    privateKeyEncoding: {
      type: "sec1",
      format: "pem",
    },
  };
  
  let ephemeralKeyPair = crypto.generateKeyPairSync("ec", options);

  return ephemeralKeyPair;
}

async function generateJKTThumbprint(publicKey) {
  const jwk = {
    kty: 'EC',
    crv: 'P-256',
    x: ephemeralKeyPair.publicKey.export({ format: 'jwk' }).x,
    y: ephemeralKeyPair.publicKey.export({ format: 'jwk' }).y
  };

  // Calculate the SHA-256 thumbprint
  const sha256Thumbprint = crypto.createHash('sha256')
    .update(JSON.stringify(jwk))
    .digest('base64');

  // Base64url encode the thumbprint
  const base64UrlThumbprint = sha256Thumbprint
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=/g, '');

  return base64UrlThumbprint;
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
    payload.ath = ath;
  }

  let privateKey = await jose.JWK.asKey(ephemeralKeyPair.privateKey, "pem");
  let jwk = (await jose.JWK.asKey(ephemeralKeyPair.publicKey, "pem")).toJSON(true);
  jwk.use = "sig";
  jwk.alg = "ES256";
  let DPoP = await jose.JWS.createSign(
    { format: "compact", fields: { typ: "dpop+jwt", jwk: jwk } },
    { key: privateKey, reference: false }
  )
    .update(JSON.stringify(payload))
    .final();

  return DPoP;
}

// jktThumbprint: base64url encoding of the JWK SHA-256 Thumbprint of the client's ephemeral public signing key used to sign the DPoP Proof JWT

async function generateClientAssertion(url, clientId, privateSigningKey, jktThumbprint) {
  let now = Math.floor((Date.now() / 1000));

  let payload = {
    'sub': clientId,
    'jti': generateRandomString(40),
    'aud': url,
    'iss': clientId,
    'iat': now,
    'exp': now + 300,
    'cnf' : {
      'jkt': jktThumbprint
    }
  };
  let jwsKey = await jose.JWK.asKey(privateSigningKey, 'pem');

  let jwtToken = await jose.JWS.createSign({ 'format': 'compact', 'fields': { 'typ': 'JWT' } }, jwsKey).update(JSON.stringify(payload)).final();
  return jwtToken;
};

module.exports = {
  generateEphemeralKeys,
  generateDpop,
  generateJKTThumbprint,
  generateClientAssertion
}