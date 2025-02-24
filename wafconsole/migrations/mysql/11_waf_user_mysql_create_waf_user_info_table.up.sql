CREATE TABLE IF NOT EXISTS waf_user_info (                                             id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,                                             created_at  DATETIME(3)  NULL,                                             updated_at  DATETIME(3)  NULL,                                             deleted_at  DATETIME(3)  NULL,                                             email       VARCHAR(50)  NOT NULL COMMENT '用户邮箱',                                             user_name   VARCHAR(20)  NULL,                                             `password`    VARCHAR(20)  NOT NULL,                                             avatar_addr VARCHAR(200) NULL COMMENT '头像地址',                                             phone       VARCHAR(20)  NULL COMMENT '手机号',                                             UNIQUE KEY uni_waf_user_info_email (email),            -- 唯一约束                                             INDEX idx_waf_user_info_deleted_at (deleted_at)        -- 普通索引) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;