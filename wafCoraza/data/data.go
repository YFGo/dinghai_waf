package data

import (
	"github.com/IBM/sarama"
	"github.com/robfig/cron/v3"
	"log/slog"
	"os"
)

type Data struct {
	kafkaProducer sarama.SyncProducer
	timeTask      *cron.Cron
	file          *os.File
}

func NewData() (*Data, func()) {
	kafkaProducer := newKafkaProducer()
	timeTask := newTimeTask()
	cleanup := func() {
		if kafkaProducer != nil {
			slog.Info("close kafka producer")
			kafkaProducer.Close()
		}
		if timeTask != nil {
			slog.Info("close time task")
			timeTask.Stop()
		}
	}
	return &Data{
		kafkaProducer: kafkaProducer,
		timeTask:      timeTask,
	}, cleanup
}

func newKafkaProducer() sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll        // 发送数据之后需要 leader 和 follow 都确认
	config.Producer.Partitioner = sarama.NewHashPartitioner // 根据hasH值选择分区
	config.Producer.Return.Successes = true                 // 成功交付的消息将在success channel 返回
	//链接kafka
	kafkaProducer, err := sarama.NewSyncProducer([]string{"string"}, config)
	if err != nil {
		slog.Error("kafka client error: ", err)
		return nil
	}
	return kafkaProducer
}

func newTimeTask() *cron.Cron {
	c := cron.New(cron.WithSeconds())
	return c
}
