require('dotenv').config();
const BankTokenStore = require('../services/bankTokenStore');
const axios = require('axios');

async function login(req, res, next) {
    const params = new URLSearchParams({
        client_id: process.env.CLIENT_ID,
        redirect_uri: process.env.REDIRECT_URI,
        response_type: 'code',
        scope: 'openid profile'
    });
    const stringParams = params.toString();
    const bankLoginUrl = `${process.env.BANKSSO_URI}/authorize?${stringParams}`;
    res.redirect(bankLoginUrl);
}

async function authCodeCallback(req, res, next) {
    const sessionId = req.sessionID;
    const authCode = req.query.code;
    try {
        const { data } = await axios({
            url: process.env.BANKSSO_URI + '/token',
            method: 'post',
            data: {
            client_id: process.env.CLIENT_ID,
            client_secret: process.env.CLIENT_SECRET,
            redirect_uri: process.env.REDIRECT_URI,
            grant_type: 'authorization_code',
            code: authCode,
            },
        });
        const accessToken = data.access_token;
        const idToken = data.id_token;
        const bankTokenStore = new BankTokenStore();
        bankTokenStore.setTokens(sessionId, accessToken, idToken);
        req.session.save((err) => {
            if (err) {
                return res.send("err while saving session information");
            }
            // TODO: decide whether there's a placeholder dashboard page to redirect to, or to redirect straight to profile page (aka userinfo page)
            res.redirect('userInfo');
        });
    } catch(error) {
        console.error('Error exchanging authorization code for access token with bank SSO:', error.message);
        res.status(500).send('Error exchanging authorization code for access token with bank SSO.');
    }
}

async function userInfo(req, res, next) {
    const sessionId = req.sessionID;
    const bankTokenStore = new BankTokenStore();
    const { accessToken } = bankTokenStore.getTokens(sessionId);
    try {
        const { data } = await axios({
            method: 'get',
            url: process.env.BANKSSO_URI + '/userinfo',
            headers: {
                Authorization: `Bearer ${accessToken}`
            }
        });
        console.log(data);
        res.send(data);
    } catch (error) {
        console.error('Error retrieving user info from bank SSO:', error.message);
        res.status(500).send('Error retrieving user info from bank SSO');
    }
}

module.exports = {
    login,
    authCodeCallback,
    userInfo
}