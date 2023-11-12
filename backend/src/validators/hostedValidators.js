const { body } = require("express-validator");

const registrationValidator = [
  body("company", "Company is empty").not().isEmpty(),
  body("email", "Email is empty").not().isEmpty(),
  body("email", "Invalid email").isEmail(),
  body("firstName", "First name is empty").not().isEmpty(),
  body("lastName", "Last name is empty").not().isEmpty(),
  body("birthdate", "Invalid birthdate").isISO8601(),
  body("password", "Password is empty").not().isEmpty(),
];

const verifyEmailValidator = [
  body("otp", "OTP is empty").not().isEmpty(),
  body("otp", "OTP must be 6-digit numeric").isNumeric().isLength({ min: 6, max: 6 })
]

const loginValidator = [
  body("company", "Company is empty").not().isEmpty(),
  body("email", "Email is empty").not().isEmpty(),
  body("email", "Invalid email").isEmail(),
  body("password", "Password is empty").not().isEmpty(),
]

const verifyOtpValidator = [
  body("otp", "OTP is empty").not().isEmpty(),
  body("otp", "OTP must be 6-digit numeric").isNumeric().isLength({ min: 6, max: 6 })
]

const changePasswordValidator = [
  body("password", "Password is empty").not().isEmpty(),
]

module.exports = {
  registrationValidator,
  verifyEmailValidator,
  loginValidator, 
  verifyOtpValidator,
  changePasswordValidator
}
