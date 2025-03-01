package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"wafconsole/app/mqconsume/internal/biz/notice"
)

type noticeUserRepo struct {
	data *Data
	log  *log.Helper
}

func NewNoticeUserRepo(data *Data, logger log.Logger) notice.UserNoticeRepo {
	return &noticeUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (n noticeUserRepo) StartData() error {
	//TODO implement me
	panic("implement me")
}
