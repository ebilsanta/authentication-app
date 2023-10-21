require('dotenv').config()
const express = require('express')
const session = require('express-session');
const cookieParser = require('cookie-parser');
const app = express();
const port = 3000;
const hostedLoginRouter = require('./src/routes/hostedLoginRoutes');
const bankSSORouter = require('./src/routes/bankSSORoutes');

const oneDay = 1000 * 60 * 60 * 24;

app.use(session({
  secret: process.env.SESSION_SECRET,
  resave: false,
  saveUninitialized: true,
  cookie: { secure: true, maxAge: oneDay },
}))

app.use(cookieParser());

app.use('/hosted', hostedLoginRouter);
app.use('/bankSSO', bankSSORouter);

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
})