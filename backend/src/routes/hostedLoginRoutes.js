const express = require('express');
const router = express.Router();
const hostedLoginController = require('../controllers/hostedLoginController');
const { validateAuthorizeRequest, getAuthHeaders } = require("../middlewares/hostedLoginMiddleware");

router.get('/login', hostedLoginController.login);

router.post('/authorize', validateAuthorizeRequest, hostedLoginController.authorize);

router.get('/user', getAuthHeaders, hostedLoginController.user);

module.exports = router;