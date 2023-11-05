const {
  requestForAuthCode,
  requestForAccessToken,
  requestForLogin,
  checkForAuthCode,
  checkForVerificationKey,
  checkForVerificationResult,
  checkForIdToken,
  checkForAccessAndRefreshToken,
  requestForRegistration,
  requestForOtpVerification,
} = require("../services/hostedServices");
const { getIdentityJwt } = require("../utils/tempUtils");
const { validationResult } = require("express-validator");

async function register(req, res, next) {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).send({ errors: errors.array() });
  }
  const sessionID = req.sessionID;
  const jsonRequest = req.body;
  try {
    const response = await requestForRegistration(jsonRequest, sessionID);

    const verificationKey = await checkForVerificationKey(sessionID);

    req.session.verificationKey = verificationKey;
    req.session.email = jsonRequest.email;

    console.log(
      "stored in session verification key:",
      req.session.verificationKey
    );

    res.send("Successful Registration");
  } catch (error) {
    res.status(500).send({ error: error.message });
  }
}

async function otp(req, res, next) {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).send({ errors: errors.array() });
  }
  const sessionID = req.sessionID;
  const otp = req.body.otp;
  const { email, verificationKey } = req.session;
  try {
    const response = await requestForOtpVerification(
      otp,
      email,
      verificationKey,
      sessionID
    );

    const verificationResult = await checkForVerificationResult(sessionID);

    res.send(verificationResult);
  } catch (error) {
    res.status(500).send({ error: error.message });
  }
}

async function login(req, res, next) {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).send({ errors: errors.array() });
  }

  const sessionID = req.sessionID;
  const jsonRequest = req.body;

  try {
    const response = await requestForLogin(jsonRequest, sessionID);

    const idToken = await checkForIdToken(sessionID);

    req.session.idToken = idToken;

    console.log("stored in session id token:", req.session.idToken);

    res.redirect(process.env.HOSTED_API_URL + "authorize");
  } catch (error) {
    res.status(500).send({ error: error.message });
  }
}

async function authorize(req, res, next) {
  const identityJwt = req.session.idToken;

  const sessionID = req.sessionID;

  try {
    const codeVerifier = await requestForAuthCode(identityJwt, sessionID);

    req.session.codeVerifier = codeVerifier;

    const authCode = await checkForAuthCode(sessionID);
    console.log("redirect url: ", process.env.HOSTED_API_URL + "token?code=" + authCode)
    res.redirect(process.env.HOSTED_API_URL + "token?code=" + authCode);
  } catch (error) {
    res.status(500).send({ error: error.message });
  }
}

async function token(req, res, next) {
  const sessionID = req.sessionID;
  const authCode = req.query.code;
  const codeVerifier = req.session.codeVerifier;

  try {
    const ephemeralKeyPair = await requestForAccessToken(codeVerifier, authCode, sessionID);
    req.session.publicKey = ephemeralKeyPair.publicKey;
    req.session.privateKey = ephemeralKeyPair.privateKey;

    const accessAndRefreshTokens = await checkForAccessAndRefreshToken(sessionID);
    const { accessToken, refreshToken } = JSON.parse(accessAndRefreshTokens);
    req.session.accessToken = accessToken;
    req.session.refreshToken = refreshToken;

    res.redirect(process.env.HOSTED_API_URL + "user");
  } catch (error) {
    res.status(500).send({ error: error.message });
  }
}

async function user(req, res, next) {
  const sessionID = req.sessionID;
  const accessToken = req.session.accessToken;
  const email = req.session.email;

  res.send({email})
}

module.exports = {
  otp,
  register,
  login,
  authorize,
  token,
  user,
};
