package service

import (
	"github.com/go-kratos/kratos/v2/log"
)

type JobInterface interface {
	GetName() string //获取cron名称
	GetSpec() string //获取cron表达式
	GetFunc() func() //获取cron执行函数
}

type JobCommonService struct {
	Name string
	Spec string
	Log  *log.Helper
}

func (d *JobCommonService) GetName() string {
	return d.Name
}

func (d *JobCommonService) GetSpec() string {
	return d.Spec
}

type CronService struct {
	jobList []JobInterface
}

func NewJobService(saveNormalHttpInfo *NormalHttpService) *CronService {
	job := &CronService{
		jobList: make([]JobInterface, 0),
	}
	// 保存正常的http请求
	job.AddJob(saveNormalHttpInfo)
	return job
}

// AddJob 添加cron任务
func (s *CronService) AddJob(job JobInterface) {
	s.jobList = append(s.jobList, job)
}

// GetJobList 获取cron任务列表
func (s *CronService) GetJobList() []JobInterface {
	return s.jobList
}
