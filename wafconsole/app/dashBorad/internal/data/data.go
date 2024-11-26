package data

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
	"wafconsole/app/dashBorad/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDataViewRepo)

// Data .
type Data struct {
	db           *gorm.DB
	clickhouseDB *gorm.DB
}

// NewData .
func NewData(s *conf.Server, bootstrap *conf.Bootstrap) (*Data, func(), error) {
	c := bootstrap.Data
	mysqlDB, err := newMysql(c.Mysql)
	if err != nil {
		return nil, nil, err
	}
	clickhouseDB, err := newClickhouse(c)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if mysqlDB != nil {
			if db, err := mysqlDB.DB(); err == nil && db != nil {
				db.Close()
			}
		}
		if clickhouseDB != nil {
			if db, err := clickhouseDB.DB(); err == nil && db != nil {
				db.Close()
			}
		}
	}
	return &Data{
		db:           mysqlDB,
		clickhouseDB: clickhouseDB,
	}, cleanup, nil
}

func newMysql(cfg *conf.Data_Mysql) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Db)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(int(cfg.MaxIdle))
	sqlDB.SetMaxOpenConns(int(cfg.MaxOpen))
	return db, nil
}

func newClickhouse(c *conf.Data) (*gorm.DB, error) {
	clickhouseDB, err := gorm.Open(clickhouse.Open(c.Clickhouse.Dsn), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect clickhouse", err)
		return nil, err
	}
	return clickhouseDB, nil
}
