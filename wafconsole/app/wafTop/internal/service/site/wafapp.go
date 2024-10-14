package service

import (
	"context"
	siteBiz "wafconsole/app/wafTop/internal/biz/site"

	pb "wafconsole/api/wafTop/v1"
)

type WafAppService struct {
	uc *siteBiz.WafAppUsecase
	pb.UnimplementedWafAppServer
}

func NewWafAppService(uc *siteBiz.WafAppUsecase) *WafAppService {
	return &WafAppService{
		uc: uc,
	}
}

func (s *WafAppService) CreateWafApp(ctx context.Context, req *pb.ChangeServerRequest) (*pb.CreateWafAppReply, error) {
	return &pb.CreateWafAppReply{}, nil
}
func (s *WafAppService) UpdateWafApp(ctx context.Context, req *pb.ChangeServerRequest) (*pb.UpdateWafAppReply, error) {
	return &pb.UpdateWafAppReply{}, nil
}
func (s *WafAppService) DeleteWafApp(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	return &pb.DeleteReply{}, nil
}
func (s *WafAppService) GetWafApp(ctx context.Context, req *pb.GetWafAppRequest) (*pb.GetWafAppReply, error) {
	return &pb.GetWafAppReply{}, nil
}
func (s *WafAppService) ListWafApp(ctx context.Context, req *pb.ListWafAppRequest) (*pb.ListWafAppReply, error) {
	return &pb.ListWafAppReply{}, nil
}
