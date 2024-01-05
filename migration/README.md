# migration

项目使用的sql

默认要带的字段

`id` bigint unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'id',
`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

时间字段
1. 某个动作的发生时间,比如pay_time,尽量使用bigint型存储时间戳  `some_action_time` bigint NOT NULL DEFAULT 0 COMMENT '各种操作的时间,可以存储unix时间戳,秒级即可'
2. 生日 ,使用 datetime "2000-xx-xx" 生日理论上只需要 月+日 ,使用2000年是因为2000年是闰年

库存字段
如果使用数据库存储库存,可以考虑使用 int unsigned 无符号整数
