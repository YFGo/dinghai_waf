CREATE TABLE IF NOT EXISTS app_waf (
     id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
     created_at DATETIME(3)  NULL,
     updated_at DATETIME(3)  NULL,
     deleted_at DATETIME(3)  NULL,
     name       VARCHAR(255) NOT NULL COMMENT '应用名称',   
     url        VARCHAR(255) NULL COMMENT 'web程序地址',
     server_id  BIGINT       NOT NULL COMMENT 'web程序所在服务器id',
     UNIQUE KEY uni_app_waf_name (name),                    
     UNIQUE KEY uni_app_waf_server_id (server_id),         
     INDEX idx_app_waf_deleted_at (deleted_at)            
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;                   