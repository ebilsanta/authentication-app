const HostedTokenStore = require("../services/hostedTokenStore");

async function login(req, res, next) {
  res.send({ login_url: process.env.AUTHORISATION_SERVER_URL });
}

async function authorize(req, res, next) {
  const identityJwt = req.body.identity_jwt;
  const sessionId = req.sessionId;

  try {
    // const authCode = await requestForAuthCode(identityJwt);

    // const { access_token, refresh_token, id_token } =
    //   await requestForAccessToken(authCode);
    const access_token = '123';
    const refresh_token = '123';
    const id_token = 'abc';
    const hostedTokenStore = new HostedTokenStore();
    hostedTokenStore.setTokens(sessionId, access_token, refresh_token, id_token);
    
    res.send({ access_token, refresh_token, id_token });
  } catch (error) {
    res.status(500).send({ error: error.message });
  }
}

async function user(req, res, next) {
  try {
    console.log(req.session);
    res.send('ok');
  } catch (error) {

  }
}

module.exports = {
  login,
  authorize,
  user
};
