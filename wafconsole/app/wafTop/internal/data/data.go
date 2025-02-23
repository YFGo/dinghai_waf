package data

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"wafconsole/app/wafTop/internal/conf"
	"wafconsole/app/wafTop/internal/hooks"
	"wafconsole/utils/migrate"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAppWafRepo, NewServerRepo, NewBuildRuleRepo, NewRuleGroupRepo, NewUserRuleRepo, NewWafStrategyRepo, NewAllowListRepo)

// Data .
type Data struct {
	db   *gorm.DB
	etcd *clientv3.Client
	log  *log.Helper
}

// NewData .
func NewData(s *conf.Server, bootstrap *conf.Bootstrap, logger log.Logger) (*Data, func(),
	error) {
	c := bootstrap.Data
	logKratos := log.NewHelper(log.With(logger, "module", "waf_top/data"))
	dbMysql, err := newMysql(c.Mysql)
	if err != nil {
		return nil, nil, err
	}
	etcd := newETCD(c.Etcd)

	// 执行全量迁移
	newMigrate(bootstrap)

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
		log:  logKratos,
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
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(int(cfg.MaxIdle))
	sqlDB.SetMaxOpenConns(int(cfg.MaxOpen))
	return db, nil
}

func newETCD(cfg *conf.Data_Etcd) *clientv3.Client {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{cfg.Host},
		DialTimeout: 2 * time.Second,
	})
	if err != nil {
		slog.Error("etcd client failed: ", err)
		panic(err)
	}
	// 设置超时时间
	hooks.InitEtcd(etcdClient, context.Background()) // 初始化键值对
	return etcdClient
}

func newMigrate(bootstrap *conf.Bootstrap) {
	cfg := bootstrap.Data
	// 从 conf.Data 中提取 MySQL 配置
	mysqlConfig := cfg.Mysql
	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Db)

	// 从 conf.Data 中提取 Redis 配置
	redisConfig := cfg.Redis
	redisAddr := redisConfig.Addr
	redisPassword := ""
	if len(redisConfig.Password) > 0 {
		redisPassword = redisConfig.Password
	}

	// 从 conf.Data 中提取 ClickHouse 配置
	clickHouseDSN := cfg.ClickHouse.Dsn

	// 构建 migrate.Config
	cfgMigrate := &migrate.Config{
		AppName:       bootstrap.Server.AppName,
		MySqlDSN:      mysqlDSN,
		ClickHouseDSN: clickHouseDSN,
		RedisAddr:     redisAddr,
		RedisPassword: redisPassword,
		RedisDB:       0,
		MigrationDir:  "wafconsole/migrate.txt",
		LockTimeout:   30 * time.Second,
	}
	migrator, err := migrate.NewDatabaseMigrator(cfgMigrate)
	if err != nil {
		slog.Error("migrate failed: ", err)
		panic(err)
	}

	defer func() {
		if err = migrator.Close(); err != nil {
			slog.Error("migrate close failed: ", err)
		}
	}()
	ctx := context.Background()
	if err = migrator.Run(ctx); err != nil {
		slog.Error("migrate run failed: ", err)
		panic(err)
	}
}
