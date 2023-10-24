async function login(req, res, next) {
  res.send({ login_url: process.env.AUTHORISATION_SERVER_URL });
}

async function authorize(req, res, next) {
  const identityJwt = req.body.identity_jwt;
  const sessionID = req.sessionID;
  console.log(sessionID);

  try {
    // const authCode = await requestForAuthCode(identityJwt);

    // const { access_token, refresh_token, id_token, ephemeral_key } =
    //   await requestForAccessToken(authCode);
    const access_token = 'abc123';
    const refresh_token = '123';
    const id_token = 'abc';
    const ephemeral_keypair = 'keypair';

    // save to session
    req.session.access_token = access_token;
    req.session.refresh_token = refresh_token;
    req.session.ephemeral_keypair = ephemeral_keypair;
    
    res.send({ access_token, refresh_token, id_token });
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
  login,
  authorize,
  user
};
