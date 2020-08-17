CREATE DATABASE `ebook`;

DROP TABLE IF EXISTS `ebook`.`ebook_user`;
CREATE TABLE `ebook`.`ebook_user` (
    `id` BIGINT(32) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id` BINARY(32) NOT NULL COMMENT '账户id',
    `username` VARCHAR(256) NOT NULL COMMENT '账户名称',
    `email` VARCHAR(256) NOT NULL COMMENT '账户邮箱',
    `password` BINARY(32) NOT NULL COMMENT '账户名称',
    `salt` BINARY(32) NOT NULL COMMENT '盐',
    `role_id` TINYINT(4) DEFAULT 2 NOT NULL COMMENT '账户角色,1:管理员,2:普通用户',
    `status` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '账户状态,0:正常,1:关闭',
    `is_delete` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '账户状态,0:正常,1:删除',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
    `update_time` DATETIME NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帐号信息表';

DROP TABLE IF EXISTS `ebook`.`ebook_ebook`;
CREATE TABLE `ebook`.`ebook_ebook` (
    `id` BIGINT(32) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id` CHAR(32) NOT NULL COMMENT '用户id',
    `ebook_name` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '书名',
    `english_name` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '书英文名',
    `alias_name` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '别名',
    `category` INT(32) NOT NULL DEFAULT 0 COMMENT '种类',
    `publish_date` DATE NOT NULL COMMENT '出版时间',
    `key_words` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '关键字',


    `preview_type` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '文件预览图类别',
    `preview_size` BIGINT(32) NOT NULL DEFAULT 0 COMMENT '预览图大小',
    `preview_dir` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '预览图路径',
    `preview_upload_name` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '文件上传名称',
    `preview_hash_value` CHAR(32) NOT NULL DEFAULT '' COMMENT '文件hash值',

    `ebook_type` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '类别',
    `ebook_size` BIGINT(64) NOT NULL DEFAULT 0 COMMENT '文件大小',
    `ebook_dir` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '文件保存路径',
    `ebook_upload_name` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '文件上传名称',
    `ebook_hash_value` CHAR(32) NOT NULL DEFAULT '' COMMENT '文件hash值',

    `is_delete` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '状态,0:正常,1:删除',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
    `update_time` DATETIME NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='电子书信息表';

DROP TABLE IF EXISTS `ebook`.`ebook_type`;
CREATE TABLE `ebook`.`ebook_type` (
    `id` INT(16) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `type` VARCHAR(16) NOT NULL COMMENT '类别',
    `is_delete` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '删除标志,0:正常,1:删除',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
    `update_time` DATETIME NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='电子书类型';

DROP TABLE IF EXISTS `ebook`.`ebook_category`;
CREATE TABLE `ebook`.`ebook_category` (
    `id` INT(16) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `category` VARCHAR(16) NOT NULL COMMENT '类别',
    `category_name` VARCHAR(32) NOT NULL COMMENT '类别名称',
    `is_delete` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '删除标志,0:正常,1:删除',
    `create_time` DATETIME NOT NULL COMMENT '创建时间',
    `update_time` DATETIME NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='书籍类别信息表';

BEGIN;
INSERT INTO `ebook`.`ebook_category`(`category`, `category_name`, `create_time`, `update_time`) VALUES 
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

DROP TABLE if exists `ebook`.`ebook_user_role_map`;
CREATE TABLE `ebook`.`ebook_user_role_map` (
  `id` BIGINT(32) UNSIGNED NOT NULL AUTO_INCREMENT comment '主键id',
  `user_id` CHAR(32) NOT NULL comment '用户id',
  `role_id` BIGINT(32) UNSIGNED NOT NULL comment '角色id',
  `is_delete` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '是否删除,0:正常,1:删除',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UK_user_is_delete`(`user_id`,`is_delete`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色关系表';

DROP TABLE if exists `ebook`.`ebook_role`;
CREATE TABLE `ebook`.`ebook_role` (
  `id` BIGINT(32) UNSIGNED NOT NULL AUTO_INCREMENT comment '主键id,角色id',
  `role_name` VARCHAR(128) NOT NULL comment '角色名称',
  `role_desc` VARCHAR(256) DEFAULT NULL comment '角色描述',
  `is_delete` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '是否删除,0:正常,1:删除',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';
BEGIN;
INSERT INTO `ebook`.`ebook_role`(`role_name`, `role_desc`, `create_time`, `update_time`) VALUES 
  ('super_admin', '超级管理员', Now(), Now()),
  ('admin', '管理员', Now(), Now()),
  ('general', '普通用户', Now(), Now()),
  ('visitor', '访客(只读)', Now(), Now());
COMMIT;

DROP TABLE if exists `ebook`.`ebook_privilege`;
CREATE TABLE `ebook`.`ebook_privilege` (
  `id` BIGINT(32) UNSIGNED NOT NULL AUTO_INCREMENT comment '主键id,权限id',
  `privilege_name` VARCHAR(128) NOT NULL comment '权限名称',
  `uri` VARCHAR(128) NOT NULL comment '权限接口uri',
  `sn` CHAR(32) NOT NULL comment 'uri md5值',
  `privilege_desc` VARCHAR(256) DEFAULT NULL comment '权限描述',
  `is_delete` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '是否删除,0:正常,1:删除',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限表';
BEGIN;
INSERT INTO `ebook`.`ebook_privilege`(`privilege_name`, `uri`, `sn`, `privilege_desc`, `create_time`, `update_time`) VALUES 
  ('首页', '/', '6666cd76f96956469e7be39d750cc7d9', '首页，所有角色都具有该权限', Now(), Now());
COMMIT;

DROP TABLE if exists `ebook`.`ebook_role_privilege_map`;
CREATE TABLE `ebook`.`ebook_role_privilege_map` (
  `id` BIGINT(32) UNSIGNED NOT NULL AUTO_INCREMENT comment '主键id',
  `role_id`  BIGINT(32) UNSIGNED NOT NULL comment '角色id',
  `privilege_id` BIGINT(32) UNSIGNED NOT NULL comment '权限id',
  `is_delete` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '是否删除,0:正常,1:删除',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UK_role_privilege_id`(`role_id`,`privilege_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色权限关系表';