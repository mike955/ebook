import * as Koa from 'koa';

export default class Router {
  async routes(koaCtx: Koa.Context, next: () => Promise<any>) {
    let url: any;
    let { method } = koaCtx.request;
    let query = koaCtx.query;
    let params = {};
    try {
      params = JSON.parse(query.params || '{}');
      if (method === 'POST') {
        const body = koaCtx.request.body;
        params = Object.assign(params, body);
      }
    } catch (error) {
      console.log('JSON.parse(params) error: ' + error);
      return 'JSON.parse(params) error';
    }
    url = koaCtx.request.url.split('?').shift();
    const urlComponentes = url.split('/');
    if (!url.startsWith('/ebook/category') === false || urlComponentes.length != 4) {
      await next();
    }
    const model = urlComponentes[2];
    const fname = urlComponentes[3];
    const Ctrl = require('../model/page/' + model + '_page').default;
    const ctrl = new Ctrl(koaCtx);
    if (ctrl.map[fname] == undefined) next();
    const fn = ctrl[ctrl.map[fname]];
    try {
      const result = await fn.call(ctrl, params);
      if (result.errno == undefined) {
        koaCtx.body = {
          errno: 0,
          errmsg: '',
          data: result,
        }
      } else {
        koaCtx.body = result
      }
    } catch (error) {
      console.log('execute fname error: ' + JSON.stringify(error));
      koaCtx.body = error;
    }
  }
}
