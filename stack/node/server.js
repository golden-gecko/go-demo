'use strict';

const express = require('express');

const HOST = '0.0.0.0';
const PORT = 2000;

const app = express();

app.get('/', (req, res) => {
    res.send('Hello World');
});

app.listen(PORT, HOST);

console.log(`Running on http://${HOST}:${PORT}`);
