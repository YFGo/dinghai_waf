package data

import (
	"context"
	"encoding/json"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"sync"
	"time"
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
	partitionConsumers map[int32]sarama.PartitionConsumer // 分区消费者
	log                *log.Helper
}

func NewNormalHttpRepo(data *Data, logger log.Logger) normalhttp.RepoNormalHttp {
	repo := &normalHttpRepo{
		data:               data,
		log:                log.NewHelper(logger),
		partitionConsumers: make(map[int32]sarama.PartitionConsumer),
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
	repo.partitionConsumers[0] = partitionConsumer
	return repo
}

var consumeMutex sync.Mutex

const (
	normalHttpBatchSize = 100
	consumeTimeout      = 55 * time.Second
)

func (n *normalHttpRepo) SaveNormalHttp2DB(ctx context.Context) error {
	consumeMutex.Lock()
	defer consumeMutex.Unlock()

	// 1. 主动拉取消息（非阻塞模式）
	var (
		batchBuffer       = make([]*v1.NormalHttpInfo, 0, normalHttpBatchSize)
		maxOffset   int64 = -1
	)

	// 设置5秒超时
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

Loop:
	for {
		select {
		case message := <-n.partitionConsumers[0].Messages():
			info, err := n.processMessage(message)
			if err != nil {
				n.log.Error("message processing failed: ", err)
				continue
			}
			batchBuffer = append(batchBuffer, info)
			maxOffset = message.Offset

			// 批量达到阈值时立即写入
			if len(batchBuffer) >= normalHttpBatchSize {
				if err := n.commitBatch(batchBuffer, maxOffset); err != nil {
					return err
				}
				batchBuffer = batchBuffer[:0]
			}

		case <-timeoutCtx.Done():
			break Loop // 主动退出拉取循环

		case err := <-n.partitionConsumers[0].Errors():
			n.log.Errorf("Kafka consumer error: %v", err)
		}
	}

	// 2. 提交剩余批次
	if len(batchBuffer) > 0 {
		return n.commitBatch(batchBuffer, maxOffset)
	}
	return nil
}

func (n *normalHttpRepo) Setup(session sarama.ConsumerGroupSession) error {
	offset, err := n.getOffsetFromRedis()
	if err != nil {
		n.log.Warn("failed to get offset from redis, using default: ", err)
		offset = 0
	}

	for partition := range session.Claims()[waftop.NormalHttpTopic] {
		session.ResetOffset(waftop.NormalHttpTopic, int32(partition), offset, "")
	}
	n.log.Info("normal_http setup completed. initial offset:", offset)
	return nil
}

func (n *normalHttpRepo) Cleanup(sarama.ConsumerGroupSession) error {
	n.log.Info("normal_http consumer cleanup completed")
	return nil
}

func (n *normalHttpRepo) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	var (
		batchBuffer              = make([]*v1.NormalHttpInfo, 0, normalHttpBatchSize)
		maxCommittedOffset int64 = -1
	)
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				return n.finalCommit(session, batchBuffer, maxCommittedOffset)
			}

			info, err := n.processMessage(message)
			if err != nil {
				n.log.Error("message processing failed: ", err)
				continue
			}

			batchBuffer = append(batchBuffer, info)
			session.MarkMessage(message, "")
			maxCommittedOffset = message.Offset

			if len(batchBuffer) >= normalHttpBatchSize {
				if err := n.commitBatch(batchBuffer, maxCommittedOffset); err != nil {
					return err
				}
				batchBuffer = batchBuffer[:0]
			}

		case <-session.Context().Done():
			return n.finalCommit(session, batchBuffer, maxCommittedOffset)
		}
	}
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
