package migrate

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/golang-migrate/migrate/v4"
	"path/filepath"
	"time"
	"wafconsole/utils/redislock"

	"github.com/golang-migrate/migrate/v4/database"
	chMigrate "github.com/golang-migrate/migrate/v4/database/clickhouse"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Config 迁移配置
type Config struct {
	MySQLDSN      string
	ClickHouseDSN string
	RedisAddr     string        // Redis服务器地址
	RedisPassword string        // Redis密码
	RedisDB       int           // Redis数据库编号
	MigrationDir  string        // 迁移文件目录
	LockTimeout   time.Duration // 分布式锁超时时间
}

// DatabaseMigrator 数据库迁移器
type DatabaseMigrator struct {
	mysqlDB      *sql.DB
	clickhouseDB *sql.DB
	redisClient  *redis.Client
	config       *Config
	lockID       string // 唯一锁标识
}

// NewDatabaseMigrator 创建新实例
func NewDatabaseMigrator(cfg *Config) (*DatabaseMigrator, error) {
	// 初始化MySQL连接
	mysqlDB, err := sql.Open("mysql", cfg.MySQLDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	// 验证MySQL连接
	if err := mysqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("MySQL ping failed: %w", err)
	}

	// 初始化ClickHouse连接
	chDB, err := sql.Open("clickhouse", cfg.ClickHouseDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ClickHouse: %w", err)
	}

	// 验证ClickHouse连接
	if err := chDB.Ping(); err != nil {
		return nil, fmt.Errorf("ClickHouse ping failed: %w", err)
	}

	// 初始化Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	// 验证Redis连接
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("Redis connection failed: %w", err)
	}

	return &DatabaseMigrator{
		mysqlDB:      mysqlDB,
		clickhouseDB: chDB,
		redisClient:  rdb,
		config:       cfg,
		lockID:       uuid.New().String(),
	}, nil
}

// Run 执行全量迁移
func (m *DatabaseMigrator) Run(ctx context.Context) error {
	lockKey := "database_migration_lock"
	rdLock := redislock.NewRedisLock(m.redisClient, m.config.LockTimeout)
	// 获取分布式锁

	if err := rdLock.AcquireLock(ctx, lockKey); err != nil {
		return fmt.Errorf("failed to acquire lock: %w", err)
	}
	defer func() {
		if err := rdLock.ReleaseLock(ctx, lockKey); err != nil {
			logrus.Errorf("Failed to release lock: %v", err)
		}
	}()

	// 执行数据库迁移
	if err := m.migrateMySQL(ctx); err != nil {
		return fmt.Errorf("MySQL migration failed: %w", err)
	}

	if err := m.migrateClickHouse(ctx); err != nil {
		return fmt.Errorf("ClickHouse migration failed: %w", err)
	}

	return nil
}

// MySQL迁移
func (m *DatabaseMigrator) migrateMySQL(ctx context.Context) error {
	driver, err := mysqlMigrate.WithInstance(m.mysqlDB, &mysqlMigrate.Config{})
	if err != nil {
		return err
	}

	return m.runMigration(ctx, driver, "mysql")
}

// ClickHouse迁移
func (m *DatabaseMigrator) migrateClickHouse(ctx context.Context) error {
	driver, err := chMigrate.WithInstance(m.clickhouseDB, &chMigrate.Config{})
	if err != nil {
		return fmt.Errorf("failed to create ClickHouse driver: %w", err)
	}

	return m.runMigration(ctx, driver, "clickhouse")
}

// 通用迁移执行方法
func (m *DatabaseMigrator) runMigration(ctx context.Context, driver database.Driver, dbType string) error {
	migratePath := filepath.Join(m.config.MigrationDir, dbType)
	sourceURL := fmt.Sprintf("file://%s", migratePath)

	migrator, err := migrate.NewWithDatabaseInstance(sourceURL, dbType, driver)
	if err != nil {
		return fmt.Errorf("failed to initialize migrator: %w", err)
	}
	defer migrator.Close()

	if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration failed: %w", err)
	}

	logrus.Infof("%s migration completed successfully", dbType)
	return nil
}

// Close 关闭资源
func (m *DatabaseMigrator) Close() error {
	var errs []error

	if err := m.mysqlDB.Close(); err != nil {
		errs = append(errs, fmt.Errorf("MySQL close error: %w", err))
	}

	if err := m.clickhouseDB.Close(); err != nil {
		errs = append(errs, fmt.Errorf("ClickHouse close error: %w", err))
	}

	if err := m.redisClient.Close(); err != nil {
		errs = append(errs, fmt.Errorf("Redis close error: %w", err))
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors occurred during shutdown: %v", errs)
	}
	return nil
}
