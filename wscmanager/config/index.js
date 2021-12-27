import dotenv from 'dotenv';
import dbconfig from './dbconfig.json';

const envFound = dotenv.config();
if (envFound.error) {
  throw new Error("Couldn't find .env file");
}

export default {
    port : parseInt(process.env.PORT, 10),
    dbconfig : dbconfig
}