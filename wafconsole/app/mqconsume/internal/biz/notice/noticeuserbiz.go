package notice

import "github.com/go-kratos/kratos/v2/log"

type UserNoticeRepo interface {
	Test() error
}

type UserNoticeUsecase struct {
	repo UserNoticeRepo
	log  *log.Helper
}

func NewUserNoticeUsecase(repo UserNoticeRepo, logger log.Logger) *UserNoticeUsecase {
	return &UserNoticeUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserNoticeUsecase) Start() error {
	err := uc.repo.Test() // 消费正常的http请求
	if err != nil {
		return err
	}
	uc.log.Info("UserNoticeUsecase Start")
	return err
}
