const { eventEmitter } = require("../services/eventEmitter");

async function register(req, res, next) {
  const sessionID = req.params.sessionId;
  try {
    const { email, message, verification_key } = req.body;
    console.log("register callback", req.body);
    if (message === "Successfully Registered!") {
      console.log("received verification key", verification_key);
      eventEmitter.emit(`verificationKey:${sessionID}`, verification_key);
    } else {
      eventEmitter.emit(`verificationKey:${sessionID}`, `error: ${message}`);
    }
  } catch (error) {
    eventEmitter.emit(`verificationKey:${sessionID}`, `error: ${error}`);
  }

  res.json({ message: "Verification key received" });
}

async function verifyEmail(req, res, next) {
  const sessionID = req.params.sessionId;
  try {
    const { details, email, status } = req.body;
    console.log("verify email callback", req.body);
    if (status === "Success") {
      console.log("received successful email verification");
      eventEmitter.emit(`verifyEmailOTP:${sessionID}`, details);
    } else {
      eventEmitter.emit(`verifyEmailOTP:${sessionID}`, `error: ${details}`);
    }
  } catch (error) {
    eventEmitter.emit(`verifyEmailOTP:${sessionID}`, `error: ${error}`);
  }

  res.json({ message: "Message received" });
}

async function login(req, res, next) {
  const sessionID = req.params.sessionId;
  try {
    const { status, idToken } = req.body;
    console.log("login callback", req.body);
    if (status === "User verified") {
      console.log("received id token", idToken);
      eventEmitter.emit(`idToken:${sessionID}`, idToken);
    } else {
      eventEmitter.emit(`idToken:${sessionID}`, `error: ${status}`);
    }
  } catch (error) {
    eventEmitter.emit(`idToken:${sessionID}`, `error: ${error}`);
  }

  res.json({ message: "Token received" });
}

async function authCode(req, res, next) {
  const sessionID = req.params.sessionId;
  try {
    console.log("auth code callback req.body", req.body);
    const response = JSON.parse(req.body.response);
    const location = response.headers.location;
    console.log("location, ", location);
    let authCode;
    if (location) {
      const params = location.split("?")[1];
      console.log("params ", params);
      if (params.startsWith("code")) {
        authCode = location.split("=")[1].split('"')[0];
        console.log("received authCode", authCode);
        eventEmitter.emit(`authCode:${sessionID}`, authCode);
      } else {
        eventEmitter.emit(`authCode:${sessionID}`, `error: ${params}`);
      }
    } else {
      eventEmitter.emit(
        `authCode:${sessionID}`,
        `error: Could not get auth code from auth server`
      );
    }
  } catch (error) {
    console.log("error", error);
  }

  res.json({ message: `Auth code received ${authCode}` });
}

async function token(req, res, next) {
  const sessionId = req.params.sessionId;
  try {
    const response = req.body.response;
    const responseObj = JSON.parse(response);
    const body = responseObj.body;
    console.log("token callback parsed body", body);

    if (!body.error) {
      const accessToken = body.access_token;
      const refreshToken = body.refresh_token;
      console.log("received access and refresh tokens");
      eventEmitter.emit(
        `accessToken:${sessionId}`,
        JSON.stringify({ accessToken, refreshToken })
      );
    } else {
      eventEmitter.emit(
        `accessToken:${sessionId}`,
        `error: ${body.error}, ${body.error_description}`
      );
    }
  } catch (error) {
    eventEmitter.emit(`accessToken:${sessionId}`, `error: ${error}`);
  }

  res.json({ message: "Token received" });
}

async function refresh(req, res, next) {
  const sessionId = req.params.sessionId;
  try {
    console.log("refresh callback", req.body);
    const response = req.body.response;
    const responseObj = JSON.parse(response);
    const body = responseObj.body;
    if (!body.error) {
      const accessToken = body.access_token;
      console.log("received refreshed access token");
      eventEmitter.emit(`refresh:${sessionId}`, accessToken);
    } else {
      eventEmitter.emit(`refresh:${sessionId}`, `error: ${body.error}`);
    }
  } catch (error) {
    eventEmitter.emit(`refresh:${sessionId}`, `error: ${error}`);
  }

  res.json({ message: "Token received" });
}

async function otp(req, res, next) {
  const sessionID = req.params.sessionId;
  try {
    const { message, verification_key } = req.body;
    console.log("otp callback", req.body);
    if (message === "OTP Sent!") {
      console.log("received otp verification key", verification_key);
      eventEmitter.emit(`verificationKey:${sessionID}`, verification_key);
    } else {
      eventEmitter.emit(`verificationKey:${sessionID}`, `error: ${message}`);
    }
  } catch (error) {
    eventEmitter.emit(`verificationKey:${sessionID}`, `error: ${error}`);
  }
  res.json({ message: "Verification key received" });
}

async function validToken(req, res, next) {
  const sessionID = req.params.sessionId;
  try {
    const { message, token } = req.body;
    console.log("valid token callback", req.body);
    if (message === "Success") {
      console.log("received valid token", token);
      eventEmitter.emit(`valid-token:${sessionID}`, token);
    } else {
      eventEmitter.emit(`valid-token:${sessionID}`, `error: ${message}`);
    }
  } catch (error) {
    eventEmitter.emit(`valid-token:${sessionID}`, `error: ${error}`);
  }
  res.json({ message: "Valid Token received" });
}

async function changePassword(req, res, next) {
  const sessionID = req.params.sessionId;
  try {
    const { status, details, email } = req.body;
    console.log("change password callback", req.body);
    if (status !== "Success") {
      eventEmitter.emit(`changePassword:${sessionID}`, `error: ${details}`);
    } else {
      console.log("received change password response: ", status);
      eventEmitter.emit(`changePassword:${sessionID}`, status);
    }
  } catch (error) {
    eventEmitter.emit(`changePassword:${sessionID}`, `error: Error changing password`);
  }
  res.json({ message: "Change password response received" });
}

module.exports = {
  register,
  verifyEmail,
  login,
  authCode,
  token,
  refresh,
  otp,
  validToken, 
  changePassword
};
