package data

import (
	"analyse/internal/conf"
	"context"
	"github.com/IBM/sarama"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"log/slog"
)

type Data struct {
	clickhouseDB       *gorm.DB
	kafkaConsumerGroup sarama.ConsumerGroup
	rdb                *redis.Client
}

func NewData(s *conf.Server, conf *conf.Bootstrap) (*Data, func(), error) {
	c := conf.Data
	clickhouseDB, err := newClickhouse(c)
	if err != nil {
		panic("failed to connect clickhouse")
	}
	kafkaConsumerGroup, err := newKafkaConsumer(c)
	if err != nil {
		panic("failed to connect kafka")
	}
	rdb := newRedis(c)
	cleanup := func() {
		if clickhouseDB != nil {
			if db, err := clickhouseDB.DB(); err == nil && db != nil {
				db.Close()
			}
		}
		if kafkaConsumerGroup != nil {
			if err := kafkaConsumerGroup.Close(); err != nil {
				slog.Error("failed to close kafka consumer group", err)
			}
		}
		if rdb != nil {
			if err := rdb.Close(); err != nil {
				slog.Error("failed to close redis", err)
			}
		}
	}
	return &Data{
		clickhouseDB:       clickhouseDB,
		kafkaConsumerGroup: kafkaConsumerGroup,
		rdb:                rdb,
	}, cleanup, nil
}

func newClickhouse(c *conf.Data) (*gorm.DB, error) {
	clickhouseDB, err := gorm.Open(clickhouse.Open(c.ClickHouse.Dsn), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect clickhouse", err)
		return nil, err
	}
	return clickhouseDB, nil
}

func newKafkaConsumer(c *conf.Data) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.ChannelBufferSize = 100
	config.Consumer.Offsets.AutoCommit.Enable = false // 禁用自动提交
	client, err := sarama.NewClient([]string{c.Kafka.Addr}, config)
	if err != nil {
		slog.Error("failed to connect kafka", err)
		return nil, err
	}
	//根据client创建消费者组
	kafkaConsumerGroup, err := sarama.NewConsumerGroupFromClient(c.Kafka.GroupId, client)
	if err != nil {
		slog.Error("failed to create consumer group", err)
		return nil, err
	}
	return kafkaConsumerGroup, err
}

func newRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		DB:       int(c.Redis.Db),
		Password: c.Redis.Password,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		slog.Error("failed to connect redis", err)
		panic(err)
	}
	return rdb
}
