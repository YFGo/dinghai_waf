package redislock

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"time"
)

type RedisLock struct {
	redisClient *redis.Client
	lockID      string
	lockTimeout time.Duration
}

func NewRedisLock(redisClient *redis.Client, lockTimeout time.Duration) *RedisLock {
	return &RedisLock{
		redisClient: redisClient,
		lockID:      uuid.New().String(),
		lockTimeout: lockTimeout,
	}
}

// AcquireLock Redis分布式锁实现
func (m *RedisLock) AcquireLock(ctx context.Context, key string) error {
	// 尝试获取锁（SET NX PX）
	result, err := m.redisClient.SetNX(ctx, key, m.lockID, m.lockTimeout).Result()
	if err != nil {
		return fmt.Errorf("redis operation failed: %w", err)
	}
	if !result {
		return errors.New("lock already acquired by another client")
	}
	return nil
}

func (m *RedisLock) ReleaseLock(ctx context.Context, key string) error {
	// 使用Lua脚本保证原子性删除
	script := `
	if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("del", KEYS[1])
	else
		return 0
	end
	`

	_, err := m.redisClient.Eval(ctx, script, []string{key}, m.lockID).Result()
	if err != nil && err != redis.Nil {
		return fmt.Errorf("failed to release lock: %w", err)
	}
	return nil
}
