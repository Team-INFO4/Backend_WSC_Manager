const express = require('express');
const logger = require('morgan');

const app = express();
const port = 3000;

app.use(logger());

app.listen(port, () => {
    console.log('WSCmanager listening at port 3000');
})