package normalhttp

import (
	"context"
	"go.opentelemetry.io/otel"

	"github.com/go-kratos/kratos/v2/log"
)

type RepoNormalHttp interface {
	SaveNormalHttp2DB(ctx context.Context) error
}

type UsercaseNormalHttp struct {
	repo RepoNormalHttp
	log  *log.Helper
}

func NewUsercaseNormalHttp(repo RepoNormalHttp, logger log.Logger) *UsercaseNormalHttp {
	return &UsercaseNormalHttp{repo: repo, log: log.NewHelper(logger)}
}

// SaveNormalHttp 保存正常http请求
func (n *UsercaseNormalHttp) SaveNormalHttp(ctx context.Context) func() {
	return func() {
		// 创建一个新的span用于跟踪此任务
		ctx, span := otel.Tracer("normalhttp").Start(ctx, "save-normal-http")
		defer span.End()

		if err := n.repo.SaveNormalHttp2DB(ctx); err != nil {
			// 使用带有trace上下文的日志记录错误
			n.log.WithContext(ctx).Errorf("save normal http err:%v", err)
			return
		}
		n.log.WithContext(ctx).Infof("save normal http success")
	}
}
