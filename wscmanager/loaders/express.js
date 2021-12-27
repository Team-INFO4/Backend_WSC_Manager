import logger from 'morgan';
import express from 'express';
import routes from '../api/index';

export default (app) => {
    app.use(logger());
    app.use(express.json());
    app.use(express.urlencoded({ extended : true }));
    app.use(routes());
}