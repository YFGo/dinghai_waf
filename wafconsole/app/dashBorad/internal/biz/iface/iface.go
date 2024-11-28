package iface

import (
	"context"
	"gorm.io/gorm"
	"wafconsole/app/dashBorad/internal/data/model"
)

type WhereOption func(*gorm.DB)
type WhereOptionWithReturn func(*gorm.DB) *gorm.DB

type Domain interface {
	model.SecLog
}

type BaseRepo[T Domain] interface {
	ListByWhere(ctx context.Context, limit, offset int64, opts ...WhereOptionWithReturn) ([]T, error)
	Count(context.Context, ...WhereOptionWithReturn) (int64, error)
	GetSecLog(ctx context.Context, logId string) (T, error)
}
