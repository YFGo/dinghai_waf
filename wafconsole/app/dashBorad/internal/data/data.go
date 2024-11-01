package data

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wafconsole/app/dashBorad/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(s *conf.Server, bootstrap *conf.Bootstrap) (*Data, func(), error) {
	c := bootstrap.Data
	mysql, err := newMysql(c.Mysql)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if mysql != nil {
			if db, err := mysql.DB(); err == nil && db != nil {
				db.Close()
			}
		}
	}
	return &Data{}, cleanup, nil
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
