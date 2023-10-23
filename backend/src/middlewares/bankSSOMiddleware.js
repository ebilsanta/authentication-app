require('dotenv').config();
const BankTokenStore = require('../services/bankTokenStore')
const jose = require('jose')

async function introspectToken (req, res, next) {
    // TODO: get bankSSO JWK from JWKS server
    var alg;
    var jwk;
    const publicKey = await jose.importJWK(jwk, alg);

    if (!req.sessionID) {
        res.status(400).send('Missing session ID.');
        return;
    }

    const sessionId = req.sessionID;
    const bankTokenStore = new BankTokenStore();
    if (!bankTokenStore.hasSession(sessionId)) {
        res.status(400).send('Invalid session ID.');
        return;
    }

    const { accessToken } = bankTokenStore.getTokens(sessionId);
    try {
        const { payload } = await jose.jwtVerify(accessToken, publicKey, {
            issuer: 'Bank App',
        });
        next();
    } catch (error) {
        console.log('Error')
        console.error(error.message);
        // general message of 'Access denied to the resource'
        res.status(403).json({
            success: false,
            message: 'Access denied to the resource'
        });
    }
}

module.exports = {
    introspectToken
}
