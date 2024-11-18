package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"log/slog"
	"time"
	ruleBiz "wafconsole/app/wafTop/internal/biz/rule"
	"wafconsole/app/wafTop/internal/data/model"
	"wafconsole/app/wafTop/internal/server/plugin"
	"wafconsole/app/wafTop/internal/utils"

	pb "wafconsole/api/wafTop/v1"
)

type RuleGroupService struct {
	uc *ruleBiz.RuleGroupUsecase
	pb.UnimplementedRuleGroupServer
}

func NewRuleGroupService(uc *ruleBiz.RuleGroupUsecase) *RuleGroupService {
	return &RuleGroupService{
		uc: uc,
	}
}

func (s *RuleGroupService) CreateRuleGroup(ctx context.Context, req *pb.CreateRuleGroupRequest) (*pb.CreateRuleGroupReply, error) {
	if req == nil {
		slog.ErrorContext(ctx, "create rule_group req is nil")
		return nil, nil
	}
	ruleGroup := model.RuleGroup{
		Name:        req.Name,
		Description: req.Description,
		IsBuildin:   uint8(req.IsBuildin),
	}
	if err := s.uc.CreateRuleGroup(ctx, ruleGroup); err != nil {
		if utils.StatusErr(err, codes.AlreadyExists) {
			return nil, plugin.RuleGroupIsExistErr()
		}
		slog.ErrorContext(ctx, "create rule_group is error: ", err)
		return nil, err
	}
	return &pb.CreateRuleGroupReply{}, nil
}
func (s *RuleGroupService) UpdateRuleGroup(ctx context.Context, req *pb.UpdateRuleGroupRequest) (*pb.UpdateRuleGroupReply, error) {
	if req == nil {
		slog.ErrorContext(ctx, "update rule_group req is nil")
		return nil, nil
	}
	ruleGroup := model.RuleGroup{
		Name:        req.Name,
		Description: req.Description,
		IsBuildin:   uint8(req.IsBuildin),
	}
	if err := s.uc.UpdateRuleGroup(ctx, req.Id, ruleGroup); err != nil {
		slog.ErrorContext(ctx, "update rule_group is error: ", err)
		return nil, err
	}
	return &pb.UpdateRuleGroupReply{}, nil
}
func (s *RuleGroupService) DeleteRuleGroup(ctx context.Context, req *pb.DeleteRuleGroupRequest) (*pb.DeleteRuleGroupReply, error) {
	if req == nil {
		slog.ErrorContext(ctx, "delete rule_group req is nil")
		return nil, nil
	}
	var ids []int64
	for _, ruleGroupInfo := range req.DeleteRuleGroupInfos {
		ids = append(ids, ruleGroupInfo.Id)
	}
	if err := s.uc.DeleteRuleGroup(ctx, ids); err != nil {
		if utils.StatusErr(err, codes.FailedPrecondition) {
			return nil, plugin.RuleGroupIsUsingErr()
		}
		slog.ErrorContext(ctx, "delete rule_group is error: ", err)
		return nil, err
	}
	return &pb.DeleteRuleGroupReply{}, nil
}
func (s *RuleGroupService) GetRuleGroup(ctx context.Context, req *pb.GetRuleGroupRequest) (*pb.GetRuleGroupReply, error) {
	if req == nil {
		slog.ErrorContext(ctx, "get rule_group is nil")
		return nil, nil
	}
	ruleGroupInfos, err := s.uc.GetRuleGroupDetail(ctx, req.Id)
	if err != nil {
		slog.ErrorContext(ctx, "get rule_group is error: ", err)
		return nil, err
	}
	res := &pb.GetRuleGroupReply{
		Name:        ruleGroupInfos.RuleGroup.Name,
		Description: ruleGroupInfos.RuleGroup.Description,
		IsBuildin:   int64(ruleGroupInfos.RuleGroup.IsBuildin),
	}
	var listRuleInfo []*pb.ListRuleInfoByGroup
	for _, rule := range ruleGroupInfos.RuleInfos {
		seclangMod := &pb.SeclangMod{
			MatchGoal:    rule.SecLangMod.MatchGoal,
			MatchAction:  rule.SecLangMod.MatchAction,
			MatchContent: rule.SecLangMod.MatchContent,
		}
		ruleInfo := &pb.ListRuleInfoByGroup{
			Id:          rule.ID,
			Name:        rule.Name,
			Description: rule.Description,
			RiskLevel:   int64(rule.RiskLevel),
			SeclangMod:  seclangMod,
		}
		listRuleInfo = append(listRuleInfo, ruleInfo)
	}
	res.ListRules = listRuleInfo
	return res, nil
}
func (s *RuleGroupService) ListRuleGroup(ctx context.Context, req *pb.ListRuleGroupRequest) (*pb.ListRuleGroupReply, error) {
	if req == nil {
		slog.ErrorContext(ctx, "list rule_group req is nil")
		return nil, nil
	}
	limit := req.PageSize
	offset := req.PageSize * (req.PageNow - 1)
	listRuleGroup, total, err := s.uc.ListRuleGroupSearch(ctx, req.Name, int8(req.IsBuildin), limit, offset)
	if err != nil {
		slog.ErrorContext(ctx, "list rule_group is error: ", err)
		return nil, err
	}
	var listRuleGroupInfo []*pb.RuleGroupInfo
	for _, ruleGroup := range listRuleGroup {
		ruleGroupInfo := &pb.RuleGroupInfo{
			Id:          int64(ruleGroup.ID),
			Name:        ruleGroup.Name,
			Description: ruleGroup.Description,
			IsBuildin:   int64(ruleGroup.IsBuildin),
			CreatedAt:   time.Unix(ruleGroup.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
			UpdatedAt:   time.Unix(ruleGroup.UpdatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
		}
		listRuleGroupInfo = append(listRuleGroupInfo, ruleGroupInfo)
	}
	return &pb.ListRuleGroupReply{
		RuleGroupInfos: listRuleGroupInfo,
		Total:          total,
	}, nil
}
