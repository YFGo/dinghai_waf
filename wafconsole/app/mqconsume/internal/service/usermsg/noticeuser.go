package usermsg

import (
	"github.com/go-kratos/kratos/v2/log"

	"wafconsole/app/mqconsume/internal/biz/notice"
)

type NoticeMsgService struct {
	logger log.Logger

	uc *notice.UserNoticeUsecase
}

func NewNoticeMsgService(logger log.Logger, uc *notice.UserNoticeUsecase) *NoticeMsgService {
	return &NoticeMsgService{
		logger: logger,
		uc:     uc,
	}
}

func (s *NoticeMsgService) Start() {
	s.uc.Start()
}
