package service

import (
	"context"
	pb "wafconsole/api/user/v1"
	"wafconsole/app/user/internal/biz"
	up "wafconsole/utils/plugin"
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
		return nil, up.ServerErr()
	}
	return &pb.CreateNewTokenReply{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenNew,
		ExpireAt:     exprieAt,
	}, nil
}
