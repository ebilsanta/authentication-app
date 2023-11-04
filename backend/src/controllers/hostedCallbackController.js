const { eventEmitter } = require("../services/eventEmitter");

async function register(req, res, next) {
  const { email, message, verification_key } = req.body;
  const sessionID = req.params.sessionId;
  if (message === 'Successfully Registered!') {
    eventEmitter.emit(`verificationKey:${sessionID}`, verification_key);
  } else {
    eventEmitter.emit(`verificationKey:${sessionID}`, `error: ${message}`);
  }

  res.send('Verification key received');
}

async function otp(req, res, next) {
  const { details, email, status } = req.body;
  const sessionID = req.params.sessionId;
  if (status === 'Success') {
    eventEmitter.emit(`otpVerification:${sessionID}`, details);
  } else {
    eventEmitter.emit(`otpVerification:${sessionID}`, `error: ${details}`);
  }

  res.send('Message received');
}


async function authCode(req, res, next) {
  const sessionId = req.params.sessionId;
  const locationHeader = req.headers.location;
  if (locationHeader) {
    const params = locationHeader.split('?')[1];
    if (params.startsWith('error')) {
      throw new Error("Error requesting auth code:", params);
    }
    const authCode = locationHeader.split('=')[1];
    eventEmitter.emit(`authCode:${sessionId}`, authCode)
    res.send("Auth code received");
  }
}

async function token(req, res, next) {
  const sessionId = req.params.sessionId;
  const locationHeader = req.headers.location;
  if (locationHeader) {
    const token = locationHeader.split('=')[1];
    eventEmitter.emit(`token:${sessionId}`, token)
    res.send("Token received");
  }
}

module.exports = {
  register,
  otp,
  authCode,
  token
};
