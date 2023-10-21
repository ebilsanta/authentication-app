const express = require('express');
const router = express.Router();
const hostedLoginController = require('../controllers/hostedLoginController');

router.get('/login', hostedLoginController.login);

router.get('/authorize', hostedLoginController.authorize);

module.exports = router;