package service

import (
	"context"
	"wafconsole/app/user/internal/biz"
	"wafconsole/app/user/internal/server/plugin"

	pb "wafconsole/api/user/v1"
)

type CommonService struct {
	uc *biz.WafUserCommonUsecase
	pb.UnimplementedCommonServer
}

func NewCommonService(uc *biz.WafUserCommonUsecase) *CommonService {
	return &CommonService{
		uc: uc,
	}
}

func (s *CommonService) CreateNewToken(ctx context.Context, req *pb.CreateNewTokenRequest) (*pb.CreateNewTokenReply, error) {
	accessToken, refreshTokenNew, exprieAt, err := s.uc.RefreshAccessToken(req.RefreshToken)
	if err != nil {
		return nil, plugin.ServerErr()
	}
	return &pb.CreateNewTokenReply{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenNew,
		ExpireAt:     exprieAt,
	}, nil
}
