-- 创建数据库 2023-01-20
CREATE DATABASE `garden_be` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 用户
-- 基础信息表 2023-01-20
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_name` varchar(255) NOT NULL COMMENT '用户名，全局唯一',
  `password` char(128) NOT NULL COMMENT '密码 SHA3-512 加密',
  `nick_name` varchar(255) NOT NULL COMMENT '昵称',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_name`(`user_name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '用户基本信息';

-- 日志表 2023-01-20
CREATE TABLE `user_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) NOT NULL COMMENT 'user_id',
  `type` varchar(255) NOT NULL COMMENT '日志类型：info, err, warning, error',
  `info` varchar(255) NOT NULL COMMENT '日志内容',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id`(`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '用户日志';

-- 文章
-- 基础信息表 2023-01-21
CREATE TABLE `article` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `uuid` char(36) NOT NULL COMMENT '文章对外暴露唯一 id',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '创建用户 id',
  `title` varchar(255) unsigned NOT NULL COMMENT '文章标题',
  `digest` varchar(1024) NOT NULL COMMENT '文章摘要',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_uuid`(`uuid`),
  KEY `idx_user_id`(`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '文章基本信息';

-- 文章内容表 2023-01-21
CREATE TABLE `article_content` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `article_id` bigint(20) unsigned NOT NULL COMMENT '文章 id',
  `content` char(128) NOT NULL COMMENT '文章内容',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_article_id`(`article_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '文章基本信息';

-- 文章标签表 2023-01-21
CREATE TABLE `article_content` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `article_id` bigint(20) unsigned NOT NULL COMMENT '文章 id',
  `label` varchar(255) NOT NULL COMMENT '文章标签',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_article_id`(`article_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '文章基本信息';