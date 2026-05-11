const express = require("express");

const app = express();
const PORT = 8080;

app.get("/ping", (req, res) => {
  res.json({ message: "pong", service: "node-express" });
});

app.get("/compute", (req, res) => {
  let sum = 0;
  for (let i = 0; i < 1000000; i++) {
    sum += i;
  }
  res.json({ result: sum, service: "node-express" });
});

app.listen(PORT, () => {
  console.log(`Node Express service running on port ${PORT}`);
});