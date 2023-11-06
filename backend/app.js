require("dotenv").config();
const express = require("express");
const session = require("express-session");
const RedisStore = require("connect-redis").default;
const cors = require('cors');

const { redisClient } = require("./src/services/redis");
const cookieParser = require("cookie-parser");
const app = express();
const port = 3000;
const hostedLoginRouter = require("./src/routes/hostedRoutes");
const bankSSORouter = require("./src/routes/bankSSORoutes");

const oneDay = 1000 * 60 * 60 * 24;

let redisStore = new RedisStore({
  client: redisClient,
  prefix: "session:",
});

app.use(cors({
  origin: ['http://localhost:3000', 'http://localhost:3001'],
  credentials: true
}));

app.use(
  session({
    store: redisStore,
    secret: process.env.SESSION_SECRET,
    resave: false,
    saveUninitialized: true,
    cookie: {
      sameSite: "strict",
    }
  })
);

app.use(cookieParser((secret = process.env.SESSION_SECRET)));

app.use(express.json());

app.use("/api/hosted", hostedLoginRouter);
app.use("/api/bankSSO", bankSSORouter);

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
