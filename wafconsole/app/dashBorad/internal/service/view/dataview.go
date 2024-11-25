package view

import (
	"context"
	"wafconsole/app/dashBorad/internal/biz/viewLogic"

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

func (s *DataViewService) GetAttackInfoFromDay(ctx context.Context, req *pb.GetAttackInfoFromDayRequest) (*pb.GetAttackInfoFromDayReply, error) {
	return &pb.GetAttackInfoFromDayReply{}, nil
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
