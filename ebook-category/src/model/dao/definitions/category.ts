import * as Sequelize from 'sequelize';

const modelName = 'ebook_category';

const attributes = {
  id: { type: Sequelize.INTEGER, allowNull: false, primaryKey: true, autoIncrement: true },
  category: { type: Sequelize.STRING, allowNull: false, comment: '类别' },
  category_name: { type: Sequelize.STRING, allowNull: false, comment: '类别名称' },
  is_delete: { type: Sequelize.TINYINT, default: 0, comment: '删除标志,0:正常,1:删除' },
};

const defineOptions = {
  updatedAt: 'update_time',
  createdAt: 'create_time',
  freezeTableName: true,
};

export interface CategoryAttributes {
  id?: number,
  category?: string,
  category_name?: string,
  is_delete?: number;
}

export default {
  modelName,
  attributes,
  defineOptions,
};
