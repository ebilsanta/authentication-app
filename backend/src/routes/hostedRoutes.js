const express = require("express");
const router = express.Router();
const hostedController = require("../controllers/hostedController");
const {
  checkIdToken, checkAuth, checkEmailAndVerificationKey, checkCodeVerifierAndAuthCode
} = require("../middlewares/hostedMiddleware");
const { registrationValidator, loginValidator, verifyEmailValidator } = require("../validators/hostedValidators");

router.get('/health', (req, res) => {
  res.status(200).send('Backend (Hosted) is healthy!');
});

router.post("/register", registrationValidator, hostedController.register);

router.post("/verify-email", verifyEmailValidator, checkEmailAndVerificationKey, hostedController.verifyEmail);


router.post("/login", loginValidator, hostedController.login);

router.get("/authorize", checkIdToken, hostedController.authorize);

router.get("/token", checkCodeVerifierAndAuthCode, hostedController.token);

router.get("/user", checkAuth, hostedController.user);

module.exports = router;

