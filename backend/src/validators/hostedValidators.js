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

const otpValidator = [
  body("otp", "OTP is empty").not().isEmpty(),
]

const loginValidator = [
  body("company", "Company is empty").not().isEmpty(),
  body("email", "Email is empty").not().isEmpty(),
  body("email", "Invalid email").isEmail(),
  body("password", "Password is empty").not().isEmpty(),
]


module.exports = {
  registrationValidator,
  otpValidator,
  loginValidator
}
