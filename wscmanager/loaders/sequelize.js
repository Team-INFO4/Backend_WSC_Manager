import dbconfig from "../config/index";
import Sequelize from "sequelize";

const env = process.env.NODE_ENV || 'development';
const config = dbconfig[env];
const db = {};

const seqerlize = new Sequelize(config.database, config.username, config.password, config);

db.sequelize = seqerlize;

export default db;