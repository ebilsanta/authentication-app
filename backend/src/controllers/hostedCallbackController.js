const { eventEmitter } = require("../services/eventEmitter");

async function register(req, res, next) {
  try {
    const { email, message, verification_key } = req.body;
    const sessionID = req.params.sessionId;
    console.log('register callback', req.body)
    if (message === 'Successfully Registered!') {
      console.log("received verification key", verification_key)
      eventEmitter.emit(`verificationKey:${sessionID}`, verification_key);
    } else {
      eventEmitter.emit(`verificationKey:${sessionID}`, `error: ${message}`);
    }
  } catch (error) {
    eventEmitter.emit(`verificationKey:${sessionID}`, `error: ${error}`);
  } 

  res.send('Verification key received');
}

async function verifyEmail(req, res, next) {
  try {
    const { details, email, status } = req.body;
    const sessionID = req.params.sessionId;
    console.log('verify email callback', req.body)
    if (status === 'Success') {
      console.log("received successful email verification")
      eventEmitter.emit(`verifyEmailOTP:${sessionID}`, details);
    } else {
      eventEmitter.emit(`verifyEmailOTP:${sessionID}`, `error: ${details}`);
    }
  } catch (error) {
    eventEmitter.emit(`verifyEmailOTP:${sessionID}`, `error: ${error}`);
  }
  
  res.send('Message received');
}

async function login(req, res, next) {
  try {
    const { status, idToken } = req.body;
    console.log('login callback', req.body)
    const sessionID = req.params.sessionId;
    if (status === 'User verified') {
      console.log('received id token', idToken)
      eventEmitter.emit(`idToken:${sessionID}`, idToken);
    } else {
      eventEmitter.emit(`idToken:${sessionID}`, `error: ${status}`);
    }
  } catch (error) {
    eventEmitter.emit(`idToken:${sessionID}`, `error: ${error}`);
  }
  
  res.send('Token received');
}

async function authCode(req, res, next) {
  try {
    const sessionID = req.params.sessionId;
    console.log('auth code callback req.body', req.body)
    console.log("headers", req.body.headers)
    const location = req.body.headers.location;
    let authCode;
    if (location) {
      const params = location.split('?')[1];
      if (params.startsWith('code')) {
        authCode = response.split('=')[1].split("\"")[0];
        console.log('received authCode', authCode)
        eventEmitter.emit(`authCode:${sessionID}`, authCode)
      } else {
        eventEmitter.emit(`authCode:${sessionID}`, `error: ${params}`);
      }    
    } else {
      eventEmitter.emit(`authCode:${sessionID}`, `error: Could not get auth code from auth server`);
    }
  } catch (error) {
    console.log('error', error)
  }
  
  res.send(`Auth code received ${authCode}`);
  
}

async function token(req, res, next) {
  try {
    const sessionId = req.params.sessionId;
    const response = req.body.response;
    const responseObj = JSON.parse(response);
    const body = responseObj.body;
    console.log('token callback parsed body', body);

    if (!body.error) {
      const accessToken = body.access_token;
      const refreshToken = body.refresh_token;
      console.log('received access and refresh tokens')
      eventEmitter.emit(`accessToken:${sessionId}`, JSON.stringify({accessToken, refreshToken}));
    } else {
      eventEmitter.emit(`accessToken:${sessionId}`, `error: ${body.error}, ${body.error_description}`);
    }
  } catch (error) {
    console.log('error', error)
  }
  
  res.send("Token received");
}

async function refresh(req, res, next) {
  try {
    const sessionId = req.params.sessionId;
    console.log('refresh callback', req.body)
    const response = req.body.response;
    const responseObj = JSON.parse(response);
    const body = responseObj.body;
    if (!body.error) {
      const accessToken = body.access_token;
      console.log('received refreshed access token')
      eventEmitter.emit(`refresh:${sessionId}`, accessToken);
    } else {
      eventEmitter.emit(`refresh:${sessionId}`, `error: ${body.error}`);
    }
  } catch (error) {
    eventEmitter.emit(`refresh:${sessionId}`, `error: ${error}`);
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
