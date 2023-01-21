-- 建表规范
-- https://github.com/yongxiongzhong/mls-group-mysql/blob/master/MySQL%E5%BB%BA%E8%A1%A8%E8%A7%84%E8%8C%83.md
-- https://zhuanlan.zhihu.com/p/342436738
-- https://juejin.cn/post/6844903928228757511
-- https://www.cnblogs.com/dragonsuc/p/6938006.html
-- sample_table_one
DROP TABLE IF EXISTS `user`;

CREATE TABLE user (
  `id` bigint(11) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(11) NOT NULL COMMENT ‘ 用户id ’ `username` varchar(45) NOT NULL COMMENT '真实姓名',
  `email` varchar(30) NOT NULL COMMENT ‘ 用户邮箱 ’,
  `nickname` varchar(45) NOT NULL COMMENT '昵称',
  `avatar` int(11) NOT NULL COMMENT '头像',
  `birthday` date NOT NULL COMMENT '生日',
  `sex` tinyint(4) DEFAULT '0' COMMENT '性别',
  `short_introduce` varchar(150) DEFAULT NULL COMMENT '一句话介绍自己，最多 50 个汉字',
  `user_resume` varchar(300) NOT NULL COMMENT '用户提交的简历存放地址',
  `user_register_ip` int NOT NULL COMMENT ‘ 用户注册时的源ip ’,
  `create_time` timestamp NOT NULL COMMENT ‘ 用户记录创建的时间 ’,
  `update_time` timestamp NOT NULL COMMENT ‘ 用户资料修改的时间 ’,
  `user_review_status` tinyint NOT NULL COMMENT ‘ 用户资料审核状态 ， 1 为通过 ， 2 为审核中 ， 3 为未通过 ， 4 为还未提交审核 ’,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_id` (`user_id`),
  KEY `idx_username`(`username`),
  KEY `idx_create_time`(`create_time`, `user_review_status`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT = '网站用户基本信息';

-- sample_table_two
CREATE TABLE `log_type` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `product_id` bigint(20) unsigned NOT NULL COMMENT '产品 ID',
  `product_name` varchar(100) NOT NULL COMMENT '产品名 (中文)',
  `type` varchar(100) NOT NULL COMMENT '接入分类',
  `log_path` TEXT NOT NULL COMMENT '日志采集路径',
  `exclude_log` TEXT COMMENT '排除的日志',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 5 DEFAULT CHARSET = utf8 ； -- sample_table_three
DROP TABLE IF EXISTS `sample_table_three`;

CREATE TABLE `sample_table_three` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `eng_name` varchar(100) NOT NULL COMMENT '应用英文名称 (唯一标识)',
  `name` varchar(100) NOT NULL COMMENT '应用中文名称',
  `desc` varchar(255) DEFAULT NULL COMMENT '描述',
  `principal` varchar(255) DEFAULT NULL COMMENT '应用负责人 (分号间隔)',
  `product_id` int(11) unsigned NOT NULL COMMENT '产品表 ID',
  `product_eng_name` varchar(100) NOT NULL COMMENT '产品英文名称',
  `creator` varchar(50) DEFAULT NULL COMMENT '创建者',
  `modifier` varchar(50) DEFAULT NULL COMMENT '更新者',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_eng_name` (`eng_name`),
  KEY `idx_name` (`name`),
  KEY `idx_product_id` (`product_id`),
  KEY `idx_project_id` (`project_id`),
  KEY `idx_app_id` (`app_id`),
  KEY `idx_status` (`status`),
  KEY `idx_rsm_status` (`rsm_status`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '应用表';