package utils

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

const (
	cronTracerName = "dinghai/trace/corn-job"
)

var cronTracer = otel.Tracer(cronTracerName)

// AsyncCorn 异步执行的定时任务
type AsyncCorn struct {
	ctx context.Context
}

type JobFunc func(ctx context.Context) error

// AsyncTaskCorn 将kafka中的攻击日志异步保存到clickhouse中
func (a *AsyncCorn) AsyncTaskCorn() {
	//taskCron := cron.New(cron.WithSeconds())
	// 每日凌晨0点 , 将攻击日志保存到clickhouse中
	//_, err := taskCron.AddFunc("0 0 0 * * ?",
	//	funcCmd(func(ctx context.Context) error {
	//	return nil
	//	})
	//if err != nil {
	//	logrus.Error(err)
	//}
}

func funcCmd(f JobFunc) func() {
	return func() {
		ctx, span := cronTracer.Start(context.Background(), "cronJob")
		defer span.End()
		var err error
		defer func() {
			if panicErr := recover(); panicErr != nil {

			}
			if err != nil {
				logrus.WithContext(ctx).WithError(err).Errorf("%+v", err)
			}
		}()
		err = f(ctx)
	}
}
