package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"time"
	"wafconsole/app/cron/internal/biz"
)

type UserCountService struct {
	JobCommonService
	uc *biz.UserUseCase
}

func NewUserCountService(logger log.Logger, uc *biz.UserUseCase) *UserCountService {
	return &UserCountService{
		JobCommonService: JobCommonService{
			name: "用户总数统计计划任务",
			spec: "*/1 * * * *",
			log:  log.NewHelper(log.With(logger, "module", "service")),
		},
		uc: uc,
	}
}

func (d *UserCountService) GetFunc() func() {
	return func() {
		fmt.Printf("当前时间 %v \n", time.Now().Format(time.DateTime))
		total, err := d.uc.Count(context.Background())
		if err != nil {
			d.log.Errorf("获取用户总数失败: %v", err)
		} else {
			fmt.Printf("当前用户总数 %d \n", total)
		}
	}
}
