package data

import (
	"context"
	"fmt"
	"log/slog"
	"net/smtp"
	"wafconsole/app/user/internal/data/types"

	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"wafconsole/app/user/internal/conf"
	"wafconsole/app/user/internal/hooks"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewWafUserRepo, NewWafUserCommonRepo)

// Data .
type Data struct {
	db       *gorm.DB
	rdb      *redis.Client
	emailCfg *types.EmailCfg
}

// NewData .
func NewData(s *conf.Server, bootstrap *conf.Bootstrap) (*Data, func(), error) {
	c := bootstrap.Data
	mysqlClient, err := newMysql(c.Mysql)
	if err != nil {
		return nil, nil, err
	}
	redisClient := newRedis(c.Redis)

	emailCfg := newEmailAuth(c.Email)

	cleanup := func() {
		if mysqlClient != nil {
			if db, err := mysqlClient.DB(); err == nil && db != nil {
				db.Close()
			}
		}
		if redisClient != nil {
			redisClient.Close()
		}
	}
	return &Data{
		db:       mysqlClient,
		rdb:      redisClient,
		emailCfg: emailCfg,
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

	hooks.CreateTable(db)

	return db, nil
}

func newRedis(cfg *conf.Data_Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		DB:       int(cfg.Db),
		Password: cfg.Password,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		slog.Error("failed to connect redis", err)
		panic(err)
	}
	return rdb
}

func newEmailAuth(cfg *conf.Data_Email) *types.EmailCfg {
	auth := smtp.PlainAuth("", cfg.SmtpUsername, cfg.SmtpPassword, cfg.SmtpServer)
	return &types.EmailCfg{
		SmtpUserName: cfg.SmtpUsername,
		SmtpPassword: cfg.SmtpPassword,
		SmtpServer:   cfg.SmtpServer,
		SmtpProt:     cfg.SmtpPort,
		Auth:         auth,
	}
}
