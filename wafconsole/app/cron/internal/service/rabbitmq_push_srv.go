package service

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"math/rand"
	"wafconsole/api/mq_consume"
	"wafconsole/app/cron/internal/biz"
)

type RabbitmqPushService struct {
	JobCommonService
	mqCase *biz.MqPushUseCase
}

func NewRabbitmqPushService(logger log.Logger, mqCase *biz.MqPushUseCase) *RabbitmqPushService {
	return &RabbitmqPushService{
		JobCommonService: JobCommonService{
			name: "rabbitmq推送测试计划任务",
			spec: "*/1 * * * *",
			log:  log.NewHelper(log.With(logger, "module", "service")),
		},
		mqCase: mqCase,
	}
}

func (d *RabbitmqPushService) GetFunc() func() {
	return func() {
		// 随机1-10之间的数
		id := rand.Intn(10) + 1
		err := d.mqCase.AddUser(&mq_consume.MqUserInfo{
			Account:  fmt.Sprintf("test%d", id),
			Password: "123456",
			Name:     "测试",
			Avatar:   "",
		})
		if err != nil {
			d.log.Errorf("rabbitmq推送测试失败: %v", err)
			return
		} else {
			d.log.Infof("rabbitmq推送测试成功")
		}
	}
}
