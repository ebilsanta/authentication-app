const { generateEphemeralKeys, generateDpop, generateJKTThumbprint, generateClientAssertion } = require("../utils/dpopUtils");
const { generateCodeVerifier, generateCodeChallenge } = require("../utils/pkceUtils");
const axios = require('axios');

async function requestForAuthCode(identityJwt) {
  try {
    const codeVerifier = generateCodeVerifier();
    const codeChallenge = generateCodeChallenge(codeVerifier);
    console.log("code verifier", codeVerifier)
    console.log("code challenge", codeChallenge, codeChallenge.length)
    const queryParams = {
      response_type: "code",
      state: codeChallenge,
      id_jwt: identityJwt,
      client_id: process.env.ALLOWED_CLIENT,
      code_challenge: codeChallenge,
      code_challenge_method: "S256",
      redirect_url: "http://localhost:8000",
      callback_url: "http://52.77.251.19/authcode"
    }
    const { data } = await axios({
      url: process.env.AUTHZ_URL + 'authcode',
      method: 'post',
      params: queryParams
    });
    const authCode = data.authCode;
    console.log(data)
    return {authCode, codeVerifier};
  } catch (error) {
    console.error('Error requesting for auth code from auth server:', error);
    throw new Error('Error requesting for auth code from auth server: ' + error.message);
  }
}

async function requestForAccessToken(authCode) {
  try {
    const ephemeralKeyPair = await generateEphemeralKeys();
    const { publicKey, privateKey } = ephemeralKeyPair;
    const dPoPProof = generateDpop(process.env.AUTHORISATION_SERVER_URL, null, 'POST', ephemeralKeyPair);
    const jktThumbprint = generateJKTThumbprint(publicKey);
    const clientAssertion = generateClientAssertion(process.env.AUTHORISATION_SERVER_URL, process.env.CLIENT_ID, privateKey, jktThumbprint);
    // to edit
    const { data } = await axios({
      url: process.env.AUTHORISATION_SERVER_URL + '/oauth/token',
      method: 'post',
      headers: {
        'DPoP': dPoPProof,
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      data: {
        grant_type: 'authorization_code',
        code: authCode,
        code_verifier: codeVerifier,
        client_assertion_type: 'urn:ietf:params:oauth:client-assertion-type:jwt-bearer',
        client_assertion: clientAssertion
      }
    });
    const { access_token, refresh_token, id_token } = data;
    return { access_token, refresh_token, id_token, ephemeral_keypair: ephemeralKeyPair };
  } catch (error) {
    throw new Error('Error requesting for access token from auth server: ' + error.message);
  }
}


module.exports = {
  requestForAuthCode,
  requestForAccessToken
}
