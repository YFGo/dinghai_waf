package view

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"wafconsole/app/dashBorad/internal/biz/viewLogic"
	"wafconsole/app/dashBorad/internal/server/plugin"

	pb "wafconsole/api/dashBorad/v1"
)

type DataViewService struct {
	uc *viewLogic.DataViewUsecase
	pb.UnimplementedDataViewServer
}

func NewDataViewService(uc *viewLogic.DataViewUsecase) *DataViewService {
	return &DataViewService{
		uc: uc,
	}
}

func calculateGrowthRate(current, previous int) float64 {
	if previous == 0 {
		return (float64(current) - 0.0001) / 0.0001 * 100 // 处理除以0的情况，可以返回一个特殊值或者错误
	}
	return (float64(current) - float64(previous)) / float64(previous) * 100
}

func (s *DataViewService) GetAttackInfoFromDay(ctx context.Context, req *pb.GetAttackInfoFromDayRequest) (*pb.GetAttackInfoFromDayReply, error) {
	const timeFormat = "2006-01-02"
	today := time.Now().Format(timeFormat)
	yesterday := time.Now().Add(-24 * time.Hour).Format(timeFormat)
	attackCount, err := s.uc.GetDayAttack(ctx, today, yesterday)
	if err != nil {
		slog.ErrorContext(ctx, "GetDayAttack err : %v", err)
		return nil, plugin.ServerErr()
	}
	attackAdd := calculateGrowthRate(attackCount.AttackCount, attackCount.AttackYesterday)
	attackIPAdd := calculateGrowthRate(attackCount.AttackIPCount, attackCount.AttackIPYesterday)
	fmt.Println(attackAdd, attackIPAdd)
	return &pb.GetAttackInfoFromDayReply{
		AttackCount:   int64(attackCount.AttackCount),
		AttackIpCount: int64(attackCount.AttackIPCount),
		AttackAdd:     float32(attackAdd),
		AttackIpAdd:   float32(attackIPAdd),
	}, nil
}
func (s *DataViewService) GetAttackInfoByTime(ctx context.Context, req *pb.GetAttackInfoByTimeRequest) (*pb.GetAttackInfoByTimeReply, error) {
	return &pb.GetAttackInfoByTimeReply{}, nil
}
func (s *DataViewService) GetAttackInfoFromServer(ctx context.Context, req *pb.GetAttackInfoFromServerRequest) (*pb.GetAttackInfoFromServerReply, error) {
	return &pb.GetAttackInfoFromServerReply{}, nil
}
func (s *DataViewService) GetAttackIpFromAddr(ctx context.Context, req *pb.GetAttackIpFromAddrRequest) (*pb.GetAttackIpFromAddrReply, error) {
	return &pb.GetAttackIpFromAddrReply{}, nil
}
