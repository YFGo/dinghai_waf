package mq

import (
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"

	"wafconsole/app/cron/internal/conf"
)

// NewKafkaProducer 初始化kafka生产者
func NewKafkaProducer(cfg *conf.Data, logger log.Logger) (sarama.SyncProducer, error) {
	// 1. 配置生产者
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.RequiredAcks = sarama.WaitForAll
	producerConfig.Producer.Retry.Max = 5
	producerConfig.Producer.Return.Successes = true

	// 2. 初始化生产者
	producer, err := sarama.NewSyncProducer(cfg.Kafka.Addrs, producerConfig)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

// NewKafkaConsumerGroup 初始化kafka消费者组
func NewKafkaConsumerGroup(cfg *conf.Data, logger log.Logger) (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.ChannelBufferSize = 100
	config.Consumer.Offsets.AutoCommit.Enable = false // 禁用自动提交
	client, err := sarama.NewClient(cfg.Kafka.Addrs, config)
	if err != nil {
		return nil, err
	}
	//根据client创建消费者组
	kafkaConsumerGroup, err := sarama.NewConsumerGroupFromClient(cfg.Kafka.GroupId, client)
	if err != nil {
		return nil, err
	}
	return kafkaConsumerGroup, err
}
