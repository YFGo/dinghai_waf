package iface

import (
	"context"
	"gorm.io/gorm"
	"wafconsole/app/user/internal/data/model"
)

type WhereOption func(*gorm.DB)
type WhereOptionWithReturn func(*gorm.DB) *gorm.DB

type Domain interface {
	model.UserInfo
}

type BaseRepo[T Domain] interface {
	Get(context.Context, int64) (T, error)
	GetByNameAndID(context.Context, string, int64) (T, error)
	Create(context.Context, T) (int64, error)
	Update(context.Context, int64, T) error
	Delete(context.Context, []int64) (int64, error)
	Count(context.Context, ...WhereOptionWithReturn) (int64, error)
	ListByWhere(ctx context.Context, limit, offset int64, opts ...WhereOptionWithReturn) ([]T, error)
}
