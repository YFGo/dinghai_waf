CREATE TABLE IF NOT EXISTS waf_strategy (                                            id          BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,                                            created_at  DATETIME(3)       NULL,                                            updated_at  DATETIME(3)       NULL,                                            deleted_at  DATETIME(3)       NULL,                                            name        VARCHAR(255)      NOT NULL COMMENT '策略名称',                                            description VARCHAR(255)      NULL COMMENT '策略描述',                                            kind        TINYINT           NULL COMMENT '策略类别-预留字段',                                            status      TINYINT DEFAULT 1 NOT NULL COMMENT '2:disable 1:enable',                                            alert_level TINYINT           NULL COMMENT '预留字段',                                            action      TINYINT DEFAULT 1 NOT NULL COMMENT '2:拦截 1:记录',                                            next_action TINYINT DEFAULT 1 NOT NULL COMMENT '2:拦截 1:记录',                                            INDEX idx_waf_strategy_deleted_at (deleted_at)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;