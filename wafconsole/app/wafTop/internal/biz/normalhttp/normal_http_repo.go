package normalhttp

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"
)

type RepoNormalHttp interface {
	iface.BaseRepo[model.NormalHttpModel]
}

type UsecaseNormalHttp struct {
	repo RepoNormalHttp
	log  *log.Helper
}

func NewUsecaseNormalHttp(repo RepoNormalHttp, logger log.Logger) *UsecaseNormalHttp {
	return &UsecaseNormalHttp{repo: repo, log: log.NewHelper(logger)}
}

func (u *UsecaseNormalHttp) SaveNormalHttpInfo(ctx context.Context,
	normalHttpInfo model.NormalHttpModel) error {
	u.repo.Create(ctx, normalHttpInfo)
	return nil
}
