const express = require('express');

const port = 8080;
const app = express();

app.use(express.static('static'));
app.listen(port);