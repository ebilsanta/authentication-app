require("dotenv").config();
const jose = require("node-jose");
const fs = require("fs");

async function introspectToken(req, res, next) {
    // TODO: whether to get bankSSO JWK from JWKS server
    if (!req.sessionID) {
        return res.status(400).send("Missing session ID.");
    }

    if (!req.session.access_token || !req.session.id_token) {
        return res.status(401).send("Not authorized by Bank SSO.");
    }

    const publicKey = fs.readFileSync(process.env.BANKSSO_KEY, "utf-8");
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

            var currentTimestamp = new Date().getTime() / 1000;
            var tokenExpired = payload.exp <= currentTimestamp;

            if (payload.iss != 'Bank App') {
                throw new Error("Not issued by Bank SSO");
            } else if (tokenExpired) {
                throw new Error("Access token expired");
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
