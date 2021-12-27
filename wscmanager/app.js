import express from "express";
import config from './config/index.js';

const app = express();

const server = app.listen(config.port, () => {
    console.log('WSC manager start at ', config.port);
})

export { server }