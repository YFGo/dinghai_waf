package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"wafconsole/app/cron/internal/biz/normalhttp"
)

type normalHttpRepo struct {
	data *Data
	log  *log.Helper
}

func NewNormalHttpRepo(data *Data, logger log.Logger) normalhttp.RepoNormalHttp {
	return &normalHttpRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (n normalHttpRepo) SaveNormalHttp2DB() error {

	return nil
}
