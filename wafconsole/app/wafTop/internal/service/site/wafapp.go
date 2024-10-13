package site

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
	return &WafAppService{uc: uc}
}

func (s *WafAppService) CreateWafApp(ctx context.Context, req *pb.CreateWafAppRequest) (*pb.CreateWafAppReply, error) {
	return &pb.CreateWafAppReply{}, nil
}
