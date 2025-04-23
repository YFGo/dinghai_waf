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
			name: "normalhttp",
			spec: "0 0/30 * * * ?",
			log:  log.NewHelper(log.With(logger, "normal_http", "service")),
		},
		uc: uc,
	}
}

func (s *NormalHttpService) GetFunc() func() {
	return s.uc.SaveNormalHttp(context.Background())
}
