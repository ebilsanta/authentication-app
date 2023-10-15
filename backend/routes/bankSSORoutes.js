const express = require('express');
const router = express.Router();
const {
    login,
    authCodeCallback
} = require('../controllers/bankSSOController');

router.get('/login', login);

router.get('/callback', authCodeCallback);

module.exports = router;