const { requestForAuthCode, requestForAccessToken, checkForAuthCode } = require('../services/hostedServices');
const { getIdentityJwt } = require("../utils/tempUtils");
const {validationResult} = require('express-validator');

async function register(req, res, next) {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).send({errors: errors.array()});
  }
  // make request to authentication server
  return res.send('verification key');
}

async function login(req, res, next) {
  res.send({ login_url: process.env.AUTHORISATION_SERVER_URL });
}

async function authorize(req, res, next) {
  // get identityJwt from request here
  // const identityJwt = req.body.identity_jwt;
  const identityJwt = getIdentityJwt();
  const sessionID = req.sessionID;

  try {
    const codeVerifier = await requestForAuthCode(identityJwt, sessionID);

    const authCode = await checkForAuthCode(sessionID);


    const { access_token, refresh_token, id_token, ephemeral_keypair } =
      await requestForAccessToken(codeVerifier, authCode);
    

    // save to session
    // req.session.access_token = access_token;
    // req.session.refresh_token = refresh_token;
    // req.session.ephemeral_keypair = ephemeral_keypair;
    res.send(authCode)
    // res.send({ access_token, refresh_token, id_token });
  } catch (error) {
    res.status(500).send({ error: error.message });
  }
}

async function user(req, res, next) {
  try {
    console.log("req.headers= ", req.headers);
    res.send('ok');
  } catch (error) {

  }
}

module.exports = {
  register, 
  login,
  authorize,
  user
};
