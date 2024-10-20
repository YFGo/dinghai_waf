package siteBiz

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log/slog"
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

// WafAppRepo is a Greater repo.
type WafAppRepo interface {
	iface.BaseRepo[model.AppWaf]
	GetAppWafByServerId(ctx context.Context, serverId int64) (appInfo model.AppWaf, err error)
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

func (w *WafAppUsecase) checkStrategyIsExist(ctx context.Context, id int64, name string) bool {
	_, err := w.repo.GetByNameAndID(ctx, name, id)
	if errors.Is(err, gorm.ErrRecordNotFound) { // app不存在 , 可以插入
		return true
	}
	return false
}

// CreateWafApp 创建应用程序
func (w *WafAppUsecase) CreateWafApp(ctx context.Context, appInfo model.AppWaf) error {
	if !w.checkStrategyIsExist(ctx, 0, appInfo.Name) {
		return errors.New("app name is exist")
	}
	_, err := w.repo.Create(ctx, appInfo)
	if err != nil {
		slog.ErrorContext(ctx, "CreateWafBiz err : %v", err)
		return err
	}
	return nil
}

// UpdateWafApp 更新应用程序
func (w *WafAppUsecase) UpdateWafApp(ctx context.Context, id int64, appInfo model.AppWaf) error {
	if !w.checkStrategyIsExist(ctx, id, appInfo.Name) {
		return errors.New("app name is exist")
	}
	err := w.repo.Update(ctx, id, appInfo)
	if err != nil {
		slog.ErrorContext(ctx, "UpdateWafApp err : %v", err)
		return err
	}
	return nil
}
