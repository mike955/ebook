import Base from '../../library/base';
import CategorytDao from '../dao/category_dao';

export default class CategoryData extends Base {
  public map:any;
  public categoryDao:any;
  constructor() {
    super();
    this.categoryDao = new CategorytDao()
  }

  async add(category:string, category_name: string) {
    let res = await this.categoryDao.add({category, category_name});
    return res;
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