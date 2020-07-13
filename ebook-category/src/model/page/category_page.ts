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
  async get(params: any) {
    console.log(params)
  }
  async getList(params: any) {
    console.log(params)
  }
  async delete(params: any) {
    console.log(params)
  }
}