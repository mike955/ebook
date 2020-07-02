DROP TABLE IF EXISTS `ebook`.`account`;
CREATE TABLE `ebook`.`account` (
    `id` BIGINT(32) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `account_id` CHAR(32) NOT NULL COMMENT '账户id',
    `account_name` VARCHAR(256) NOT NULL COMMENT '账户名称',
    `account_email` VARCHAR(256) NOT NULL COMMENT '账户邮箱',
    `account_password` BINARY(64) NOT NULL COMMENT '账户名称',
    `salt` BINARY(64) NOT NULL COMMENT '盐',
    `account_role` TINYINT(4) DEFAULT 2 NOT NULL COMMENT '账户角色,1:管理员,2:普通用户',
    `status` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '账户状态,0:正常,1:关闭',
    `is_delete` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '账户状态,0:正常,1:删除',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
    `update_time` DATETIME NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帐号信息表';

DROP TABLE IF EXISTS `ebook`.`book`;
CREATE TABLE `ebook`.`book` (
    `id` BIGINT(32) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `book_name` VARCHAR(32) NOT NULL COMMENT '书名',
    `english_name` VARCHAR(256) NOT NULL COMMENT '书英文名',
    `alias_name` VARCHAR(256) DEFAULT 2 NOT NULL COMMENT '别名',
    `category` TINYINT(4) DEFAULT 2 NOT NULL COMMENT '种类',
    `publish_time` DATE NOT NULL COMMENT '出版时间',
    `key_words` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '关键字',
    `is_delete` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '状态,0:正常,1:删除',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
    `update_time` DATETIME NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='电子书信息表';

DROP TABLE IF EXISTS `ebook`.`book_category`;
CREATE TABLE `ebook`.`book` (
    `id` BIGINT(32) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `category` VARCHAR(16) NOT NULL COMMENT '类别',
    `category_name` VARCHAR(32) NOT NULL COMMENT '类别名称',
    `is_delete` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '状态,0:正常,1:删除',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
    `update_time` DATETIME NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='书籍类别信息表';

BEGIN;
INSERT INTO `ebook`.`book_category`(`category`, `category_name`, `create_time`, `update_time`) VALUES 
  ('M', '专著（含古籍中的史、志论著）', Now(), Now()), 
  ('C', '论文集', Now(), Now()), 
  ('N', '报纸文章', Now(), Now()), 
  ('J', '期刊文章', Now(), Now()), 
  ('D', '学位论文',  Now(), Now()), 
  ('R', '研究报告', Now(), Now()), 
  ('S', '标准', Now(), Now()), 
  ('P', '专利', Now(), Now()), 
  ('A', '专著、论文集中的析出文献', Now(), Now()), 
  ('Z', '其他', Now(), Now());
COMMIT;

DROP TABLE if exists `ebook`.`account_role`;
CREATE TABLE ebook.pri_user (
  `id` BIGINT(32) UNSIGNED NOT NULL AUTO_INCREMENT comment '主键id',
  `account_id` CHAR(32) NOT NULL comment '用户id',
  `role_id` CHAR(16) UNSIGNED NOT NULL comment '角色id',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UK_user_role_id`(`user_id`,`role_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色关系表';

DROP TABLE if exists `ebook`.`role`;
CREATE TABLE `ebook`.`pri_user` (
  `id` BIGINT(32) UNSIGNED NOT NULL AUTO_INCREMENT comment '主键id',
  `role_id` CHAR(16) UNSIGNED NOT NULL comment '角色id',
  `role_name` VARCHAR(128) UNSIGNED NOT NULL comment '角色id',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UK_role_id`(`role_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

DROP TABLE if exists `ebook`.`role_privilege`;
CREATE TABLE ebook.pri_user (
  `id` BIGINT(32) UNSIGNED NOT NULL AUTO_INCREMENT comment '主键id',
  `role_id` CHAR(16) UNSIGNED NOT NULL comment '角色id',
  `privilege_id` CHAR(16) UNSIGNED NOT NULL comment '权限id',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UK_role_privilege_id`(`role_id`,`privilege_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色权限关系表';

DROP TABLE if exists `ebook`.`pri_privilege`;
CREATE TABLE `ebook`.`pri_privilege` (
  `id` BIGINT(32) UNSIGNED NOT NULL AUTO_INCREMENT comment '主键id',
  `privilege_id` CHAR(16) UNSIGNED NOT NULL comment '权限id',
  `privilege_name` VARCHAR(128) UNSIGNED NOT NULL comment '权限名称',
  `uri` VARCHAR(128) UNSIGNED NOT NULL comment '权限接口',
  `sn` CHAR(32) UNSIGNED NOT NULL comment 'uri md5值',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UK_privilege_id_uri`(`privilege_id`,`uri`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限表';