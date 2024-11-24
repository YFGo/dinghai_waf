package data

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/gocarina/gocsv"
	"log/slog"
	"os"
	"strconv"
	"sync"
	"wafCoraza/biz"
	"wafCoraza/data/model"
)

type saveAttackEventRepo struct {
	data *Data
	mu   sync.Mutex
}

func NewSaveAttackEventRepo(data *Data) biz.AttackEventRepo {
	return &saveAttackEventRepo{data: data}
}

// ReadAttackEvent 读取csv文件中的数据
func (s saveAttackEventRepo) ReadAttackEvent() []model.AttackEvent {
	file, err := os.OpenFile("wafCoraza/waf_log/attack_events.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		slog.Error("ReadAttackEvent Error opening file: ", err)
		return nil
	}
	defer file.Close()
	var attackEvents []model.AttackEvent
	if err := gocsv.UnmarshalFile(file, &attackEvents); err != nil {
		slog.Error("ReadAttackEvent Error unmarshaling file: ", err)
		return nil
	}
	return attackEvents
}

// AppendToFile 将新数据写入csv文件
func (s saveAttackEventRepo) AppendToFile(attackEvent []model.AttackEvent) {
	path := "wafCoraza/waf_log/attack_events.csv"
	var file *os.File
	//判断此文件是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644) //文件不存在 , 需要创建文件
		defer func(file *os.File) {
			if err := file.Close(); err != nil {
				slog.Error("Error closing file: ", err)
			}
		}(file)
		if err != nil {
			slog.Error("Error opening file: ", err)
			return
		}
		// 写入标头和数据
		err = gocsv.MarshalFile(attackEvent, file)
		if err != nil {
			slog.Error("Error marshaling file: ", err)
			return
		}
	} else {
		file, err = os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		defer func(file *os.File) {
			if err := file.Close(); err != nil {
				slog.Error("Error closing file: ", err)
			}
		}(file)
		if err != nil {
			slog.Error("Error opening file: ", err)
			return
		}
		//文件存在 , 不写入标头
		err = gocsv.MarshalWithoutHeaders(&attackEvent, file)
		if err != nil {
			slog.Error("Error marshaling file: ", err)
			return
		}
	}
}

// CollectionAttackEvent 将数据写入kafka
func (s saveAttackEventRepo) CollectionAttackEvent() {
	events := s.ReadAttackEvent()
	if len(events) == 0 {
		slog.Info("no attack event")
		return
	}
	var producerMessages []*sarama.ProducerMessage
	for _, event := range events {
		eventJson, err := json.Marshal(&event)
		if err != nil {
			slog.Error("json marshal error: ", err)
			return
		}
		msg := &sarama.ProducerMessage{
			Topic: model.AttackEventLogTopic,
			Value: sarama.StringEncoder(eventJson),
			Key:   sarama.StringEncoder(strconv.Itoa(event.RuleId)),
		}
		producerMessages = append(producerMessages, msg)
	}
	err := s.data.kafkaProducer.SendMessages(producerMessages)
	if err != nil {
		slog.Info("send messages to kafka: ", err)
		return
	}
	//写入成功之后 , 删除json文件
	if err := os.Remove("wafCoraza/waf_log/attack_events.csv"); err != nil {
		slog.Error("remove attack_events.csv error: ", err)
		return
	}
	slog.Info("write kafka success")
}

func (s saveAttackEventRepo) WriteEventTask() {
	spec := "0 0/5 * * * ?" // 每隔5分钟执行一次
	// 添加一个任务
	eventTaskID, err := s.data.timeTask.AddFunc(spec, s.CollectionAttackEvent)
	if err != nil {
		slog.Error("add task error: ", err, "task id: ", eventTaskID)
		return
	}
	s.data.timeTask.Start()
}
