const { generateEphemeralKeys, generateDpop, generateJKTThumbprint, generateClientAssertion } = require("../utils/dpopUtils");
const { generateCodeVerifier, generateCodeChallenge } = require("../utils/pkceUtils");
const axios = require('axios');
const { eventEmitter } = require("../services/eventEmitter")

async function requestForAuthCode(identityJwt, sessionID) {
  try {
    const codeVerifier = generateCodeVerifier();
    const codeChallenge = generateCodeChallenge(codeVerifier);
    const queryParams = {
      response_type: "code",
      state: codeChallenge,
      id_jwt: identityJwt,
      client_id: process.env.ALLOWED_CLIENT,
      code_challenge: codeChallenge,
      code_challenge_method: "S256",
      redirect_url: "http://localhost:8000",
      // redirect_url: "http://localhost:8000/api/hosted/callback/authcode/" + sessionID,
      callback_url: process.env.TEMP_CALLBACK_URL + "authcode"
      // callback_url: process.env.HOSTED_CALLBACK_URL + "authcode",
    }
    const response = await axios({
      url: process.env.AUTHZ_URL + 'authcode',
      method: 'post',
      params: queryParams
    });
    console.log(queryParams)
    return codeVerifier;
  } catch (error) {
    console.error('Error requesting for auth code from auth server:', error);
    throw new Error('Error requesting for auth code from auth server: ' + error.message);
  }
}

async function requestForAccessToken(codeVerifier, authCode) {
  try {
    const ephemeralKeyPair = await generateEphemeralKeys();
    const { publicKey, privateKey } = ephemeralKeyPair;
    const dPoPProof = generateDpop(process.env.AUTHZ_URL, null, 'POST', ephemeralKeyPair);
    const jktThumbprint = generateJKTThumbprint(publicKey);
    const clientAssertion = generateClientAssertion(process.env.AUTHZ_URL, process.env.CLIENT_ID, privateKey, jktThumbprint);
    // to edit
    const { data } = await axios({
      url: process.env.AUTHZ_URL + 'token',
      method: 'post',
      data: {
        grant_type: 'authorization_code',
        authcode: authCode,
        dpop: dPoPProof,
        client_assertion: clientAssertion,
        redirect_url: "http://localhost:8000/api/hosted/callback/token/" + sessionID,
        code_verifier: codeVerifier,
      }
    });
    const { access_token, refresh_token, id_token } = data;
    return { access_token, refresh_token, id_token, ephemeral_keypair: ephemeralKeyPair };
  } catch (error) {
    throw new Error('Error requesting for access token from auth server: ' + error.message);
  }
}

function checkForAuthCode(sessionID) {
  console.log("waiting for auth code", sessionID)
  return new Promise((resolve, reject) => {
    let timeout;

    eventEmitter.on(`authCode:${sessionID}`, (authCode) => {
      clearTimeout(timeout); // Clear the timeout since event was received
      console.log(`Value of key 'authCode:${sessionID}': ${authCode}`);
      resolve(authCode);
    });

    // Set a timeout to reject the promise if event is not received in 5 seconds
    timeout = setTimeout(() => {
      eventEmitter.removeAllListeners(`authCode:${sessionID}`);
      reject(new Error(`Timeout waiting for auth code`));
    }, 50000); // 5 seconds
  });
}

module.exports = {
  requestForAuthCode,
  requestForAccessToken,
  checkForAuthCode
}
