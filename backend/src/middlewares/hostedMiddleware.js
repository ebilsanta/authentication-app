async function getAuthHeaders(req, res, next) {
  if (!req.sessionID) {
    return res.status(400).send('Missing session ID.');
  }
  if (!req.session.access_token || !req.session.refresh_token || !req.session.ephemeral_keypair) {
    return res.status(401).send('Not authorized.');
  }

  const { access_token, refresh_token, ephemeral_keypair } = req.session;
  // check that access_token is valid, if not refresh it.

  // create dpop and attach to request
  req.headers['Authorization'] = 'Bearer ' + access_token;
  
  next();
}

module.exports = {
  getAuthHeaders, 
}