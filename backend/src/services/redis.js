const { createClient } = require("redis");

let redisClient = createClient({url: process.env.REDIS_URL});
redisClient.connect().catch(console.error);

module.exports = { redisClient }