const crypto = require('node:crypto');

// // The client generates a random string of bytes 
// function generateCodeVerifier() {
//   return base64URLEncode(crypto.randomBytes(32));
// }

// // The client hashes these bytes and encodes them for use in a url
// function generateCodeChallenge(codeVerifier) {
//   return base64URLEncode(crypto.createHash('sha256').update(codeVerifier).digest());
// }

function base64URLEncode(str) {
  return str.toString('base64')
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=/g, '');
  // urlSafe = str.toString('base64url')
  // while (urlSafe.length % 4 !== 0) {
  //   urlSafe += '=';
  // }
  // return urlSafe;
  
}
function generateCodeVerifier() {
  const length = Math.floor(Math.random() * (128 - 43 + 1)) + 43;
  const randomBytes = crypto.randomBytes(Math.ceil(length / 2));
  const codeVerifier = randomBytes.toString('hex').slice(0, length);
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
