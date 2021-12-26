const express = require("express");
const services = require("../services");
const utils = require("../utils");
const { validationResult, checkSchema } = require("express-validator");

const router = express.Router();
const schema = require("./schema.json");

router.post("/chart/demo", checkSchema(schema), async (req, res) => {
  //express-validator
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    return res.status(400).json({ message: errors.array(), success: false });
  }
  //handmade validator in utils
  try {
    utils.register("/chart/demo",req.body)
  } catch (e) {
    return res.status(400).json({ message: "Data have invalid value", success: false });
  }

  //builchart service
  try {
    const path = await services.buildChart(req.body);

    res.status(200).send({
      message: path,
      success: true,
    });
  } catch (err) {
    console.error(err);
    res.status(500).send({
      message: "Chart build failed",
    });
  }
});

module.exports = router;
