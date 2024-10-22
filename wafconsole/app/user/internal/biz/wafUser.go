package biz

import (
	"wafconsole/app/user/internal/biz/iface"
	"wafconsole/app/user/internal/data/model"
)

type WafUserRepo interface {
	iface.BaseRepo[model.UserInfo]
}

type WafUserUsecase struct {
	repo WafUserRepo
}

func NewWafUserUsecase(repo WafUserRepo) *WafUserUsecase {
	return &WafUserUsecase{repo: repo}
}
