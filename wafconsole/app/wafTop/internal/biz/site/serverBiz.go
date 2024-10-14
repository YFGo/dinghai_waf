package siteBiz

import (
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"
)

// ServerRepo 服务器上层实现
type ServerRepo interface {
	iface.BaseRepo[model.ServerWaf]
}

type ServerUsecase struct {
	repo ServerRepo
}

func NewServerUsecase(repo ServerRepo) *ServerUsecase {
	return &ServerUsecase{repo: repo}
}
