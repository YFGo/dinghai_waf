package normalhttp

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type RepoNormalHttp interface {
	SaveNormalHttp2DB() error
}

type UsercaseNormalHttp struct {
	repo RepoNormalHttp
	log  *log.Helper
}

func NewUsercaseNormalHttp(repo RepoNormalHttp, logger log.Logger) *UsercaseNormalHttp {
	return &UsercaseNormalHttp{repo: repo, log: log.NewHelper(logger)}
}

// SaveNormalHttp 保存正常http请求
func (n *UsercaseNormalHttp) SaveNormalHttp(ctx context.Context) {
	if err := n.repo.SaveNormalHttp2DB(); err != nil {
		n.log.WithContext(ctx).Errorf("save normal http error: %v", err)
	}
	return
}
