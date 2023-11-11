const {
  requestForAuthCode,
  requestForAccessToken,
  requestForLogin,
  requestForRegistration,
  requestToVerifyEmail,
  requestForOtp,
  requestToChangePassword,
  waitForEvent,
} = require("../services/hostedServices");
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

    const verificationKey = await waitForEvent("verificationKey", sessionID);

    req.session.verificationKey = verificationKey;
    req.session.email = jsonRequest.email;

    res.json({ message: "Successful Registration" });
  } catch (error) {
    let message = error.message;
    if (message.includes(":")) {
      message = message.split(":")[1].trim();
    }
    res.status(500).send({ error: message });
  }
}

async function verifyEmail(req, res, next) {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).send({ errors: errors.array() });
  }
  const sessionID = req.sessionID;
  const otp = req.body.otp;
  const { email, verificationKey } = req.session;
  try {
    const response = await requestToVerifyEmail(
      otp,
      email,
      verificationKey,
      sessionID
    );

    const verificationResult = await waitForEvent("verifyEmailOTP", sessionID);

    res.json({ message: verificationResult });
  } catch (error) {
    let message = error.message;
    if (message.includes(":")) {
      message = message.split(":")[1].trim();
    }
    res.status(500).json({ error: message });
  }
}

async function login(req, res, next) {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).json({ errors: errors.array() });
  }

  const sessionID = req.sessionID;
  const jsonRequest = req.body;
  req.session.email = jsonRequest.email;

  try {
    const response = await requestForLogin(jsonRequest, sessionID);

    const idToken = await waitForEvent("idToken", sessionID);

    req.session.idToken = idToken;

    res.redirect(process.env.CLIENT_HOSTED_URL + "authorize");
  } catch (error) {
    let message = error.message;
    if (message.includes(":")) {
      message = message.split(":")[1].trim();
    }
    res.status(500).json({ error: message });
  }
}

async function authorize(req, res, next) {
  const identityJwt = req.session.idToken;

  const sessionID = req.sessionID;

  try {
    const codeVerifier = await requestForAuthCode(identityJwt, sessionID);

    req.session.codeVerifier = codeVerifier;
    const authCode = await waitForEvent("authCode", sessionID);

    res.redirect(process.env.CLIENT_HOSTED_URL + "token?code=" + authCode);
  } catch (error) {
    let message = error.message;
    if (message.includes(":")) {
      message = message.split(":")[1].trim();
    }
    res.status(500).json({ error: message });
  }
}

async function token(req, res, next) {
  const sessionID = req.sessionID;
  const authCode = req.query.code;
  const codeVerifier = req.session.codeVerifier;

  try {
    const ephemeralKeyPair = await requestForAccessToken(
      codeVerifier,
      authCode,
      sessionID
    );
    req.session.publicKey = ephemeralKeyPair.publicKey;
    req.session.privateKey = ephemeralKeyPair.privateKey;

    const accessAndRefreshTokens = await waitForEvent("accessToken", sessionID);

    const { accessToken, refreshToken } = JSON.parse(accessAndRefreshTokens);
    req.session.accessToken = accessToken;
    req.session.refreshToken = refreshToken;

    res.redirect(process.env.CLIENT_HOSTED_URL + "user");
  } catch (error) {
    let message = error.message;
    if (message.includes(":")) {
      message = message.split(":")[1].trim();
    }
    res.status(500).json({ error: message });
  }
}

async function user(req, res, next) {
  const sessionID = req.sessionID;
  const accessToken = req.session.accessToken;
  const email = req.session.email;

  res.json({ email });
}

async function otp(req, res, next) {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).send({ errors: errors.array() });
  }
  const { company, email } = req.body;
  const sessionID = req.sessionID;
  try {
    const response = await requestForOtp(company, email, sessionID);

    const verificationKey = await waitForEvent("verificationKey", sessionID);

    req.session.verificationKey = verificationKey;
    req.session.email = email;

    res.json({ message: "OTP Sent!" });
  } catch (error) {
    let message = error.message;
    if (message.includes(":")) {
      message = message.split(":")[1].trim();
    }
    res.status(500).json({ error: message });
  }
}

async function changePassword(req, res, next) {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).send({ errors: errors.array() });
  }
  const { email, company, password, otp } = req.body;
  const sessionID = req.sessionID;
  const verificationKey = req.session.verificationKey;
  try {
    const response = await requestToChangePassword(
      verificationKey,
      company,
      email,
      otp,
      password,
      sessionID
    );

    const changePasswordResult = await waitForEvent(
      "changePassword",
      sessionID
    );

    res.json({ message: "Change password successfully" });
  } catch (error) {
    let message = error.message;
    if (message.includes(":")) {
      message = message.split(":")[1].trim();
    }
    res.status(500).json({ error: message });
  }
}

module.exports = {
  verifyEmail,
  register,
  login,
  authorize,
  token,
  user,
  otp,
  changePassword,
};
