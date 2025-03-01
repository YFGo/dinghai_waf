package notice

import "github.com/go-kratos/kratos/v2/log"

type UserNoticeRepo interface {
	StartData() error
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

func (uc *UserNoticeUsecase) Start() {
	err := uc.repo.StartData()
	if err != nil {
		return
	}
	uc.log.Info("UserNoticeUsecase Start")
}
