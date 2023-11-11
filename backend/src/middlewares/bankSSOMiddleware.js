require("dotenv").config();
const jose = require("node-jose");

async function introspectToken(req, res, next) {
    if (!req.sessionID) {
        return res.status(400).send("Missing session ID.");
    }

    if (!req.session.access_token || !req.session.id_token) {
        return res.status(401).send("Not authorized by Bank SSO.");
    }

    const publicKey = process.env.BANKSSO_KEY;
    const pubKey = publicKey.replace(/\\n/g, "\n");

    let keystore = jose.JWK.createKeyStore();
    keystore
        .add(pubKey, "pem")
        .then((jwkKey) => {
            return jose.JWS.createVerify(jwkKey).verify(req.session.access_token);
        })
        .then((result) => {
            var payload = JSON.parse(Buffer.from(result.payload).toString());
            console.log(payload);
            var header = result.header;

            var currentTimestamp = new Date().getTime() / 1000;
            var tokenExpired = payload.exp <= currentTimestamp;

            if (payload.iss != 'Bank App') {
                throw new Error("Not issued by Bank SSO");
            } else if (tokenExpired) {
                throw new Error("Access token expired");
            } else if (header.kid != process.env.BANKSSO_CLIENT_ID) {
                throw new Error("Access token not meant for our client.")
            } 
            next();
        })
        .catch((error) => {
            console.log("Error");
            console.error(error.message);
            return res.status(403).end("Access denied to the resource!");
        });
}

module.exports = {
    introspectToken,
};
