package service

import (
	"context"
	"time"
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
		return nil, nil
	}
	buildinRule, err := s.uc.GetBuildinRuleDetailById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetBuildRuleReply{
		Name:        buildinRule.Name,
		Description: buildinRule.Description,
		RiskLevel:   int64(buildinRule.RiskLevel),
		GroupId:     buildinRule.GroupId,
	}, nil
}
func (s *BuildRuleService) ListBuildRule(ctx context.Context, req *pb.ListBuildRuleRequest) (*pb.ListBuildRuleReply, error) {
	if req == nil {
		return nil, nil
	}
	limit := req.PageSize
	offset := req.PageSize * (req.PageNow - 1)
	buildinRules, total, err := s.uc.ListBuildinRules(ctx, req.Name, limit, offset)
	if err != nil {
		return nil, err
	}
	var buildinRuleRes []*pb.BuildinRule
	for _, buildinRule := range buildinRules {
		buildinRes := &pb.BuildinRule{
			Id:          int64(buildinRule.ID),
			RiskLevel:   int64(buildinRule.RiskLevel),
			Name:        buildinRule.Name,
			Description: buildinRule.Description,
			GroupId:     buildinRule.GroupId,
			CreatedAt:   time.Unix(buildinRule.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
			UpdatedAt:   time.Unix(buildinRule.UpdatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
		}
		buildinRuleRes = append(buildinRuleRes, buildinRes)
	}
	return &pb.ListBuildRuleReply{
		Total:        total,
		BuildinRules: buildinRuleRes,
	}, nil
}
