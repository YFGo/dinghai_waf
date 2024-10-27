package service

import (
	"context"
	"wafconsole/app/user/internal/biz"
	"wafconsole/app/user/internal/data/model"

	pb "wafconsole/api/user/v1"
)

type WafUserService struct {
	uc *biz.WafUserUsecase
	pb.UnimplementedWafUserServer
}

func NewWafUserService(uc *biz.WafUserUsecase) *WafUserService {
	return &WafUserService{
		uc: uc,
	}
}

func (s *WafUserService) CreateWafUser(ctx context.Context, req *pb.CreateWafUserRequest) (*pb.CreateWafUserReply, error) {
	userInfo := model.UserInfo{
		Email:    req.Email,
		Password: req.Password,
	}
	err := s.uc.SignUp(ctx, userInfo)
	if err != nil {
		return nil, err
	}
	return &pb.CreateWafUserReply{}, nil
}
func (s *WafUserService) UpdateWafUser(ctx context.Context, req *pb.UpdateWafUserRequest) (*pb.UpdateWafUserReply, error) {
	return &pb.UpdateWafUserReply{}, nil
}
func (s *WafUserService) DeleteWafUser(ctx context.Context, req *pb.DeleteWafUserRequest) (*pb.DeleteWafUserReply, error) {
	return &pb.DeleteWafUserReply{}, nil
}
func (s *WafUserService) GetWafUser(ctx context.Context, req *pb.GetWafUserRequest) (*pb.GetWafUserReply, error) {
	return &pb.GetWafUserReply{}, nil
}
func (s *WafUserService) Login(ctx context.Context, req *pb.LoginUserInfoRequest) (*pb.LoginUserInfoReply, error) {
	return &pb.LoginUserInfoReply{}, nil
}
func (s *WafUserService) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordReply, error) {
	return &pb.UpdatePasswordReply{}, nil
}
