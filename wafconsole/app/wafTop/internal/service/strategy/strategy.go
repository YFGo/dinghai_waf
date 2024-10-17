package service

import (
	"context"
	strategyBiz "wafconsole/app/wafTop/internal/biz/strategy"

	pb "wafconsole/api/wafTop/v1"
)

type StrategyService struct {
	uc *strategyBiz.WafStrategyUsecase
	pb.UnimplementedStrategyServer
}

func NewStrategyService(uc *strategyBiz.WafStrategyUsecase) *StrategyService {
	return &StrategyService{
		uc: uc,
	}
}

func (s *StrategyService) CreateStrategy(ctx context.Context, req *pb.CreateStrategyRequest) (*pb.CreateStrategyReply, error) {
	return &pb.CreateStrategyReply{}, nil
}
func (s *StrategyService) UpdateStrategy(ctx context.Context, req *pb.UpdateStrategyRequest) (*pb.UpdateStrategyReply, error) {
	return &pb.UpdateStrategyReply{}, nil
}
func (s *StrategyService) DeleteStrategy(ctx context.Context, req *pb.DeleteStrategyRequest) (*pb.DeleteStrategyReply, error) {
	return &pb.DeleteStrategyReply{}, nil
}
func (s *StrategyService) GetStrategy(ctx context.Context, req *pb.GetStrategyRequest) (*pb.GetStrategyReply, error) {
	return &pb.GetStrategyReply{}, nil
}
func (s *StrategyService) ListStrategy(ctx context.Context, req *pb.ListStrategyRequest) (*pb.ListStrategyReply, error) {
	return &pb.ListStrategyReply{}, nil
}
