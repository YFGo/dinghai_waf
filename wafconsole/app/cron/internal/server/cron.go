package server

import (
	"context"

	"wafconsole/app/cron/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"
)

type CronServer struct {
	sche *cron.Cron
	log  *log.Helper
}

func NewCronServer(jobSrc *service.CronService, logger log.Logger) *CronServer {
	l := log.NewHelper(log.With(logger, "module", "cron"))
	ser := &CronServer{
		sche: cron.New(),
	}
	if jobSrc != nil {
		for _, job := range jobSrc.GetJobList() {
			cronId, err := ser.sche.AddFunc(job.GetSpec(), job.GetFunc())
			if err != nil {
				l.Errorf("add cron job error: %v", err)
			} else {
				l.Infof("add cron job success, id: %d, name: %s", cronId, job.GetName())
			}
		}
	}

	return ser
}

func (s *CronServer) Start(c context.Context) error {
	s.sche.Start()
	return nil
}

func (s *CronServer) Stop(c context.Context) error {
	s.sche.Stop()
	return nil
}
