package iface

import (
	"analyse/internal/data/model"
	"gorm.io/gorm"
)

type WhereOption func(*gorm.DB)

type Domain interface {
	model.SecLog
}

type BaseRepo[T Domain] interface {
	Consumer() func()
	Save(secLogList []T)
}
