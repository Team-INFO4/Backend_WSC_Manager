import logger from 'morgan';

export default (app) => {
    app.use(logger());
}