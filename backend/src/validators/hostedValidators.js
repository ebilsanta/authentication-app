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

const otpValidator = [
  body("email", "Email is empty").not().isEmpty(),
  body("email", "Invalid email").isEmail(),
  body("company", "Company is empty").not().isEmpty()
]

const changePasswordValidator = [
  body("company", "Company is empty").not().isEmpty(),
  body("password", "Password is empty").not().isEmpty(),
  body("otp", "OTP is empty").not().isEmpty(),
  body("otp", "OTP must be 6-digit numeric").isNumeric().isLength({ min: 6, max: 6 })
]


module.exports = {
  registrationValidator,
  verifyEmailValidator,
  loginValidator, 
  otpValidator,
  changePasswordValidator
}
