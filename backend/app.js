require('dotenv').config()
const express = require('express')
const session = require('express-session');
const app = express();
const port = 3000;
const hostedLoginRouter = require('./routes/hostedLoginRoutes');
const bankSSORouter = require('./routes/bankSSORoutes');

app.use(session({
  secret: process.env.SESSION_SECRET,
  resave: false,
  saveUninitialized: true
}))

app.use('/hosted', hostedLoginRouter);
app.use('/bankSSO', bankSSORouter);

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
})