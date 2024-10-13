package siteBiz

import (
	"context"
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

// WafAppRepo is a Greater repo.
type WafAppRepo interface {
	iface.BaseRepo[model.AppWaf]
}

// WafAppUsecase is a Greeter usecase.
type WafAppUsecase struct {
	repo WafAppRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo WafAppRepo, logger log.Logger) *WafAppUsecase {
	return &WafAppUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (w *WafAppUsecase) CreateWafBiz(appInfo model.AppWaf) {
	w.repo.Create(context.Background(), appInfo)
}
