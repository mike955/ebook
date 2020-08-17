import * as Koa from 'koa';
import * as body_parser from 'koa-bodyparser';
import * as logger from 'koa-morgan';
import Router from "./router";
// import Action from './library/action';

const app = new Koa();
const router = new Router();

// logger
app.use(logger('ebook-category'));

// error handling
app.use(async (ctx, next) => {
  try {
      await next();
  } catch (err) {
      console.error(err);
      ctx.status = 200;
      ctx.body = JSON.stringify({
          errno: 1010,
          errcode: 'unknown_route_error',
          errmsg: 'unknown route error',
      });
  }
});

// body parser
app.use(
  body_parser({
      enableTypes: ['json'],
      jsonLimit: '10240kb',
      onerror: function (err, ctx) {
          console.error(err);
          ctx.throw('body parse error');
      },
  }),
);

// router
app.use(router.routes)

app.use(async function (ctx: Koa.Context) {
  ctx.status = 404;
  ctx.body = 'Not Found';
});

// console.log("server start at 3000");
// app.listen("3000")

export = app;