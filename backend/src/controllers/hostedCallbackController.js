const { eventEmitter } = require("../services/eventEmitter");

async function register(req, res, next) {
  const { email, message, verification_key } = req.body;
  const sessionID = req.params.sessionId;
  if (message === 'Successfully Registered!') {
    console.log("received verification key", verification_key)
    eventEmitter.emit(`verificationKey:${sessionID}`, verification_key);
  } else {
    eventEmitter.emit(`verificationKey:${sessionID}`, `error: ${message}`);
  }

  res.send('Verification key received');
}

async function verifyEmail(req, res, next) {
  const { details, email, status } = req.body;
  const sessionID = req.params.sessionId;
  if (status === 'Success') {
    console.log("received successful email verification")
    eventEmitter.emit(`verifyEmailOTP:${sessionID}`, details);
  } else {
    eventEmitter.emit(`verifyEmailOTP:${sessionID}`, `error: ${details}`);
  }

  res.send('Message received');
}

async function login(req, res, next) {
  const { status, idToken } = req.body;
  const sessionID = req.params.sessionId;
  if (status === 'User verified') {
    console.log('received id token', idToken)
    eventEmitter.emit(`idToken:${sessionID}`, idToken);
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
      console.log('received authCode', authCode)
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
    console.log('received access and refresh tokens')
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
    console.log('received refreshed access token')
    eventEmitter.emit(`refresh:${sessionId}`, accessToken);
  } else {
    eventEmitter.emit(`refresh:${sessionId}`, `error: ${body.error}`);
  }
  res.send("Token received");
}



module.exports = {
  register,
  verifyEmail,
  login,
  authCode,
  token,
  refresh
};
