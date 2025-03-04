package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"wafconsole/app/cron/internal/biz/normalhttp"
)

type NormalHttpService struct {
	JobCommonService
	uc *normalhttp.UsercaseNormalHttp
}

func NewServiceNormalHttp(logger log.Logger, uc *normalhttp.UsercaseNormalHttp) *NormalHttpService {
	return &NormalHttpService{
		JobCommonService: JobCommonService{
			Name: "normalhttp",
			Spec: "0 */10 * * * ?",
			Log:  log.NewHelper(logger),
		},
		uc: uc,
	}
}

func (s *NormalHttpService) GetFunc() func() {
	return s.uc.SaveNormalHttp(context.Background())
}
