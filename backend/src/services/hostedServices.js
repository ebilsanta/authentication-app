const { generateEphemeralKeys, generateDpop, generateJKTThumbprint, generateClientAssertion } = require("../utils/dpopUtils");
const { generateCodeVerifier, generateCodeChallenge } = require("../utils/pkceUtils");
const axios = require('axios');
const { eventEmitter } = require("../services/eventEmitter")

async function requestForRegistration(jsonRequest, sessionID) {
  const { company, email, firstName, lastName, birthdate, password } = jsonRequest;
  const data = {
    company,
    email,
    firstName,
    lastName,
    birthdate,
    password,
    callback: process.env.TEMP_CALLBACK_URL + "register/" + sessionID,
  }
  try {
    const response = await axios({
      url: process.env.API_URL + 'register',
      method: 'post',
      data: data
    });
    return response.data;
  } catch (error) {
    console.error('Error registering user:', error);
    throw new Error('Error registering user: ' + error.message);
  }
}

async function requestForOtpVerification(otp, email, verificationKey, sessionID) {
  const data = {
    otp,
    email,
    verificationKey,
    callback: process.env.TEMP_CALLBACK_URL + "otp/" + sessionID,
  }
  try {
    const response = await axios({
      url: process.env.API_URL + 'verify-email',
      method: 'post',
      data: data
    });
    return response.data;
  } catch (error) {
    console.error('Error verifying otp:', error);
    throw new Error('Error verifying otp: ' + error.message);
  }
}

async function requestForLogin(jsonRequest, sessionID) {
  const { company, email, password } = jsonRequest;
  const data = {
    company,
    email,
    password,
    callback: process.env.TEMP_CALLBACK_URL + "login/" + sessionID,
  }
  try {
    const response = await axios({
      url: process.env.API_URL + 'login',
      method: 'post',
      data: data
    });
    return response.data;
  } catch (error) {
    console.error('Error logging in:', error);
    throw new Error('Error logging in: ' + error.message);
  }
}

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
      callback_url: process.env.TEMP_CALLBACK_URL + "authcode/" + sessionID,
      // callback_url: process.env.HOSTED_CALLBACK_URL + "authcode",
    }
    const response = await axios({
      url: process.env.API_URL + 'hosted/authcode',
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

function checkForVerificationKey(sessionID) {
  console.log("waiting for verification key", sessionID)
  return new Promise((resolve, reject) => {
    let timeout;

    eventEmitter.on(`verificationKey:${sessionID}`, (authCode) => {
      clearTimeout(timeout); 
      console.log(`Value of key 'verificationKey:${sessionID}': ${authCode}`);
      if (authCode.startsWith('error')) {
        reject(new Error(authCode));
      }
      resolve(authCode);
    });

    timeout = setTimeout(() => {
      eventEmitter.removeAllListeners(`verificationKey:${sessionID}`);
      reject(new Error(`Timeout waiting for verification key`));
    }, 50000); 
  });
}

function checkForVerificationResult(sessionID) {
  console.log("waiting for verification result", sessionID)
  return new Promise((resolve, reject) => {
    let timeout;

    eventEmitter.on(`otpVerification:${sessionID}`, (details) => {
      clearTimeout(timeout); 
      console.log(`Value of key 'otpVerification:${sessionID}': ${details}`);
      if (details.startsWith('error')) {
        reject(new Error(details));
      }
      resolve(details);
    });

    timeout = setTimeout(() => {
      eventEmitter.removeAllListeners(`verificationResult:${sessionID}`);
      reject(new Error(`Timeout waiting for verification result`));
    }, 50000); 
  });
}

function checkForIdToken(sessionID) {
  console.log("waiting for id token", sessionID)
  return new Promise((resolve, reject) => {
    let timeout;

    eventEmitter.on(`idToken:${sessionID}`, (details) => {
      clearTimeout(timeout); 
      console.log(`Value of key 'idToken:${sessionID}': ${details}`);
      if (details.startsWith('error')) {
        reject(new Error(details));
      }
      resolve(details);
    });

    timeout = setTimeout(() => {
      eventEmitter.removeAllListeners(`loginResult:${sessionID}`);
      reject(new Error(`Timeout waiting for login result`));
    }, 50000); 
  });
}

module.exports = {
  requestForRegistration,
  requestForOtpVerification,
  requestForLogin,
  requestForAuthCode,
  requestForAccessToken,
  checkForVerificationKey,
  checkForIdToken,
  checkForAuthCode,
  checkForVerificationResult
}
