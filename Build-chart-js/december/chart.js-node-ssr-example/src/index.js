const express = require("express");
//middleware
const cors = require("cors");

const router = require("./api");

const app = express();

app.use(express.json());
app.use(cors());

app.use("/api", router);

const server = app.listen(process.env.PORT || 3000, () => {
  console.log("Server is running on port", server.address().port);
});
