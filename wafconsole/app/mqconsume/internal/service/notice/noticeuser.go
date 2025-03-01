package notice

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	v1 "wafconsole/api/mqconsume/v1"

	"wafconsole/app/mqconsume/internal/biz/notice"
)

type MsgNoticeService struct {
	logger log.Logger

	uc *notice.UserNoticeUsecase
}

func NewNoticeMsgService(logger log.Logger, uc *notice.UserNoticeUsecase) *MsgNoticeService {
	return &MsgNoticeService{
		logger: logger,
		uc:     uc,
	}
}

// SaveUser save user.
func (m *MsgNoticeService) SaveUser(ctx context.Context, _ string, _ broker.Headers,
	msg *v1.MqUserInfo) error {
	return m.uc.Start()
}
