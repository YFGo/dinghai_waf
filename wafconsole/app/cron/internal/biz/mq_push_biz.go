package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	v1 "wafconsole/api/mqconsume/v1"
)

type MqPushRepo interface {
	PushUserInfo(info *v1.MqUserInfo) error
}

// MqPushUseCase .
type MqPushUseCase struct {
	repo MqPushRepo
	log  *log.Helper
}

func NewMqPushUseCase(repo MqPushRepo, logger log.Logger) *MqPushUseCase {
	return &MqPushUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (u *MqPushUseCase) AddUser(info *v1.MqUserInfo) error {
	return u.repo.PushUserInfo(info)
}
