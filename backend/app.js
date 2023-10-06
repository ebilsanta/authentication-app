require('dotenv').config()
const express = require('express')
const session = require('express-session');
const app = express();
const port = 3000;
const hostedLoginRouter = require('./routes/hostedLoginRoutes');

app.use(session({
  secret: process.env.SESSION_SECRET,
  resave: false,
  saveUninitialized: true
}))

app.use('/hosted', hostedLoginRouter);

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
})