CREATE TABLE IF NOT EXISTS waf_server_allow (                                                id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,                                                created_at DATETIME(3) NULL,                                                updated_at DATETIME(3) NULL,                                                deleted_at DATETIME(3) NULL,                                                allow_id   BIGINT      NOT NULL COMMENT '白名单id',                                                server_id  BIGINT      NOT NULL COMMENT 'server_id',                                                INDEX idx_waf_server_allow_deleted_at (deleted_at)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;