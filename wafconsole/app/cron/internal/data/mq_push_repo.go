package data

import (
	"context"
	"wafconsole/app/cron/internal/conf"

	"wafconsole/api/mq_consume"
	"wafconsole/app/cron/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type mqPushRepo struct {
	data *Data
	c    *conf.Data
	log  *log.Helper
}

// NewMqPushRepo . Create a new mq push repo.
func NewMqPushRepo(data *Data, cfg *conf.Data, logger log.Logger) biz.MqPushRepo {
	return &mqPushRepo{data: data, c: cfg, log: log.NewHelper(logger)}
}

// AddUser . Add user.
func (m *mqPushRepo) PushUserInfo(info *mq_consume.MqUserInfo) error {
	_ = m.data.rabbitmqBroker.Publish(context.Background(), m.c.Rabbitmq.Routing, info)
	return nil
}
