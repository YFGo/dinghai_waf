package data

import (
	"context"
	"encoding/json"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	v1 "wafconsole/api/wafTop/v1"
	"wafconsole/app/cron/internal/types"

	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/log"

	"wafconsole/app/cron/internal/biz/normalhttp"
	"wafconsole/utils/const/cron"
	"wafconsole/utils/const/waftop"
)

type normalHttpRepo struct {
	data *Data
	log  *log.Helper
}

func NewNormalHttpRepo(data *Data, logger log.Logger) normalhttp.RepoNormalHttp {
	return &normalHttpRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (n *normalHttpRepo) SaveNormalHttp2DB(ctx context.Context) error {
	if err := n.data.kafkaConsumeGroup.Consume(ctx, []string{waftop.NormalHttpTopic},
		n); err != nil {
		n.log.WithContext(ctx).Errorf("kafka consume error: %v", err)
		return err
	}
	return nil
}

func (n *normalHttpRepo) Setup(session sarama.ConsumerGroupSession) error {
	// 获取最新的偏移量
	var offset int
	offsetStr, err := n.data.rdb.Get(context.Background(), cron.NormalHttpOffsetKey).Result()
	if err != nil {
		n.log.Warn("normal_http offset redis get error: ", err)
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			n.log.Warn("normal_http setup redis get error: ", err)
			return err
		}
	}
	session.ResetOffset(waftop.NormalHttpTopic, 0, int64(offset), "")
	log.Info(session.Claims())
	return nil
}

func (n *normalHttpRepo) Cleanup(session sarama.ConsumerGroupSession) error {
	//TODO implement me
	panic("implement me")
}

func (n *normalHttpRepo) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var normalHttpList []*v1.NormalHttpInfo
	var offsetNew int64
	for message := range claim.Messages() {
		var normalHttp types.NormalHttpInfo
		err := json.Unmarshal(message.Value, &normalHttp)
		if err != nil {
			n.log.Error("normal_http unmarshal error: ", err)
			return err
		}
		ts := timestamppb.New(normalHttp.RequestTime)
		normalHttpList = append(normalHttpList, &v1.NormalHttpInfo{
			Id:            normalHttp.ID,
			Ip:            normalHttp.IP,
			RequestUri:    normalHttp.RequestURI,
			RequestTime:   ts,
			RequestMethod: normalHttp.RequestMethod,
			Protocol:      normalHttp.Protocol,
			RequestBody:   normalHttp.RequestBody,
		})
		session.MarkMessage(message, "")
		offsetNew = message.Offset
	}
	// 写入clickhouse
	_, err := n.data.normalHttpRpc.CreateNormalHttpBatch(context.Background(), &v1.CreateNormalHttpRequest{NormalHttpInfos: normalHttpList})
	if err != nil {
		n.log.Error("normal_http rpc error: ", err)
		return err
	}
	session.Commit() // 提交偏移量
	// 将最新的偏移量写入redis
	err = n.data.rdb.Set(context.Background(), cron.NormalHttpOffsetKey, offsetNew, 0).Err()
	if err != nil {
		n.log.Error("normal_http offset redis set error: ", err)
		return err
	}
	return nil
}
