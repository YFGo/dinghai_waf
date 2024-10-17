package service

import (
	"context"
	"log/slog"
	ruleBiz "wafconsole/app/wafTop/internal/biz/rule"
	"wafconsole/app/wafTop/internal/data/model"

	pb "wafconsole/api/wafTop/v1"
)

type UserRuleService struct {
	uc *ruleBiz.UserRuleUsecase
	pb.UnimplementedUserRuleServer
}

func NewUserRuleService(uc *ruleBiz.UserRuleUsecase) *UserRuleService {
	return &UserRuleService{
		uc: uc,
	}
}

func (s *UserRuleService) CreateUserRule(ctx context.Context, req *pb.CreateUserRuleRequest) (*pb.CreateUserRuleReply, error) {
	userRuleInfo := model.UserRule{
		Name:        req.Name,
		Description: req.Description,
		RiskLevel:   uint8(req.RiskLevel),
		Status:      uint8(req.Status),
		GroupId:     req.GroupId,
		SeclangMod: model.SeclangMod{
			MatchGoal:    req.SeclangMod.MatchGoal,
			MatchAction:  req.SeclangMod.MatchAction,
			MatchContent: req.SeclangMod.MatchContent,
		},
	}
	err := s.uc.CreateUserRule(ctx, userRuleInfo)
	if err != nil {
		slog.ErrorContext(ctx, "CreateUserRule err : %v", err)
		return nil, err
	}
	return &pb.CreateUserRuleReply{}, nil
}
func (s *UserRuleService) UpdateUserRule(ctx context.Context, req *pb.UpdateUserRuleRequest) (*pb.UpdateUserRuleReply, error) {
	userRuleInfo := model.UserRule{
		Name:        req.Name,
		Description: req.Description,
		RiskLevel:   uint8(req.RiskLevel),
		Status:      uint8(req.Status),
		GroupId:     req.GroupId,
		SeclangMod: model.SeclangMod{
			MatchGoal:    req.SeclangMod.MatchGoal,
			MatchAction:  req.SeclangMod.MatchAction,
			MatchContent: req.SeclangMod.MatchContent,
		},
	}
	err := s.uc.UpdateUserRule(ctx, req.Id, userRuleInfo)
	if err != nil {
		slog.ErrorContext(ctx, "UpdateUserRule err : %v", err)
		return nil, err
	}
	return &pb.UpdateUserRuleReply{}, nil
}
func (s *UserRuleService) DeleteUserRule(ctx context.Context, req *pb.DeleteUserRuleRequest) (*pb.DeleteUserRuleReply, error) {
	err := s.uc.DeleteUserRule(ctx, req.Ids)
	if err != nil {
		slog.ErrorContext(ctx, "DeleteUserRule err : %v", err)
		return nil, err
	}
	return &pb.DeleteUserRuleReply{}, nil
}
