const jwt = require('jsonwebtoken');
const { requestToRefreshToken } = require('../services/hostedServices');

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
    const currentTimestamp = Math.floor(Date.now() / 1000)
    if (true || decodedAccessToken.exp < currentTimestamp) {
      const decodedRefreshToken = jwt.verify(refreshToken, authZPubKey);
      if (decodedRefreshToken.exp < currentTimestamp) {
        return res.status(401).send('Not authorized.');
      }

      const response = await requestToRefreshToken(refreshToken, publicKey, privateKey, sessionID);
      return res.send(response);
    }

    
  } catch (error) {
    console.error(error);
    return res.status(401).send('Not authorized.');
  }
  // check that access_token is valid, if not refresh it.

  // create dpop and attach to request
    
  next();
}

module.exports = {
  checkIdToken,
  checkCodeVerifier,
  checkAuth, 
}