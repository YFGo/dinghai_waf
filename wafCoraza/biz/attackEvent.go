package biz

import (
	"github.com/corazawaf/coraza/v3/types"
	uuid "github.com/satori/go.uuid"
	"log/slog"
	"net"
	"net/http"
	"strconv"
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
}

func NewAttackEventUsercase(repo AttackEventRepo) *AttackEventUsercase {
	return &AttackEventUsercase{
		repo: repo,
	}
}

// LogAttackEvent 记录每次的攻击事件到JSON文件
func (uc *AttackEventUsercase) LogAttackEvent(matchRules []types.MatchedRule, req *http.Request, requestBody []byte, action, nextAction uint8) {
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
	endRuleInfo := matchRules[len(matchRules)-1] // 获取生效的最后一条规则信息
	attackId := uuid.NewV4()                     //生成攻击事件的唯一id
	event := model.AttackEvent{
		RuleId:        endRuleInfo.Rule().ID(),
		Port:          port,
		Timestamp:     time.Now(),
		IP:            endRuleInfo.ClientIPAddress(),
		ID:            attackId.String(),
		RequestMethod: req.Method,
		RequestURI:    req.RequestURI,
		Action:        strconv.Itoa(int(action)),
		NextAction:    strconv.Itoa(int(nextAction)),
		Message:       endRuleInfo.Message(),
		Protocol:      req.Proto,
		RuleName:      endRuleInfo.Message(),
		RuleDesc:      endRuleInfo.Data(),
		Request:       string(requestBody),
	}
	eventCsv := []model.AttackEvent{event} // csv格式数据正确处理
	uc.repo.AppendToFile(eventCsv)
}

// StartTimeTask 开启定时任务
func (uc *AttackEventUsercase) StartTimeTask() func() {
	return uc.repo.WriteEventTask
}
