const { generateEphemeralKeys, generateDpop, generateClientAssertion } = require("../utils/dpopUtils");

// shld be called to get access token on each api request
function requestForAccessToken(req, res, next) {
  const ephemeralKeyPair = generateEphemeralKeys();
  const { publicKey, privateKey } = ephemeralKeyPair;
  const dPopProof = generateDpop(req.url, null, req.method, ephemeralKeyPair);
  const clientAssertion = generateClientAssertion(req.url, process.env.CLIENT_ID, privateKey, publicKey);
}

module.exports = {
  requestForAccessToken
}
