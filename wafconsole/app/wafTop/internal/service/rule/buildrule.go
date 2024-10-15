package service

import (
	"context"
	"log/slog"
	"wafconsole/app/wafTop/internal/biz/rule"

	pb "wafconsole/api/wafTop/v1"
)

type BuildRuleService struct {
	uc *ruleBiz.BuildRuleUsecase
	pb.UnimplementedBuildRuleServer
}

func NewBuildRuleService(uc *ruleBiz.BuildRuleUsecase) *BuildRuleService {
	return &BuildRuleService{
		uc: uc,
	}
}

func (s *BuildRuleService) GetBuildRule(ctx context.Context, req *pb.GetBuildRuleRequest) (*pb.GetBuildRuleReply, error) {
	if req == nil {
		slog.ErrorContext(ctx, "get buildin_rule req failed: ", req)
		return nil, nil
	}
	buildinRule, err := s.uc.GetBuildinRuleDetailById(ctx, req.Id)
	if err != nil {
		slog.ErrorContext(ctx, "uc get buildin_rule failed: ", err)
		return nil, err
	}
	return &pb.GetBuildRuleReply{
		Name:        buildinRule.Name,
		Description: buildinRule.Description,
		RiskLevel:   uint32(buildinRule.RiskLevel),
		Status:      uint32(buildinRule.Status),
	}, nil
}
func (s *BuildRuleService) ListBuildRule(ctx context.Context, req *pb.ListBuildRuleRequest) (*pb.ListBuildRuleReply, error) {
	if req == nil {
		slog.ErrorContext(ctx, "list buildin_rules req failed: ", req)
		return nil, nil
	}
	limit := req.PageSize
	offset := req.PageSize * (req.PageNow - 1)
	buildinRules, total, err := s.uc.ListBuildinRules(ctx, req.Name, limit, offset)
	if err != nil {
		slog.ErrorContext(ctx, "uc list buildin_rules failed: ", err)
		return nil, err
	}
	var buildinRuleRes []*pb.BuildinRule
	for _, buildinRule := range buildinRules {
		buildinRes := &pb.BuildinRule{
			Id:          int64(buildinRule.ID),
			RiskLevel:   uint32(buildinRule.RiskLevel),
			Status:      uint32(buildinRule.Status),
			Name:        buildinRule.Name,
			Description: buildinRule.Description,
		}
		buildinRuleRes = append(buildinRuleRes, buildinRes)
	}
	return &pb.ListBuildRuleReply{
		Total:        total,
		BuildinRules: buildinRuleRes,
	}, nil
}
