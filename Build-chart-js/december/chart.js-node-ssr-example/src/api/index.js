const express = require("express");

const router = express.Router();

router.use("/", require("./chart"));

module.exports = router;
