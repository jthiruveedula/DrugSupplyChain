var express = require('express');
var app = express();


var swaggerUi = require('swagger-ui-express');
var swaggerDocument = require('./swagger.json');


var DSCTController = require('./DSCTController');

app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerDocument));
app.use('/dsct', DSCTController);

module.exports = app;