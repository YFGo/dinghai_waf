package service

import (
	"context"
	"google.golang.org/grpc/codes"
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

func (s *CommonService) GetCaptcha(ctx context.Context, req *pb.GetCaptchaRequest) (*pb.GetCaptchaReply, error) {
	captchaId, masterImg, thumbImg, err := s.uc.GetCaptcha(ctx)
	if err != nil {
		return nil, up.ServerErr()
	}

	return &pb.GetCaptchaReply{
		CaptchaId:   captchaId,
		MasterImage: masterImg,
		ThumbImage:  thumbImg,
	}, nil
}

func (s *CommonService) VerifyCaptcha(ctx context.Context, req *pb.VerifyCaptchaRequest) (*pb.VerifyCaptchaReply, error) {
	err := s.uc.VerifyCaptchaInfo(ctx, req.CaptchaId, float64(req.UserAngle))
	if err != nil {
		if up.StatusErr(err, codes.PermissionDenied) || up.StatusErr(err, codes.Canceled) {
			return nil, up.CaptchaErr()
		}
		return nil, up.ServerErr()
	}
	return &pb.VerifyCaptchaReply{}, nil
}

func (s *CommonService) SendCode(ctx context.Context, req *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	err := s.uc.SendCode(ctx, req.UserEmail)
	if err != nil {
		return nil, up.ServerErr()
	}
	return &pb.SendEmailReply{}, nil
}
