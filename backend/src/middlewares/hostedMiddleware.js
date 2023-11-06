const jwt = require('jsonwebtoken');
const { requestToRefreshToken, checkForRefreshedAccessToken } = require('../services/hostedServices');

async function checkIdToken(req, res, next) {
  if (!req.session.idToken) {
    return res.status(400).send({error: 'User is not authenticated'});
  }
  next();
}

async function checkEmailAndVerificationKey(req, res, next) {
  if (!req.session.email || !req.session.verificationKey) {
    return res.status(400).send({error: 'Missing email or verification key in session.'});
  }
  next();
}

async function checkCodeVerifierAndAuthCode(req, res, next) {
  if (!req.query.code) {
    return res.status(400).send({error: 'Missing auth code in query.'});
  }
  if (!req.session.codeVerifier) {
    return res.status(400).send({error: 'Missing code verifier in session.'});
  }
  next();
}

async function checkAuth(req, res, next) {
  if (!req.session.accessToken || !req.session.refreshToken || !req.session.privateKey || !req.session.publicKey) {
    return res.status(401).send({error: 'No tokens in session.'});
  }

  const { accessToken, refreshToken, privateKey, publicKey } = req.session;
  const sessionID = req.sessionID;

  const authZPubKey = process.env.AUTHZ_PUB_KEY.replace(/\\n/g, '\n');
  // try {

  //   const decodedAccessToken = jwt.verify(accessToken, authZPubKey);

  // } catch (error) {
  //   try {

  //     const decodedRefreshToken = jwt.verify(refreshToken, authZPubKey);

  //   } catch (error) {

  //     return res.status(401).send({error: 'Refresh token expired.'});

  //   }

    try {
      const response = await requestToRefreshToken(refreshToken, publicKey, privateKey, sessionID);

      const newAccessToken = await checkForRefreshedAccessToken(sessionID);

      req.session.accessToken = newAccessToken;

    } catch (error) {
        
      return res.status(401).send({error: 'Error refreshing access token.'});

    }
    
  // }
    
  next();
}

module.exports = {
  checkIdToken,
  checkEmailAndVerificationKey,
  checkCodeVerifierAndAuthCode,
  checkAuth, 
}