require("dotenv").config();
const express = require("express");
const session = require("express-session");
const app = express();
const port = 3000;
const hostedLoginRouter = require("./src/routes/hostedLoginRoutes");
const bankSSORouter = require("./src/routes/bankSSORoutes");

app.use(
  session({
    secret: process.env.SESSION_SECRET,
    resave: false,
    saveUninitialized: true,
  })
);

app.use(express.json());

app.use("/hosted", hostedLoginRouter);
app.use("/bankSSO", bankSSORouter);

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});