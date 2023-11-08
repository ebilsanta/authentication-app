const express = require("express");
const router = express.Router();
const callbackController = require("../controllers/hostedCallbackController");

router.post("/register/:sessionId", callbackController.register);

router.post("/verify-email/:sessionId", callbackController.verifyEmail);

router.post("/login/:sessionId", callbackController.login);

router.post("/authcode/:sessionId", callbackController.authCode);

router.post("/token/:sessionId", callbackController.token);

router.post("/refresh/:sessionId", callbackController.refresh);

router.post("/otp/:sessionId", callbackController.otp);

router.post("/change-password/:sessionId", callbackController.changePassword);

module.exports = router;
