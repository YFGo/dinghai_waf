package iface

import (
	"gorm.io/gorm"
)

type WhereOption func(*gorm.DB)
type WhereOptionWithReturn func(*gorm.DB) *gorm.DB
