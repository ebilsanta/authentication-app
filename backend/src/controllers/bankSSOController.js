require('dotenv').config();
const axios = require('axios');

async function login(req, res, next) {
    const params = new URLSearchParams({
        client_id: process.env.BANKSSO_CLIENT_ID,
        redirect_uri: process.env.BANKSSO_REDIRECT_URI,
        response_type: 'code',
        scope: 'openid profile'
    });
    const stringParams = params.toString();
    const bankLoginUrl = `${process.env.BANKSSO_URI}/authorize?${stringParams}`;
    res.redirect(bankLoginUrl);
}

async function authCodeCallback(req, res, next) {
    const authCode = req.query.code;
    try {
        const { data } = await axios({
            url: process.env.BANKSSO_URI + '/token',
            method: 'post',
            data: {
            client_id: process.env.BANKSSO_CLIENT_ID,
            client_secret: process.env.BANKSSO_CLIENT_SECRET,
            redirect_uri: process.env.BANKSSO_REDIRECT_URI,
            grant_type: 'authorization_code',
            code: authCode,
            },
        });
        req.session.access_token = data.access_token;
        req.session.id_token = data.id_token;
        req.session.save((err) => {
            if (err) {
                return res.send("error while saving session information");
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
    try {
        const { data } = await axios({
            method: 'get',
            url: process.env.BANKSSO_URI + '/userinfo',
            headers: {
                Authorization: `Bearer ${req.session.access_token}`
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