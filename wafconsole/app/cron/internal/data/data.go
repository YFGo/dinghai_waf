package data

import (
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	v1 "wafconsole/api/wafTop/v1"
	"wafconsole/app/cron/internal/conf"
	"wafconsole/app/cron/internal/data/discovery"
	"wafconsole/app/cron/internal/data/mq"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, discovery.NewDiscovery, discovery.NewSiteServerRpc, NewNormalHttpRepo)

// Data .
type Data struct {
	log               *log.Helper
	kafkaConsumeGroup sarama.ConsumerGroup
	kafkaProducer     sarama.SyncProducer
	siteServerRpc     v1.ServerClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, siteServerRpc v1.ServerClient) (*Data, func(), error) {
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
	cleanup := func() {
		l.Info("closing the data resources")
		if err := kafkaConsumeGroup.Close(); err != nil {
			l.Error(err)
		}
		if err := kafkaConsumeGroup.Close(); err != nil {
			l.Error(err)
		}
	}
	return &Data{
		log:               l,
		kafkaConsumeGroup: kafkaConsumeGroup,
		kafkaProducer:     kafkaProducer,
		siteServerRpc:     siteServerRpc,
	}, cleanup, nil
}
