require('dotenv').config();
const axios = require('axios');
const { v4: uuidv4 } = require('uuid');

var idToAuthCode = {};

async function login(req, res, next) {
    const uuid = uuidv4();
    req.session.uuid = uuid;
    const params = new URLSearchParams({
        client_id: process.env.CLIENT_ID,
        // TODO: change callback uri
        redirect_uri: 'http://localhost:3000/bankSSO/callback',
        response_type: 'code',
        scope: 'openid profile'
    });
    const stringParams = params.toString();
    const bankLoginUrl = `https://smurnauth-production.fly.dev/oauth/authorize?${stringParams}`;
    res.redirect(bankLoginUrl);
}

async function authCodeCallback(req, res, next) {
    const authCode = req.query.code;
    const uuid = req.session.uuid
    idToAuthCode[uuid] = authCode;
    try {
        const { data } = await axios({
            url: 'https://smurnauth-production.fly.dev/oauth/token',
            method: 'post',
            data: {
            client_id: process.env.CLIENT_ID,
            client_secret: process.env.CLIENT_SECRET,
            // TODO: change callback uri
            redirect_uri: 'http://localhost:3000/bankSSO/callback',
            grant_type: 'authorization_code',
            code: authCode,
            },
        });
        const access_token = data.access_token;
        const expiry = data.expires_in;
        const id_token = data.id_token;
        console.log(access_token, id_token, expiry);
        // TODO: change where to redirect to , e.g. dashboard/homepage?
        res.redirect('https://www.google.com');
    } catch(error) {
        console.error('Error exchanging authorization code for access token with bank SSO:', error.message);
        res.status(500).send('Error exchanging authorization code for access token with bank SSO.');
    }
}

module.exports = {
    login,
    authCodeCallback
}
