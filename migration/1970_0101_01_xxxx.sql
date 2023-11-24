CREATE TABLE `sample` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'id',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `bigint_type` bigint NOT NULL DEFAULT 0 COMMENT 'bigint 类型的数据,一般是其他表的主键,在代码中,使用int64',
  `bool_type` boolean NOT NULL DEFAULT FALSE COMMENT '布尔型',
  `varchar_type` varchar(16) NOT NULL DEFAULT '' COMMENT '字符串类型',
  `json_type` varchar(400) NOT NULL DEFAULT '' COMMENT 'json序列化成字符串,再存储到数据库,尽量不要使用数据库的json type',
  `inventory` bigint unsigned NOT NULL DEFAULT 0 COMMENT '库存 可以使用数据库的 unsigned 特性 (但是这仅限于mysql pg无此特性)',
  `some_action_time` bigint NOT NULL DEFAULT 0 COMMENT '各种操作的时间,可以存储unix时间戳,秒级即可',
  INDEX idx_sample(`user_id`,`is_default`),
  UNIQUE uniq_sample(`user_id`,`is_default`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户地址信息表';
