package migrate

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"time"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/go-redis/redis/v8"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	chMigrate "github.com/golang-migrate/migrate/v4/database/clickhouse"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"wafconsole/utils/redislock"

	_ "github.com/go-sql-driver/mysql"
)

// Config 迁移配置
type Config struct {
	AppName       string
	MySqlDSN      string
	ClickHouseDSN string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	MigrationDir  string
	LockTimeout   time.Duration
}

// DatabaseMigrator 数据库迁移器
type DatabaseMigrator struct {
	mysqlDB      *sql.DB
	clickhouseDB *sql.DB
	redisClient  *redis.Client
	config       *Config
	lockID       string
}

// NewDatabaseMigrator 创建新实例
func NewDatabaseMigrator(cfg *Config) (*DatabaseMigrator, error) {
	// 初始化MySQL连接
	mysqlDB, err := sql.Open("mysql", cfg.MySqlDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	// 验证MySQL连接
	if err = mysqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("MySQL ping failed: %w", err)
	}

	// 初始化ClickHouse连接（修正错误信息）
	clickhouseDB, err := sql.Open("clickhouse", cfg.ClickHouseDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ClickHouse: %w", err) // 修正错误信息
	}

	// 验证ClickHouse连接（修正错误信息）
	if err = clickhouseDB.Ping(); err != nil {
		return nil, fmt.Errorf("ClickHouse ping failed: %w", err) // 修正错误信息
	}

	// 初始化Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	// 验证Redis连接
	if err = rdb.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("redis connection failed: %w", err)
	}

	return &DatabaseMigrator{
		mysqlDB:      mysqlDB,
		clickhouseDB: clickhouseDB,
		redisClient:  rdb,
		config:       cfg,
		lockID:       uuid.New().String(),
	}, nil
}

// Run 执行全量迁移
func (m *DatabaseMigrator) Run(ctx context.Context) error {
	lockKey := "database_migration_lock"
	rdLock := redislock.NewRedisLock(m.redisClient, m.config.LockTimeout)

	if err := rdLock.AcquireLock(ctx, lockKey); err != nil {
		return fmt.Errorf("failed to acquire lock: %w", err)
	}
	defer func() {
		if err := rdLock.ReleaseLock(ctx, lockKey); err != nil {
			logrus.Errorf("Failed to release lock: %v", err)
		}
	}()

	line2MigrateFilesMap, err := DistinguishModules(m.config.MigrationDir)
	if err != nil {
		return err
	}
	fmt.Println(line2MigrateFilesMap)

	if err := m.migrateMySQL(ctx); err != nil {
		return fmt.Errorf("MySQL migration failed: %w", err)
	}

	if err := m.migrateClickHouse(ctx); err != nil {
		return fmt.Errorf("ClickHouse migration failed: %w", err)
	}

	return nil
}

// MySQL迁移（保持不变）
func (m *DatabaseMigrator) migrateMySQL(ctx context.Context) error {
	driver, err := mysqlMigrate.WithInstance(m.mysqlDB, &mysqlMigrate.Config{})
	if err != nil {
		return err
	}
	return m.runMigration(ctx, driver, "mysql", "DingHaiWAF")
}

// ClickHouse迁移（保持不变）
func (m *DatabaseMigrator) migrateClickHouse(ctx context.Context) error {
	driver, err := chMigrate.WithInstance(m.clickhouseDB, &chMigrate.Config{})
	if err != nil {
		return fmt.Errorf("failed to create ClickHouse driver: %w", err)
	}
	return m.runMigration(ctx, driver, "clickhouse", "waf")
}

// 通用迁移执行方法（保持不变）
func (m *DatabaseMigrator) runMigration(ctx context.Context, driver database.Driver,
	dbType string, dbName string) error {
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

// Close 关闭资源（保持不变）
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
