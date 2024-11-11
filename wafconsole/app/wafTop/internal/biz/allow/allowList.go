package allow

import (
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"
)

type ListAllowRepo interface {
	iface.BaseRepo[model.AllowList]
}

type ListAllowUsecase struct {
	repo ListAllowRepo
}

func NewListAllowUsecase(repo ListAllowRepo) *ListAllowUsecase {
	return &ListAllowUsecase{repo: repo}
}
