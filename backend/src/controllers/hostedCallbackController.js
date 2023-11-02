const { eventEmitter } = require("../services/eventEmitter");

async function authCode(req, res, next) {
  const sessionId = req.params.sessionId;
  const locationHeader = req.headers.location;
  if (locationHeader) {
    const params = locationHeader.split('?')[1];
    if (params.startsWith('error')) {
      throw new Error("Error requesting auth code:", params);
    }
    const authCode = locationHeader.split('=')[1];
    eventEmitter.emit(`authCode:${sessionId}`, authCode)
    res.send("Auth code received");
  }
}

async function token(req, res, next) {
  const sessionId = req.params.sessionId;
  const locationHeader = req.headers.location;
  if (locationHeader) {
    const token = locationHeader.split('=')[1];
    eventEmitter.emit(`token:${sessionId}`, token)
    res.send("Token received");
  }
}

module.exports = {
  authCode,
  token
};
