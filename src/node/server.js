'use strict';

const express = require('express');

const HOST = '0.0.0.0';
const PORT = 2000;

const app = express();

app.get('/', (req, res) => {
    console.log('GET /');

    res.send('Hello World\n');
});

app.listen(PORT, HOST);

console.log(`Running on http://${HOST}:${PORT}`);
