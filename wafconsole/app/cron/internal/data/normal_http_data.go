package data

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"sync"
	"wafconsole/app/cron/internal/biz/normalhttp"
	"wafconsole/utils/const/waftop"
)

type normalHttpRepo struct {
	data  *Data
	ready chan bool
	log   *log.Helper
}

func NewNormalHttpRepo(data *Data, logger log.Logger) normalhttp.RepoNormalHttp {
	return &normalHttpRepo{
		data:  data,
		ready: make(chan bool),
		log:   log.NewHelper(logger),
	}
}

func (n *normalHttpRepo) SaveNormalHttp2DB(ctx context.Context) func() {
	appCtx, cancel := context.WithCancel(ctx)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := n.data.kafkaConsumeGroup.Consume(appCtx, []string{waftop.NormalHttpTopic},
				n); err != nil {
				n.log.WithContext(appCtx).Errorf("kafka consume error: %v", err)
				return
			}
			if ctx.Err() != nil {
				n.log.WithContext(appCtx).Errorf("kafka cancel error: %v", ctx.Err())
				return
			}
			n.ready = make(chan bool)
		}
	}()
	<-n.ready
	n.log.WithContext(appCtx).Info("kafka consume ready")
	return func() {
		n.log.WithContext(appCtx).Info("kafka consume cancel")
		cancel()
		wg.Wait()
	}

}

func (n *normalHttpRepo) Setup(session sarama.ConsumerGroupSession) error {
	session.ResetOffset(waftop.NormalHttpTopic, 0, 13, "")
	log.Info(session.Claims())
	// Mark the consumer as ready
	close(n.ready)
	return nil
}

func (n *normalHttpRepo) Cleanup(session sarama.ConsumerGroupSession) error {
	//TODO implement me
	panic("implement me")
}

func (n *normalHttpRepo) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Println(message)
		// 更新位移
		session.MarkMessage(message, "")
	}
	// 消费完成
	return nil
}
