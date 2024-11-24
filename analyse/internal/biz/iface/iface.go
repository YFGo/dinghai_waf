package iface

import (
	"analyse/internal/data/model"
	"github.com/IBM/sarama"
	"gorm.io/gorm"
)

type WhereOption func(*gorm.DB)

type Domain interface {
	model.SecLog
}

type BaseRepo[T Domain] interface {
	Consumer() func()
	Save(secLogList []T, session sarama.ConsumerGroupSession)
}
