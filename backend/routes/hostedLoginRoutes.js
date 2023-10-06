const express = require('express');
const router = express.Router();
const hostedLoginController = require('../controllers/hostedLoginController');

router.get('/login', hostedLoginController.login);

router.get('/authcode', hostedLoginController.authCodeCallback);

module.exports = router;