const crypto = require('crypto');

function base64URLEncode(str) {
  return str.toString('base64')
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=/g, '');
}

// The client generates a random string of bytes 
function generateCodeVerifier() {
  return base64URLEncode(crypto.randomBytes(32));
}

// The client hashes these bytes and encodes them for use in a url
function generateCodeChallenge(codeVerifier) {
  return base64URLEncode(crypto.createHash('sha256').update(codeVerifier).digest());
}

module.exports = {
  generateCodeVerifier,
  generateCodeChallenge
};
