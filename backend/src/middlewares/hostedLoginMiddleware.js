const HostedTokenStore = require('../services/hostedTokenStore');

async function validateAuthorizeRequest(req, res, next) {
  if (!req.body.identity_jwt) {
    res.status(400).send('Missing identity_jwt in request body.');
  }
  next();
}

async function getAuthHeaders(req, res, next) {
  // console.log(req.sessionID)
  console.log(req.session.cookie)
  if (!req.sessionID) {
    res.status(400).send('Missing session ID.');
    return;
  }
  const sessionId = req.sessionId;
  const hostedTokenStore = new HostedTokenStore();
  if (!hostedTokenStore.hasSession(sessionId)) {
    res.status(400).send('Invalid session ID.');
    return;
  }
  const { accessToken, refreshToken, dpopToken } = hostedTokenStore.getTokens(sessionId);
  req.headers['Authorization'] = 'Bearer ' + accessToken;
  req.headers['DPoP'] = dpopToken;
  next();
}

module.exports = {
  getAuthHeaders, 
  validateAuthorizeRequest
}