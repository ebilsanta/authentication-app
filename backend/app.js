require("dotenv").config();
const express = require("express");
const session = require("express-session");
const RedisStore = require("connect-redis").default;
const cors = require("cors");
const { redisClient } = require("./src/services/redis");
const app = express();
const port = 3000;
const hostedLoginRouter = require("./src/routes/hostedRoutes");
const bankSSORouter = require("./src/routes/bankSSORoutes");

const oneDay = 1000 * 60 * 60 * 24;

let redisStore = new RedisStore({
  client: redisClient,
  prefix: "session:",
});

const allowedOrigins = [
  "http://localhost:3000",
  "http://localhost:3001",
  "https://localhost:3000",
  "https://localhost:3001",
  process.env.FRONTEND_URL,
];

app.use(
  cors({
    origin: allowedOrigins,
    credentials: true,
  })
);

const cookieSettings = {
  sameSite: "none",
  secure: true,
  maxAge: oneDay,
  httpOnly: true,
};

app.use(
  session({
    store: redisStore,
    secret: process.env.SESSION_SECRET,
    resave: false,
    saveUninitialized: true,
    cookie: cookieSettings,
  })
);
console.log("cookie settings, ", cookieSettings);

app.use(express.json());

app.use("/api/hosted", hostedLoginRouter);
app.use("/api/bankSSO", bankSSORouter);

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
