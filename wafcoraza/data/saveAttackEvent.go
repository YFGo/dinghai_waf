package data

import (
	"encoding/json"
	"log/slog"
	"strconv"
	"sync"

	"github.com/IBM/sarama"
	"wafcoraza/biz"
	"wafcoraza/data/model"
)

type saveAttackEventRepo struct {
	data            *Data
	mu              sync.Mutex
	attackEventFile string
}

func NewSaveAttackEventRepo(data *Data) biz.AttackEventRepo {
	return &saveAttackEventRepo{
		data:            data,
		attackEventFile: "waf_log/attack_events.csv",
	}
}

func (s *saveAttackEventRepo) AppendToKafka(attackEvent []model.AttackEvent) {
	var producerMessages []*sarama.ProducerMessage
	for _, event := range attackEvent {
		eventJson, err := json.Marshal(&event)
		if err != nil {
			slog.Error("json marshal error: ", err)
			return
		}
		msg := &sarama.ProducerMessage{
			Topic: model.AttackEventLogTopic,
			Value: sarama.StringEncoder(eventJson),
			Key:   sarama.StringEncoder(strconv.Itoa(event.RuleId)),
		}
		producerMessages = append(producerMessages, msg)
	}
	err := s.data.kafkaProducer.SendMessages(producerMessages)
	if err != nil {
		slog.Info("send messages to kafka: ", err)
		return
	}
}
