const {
  requestForAuthCode,
  requestForAccessToken,
  checkForAuthCode,
  checkForVerificationKey,
  checkForVerificationResult,
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
    const response = await requestForOtpVerification(otp, email, verificationKey, sessionID);

    const verificationResult = await checkForVerificationResult(sessionID);

    res.redirect(verificationResult);
  } catch (error) {
    res.status(500).send({ error: error.message });
  }
}

async function login(req, res, next) {
  
}

async function authorize(req, res, next) {
  // get identityJwt from request here
  // const identityJwt = req.body.identity_jwt;
  const identityJwt = getIdentityJwt();
  const sessionID = req.sessionID;

  try {
    const codeVerifier = await requestForAuthCode(identityJwt, sessionID);

    const authCode = await checkForAuthCode(sessionID);

    const { access_token, refresh_token, id_token, ephemeral_keypair } =
      await requestForAccessToken(codeVerifier, authCode);

    // save to session
    // req.session.access_token = access_token;
    // req.session.refresh_token = refresh_token;
    // req.session.ephemeral_keypair = ephemeral_keypair;
    res.send(authCode);
    // res.send({ access_token, refresh_token, id_token });
  } catch (error) {
    res.status(500).send({ error: error.message });
  }
}

async function user(req, res, next) {
  try {
    console.log("req.headers= ", req.headers);
    res.send("ok");
  } catch (error) {}
}

module.exports = {
  otp,
  register,
  login,
  authorize,
  user,
};
