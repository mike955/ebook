import Base from '../../library/base';
import CategorytDao from '../dao/category_dao';
import * as sequelize from 'sequelize';

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

  async get(conditions: { id: number; category: string; category_name:string}) {
    let where = {};
    if (conditions.id != undefined) where['id'] = conditions.id;
    if (conditions.category != undefined) where['category'] = conditions.category;
    if (conditions.category_name != undefined) {
      where['category_name'] = { [sequelize.Op.like]: `%${conditions.category_name}%` }
    }
    let res = await this.categoryDao.getByConds({where});
    return res;
  }

  async getList(conditions: { page: number; count: string; ids:string}) {
    let getParams = { page: conditions.page, count: conditions.count };
    if (conditions.ids != undefined) {
      getParams['where'] = { id: { [sequelize.Op.in]: conditions.ids }}
    }
    let res = await this.categoryDao.getPageListByConds(getParams);
    return res;
  }

  async delete(id: number) {
    let res = await this.categoryDao.updateById(id, {is_delete: 1});
    return res.length;
  }
}