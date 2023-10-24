require("dotenv").config();
const express = require("express");
const session = require("express-session");
const RedisStore = require("connect-redis").default;
const { createClient } = require("redis");
const cookieParser = require("cookie-parser");
const app = express();
const port = 3000;
const hostedLoginRouter = require("./src/routes/hostedLoginRoutes");
const bankSSORouter = require("./src/routes/bankSSORoutes");

const oneDay = 1000 * 60 * 60 * 24;

let redisClient = createClient({url: process.env.REDIS_URL});
redisClient.connect().catch(console.error);

let redisStore = new RedisStore({
  client: redisClient,
  prefix: "session:"
})

app.use(
  session({
    store: redisStore, 
    secret: process.env.SESSION_SECRET,
    resave: false,
    saveUninitialized: false,
  })
);

app.use(cookieParser(secret = process.env.SESSION_SECRET));

app.use(express.json());

app.use("/hosted", hostedLoginRouter);
app.use("/bankSSO", bankSSORouter);

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});