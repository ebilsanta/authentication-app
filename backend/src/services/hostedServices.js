const { generateEphemeralKeys, generateDpop, generateClientAssertion } = require("../utils/dpopUtils");
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
    callback: process.env.CLIENT_HOSTED_CALLBACK_URL + "register/" + sessionID,
  }
  console.log("registration req", data)
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

async function requestToVerifyEmail(otp, email, verificationKey, sessionID) {
  const data = {
    otp,
    email,
    verificationKey,
    callback: process.env.CLIENT_HOSTED_CALLBACK_URL + "verify-email/" + sessionID,
  }
  console.log("verify email req", data)
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
    callback: process.env.CLIENT_HOSTED_CALLBACK_URL + "login/" + sessionID,
  }
  console.log("login request", data)
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
      callback_url: process.env.CLIENT_HOSTED_CALLBACK_URL + "authcode/" + sessionID,
    }
    console.log("auth code request", queryParams)
    const response = await axios({
      url: process.env.API_URL + 'hosted/authcode',
      method: 'post',
      params: queryParams
    });
    return codeVerifier;
  } catch (error) {
    console.error('Error requesting for auth code from auth server:', error);
    throw new Error('Error requesting for auth code from auth server: ' + error.message);
  }
}

async function requestForAccessToken(codeVerifier, authCode, sessionID) {
  try {
    const ephemeralKeyPair = await generateEphemeralKeys();
    const tokenEndpoint = process.env.API_URL + 'hosted/token';
    const dPoPProof = await generateDpop(tokenEndpoint, null, 'POST', ephemeralKeyPair);
    
    const clientPrivateKey = process.env.ALLOWED_CLIENT_PVT_KEY.replace(/\\n/g, '\n');

    const clientAssertion = await generateClientAssertion(process.env.AUDIENCE, process.env.ALLOWED_CLIENT, clientPrivateKey);

    const data = {
      grant_type: 'authorization_code',
      authcode: authCode,
      dpop: dPoPProof,
      client_assertion: clientAssertion,
      redirect_url: "http://localhost:8000",
      code_verifier: codeVerifier,
      callback_url: process.env.CLIENT_HOSTED_CALLBACK_URL + "token/" + sessionID,
    }
    console.log("request access token", data)

    const response = await axios({
      url: tokenEndpoint,
      method: 'post',
      data: data
    });
    return ephemeralKeyPair;
  } catch (error) {
    console.error(error);
    throw new Error('Error requesting for access token from auth server: ' + error.message);
  }
}

async function requestToRefreshToken(refreshToken, publicKey, privateKey, sessionID) {
  try {
    const tokenEndpoint = process.env.API_URL + 'hosted/refresh';
    const ephemeralKeyPair = { publicKey, privateKey };
    const dPoPProof = await generateDpop(tokenEndpoint, refreshToken, 'POST', ephemeralKeyPair);
    const data = {
      grant_type: 'authorization_code',
      refresh_token: refreshToken,
      dpop: dPoPProof,
      callback_url: process.env.CLIENT_HOSTED_CALLBACK_URL + "refresh/" + sessionID,
    }
    console.log("request refresh token", data)
    const response = await axios({
      url: tokenEndpoint,
      method: 'post',
      data: data
    });
    return response.data;
  } catch (error) {
    console.error('Error refreshing access token from auth server:', error);
    throw new Error('Error refreshing access token from auth server: ' + error.message);
  }
}


function checkForAuthCode(sessionID) {
  console.log("waiting for auth code", sessionID)
  return new Promise((resolve, reject) => {
    let timeout;

    eventEmitter.on(`authCode:${sessionID}`, (authCode) => {
      clearTimeout(timeout); 
      console.log(`Value of key 'authCode:${sessionID}': ${authCode}`);
      if (authCode.startsWith('error')) {
        reject(new Error(authCode));
      }
      resolve(authCode);
    });

    timeout = setTimeout(() => {
      eventEmitter.removeAllListeners(`authCode:${sessionID}`);
      reject(new Error(`Timeout waiting for auth code`));
    }, 20000); 
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
    }, 20000); 
  });
}

function checkForEmailVerification(sessionID) {
  console.log("waiting for verify email result", sessionID)
  return new Promise((resolve, reject) => {
    let timeout;

    eventEmitter.on(`verifyEmailOTP:${sessionID}`, (details) => {
      clearTimeout(timeout); 
      console.log(`Value of key verifyEmailOTP':${sessionID}': ${details}`);
      if (details.startsWith('error')) {
        reject(new Error(details));
      }
      resolve(details);
    });

    timeout = setTimeout(() => {
      eventEmitter.removeAllListeners(`verifyEmailOTP:${sessionID}`);
      reject(new Error(`Timeout waiting for email verification result`));
    }, 20000); 
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
    }, 20000); 
  });
}

function checkForAccessAndRefreshToken(sessionID) {
  console.log("waiting for access and refresh token", sessionID)
  return new Promise((resolve, reject) => {
    let timeout;

    eventEmitter.on(`accessToken:${sessionID}`, (details) => {
      clearTimeout(timeout); 
      console.log(`Value of key 'accessToken:${sessionID}': ${details}`);
      if (details.startsWith('error')) {
        reject(new Error(details));
      }
      resolve(details);
    });

    timeout = setTimeout(() => {
      eventEmitter.removeAllListeners(`accessToken:${sessionID}`);
      reject(new Error(`Timeout waiting for access token`));
    }, 20000); 
  });
}

function checkForRefreshedAccessToken(sessionID) {
  console.log("waiting for refreshed access token", sessionID)
  return new Promise((resolve, reject) => {
    let timeout;

    eventEmitter.on(`refresh:${sessionID}`, (details) => {
      clearTimeout(timeout); 
      console.log(`Value of key 'refresh:${sessionID}': ${details}`);
      if (details.startsWith('error')) {
        reject(new Error(details));
      }
      resolve(details);
    });

    timeout = setTimeout(() => {
      eventEmitter.removeAllListeners(`refresh:${sessionID}`);
      reject(new Error(`Timeout waiting for refreshed access token`));
    }, 20000); 
  });
}

module.exports = {
  requestForRegistration,
  requestToVerifyEmail,
  requestForLogin,
  requestForAuthCode,
  requestForAccessToken,
  requestToRefreshToken,
  checkForVerificationKey,
  checkForIdToken,
  checkForAuthCode,
  checkForEmailVerification,
  checkForAccessAndRefreshToken,
  checkForRefreshedAccessToken
}
