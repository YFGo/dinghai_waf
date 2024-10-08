package biz

import (
	"github.com/corazawaf/coraza/v3/types"
	uuid "github.com/satori/go.uuid"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
	"wafCoraza/data/model"
)

type AttackEventRepo interface {
	ReadAttackEvent() []model.AttackEvent
	AppendToFile(attackEvent []model.AttackEvent)
	WriteEventTask()
	CollectionAttackEvent()
}

type AttackEventUsercase struct {
	repo AttackEventRepo
	mu   sync.Mutex
}

func NewAttackEventUsercase(repo AttackEventRepo) *AttackEventUsercase {
	return &AttackEventUsercase{
		repo: repo,
	}
}

// LogAttackEvent 记录每次的攻击事件到JSON文件
func (uc *AttackEventUsercase) LogAttackEvent(matchRules []types.MatchedRule, req *http.Request, requestBody []byte) {
	// 获取真实的客户端 IP 地址和端口
	_, clientPort, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		slog.Error("LogAttackEvent SplitHostPort: ", err)
		return
	}
	port, err := strconv.Atoi(clientPort)
	if err != nil {
		slog.Error("LogAttackEvent Error converting port to int: ", err)
		return
	}
	uc.mu.Lock()
	defer func() {
		uc.mu.Unlock()
		if err := recover(); err != nil {
			slog.Error("LogAttackEvent recover: ", err)
		}
	}()
	//生成攻击事件的唯一id
	var attackEvents []model.AttackEvent
	attackId := uuid.NewV4()
	for i := 0; i < len(matchRules); i++ {
		rule := matchRules[i]
		var nextAction string
		if i+1 < len(matchRules) {
			nextAction = matchRules[i+1].Rule().Raw()
		}
		event := model.AttackEvent{
			RuleId:        rule.Rule().ID(),
			Port:          port,
			Timestamp:     time.Now(),
			IP:            rule.ClientIPAddress(),
			ID:            attackId.String(),
			RequestMethod: req.Method,
			RequestURI:    rule.URI(),
			Message:       rule.Message(),
			Protocol:      req.Proto,
			RuleName:      rule.Message(),
			RuleDesc:      rule.Data(),
			Request:       string(requestBody),
			Action:        rule.Rule().Raw(),
			NextAction:    nextAction,
		}
		attackEvents = append(attackEvents, event)
	}
	// 写入csv文件
	uc.repo.AppendToFile(attackEvents)
}

// StartTimeTask 开启定时任务
func (uc *AttackEventUsercase) StartTimeTask() func() {
	return uc.repo.WriteEventTask
}
