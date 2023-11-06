const jwt = require('jsonwebtoken');
const { requestToRefreshToken, checkForRefreshedAccessToken } = require('../services/hostedServices');

async function checkIdToken(req, res, next) {
  if (!req.session.idToken) {
    return res.status(400).send('User is not authenticated');
  }
  next();
}

async function checkCodeVerifier(req, res, next) {
  if (!req.session.codeVerifier) {
    return res.status(400).send('Missing code verifier in session.');
  }
  next();
}

async function checkAuth(req, res, next) {
  if (!req.sessionID) {
    return res.status(400).send('Missing session ID.');
  }
  if (!req.session.accessToken || !req.session.refreshToken || !req.session.privateKey || !req.session.publicKey) {
    return res.status(401).send('Not authorized.');
  }

  const { accessToken, refreshToken, privateKey, publicKey } = req.session;
  const sessionID = req.sessionID;

  const authZPubKey = process.env.AUTHZ_PUB_KEY.replace(/\\n/g, '\n');
  try {

    const decodedAccessToken = jwt.verify(accessToken, authZPubKey);

  } catch (error) {
    try {

      const decodedRefreshToken = jwt.verify(refreshToken, authZPubKey);

    } catch (error) {

      return res.status(401).send('Refresh token expired.');

    }

    try {
      const response = await requestToRefreshToken(refreshToken, publicKey, privateKey, sessionID);

      const newAccessToken = await checkForRefreshedAccessToken(sessionID);

      req.session.accessToken = newAccessToken;

    } catch (error) {
        
      return res.status(401).send('Error refreshing access token.');

    }
    
  }
    
  next();
}

module.exports = {
  checkIdToken,
  checkCodeVerifier,
  checkAuth, 
}