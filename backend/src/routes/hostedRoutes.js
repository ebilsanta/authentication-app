const express = require("express");
const router = express.Router();
const { validationResult } = require("express-validator");
const hostedController = require("../controllers/hostedController");
const {
  validateAuthorizeRequest,
  getAuthHeaders,
} = require("../middlewares/hostedMiddleware");
const { registrationValidator } = require("../validators/hostedValidators");

router.get("/login", hostedController.login);

router.post("/authorize", validateAuthorizeRequest, hostedController.authorize);

router.get("/user", getAuthHeaders, hostedController.user);

router.post("/register", registrationValidator, hostedController.register);

module.exports = router;
