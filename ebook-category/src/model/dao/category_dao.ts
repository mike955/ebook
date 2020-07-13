import DaoMySQL from '../../library/dao_mysql';
import definition, { CategoryAttributes } from './definitions/category';
// import CategorytData from '../data/category_data';


export default class CategoryDao extends DaoMySQL<CategoryAttributes> {
  constructor(){
    super();
    this.model = this.db(definition);
  }
}

// export default class CategoryDao extends Base {
//   public map:any;
//   public categoryData:any;
//   constructor() {
//     super();
//     this.map = {
//       add: 'add',
//       get: 'get',
//       getList: 'getList',
//       delete: 'delete',
//     }
//     // this.categoryData = new CategorytData()
//   }

//   async add(params: any) {
//     this.validate_params(params, {
//       category: v => ['admin'].indexOf(v) > -1,
//       category_name: v => [1, 2, 3].indexOf(v) > -1,
//     });
//     // const { category, category } = params;
//     console.log(params)
//   }
//   async get(params: any) {
//     console.log(params)
//   }
//   async getList(params: any) {
//     console.log(params)
//   }
//   async delete(params: any) {
//     console.log(params)
//   }
// }