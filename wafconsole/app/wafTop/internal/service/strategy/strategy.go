package service

import (
	"context"
	"log/slog"
	"time"
	strategyBiz "wafconsole/app/wafTop/internal/biz/strategy"
	"wafconsole/app/wafTop/internal/data/model"

	pb "wafconsole/api/wafTop/v1"
)

type StrategyService struct {
	uc *strategyBiz.WafStrategyUsecase
	pb.UnimplementedStrategyServer
}

func NewStrategyService(uc *strategyBiz.WafStrategyUsecase) *StrategyService {
	return &StrategyService{
		uc: uc,
	}
}

func (s *StrategyService) CreateStrategy(ctx context.Context, req *pb.CreateStrategyRequest) (*pb.CreateStrategyReply, error) {
	strategyInfo := model.Strategy{
		Name:        req.Name,
		Description: req.Description,
		Status:      uint8(req.Status),
		Action:      uint8(req.Action),
		NextAction:  uint8(req.NextAction),
	}
	var listStrategyConfig []model.StrategyConfig
	for _, groupId := range req.GroupId {
		strategyConfig := model.StrategyConfig{
			RuleGroupID: groupId,
		}
		listStrategyConfig = append(listStrategyConfig, strategyConfig)
	}
	strategyInfo.StrategyConfig = listStrategyConfig
	err := s.uc.CreateStrategy(ctx, strategyInfo)
	if err != nil {
		slog.ErrorContext(ctx, "CreateStrategy err : %v", err)
		return nil, err
	}
	return &pb.CreateStrategyReply{}, nil
}
func (s *StrategyService) UpdateStrategy(ctx context.Context, req *pb.UpdateStrategyRequest) (*pb.UpdateStrategyReply, error) {
	strategyInfo := model.Strategy{
		Name:        req.Name,
		Description: req.Description,
		Status:      uint8(req.Status),
		Action:      uint8(req.Action),
		NextAction:  uint8(req.NextAction),
	}
	var listStrategyConfig []model.StrategyConfig
	for _, groupId := range req.GroupId {
		strategyConfig := model.StrategyConfig{
			StrategyId:  req.Id,
			RuleGroupID: groupId,
		}
		listStrategyConfig = append(listStrategyConfig, strategyConfig)
	}
	strategyInfo.StrategyConfig = listStrategyConfig
	err := s.uc.UpdateStrategy(ctx, req.Id, strategyInfo)
	if err != nil {
		slog.ErrorContext(ctx, "UpdateStrategy err : %v", err)
		return nil, err
	}
	return &pb.UpdateStrategyReply{}, nil
}
func (s *StrategyService) DeleteStrategy(ctx context.Context, req *pb.DeleteStrategyRequest) (*pb.DeleteStrategyReply, error) {
	if err := s.uc.DeleteStrategy(ctx, req.Ids); err != nil {
		slog.ErrorContext(ctx, "DeleteStrategy err : %v", err)
		return nil, err
	}
	return &pb.DeleteStrategyReply{}, nil
}
func (s *StrategyService) GetStrategy(ctx context.Context, req *pb.GetStrategyRequest) (*pb.GetStrategyReply, error) {
	strategyInfo, err := s.uc.GetStrategyDetail(ctx, req.Id)
	if err != nil {
		slog.ErrorContext(ctx, "GetStrategy err : %v", err)
		return nil, err
	}
	var ruleGroupList []*pb.RuleGroupInfo
	for _, rule := range strategyInfo.RuleGroupInfo {
		ruleGroup := &pb.RuleGroupInfo{
			Id:          int64(rule.ID),
			Name:        rule.Name,
			Description: rule.Description,
			IsBuildin:   int64(rule.IsBuildin),
			CreatedAt:   time.Unix(rule.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
			UpdatedAt:   time.Unix(rule.UpdatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
		}
		ruleGroupList = append(ruleGroupList, ruleGroup)
	}
	return &pb.GetStrategyReply{
		Name:           strategyInfo.Name,
		Description:    strategyInfo.Description,
		Status:         int64(strategyInfo.Status),
		Action:         int64(strategyInfo.Action),
		NextAction:     int64(strategyInfo.NextAction),
		RuleGroupInfos: ruleGroupList,
	}, nil
}
func (s *StrategyService) ListStrategy(ctx context.Context, req *pb.ListStrategyRequest) (*pb.ListStrategyReply, error) {
	limit := req.PageSize
	offset := req.PageSize * (req.PageNow - 1)
	listStrategy, total, err := s.uc.ListStrategyInfo(ctx, limit, offset, req.Status, req.Name)
	if err != nil {
		slog.ErrorContext(ctx, "ListStrategy err :", err)
		return nil, err
	}
	var listStrategyInfo []*pb.StrategyInfo
	for _, strategy := range listStrategy {
		strategyInfo := &pb.StrategyInfo{
			Id:          int64(strategy.ID),
			Name:        strategy.Name,
			Description: strategy.Description,
			Status:      int64(strategy.Status),
			Action:      int64(strategy.Action),
			NextAction:  int64(strategy.NextAction),
			CreatedAt:   time.Unix(strategy.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
			UpdatedAt:   time.Unix(strategy.UpdatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
		}
		listStrategyInfo = append(listStrategyInfo, strategyInfo)
	}
	return &pb.ListStrategyReply{
		Total:      total,
		Strategies: listStrategyInfo,
	}, nil
}
