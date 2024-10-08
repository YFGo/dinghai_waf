package iface

import (
	"analyse/internal/data/model"
	"github.com/IBM/sarama"
	"gorm.io/gorm"
	"sync"
)

type WhereOption func(*gorm.DB)

type Domain interface {
	model.SecLog
}

type BaseRepo[T Domain] interface {
	Consumer() func()
	Save(secLog *sync.Map, session sarama.ConsumerGroupSession)
}
