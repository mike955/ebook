import * as _ from 'lodash';
import logger, { Logger } from './logger';
// import logger from './logger';
const ErrCode = require('../config/error_code')

export default class Base {
  public _logger: Logger;
  constructor() {
    // this._logger = logger('ebook-category');
  }

  validate_params(params, def) {
    Object.keys(def).forEach(k => {
      if (!def[k](params[k])) {
        this.throw_params_error(k);
      }
    });
  }

  private throw_params_error(prop_name) {
    console.log(`error params.${prop_name}`);
    this.logger.error(`error params.${prop_name}`);
    throw this.new_sys_error('PARAMS_ERR', `error params.${prop_name}`);
  }

  throw_sys_error(e: string, err_msg?: string): never {
    if (typeof e == 'string') {
      let err = ErrCode[e];
      if (!err) {
        throw new Error('Invalid Error Code');
      }
      if (err_msg) {
        if (_.isString(err_msg)) {
          err.errmsg = err_msg;
        } else {
          err.errmsg = JSON.stringify(err_msg);
        }
      }
      throw err;
    } else {
      console.error(e);
      this.throw_sys_error('FORMAT_ERR', 'Invalid Error Format');
    }
  }

  new_sys_error(
    e: string | { errno: number; errmsg: string; errcode: string },
    err_msg?: string,
  ) {
    if (typeof e === 'string') {
      let err = ErrCode[e];
      if (!err) {
        throw new Error('Invalid Error Code');
      }
      if (err_msg) {
        if (_.isString(err_msg)) {
          err.errmsg = err_msg;
        } else {
          err.errmsg = JSON.stringify(err_msg);
        }
      }
      return err;
    } else {
      let clone = _.clone(e);
      return clone;
    }
  }

  get logger() {
    return this._logger = logger('ebook-category')
  }

}