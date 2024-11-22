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
	data  *Data
	ready chan bool
}

func NewAttackLogRepo(data *Data) attack_log.AttackLogRepo {
	return &attackLogRepo{
		data:  data,
		ready: make(chan bool),
	}
}

const attackRedisOffsetKey = "attack_offset"

var (
	AttackMap       = new(sync.Map) //当map中存储的消息达到指定长度 , 将其写入到clickhouse
	AttackMapLength int
)

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
		slog.Error("set_up redis get error: ", err)
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
		//将此消息写入到全局map中
		var attackRule interface{}
		attackRule, ok := AttackMap.Load(attackEvent.ID)
		if ok {
			attackRuleSlice, ok := attackRule.([]model.SecLog)
			if !ok {
				// 如果类型不匹配，处理错误
				return err
			}
			// 追加新元素
			attackRuleSlice = append(attackRuleSlice, secLog)
			// 重新存储
			AttackMap.Store(attackEvent.ID, attackRuleSlice)
		} else {
			// 如果不存在，则初始化为空切片
			attackRule = []model.SecLog{secLog}
			AttackMap.Store(attackEvent.ID, attackRule)
		}
		AttackMapLength++
		if AttackMapLength >= 18 {
			a.Save(AttackMap, session) //保存到clickhouse中
			AttackMapLength = 0
		}
		session.MarkMessage(message, "")
		a.data.rdb.Set(context.Background(), attackRedisOffsetKey, message.Offset, 0) // 获取当前消费消息的偏移量 并存入redis
		session.Commit()                                                              // 手动提交
	}
	return nil
}

// Save 将消息队列中的消息保存到 clickhouse中
func (a *attackLogRepo) Save(secLog *sync.Map, session sarama.ConsumerGroupSession) {
	err := a.data.clickhouseDB.Transaction(func(tx *gorm.DB) error {
		var errClickhouse error
		// 遍历map中的所有消息
		secLog.Range(func(key, value interface{}) bool {
			secLogSlice, ok := value.([]model.SecLog)
			if !ok {
				// 如果类型不匹配，处理错误
				return false
			}
			a.data.clickhouseDB.CreateInBatches(secLogSlice, len(secLogSlice))
			//删除map中的消息
			secLog.Delete(key)
			return true
		})
		if errClickhouse != nil {
			return errClickhouse
		}
		return nil
	})
	if err != nil {
		slog.Error("clickhouse start save data error", err)
	}
}
