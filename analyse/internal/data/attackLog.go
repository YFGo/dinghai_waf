package data

import (
	"analyse/internal/biz/attack_log"
	"analyse/internal/data/model"
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"gorm.io/gorm"
	"log/slog"
	"strconv"
	"sync"
)

type attackLogRepo struct {
	data       *Data
	ready      chan bool
	secLogList []model.SecLog
}

func NewAttackLogRepo(data *Data) attack_log.AttackLogRepo {
	return &attackLogRepo{
		data:       data,
		ready:      make(chan bool),
		secLogList: make([]model.SecLog, 0),
	}
}

const attackRedisOffsetKey = "attack_offset"

// Consumer 消费kafka中 的消息 , 当达到指定长度 , 存入clickhouse中
func (a *attackLogRepo) Consumer() func() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := a.data.kafkaConsumerGroup.Consume(ctx, []string{model.AttackEventsTopic}, a); err != nil {
				//当Setup失败时 , error返回在这里
				slog.ErrorContext(ctx, "consume error", err)
				return
			}
			// 检测这个上下文通道是否被取消
			if ctx.Err() != nil {
				return
			}
			a.ready = make(chan bool)
		}
	}()
	<-a.ready
	slog.InfoContext(ctx, "Sarama consumer up and running!...")
	//确保系统退出时,通道中的消息被消费
	return func() {
		cancel()
		wg.Wait()
	}
}

// Setup 在新的会话开始时运行
func (a *attackLogRepo) Setup(session sarama.ConsumerGroupSession) error {
	// 获取最新的偏移量
	var offset int
	offsetStr, err := a.data.rdb.Get(context.Background(), attackRedisOffsetKey).Result()
	if err != nil {
		slog.Warn("set_up redis get error: ", err)
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			slog.Error("set_up strconv_atoi error: ", err)
			return err
		}
	}
	session.ResetOffset(model.AttackEventsTopic, 0, int64(offset), "")
	close(a.ready)
	return nil
}

func (a *attackLogRepo) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (a *attackLogRepo) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	//具体消费消息
	for message := range claim.Messages() {
		attackEvent := model.AttackEvent{}
		err := json.Unmarshal(message.Value, &attackEvent)
		if err != nil {
			slog.Error("ConsumeClaim JSON Unmarshal error: ", err)
			return err
		}
		// 更新位移
		secLog := model.SecLog{
			RuleID:        attackEvent.RuleId,
			LogID:         attackEvent.ID,
			Ctime:         attackEvent.Timestamp,
			URI:           attackEvent.RequestURI,
			Protocol:      attackEvent.Protocol,
			Request:       attackEvent.Request,
			RequestMethod: attackEvent.RequestMethod,
			ClientIP:      attackEvent.IP,
			ClientPort:    int32(attackEvent.Port),
			RuleName:      attackEvent.RuleName,
			RuleDesc:      attackEvent.RuleDesc,
			Action:        attackEvent.Action,
			NextAction:    attackEvent.NextAction,
		}
		session.MarkMessage(message, "")                                              // 标记
		a.data.rdb.Set(context.Background(), attackRedisOffsetKey, message.Offset, 0) // 获取当前消费消息的偏移量 并存入redis
		a.secLogList = append(a.secLogList, secLog)
		if len(a.secLogList) >= 20 {
			a.Save(a.secLogList, session)
		}
		session.Commit()
	}
	return nil
}

// Save 将消息队列中的消息保存到 clickhouse中
func (a *attackLogRepo) Save(secLogList []model.SecLog, session sarama.ConsumerGroupSession) {
	err := a.data.clickhouseDB.Transaction(func(tx *gorm.DB) error {
		if err := a.data.clickhouseDB.CreateInBatches(secLogList, len(secLogList)).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		slog.Error("clickhouse start save data error", err)
		return
	}
	a.secLogList = make([]model.SecLog, 0)
}
