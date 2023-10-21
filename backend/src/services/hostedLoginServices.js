const { generateEphemeralKeys, generateDpop, generateJKTThumbprint, generateClientAssertion } = require("../utils/dpopUtils");
const https = require('https');

// shld be called to get access token on each api request
function generateDPoPAndClientAssertion() {
  const ephemeralKeyPair = generateEphemeralKeys();
  const { publicKey, privateKey } = ephemeralKeyPair;
  const dPoPProof = generateDpop(process.env.AUTHORISATION_SERVER_URL, null, 'POST', ephemeralKeyPair);
  const jktThumbprint = generateJKTThumbprint(publicKey);
  const clientAssertion = generateClientAssertion(process.env.AUTHORISATION_SERVER_URL, process.env.CLIENT_ID, privateKey, jktThumbprint);
  return {
    dPoPProof, clientAssertion
  }

}

async function requestForAccessToken() {

}

module.exports = {
  generateDPoPAndClientAssertion,
  requestForAccessToken
}
