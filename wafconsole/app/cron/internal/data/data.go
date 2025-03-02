package data

import (
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/gorm"
	"wafconsole/app/cron/internal/data/db"

	v1 "wafconsole/api/wafTop/v1"
	"wafconsole/app/cron/internal/conf"
	"wafconsole/app/cron/internal/data/discovery"
	"wafconsole/app/cron/internal/data/mq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, discovery.NewDiscovery, discovery.NewNormalHttpRpc, NewNormalHttpRepo)

// Data .
type Data struct {
	log               *log.Helper
	kafkaConsumeGroup sarama.ConsumerGroup
	kafkaProducer     sarama.SyncProducer
	normalHttpRpc     v1.NormalHttpClient
	clickhouseDB      *gorm.DB
	rdb               *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, normalHttpRpc v1.NormalHttpClient) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	kafkaConsumeGroup, err := mq.NewKafkaConsumerGroup(c, logger)
	if err != nil {
		l.Errorf("创建Kafka消费者组失败: %v", err)
		return nil, nil, err
	}
	kafkaProducer, err := mq.NewKafkaProducer(c, logger)
	if err != nil {
		l.Errorf("创建Kafka生产者失败: %v", err)
		return nil, nil, err
	}
	clickhouseDB, err := db.NewClickHouse(c.Clickhouse, l)
	if err != nil {
		l.Errorf("创建ClickHouse连接失败: %v", err)
		return nil, nil, err
	}
	redisDB, err := db.NewRedis(c, l)
	if err != nil {
		l.Errorf("创建Redis连接失败: %v", err)
		return nil, nil, err
	}
	cleanup := func() {
		l.Info("closing the data resources")
		if err := kafkaConsumeGroup.Close(); err != nil {
			l.Error(err)
		}
		if err := kafkaConsumeGroup.Close(); err != nil {
			l.Error(err)
		}
		if err := redisDB.Close(); err != nil {
			l.Error(err)
		}
		if clickhouseDB != nil {
			if ck, err := clickhouseDB.DB(); err == nil && ck != nil {
				ck.Close()
			}
			l.Error(err)
		}
	}
	return &Data{
		log:               l,
		kafkaConsumeGroup: kafkaConsumeGroup,
		kafkaProducer:     kafkaProducer,
		normalHttpRpc:     normalHttpRpc,
		clickhouseDB:      clickhouseDB,
		rdb:               redisDB,
	}, cleanup, nil
}
