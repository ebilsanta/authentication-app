const express = require("express");
const router = express.Router();
const hostedController = require("../controllers/hostedController");
const callbackController = require("../controllers/hostedCallbackController");
const {
  validateAuthorizeRequest,
  getAuthHeaders,
} = require("../middlewares/hostedMiddleware");
const { registrationValidator } = require("../validators/hostedValidators");

router.get("/login", hostedController.login);

router.post("/authorize", validateAuthorizeRequest, hostedController.authorize);

router.get("/user", getAuthHeaders, hostedController.user);

router.post("/register", registrationValidator, hostedController.register);

router.get("/callback/authcode/:sessionId", callbackController.authCode);

router.get("/callback/token/:sessionId", callbackController.token);

module.exports = router;
