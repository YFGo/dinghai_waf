package data

import (
	"log/slog"
	"time"

	"github.com/IBM/sarama"
	"github.com/robfig/cron/v3"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gopkg.in/ini.v1"
)

type Data struct {
	kafkaProducer sarama.SyncProducer
	timeTask      *cron.Cron
	etcdClient    *clientv3.Client
}

func NewConfFile() *ini.File {
	file, err := ini.Load("wafCoraza/conf/conf_dev.ini")
	if err != nil {
		slog.Error("load config file error: ", err)
		panic(err)
	}
	return file
}

func NewData(c *ini.File) (*Data, func()) {
	kafkaProducer := newKafkaProducer(c.Section("kafka").Key("address").String())
	timeTask := newTimeTask()
	etcdClient := newETCD(c.Section("etcd").Key("address").String())
	cleanup := func() {
		if kafkaProducer != nil {
			slog.Info("close kafka producer")

			kafkaProducer.Close()
		}
		if timeTask != nil {
			slog.Info("close time task")
			timeTask.Stop()
		}
		if etcdClient != nil {
			slog.Info("close etcd client")
			etcdClient.Close()
		}
	}
	return &Data{
		//kafkaProducer: kafkaProducer,
		timeTask:   timeTask,
		etcdClient: etcdClient,
	}, cleanup
}

func newKafkaProducer(address string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll        // 发送数据之后需要 leader 和 follow 都确认
	config.Producer.Partitioner = sarama.NewHashPartitioner // 根据hasH值选择分区
	config.Producer.Return.Successes = true                 // 成功交付的消息将在success channel 返回
	//链接kafka
	kafkaProducer, err := sarama.NewSyncProducer([]string{address}, config)
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

func newETCD(address string) *clientv3.Client {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{address},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		slog.Error("etcd client failed: ", err)
		panic(err)
	}
	return etcdClient
}
