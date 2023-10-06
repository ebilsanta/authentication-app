const { generateCodeChallenge, generateCodeVerifier } = require("../utils/pkceUtils");
const { v4: uuidv4 } = require('uuid');

// mock storage of client id to code verifier
var idToCodeVerifier = {};

var idToAuthCode = {};

async function login(req, res, next) {
  const codeVerifier = generateCodeVerifier();
  const codeChallenge = generateCodeChallenge(codeVerifier);
  const uuid = uuidv4();
  idToCodeVerifier[uuid] = codeVerifier;
  req.session.uuid = uuid;
  console.log("uuid", uuid);
  console.log("codeVerifier", codeVerifier);
  console.log("codeChallenge", codeChallenge);
  // call to Authorization server here, receive redirect link, forward this to browser
  res.redirect("https://www.google.com");
}

async function authCodeCallback(req, res, next) {
  const code = req.query.code;
  const uuid = req.session.uuid;
  idToAuthCode[uuid] = code;
  const codeVerifier = idToCodeVerifier[uuid];
  console.log("uuid", uuid);
  console.log("codeVerifier", codeVerifier);
  res.send({codeVerifier, code, uuid});
}

module.exports = {
  login,
  authCodeCallback
}