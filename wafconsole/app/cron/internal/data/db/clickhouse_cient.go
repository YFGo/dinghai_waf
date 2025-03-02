package db

import (
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"

	"wafconsole/app/cron/internal/conf"
)

func NewClickHouse(cfg *conf.Data_ClickHouse, logger *log.Helper) (*gorm.DB, error) {
	clickhouseDB, err := gorm.Open(clickhouse.Open(cfg.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return clickhouseDB, nil
}
