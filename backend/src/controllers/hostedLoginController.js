const session = require("express-session");
const { generateCodeChallenge, generateCodeVerifier } = require("../utils/pkceUtils");
const { generateDPoPAndClientAssertion } = require("../services/hostedLoginServices");
const { v4: uuidv4 } = require('uuid');

// mock storage of client id to code verifier
var idToCodeVerifier = {};

var idToAuthCode = {};

async function login(req, res, next) {
  // redirect to authentication server here
  res.redirect(process.env.AUTHORISATION_SERVER_URL);
}

async function authorize(req, res, next) {
  const codeVerifier = generateCodeVerifier();
  const codeChallenge = generateCodeChallenge(codeVerifier);
  const uuid = uuidv4();
  idToCodeVerifier[uuid] = codeVerifier;
  // call to Authorization server here and receive auth code

  const { dPoPProof, clientAssertion } = generateDPoPAndClientAssertion(codeVerifier, uuid);

  // call to Authorization server here and receive ID Token, Access Token, Refresh Token

  // res.redirect("https://www.google.com");
  res.send('OK');
}

module.exports = {
  login,
  authorize,
}