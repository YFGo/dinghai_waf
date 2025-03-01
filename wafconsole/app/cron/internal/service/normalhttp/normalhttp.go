package normalhttp

import (
	"github.com/go-kratos/kratos/v2/log"
	"wafconsole/app/cron/internal/biz/normalhttp"
	"wafconsole/app/cron/internal/service"
)

type ServiceNormalHttp struct {
	service.JobCommonService
	uc *normalhttp.UsercaseNormalHttp
}

func NewServiceNormalHttp(logger log.Logger, uc *normalhttp.UsercaseNormalHttp) *ServiceNormalHttp {
	return &ServiceNormalHttp{
		JobCommonService: service.JobCommonService{
			Name: "normalhttp",
			Spec: "0 */1 * * * *",
			Log:  log.NewHelper(logger),
		},
		uc: uc,
	}
}

func (s *ServiceNormalHttp) GetFunc() func() {
	return func() {
		s.uc.SaveNormalHttp()
	}
}
