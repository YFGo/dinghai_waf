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
	data               *Data
	partitionConsumers sarama.PartitionConsumer // 分区消费者
	log                *log.Helper
}

func NewNormalHttpRepo(data *Data, logger log.Logger) normalhttp.RepoNormalHttp {
	repo := &normalHttpRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
	offset, _ := repo.getOffsetFromRedis()
	partitionConsumer, err := data.kafkaConsumer.ConsumePartition(
		waftop.NormalHttpTopic,
		0,
		offset,
	)
	if err != nil {
		panic(err)
	}
	repo.partitionConsumers = partitionConsumer
	return repo
}

func (n *normalHttpRepo) SaveNormalHttp2DB(ctx context.Context) error {
	normalHttpInfoList := make([]*v1.NormalHttpInfo, 0)
	var maxOffset int64 = 0
	for msg := range n.partitionConsumers.Messages() {
		normalHttpInfo, err := n.processMessage(msg)
		if err != nil {
			n.log.WithContext(ctx).Error(err)
			return err
		}
		normalHttpInfoList = append(normalHttpInfoList, normalHttpInfo)
		maxOffset = msg.Offset
	}
	if err := n.commitBatch(normalHttpInfoList, maxOffset); err != nil {
		n.log.WithContext(ctx).Error(err)
		return err
	}
	return nil
}

// processMessage 处理消息
func (n *normalHttpRepo) processMessage(message *sarama.ConsumerMessage) (*v1.NormalHttpInfo, error) {
	var normalHttp types.NormalHttpInfo
	if err := json.Unmarshal(message.Value, &normalHttp); err != nil {
		return nil, err
	}

	return &v1.NormalHttpInfo{
		Id:            normalHttp.ID,
		Ip:            normalHttp.IP,
		RequestUri:    normalHttp.RequestURI,
		RequestTime:   timestamppb.New(normalHttp.RequestTime),
		RequestMethod: normalHttp.RequestMethod,
		Protocol:      normalHttp.Protocol,
		RequestBody:   normalHttp.RequestBody,
	}, nil
}

// commitBatch 写入ClickHouse并提交Offset
func (n *normalHttpRepo) commitBatch(batch []*v1.NormalHttpInfo, offset int64) error {
	if _, err := n.data.normalHttpRpc.CreateNormalHttpBatch(context.Background(), &v1.CreateNormalHttpRequest{
		NormalHttpInfos: batch,
	}); err != nil {
		n.log.Error("batch create failed: ", err)
		return err
	}

	// 提交Offset到Redis（注意：需按分区存储）
	if offset != -1 {
		if err := n.saveOffsetToRedis(offset); err != nil {
			n.log.Error("offset save failed: ", err)
			return err
		}
		n.log.Infof("committed %d messages, offset: %d", len(batch), offset)
	}
	return nil
}

func (n *normalHttpRepo) finalCommit(session sarama.ConsumerGroupSession, batch []*v1.NormalHttpInfo, offset int64) error {
	if err := n.commitBatch(batch, offset); err != nil {
		return err
	}
	n.log.Info("final commit completed")
	return nil
}

// getOffsetFromRedis 获取最新的偏移量
func (n *normalHttpRepo) getOffsetFromRedis() (int64, error) {
	offsetStr, err := n.data.rdb.Get(context.Background(), cron.NormalHttpOffsetKey).Result()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(offsetStr, 10, 64)
}

// 保存偏移量
func (n *normalHttpRepo) saveOffsetToRedis(offset int64) error {
	return n.data.rdb.Set(context.Background(), cron.NormalHttpOffsetKey, offset, 0).Err()
}
