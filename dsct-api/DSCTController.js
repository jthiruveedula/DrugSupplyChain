var express = require("express");
var router = express.Router();
var bodyParser = require("body-parser");

router.use(bodyParser.urlencoded({ extended: true }));
router.use(bodyParser.json());

const txSubmit = require("./invoke");
const txFetch = require("./query");

//var TFBC = require("./FabricHelper");

// Creation of Drug
router.post("/DrugCreation", async function(req, res) {
  try {
    let result = await txSubmit.invoke("DrugCreation", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

// IssueAuth 
router.post("/IssueAuth", async function(req, res) {
  try {
    let result = await txSubmit.invoke("IssueAuth", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

// Order_Delear
router.post("/Order_Delear", async function(req, res) {
  try {
    let result = await txSubmit.invoke("Order_Delear", JSON.stringify(req.body));
    res.send(result);
  } catch (err) {
    res.status(500).send(err);
  }
});

// Order_Hosp
router.post("/Order_Hosp", async function(req, res) {
  //TFBC.getLC(req, res); req.body.lcId
  try {
    let result = await txFetch.query("Order_Hosp", req.body.lcId);
    res.send(JSON.parse(result));
  } catch (err) {
    res.status(500).send(err);
  }
});

// DrugHist
router.post("/DrugHist", async function(req, res) {
  try {
    let result = await txFetch.query("DrugHist", req.body.lcId);
    res.send(JSON.parse(result));
  } catch (err) {
    res.status(500).send(err);
  }
});

module.exports = router;
