
import Base from './base';
import * as sequelize from 'sequelize';
import * as _ from 'lodash';

let _map: { [key: string]: sequelize.Sequelize } = {}

export default class DaoMySQL<TAttributes> extends Base {
  protected model: any;
  // public model: sequelize.Model<TAttributes extends {}, TAttributes>;

  constructor() {
    super();
  }

  // db(definition): sequelize.Model<sequelize.ModelAttributes & TAttributes, TAttributes> {
  db(definition): any {
    const config = require('../config/global');
    const key = `${config.mysqlHost}/${config.mysqlPort}/${config.mysqlDatabase}`
    const instance = _map[key];
    if (instance == undefined) {
      const newSequelize = new sequelize.Sequelize(config.mysqlDatabase, config.mysqlUser, config.mysqlPassword, {
        host: config.mysqlHost,
        port: config.mysqlPort,
        dialect: config.dialect || 'mysql',
        pool: {
          max: config.maxConnections || 10,
          min: config.minConnections || 2,
          idle: config.maxIdleTime || 20000,
        },
        timezone: '+08:00',
        // logging: !!config.mysqlLogging,
        dialectOptions: { decimalNumbers: true }
      });
      _map[key] = newSequelize;
    }
    const model = _map[key].define(definition.modelName, definition.attributes, definition.defineOptions);
    return model;
  }

  getDb() {
    return (this.model as any).__db;
  }

  protected async transaction() {
    return (await this.getDb().transaction()) as sequelize.Transaction;
  }

  // protected query(
  //   sql: string | { query: string; values: any[] },
  //   options?: sequelize.QueryOptions,
  // ): Promise<any> {
  //   const db = this.getDb();
  //   return db.query.apply(db, arguments);
  // }

  getWhere(where) {
    var like = where._like;
    for (var i in like) {
      if (where[like[i]]) {
        where[like[i]] = { $like: '%' + where[like[i]] + '%' };
      }
    }
    return _.omit(where, '_like');
  }

  async batch_add(arrFields_list) {
    try {
      let list = await this.model.bulkCreate(arrFields_list);
      return _.map(list, 'dataValues');
    } catch (err) {
      this.logger.error(err);
      throw err;
    }
  }

  async add(arrFields: TAttributes): Promise<TAttributes> {
    try {
      let result = (await this.model.create(arrFields)) as any;
      return result.dataValues;
    } catch (err) {
      this.logger.error(err);
      throw err;
    }
  }

  async getById(id) {
    try {
      let where = { id } as any;
      return await this.model.findOne({ where: where, raw: true });
    } catch (err) {
      this.logger.error(err);
      throw err;
    }
  }

  async getListByIds(arrIds) {
    try {
      let where = { id: { $in: arrIds } } as any;
      return await this.model.findAll({ where: where, raw: true });
    } catch (err) {
      this.logger.error(err);
      throw err;
    }
  }

  async updateByConds(where, arrFields) {
    try {
      return await this.model.update(arrFields, { where: where, raw: true } as any);
    } catch (err) {
      this.logger.error(err);
      throw err;
    }
  }

  getCondsByOpt(opt) {
    try {
      var conds = {};
      opt = opt || {};
      var where = opt.where;
      if (where) {
        conds['where'] = this.getWhere(where);
      }
      var count = opt.count;
      var page = opt.page;
      if (count && page) {
        conds['limit'] = count;
        conds['offset'] = (page - 1) * count;
      }
      var limit = opt.limit;
      if (limit) conds['limit'] = limit;
      var offset = opt.offset;
      if (offset) conds['offset'] = offset;
      var order = opt.order;
      if (order) conds['order'] = order;
      var group = opt.group;
      if (group) conds['group'] = group;
      var include_fields = opt.include_fields;
      if (include_fields) conds['attributes'] = include_fields;
      conds['raw'] = true;
      return conds;
    } catch (err) {
      this.logger.error(err);
      throw err;
    }
  }

  async getByConds(opt) {
    try {
      return await this.model.findOne(this.getCondsByOpt(opt));
    } catch (err) {
      this.logger.error(err);
      throw err;
    }
  }

  async getListByConds(opt) {
    try {
      return await this.model.findAll(this.getCondsByOpt(opt));
    } catch (err) {
      this.logger.error(err);
      throw err;
    }
  }

  async count(opt) {
    try {
      var conds = {};
      opt = opt || {};
      var where = opt.where;
      if (where) {
        conds['where'] = this.getWhere(where);
      }
      var group = opt.group;
      if (group) conds['group'] = group;
      var include_fields = opt.include_fields;
      if (include_fields) conds['attributes'] = include_fields;
      return await this.model.count(conds);
    } catch (err) {
      this.logger.error(err);
      throw err;
    }
  }

  async updateById(id, arrFields) {
    try {
      console.log("======== arrFields: ", arrFields)
      return await this.model.update(arrFields, { where: { id: id } });
    } catch (err) {
      this.logger.error(err);
      throw err;
    }
  }

  async delById(id) {
    try {
      return await this.model.destroy({ where: { id: id } });
    } catch (err) {
      this.logger.error(err);
      throw err;
    }
  }

  async delByIds(arrIds) {
    try {
      return await this.model.destroy({ where: { id: { $in: arrIds } } });
    } catch (err) {
      this.logger.error(err);
      throw err;
    }
  }

  async delByCondition(where) {
    try {
      return await this.model.destroy({ where: where });
    } catch (err) {
      this.logger.error(err);
      throw err;
    }
  }

  async getPageListByConds(opt) {
    try {
      var conds = {};
      opt = opt || {};
      var where = opt.where;
      console.log(JSON.stringify(opt.where))
      if (where) {
        conds['where'] = this.getWhere(where);
      }
      var count = opt.count;
      var page = opt.page;
      if (count && page) {
        conds['limit'] = count;
        conds['offset'] = (page - 1) * count;
      }
      var limit = opt.limit;
      if (limit) conds['limit'] = limit;
      var offset = opt.offset;
      if (offset) conds['offset'] = offset;
      var order = opt.order;
      if (order) conds['order'] = order;
      var group = opt.group;
      if (group) conds['group'] = group;
      var include_fields = opt.include_fields;
      if (include_fields) conds['attributes'] = include_fields;
      conds['raw'] = true;
      console.log("=========")
      console.log(JSON.stringify(conds))
      return await this.model.findAndCountAll(conds);
    } catch (err) {
      this.logger.error(err);
      throw err;
    }
  }
}