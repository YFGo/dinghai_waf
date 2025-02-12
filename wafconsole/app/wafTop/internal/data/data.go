package data

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
	"time"
	"wafconsole/app/wafTop/internal/conf"
	"wafconsole/app/wafTop/internal/hooks"
	"wafconsole/utils/context"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAppWafRepo, NewServerRepo, NewBuildRuleRepo, NewRuleGroupRepo, NewUserRuleRepo, NewWafStrategyRepo, NewAllowListRepo)

// Data .
type Data struct {
	db   *gorm.DB
	etcd *clientv3.Client
}

// NewData .
func NewData(s *conf.Server, bootstrap *conf.Bootstrap) (*Data, func(), error) {
	c := bootstrap.Data
	appCtx := utils.NewAppCtx(context.Background())
	dbMysql, err := newMysql(c.Mysql, appCtx)
	if err != nil {
		return nil, nil, err
	}
	etcd := newETCD(c.Etcd, appCtx)
	cleanup := func() {
		if dbMysql != nil {
			if db, err := dbMysql.DB(); err == nil && db != nil {
				db.Close()
			}
		}
		if etcd != nil {
			etcd.Close()
		}
	}

	return &Data{
		db:   dbMysql,
		etcd: etcd,
	}, cleanup, nil
}

func newMysql(cfg *conf.Data_Mysql, ctx *utils.CustomContext) (*gorm.DB, error) {
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
	// 自动 建表
	hooks.CreateTable(db)
	return db, nil
}

func newETCD(cfg *conf.Data_Etcd, ctx *utils.CustomContext) *clientv3.Client {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{cfg.Host},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		slog.Error("etcd client failed: ", err)
		panic(err)
	}
	// 设置超时时间
	hooks.InitEtcd(etcdClient, ctx) // 初始化键值对
	return etcdClient
}
