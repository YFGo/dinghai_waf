package view

import (
	"context"
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

const timeFormat = "2006-01-02"

func calculateGrowthRate(current, previous int) float64 {
	if previous == 0 {
		return (float64(current) - 0.0001) / 0.0001 * 100 // 处理除以0的情况，可以返回一个特殊值或者错误
	}
	return (float64(current) - float64(previous)) / float64(previous) * 100
}

func (s *DataViewService) GetAttackInfoFromDay(ctx context.Context, req *pb.GetAttackInfoFromDayRequest) (*pb.GetAttackInfoFromDayReply, error) {

	today := time.Now().Format(timeFormat)
	yesterday := time.Now().Add(-24 * time.Hour).Format(timeFormat)
	attackCount, err := s.uc.GetDayAttack(ctx, today, yesterday)
	if err != nil {
		slog.ErrorContext(ctx, "GetDayAttack err : %v", err)
		return nil, plugin.ServerErr()
	}
	attackAdd := calculateGrowthRate(attackCount.AttackCount, attackCount.AttackYesterday)
	attackIPAdd := calculateGrowthRate(attackCount.AttackIPCount, attackCount.AttackIPYesterday)
	return &pb.GetAttackInfoFromDayReply{
		AttackCount:   int64(attackCount.AttackCount),
		AttackIpCount: int64(attackCount.AttackIPCount),
		AttackAdd:     float32(attackAdd),
		AttackIpAdd:   float32(attackIPAdd),
	}, nil
}
func (s *DataViewService) GetAttackInfoByTime(ctx context.Context, req *pb.GetAttackInfoByTimeRequest) (*pb.GetAttackInfoByTimeReply, error) {
	attackInfoCount, err := s.uc.GetAttackByTime(ctx, req.StartTime, req.EndTime)
	if err != nil {
		slog.ErrorContext(ctx, "GetAttackByTime err : %v", err)
		return nil, plugin.ServerErr()
	}
	attackInfoCountRespList := make([]*pb.AttackInfoByTime, 0)
	for _, value := range attackInfoCount {
		attackInfoCountRespList = append(attackInfoCountRespList, &pb.AttackInfoByTime{
			Time:          value.Time,
			AttackCount:   int64(value.AttackCount),
			AttackIpCount: int64(value.AttackIPCount),
		})
	}
	return &pb.GetAttackInfoByTimeReply{
		AttackInfoByTimes: attackInfoCountRespList,
	}, nil
}
func (s *DataViewService) GetAttackInfoFromServer(ctx context.Context, req *pb.GetAttackInfoFromServerRequest) (*pb.GetAttackInfoFromServerReply, error) {
	limit := req.PageSize
	offset := (req.PageNow - 1) * limit
	attackList, total, err := s.uc.ListAttackLog(ctx, limit, offset, req.StartTime, req.EndTime)
	if err != nil {
		slog.ErrorContext(ctx, "ListAttackLog err : %v", err)
		return nil, plugin.ServerErr()
	}
	res := make([]*pb.GetAttackInfoFormServer, 0, len(attackList))
	for _, attack := range attackList {
		temp := &pb.GetAttackInfoFormServer{
			LogId:    attack.LogID,
			Uri:      attack.URI,
			Ctime:    time.Unix(attack.Ctime.Unix(), 0).Format(timeFormat),
			RuleName: attack.RuleName,
		}
		res = append(res, temp)
	}

	return &pb.GetAttackInfoFromServerReply{
		GetAttackInfoFormServer: res,
		Total:                   total,
	}, nil
}
func (s *DataViewService) GetAttackIpFromAddr(ctx context.Context, req *pb.GetAttackIpFromAddrRequest) (*pb.GetAttackIpFromAddrReply, error) {
	return &pb.GetAttackIpFromAddrReply{}, nil
}

func (s *DataViewService) GetAttackDetail(ctx context.Context, req *pb.GetAttackDetailRequest) (*pb.GetAttackDetailReply, error) {
	attackDetail, err := s.uc.GetAttackLogDetail(ctx, req.LogId)
	if err != nil {
		slog.ErrorContext(ctx, "GetAttackDetail err : %v", err)
		return nil, plugin.ServerErr()
	}
	return &pb.GetAttackDetailReply{
		Uri:           attackDetail.URI,
		Ctime:         time.Unix(attackDetail.Ctime.Unix(), 0).Format(timeFormat),
		Protocol:      attackDetail.Protocol,
		RuleDesc:      attackDetail.RuleDesc,
		ClientIp:      attackDetail.ClientIP,
		ClientPort:    int64(attackDetail.ClientPort),
		RequestMethod: attackDetail.RequestMethod,
		Request:       attackDetail.Request,
		RuleName:      attackDetail.RuleName,
	}, nil
}
