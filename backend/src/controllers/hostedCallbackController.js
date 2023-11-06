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

async function login(req, res, next) {
  const { status, idToken } = req.body;
  const sessionID = req.params.sessionId;
  if (status === 'User verified') {
    eventEmitter.emit(`idToken:${sessionID}`, idToken);
    console.log("emitted:", `idToken:${sessionID}`)
  } else {
    eventEmitter.emit(`idToken:${sessionID}`, `error: ${status}`);
  }

  res.send('Token received');
}

async function authCode(req, res, next) {
  const sessionID = req.params.sessionId;
  const response = req.body.headers.location;
  let authCode;
  if (response) {
    const params = response.split('?')[1];
    if (params.startsWith('code')) {
      authCode = response.split('=')[1];
      eventEmitter.emit(`authCode:${sessionID}`, authCode)
    } else {
      eventEmitter.emit(`authCode:${sessionID}`, `error: ${params}`);
    }    
  }
  res.send(`Auth code received ${authCode}`);
  
}

async function token(req, res, next) {
  const sessionId = req.params.sessionId;
  const body = req.body.body;
  if (!body.error) {
    const accessToken = body.access_token;
    const refreshToken = body.refresh_token;
    eventEmitter.emit(`accessToken:${sessionId}`, JSON.stringify({accessToken, refreshToken}));
  } else {
    eventEmitter.emit(`accessToken:${sessionId}`, `error: ${body.error}`);
  }
  res.send("Token received");
}

async function refresh(req, res, next) {
  const sessionId = req.params.sessionId;
  const body = req.body.body;
  if (!body.error) {
    const accessToken = body.access_token;
    eventEmitter.emit(`refresh:${sessionId}`, accessToken);
  } else {
    eventEmitter.emit(`refresh:${sessionId}`, `error: ${body.error}`);
  }
  res.send("Token received");
}



module.exports = {
  register,
  otp,
  login,
  authCode,
  token,
  refresh
};
