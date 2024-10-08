package data

import (
	"analyse/internal/conf"
	"github.com/IBM/sarama"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"log/slog"
)

type Data struct {
	clickhouseDB       *gorm.DB
	kafkaConsumerGroup sarama.ConsumerGroup
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
	}
	return &Data{
		clickhouseDB:       clickhouseDB,
		kafkaConsumerGroup: kafkaConsumerGroup,
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
