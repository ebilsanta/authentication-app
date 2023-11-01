const { body } = require("express-validator");

const registrationValidator = [
  body("email", "Email is empty").not().isEmpty(),
  body("email", "Invalid email").isEmail(),
  body("firstName", "First name is empty").not().isEmpty(),
  body("lastName", "Last name is empty").not().isEmpty(),
  body("birthdate", "Invalid birthdate").isISO8601(),
];

module.exports = {
  registrationValidator
}
