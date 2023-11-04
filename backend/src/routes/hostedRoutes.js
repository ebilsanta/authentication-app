const express = require("express");
const router = express.Router();
const hostedController = require("../controllers/hostedController");
const callbackController = require("../controllers/hostedCallbackController");
const {
  getAuthHeaders,
} = require("../middlewares/hostedMiddleware");
const { registrationValidator, loginValidator, otpValidator } = require("../validators/hostedValidators");


router.post("/register", registrationValidator, hostedController.register);

router.post("/otp", otpValidator, hostedController.otp);


router.post("/login", loginValidator, hostedController.login);

router.get("/authorize", hostedController.authorize);

router.get("/user", getAuthHeaders, hostedController.user);


router.post("/callback/register/:sessionId", callbackController.register);

router.post("/callback/otp/:sessionId", callbackController.otp);

router.post("/callback/login/:sessionId", callbackController.login);

router.get("/callback/authcode/:sessionId", callbackController.authCode);

router.get("/callback/token/:sessionId", callbackController.token);

module.exports = router;
