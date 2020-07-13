import * as _ from 'lodash';

import Base from '../../library/base';
import CategoryData from '../data/category_data';

export default class CategoryPage extends Base {
  public map:any;
  public categoryData:any;
  constructor() {
    super();
    this.map = {
      add: 'add',
      get: 'get',
      getList: 'getList',
      delete: 'delete',
    }
    this.categoryData = new CategoryData()
  }

  async add(params: any) {
    this.validate_params(params, {
      category: v => ['M', 'C', 'N', 'J', 'D', 'R', 'S', 'P', 'A', 'Z'].indexOf(v) > -1,
      category_name: v => _.isString(v),
    });
    const { category, category_name } = params;
    return await this.categoryData.add(category, category_name);
  }

  async get(params: { id: number; category: string; category_name:string}) {
    this.validate_params(params, {
      id: v => _.isNumber(v) || _.isUndefined(v),
      category: v => ['M', 'C', 'N', 'J', 'D', 'R', 'S', 'P', 'A', 'Z'].indexOf(v) > -1  || _.isUndefined(v),
      category_name: v => _.isString(v) || _.isUndefined(v),
    });
    if (params.id == undefined && params.category == undefined && params.category_name == undefined) {
      this.throw_sys_error('PARAMS_ERR', `error params.id`);
    }
    let conditions = _.pick(params, ['id', 'category', 'category_name']);
    return await this.categoryData.get(conditions);
  }

  async getList(params: { page: number; count: string; ids:string}) {
    this.validate_params(params, {
      page: v => _.isNumber(v) && v >= 1,
      count: v => _.isNumber(v) && v >= 0,
      ids: v => _.isArray(v)|| _.isUndefined(v),
    });
    let conditions = _.pick(params, ['page', 'count', 'ids']);
    return await this.categoryData.getList(conditions);
  }

  async delete(params: any) {
    this.validate_params(params, {
      id: v => _.isNumber(v) && v >= 0,
    });
    return await this.categoryData.delete(params.id);
  }
}