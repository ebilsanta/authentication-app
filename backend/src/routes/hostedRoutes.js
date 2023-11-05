const express = require("express");
const router = express.Router();
const hostedController = require("../controllers/hostedController");
const callbackController = require("../controllers/hostedCallbackController");
const {
  getAuthHeaders, checkIdToken, checkCodeVerifier, checkAuth
} = require("../middlewares/hostedMiddleware");
const { registrationValidator, loginValidator, otpValidator } = require("../validators/hostedValidators");


router.post("/register", registrationValidator, hostedController.register);

router.post("/otp", otpValidator, hostedController.otp);


router.post("/login", loginValidator, hostedController.login);

router.get("/authorize", checkIdToken, hostedController.authorize);

router.get("/token", checkCodeVerifier, hostedController.token);

router.get("/user", checkAuth, hostedController.user);


router.post("/callback/register/:sessionId", callbackController.register);

router.post("/callback/otp/:sessionId", callbackController.otp);

router.post("/callback/login/:sessionId", callbackController.login);

router.post("/callback/authcode/:sessionId", callbackController.authCode);

router.post("/callback/token/:sessionId", callbackController.token);

router.post("/callback/refresh/:sessionId", callbackController.refresh)

module.exports = router;
