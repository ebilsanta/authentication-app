const express = require('express');
const router = express.Router();
const {
    login,
    authCodeCallback,
    userInfo
} = require('../controllers/bankSSOController');

const { introspectToken } = require('../middlewares/bankSSOMiddleware');

router.get('/login', login);

router.get('/callback', authCodeCallback);

router.get('/userInfo', introspectToken, userInfo);

module.exports = router;