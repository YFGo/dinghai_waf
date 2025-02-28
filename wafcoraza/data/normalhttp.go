package data

import (
	"context"

	"github.com/IBM/sarama"

	"wafcoraza/biz"
	"wafcoraza/data/types"
)

type normalHttpRepo struct {
	data *Data
}

func NewNormalHttpRepo(data *Data) biz.NormalHttpRepo {
	return &normalHttpRepo{
		data: data,
	}
}

func (n normalHttpRepo) SaveNormalHttpToKafka(ctx context.Context,
	normalHttpInfoJson, id string) error {
	msg := &sarama.ProducerMessage{
		Topic: types.NormalHttpTopic,
		Value: sarama.StringEncoder(normalHttpInfoJson),
		Key:   sarama.StringEncoder(id),
	}
	_, _, err := n.data.kafkaProducer.SendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
