import { createLogger, format, transports } from 'winston';

const { combine, timestamp, printf } = format;
let _self_define_logger = {};

export interface Logger {
  debug(msg: any);
  info(msg: any);
  warn(msg: any);
  error(msg: any);
}

const appFormat = printf(info => {
  let str = info.message;
  if(typeof info.message === 'object'){
      str = JSON.stringify(str);
  }
  return `[${info.timestamp}] [${info.level}]: ${str}`;
});

export default function(app, log_level = 'info') {
  const key = app;
  if(process.env.FORCE_LOG_LEVEL){
      log_level = process.env.FORCE_LOG_LEVEL;
  }
  if (!_self_define_logger[key]) {
      const key_logger = createLogger({
          level: log_level,
          format: combine(timestamp(), appFormat),
          transports: [
              new transports.Console(),
          ],
      });
      _self_define_logger[key] = key_logger;
  }
  return _self_define_logger[key];
}