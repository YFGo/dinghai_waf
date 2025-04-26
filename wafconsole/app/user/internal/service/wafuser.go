package service

import (
	"context"
	"log/slog"
	"wafconsole/app/user/internal/service/validation"

	pb "wafconsole/api/user/v1"
	"wafconsole/app/user/internal/biz"
	"wafconsole/app/user/internal/data/types"
	up "wafconsole/utils/plugin"

	"google.golang.org/grpc/codes"
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
	userInfo := types.UserInfo{
		Email:    req.Email,
		Password: req.Password,
	}
	err := s.uc.SignUp(ctx, userInfo, req.Code)
	if err != nil {
		if up.StatusErr(err, codes.PermissionDenied) {
			return nil, up.UserCodeErr()
		}
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
	userId := req.Id
	if userId == 0 { //查询自己
		userId = int64(ctx.Value(up.UserIDMid).(uint64))
	}
	userInfo, err := s.uc.GetUserInfoByID(ctx, userId)
	if err != nil {
		slog.ErrorContext(ctx, "GetUserInfoByID err : %v", err)
		return nil, up.ServerErr()
	}
	return &pb.GetWafUserReply{
		Email:      userInfo.Email,
		UserName:   userInfo.UserName,
		AvatarAddr: userInfo.AvatarAddr,
		Phone:      userInfo.Phone,
	}, nil
}
func (s *WafUserService) Login(ctx context.Context, req *pb.LoginUserInfoRequest) (*pb.LoginUserInfoReply, error) {
	loginInfo, isExist := validation.LoginMethod(uint8(req.LoginMethod), req.Email, req.Phone, req.Code, req.Password)
	if !isExist {
		return nil, up.LoginMethodErr()
	}
	accessToken, refreshToken, avatarAddr, userId, err := s.uc.LoginByEmailPassword(ctx, loginInfo.Account, loginInfo.AccountCheck, uint8(req.LoginMethod))
	if err != nil {
		if up.StatusErr(err, codes.NotFound) {
			return nil, up.UserNotFoundErr()
		}
		return nil, up.ServerErr()
	}
	return &pb.LoginUserInfoReply{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AvatarAddr:   avatarAddr,
		UserId:       uint64(userId),
	}, nil
}
func (s *WafUserService) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordReply, error) {
	return &pb.UpdatePasswordReply{}, nil
}
