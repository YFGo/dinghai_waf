package service

import (
	"context"
	"log/slog"
	siteBiz "wafconsole/app/wafTop/internal/biz/site"
	"wafconsole/app/wafTop/internal/data/model"

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

func (s *WafAppService) CreateWafApp(ctx context.Context, req *pb.CreateWafAppRequest) (*pb.CreateWafAppReply, error) {
	appInfo := model.AppWaf{
		Name:     req.Name,
		Url:      req.Url,
		ServerID: req.ServerId,
	}
	err := s.uc.CreateWafApp(ctx, appInfo)
	if err != nil {
		slog.ErrorContext(ctx, "create_waf_app  err : %v", err)
		return nil, err
	}
	return &pb.CreateWafAppReply{}, nil
}
func (s *WafAppService) UpdateWafApp(ctx context.Context, req *pb.UpdateWafAppRequest) (*pb.UpdateWafAppReply, error) {
	appInfo := model.AppWaf{
		Name:     req.Name,
		Url:      req.Url,
		ServerID: req.ServerId,
	}
	err := s.uc.UpdateWafApp(ctx, req.Id, appInfo)
	if err != nil {
		slog.ErrorContext(ctx, "update_waf_app  err : %v", err)
		return nil, err
	}
	return &pb.UpdateWafAppReply{}, nil
}
