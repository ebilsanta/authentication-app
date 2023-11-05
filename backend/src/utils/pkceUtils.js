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

module.exports = {
  base64URLEncode,
  generateCodeVerifier,
  generateCodeChallenge
};
