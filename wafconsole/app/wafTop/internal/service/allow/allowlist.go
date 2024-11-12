package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"log/slog"
	"time"
	"wafconsole/app/wafTop/internal/biz/allow"
	"wafconsole/app/wafTop/internal/data/model"
	"wafconsole/app/wafTop/internal/server/plugin"
	"wafconsole/app/wafTop/internal/utils"

	pb "wafconsole/api/wafTop/v1"
)

type AllowListService struct {
	uc *allow.ListAllowUsecase
	pb.UnimplementedAllowListServer
}

func NewAllowListService(uc *allow.ListAllowUsecase) *AllowListService {
	return &AllowListService{
		uc: uc,
	}
}

func (s *AllowListService) CreateAllowList(ctx context.Context, req *pb.CreateAllowListRequest) (*pb.CreateAllowListReply, error) {
	allowInfo := model.AllowList{
		Name:        req.Name,
		Description: req.Description,
		Key:         req.Key,
		Value:       req.Value,
	}
	err := s.uc.CreateAllow(ctx, allowInfo)
	if err != nil {
		if utils.StatusErr(err, codes.AlreadyExists) {
			return nil, plugin.AllowExistErr()
		}
		slog.ErrorContext(ctx, "create allowlist error : ", err)
		return nil, err
	}
	return &pb.CreateAllowListReply{}, nil
}
func (s *AllowListService) UpdateAllowList(ctx context.Context, req *pb.UpdateAllowListRequest) (*pb.UpdateAllowListReply, error) {
	allowInfo := model.AllowList{
		Name:        req.Name,
		Description: req.Description,
		Key:         req.Key,
		Value:       req.Value,
	}
	err := s.uc.UpdateAllow(ctx, req.Id, allowInfo)
	if err != nil {
		if utils.StatusErr(err, codes.AlreadyExists) {
			return nil, plugin.AllowExistErr()
		}
		slog.ErrorContext(ctx, "update allowlist error : ", err)
		return nil, err
	}
	return &pb.UpdateAllowListReply{}, nil
}
func (s *AllowListService) DeleteAllowList(ctx context.Context, req *pb.DeleteAllowListRequest) (*pb.DeleteAllowListReply, error) {
	err := s.uc.DeleteAllow(ctx, req.Ids)
	if err != nil {
		slog.ErrorContext(ctx, "delete allowlist error : ", err)
		return nil, err
	}
	return &pb.DeleteAllowListReply{}, nil
}
func (s *AllowListService) GetAllowList(ctx context.Context, req *pb.GetAllowListRequest) (*pb.GetAllowListReply, error) {
	allowDetail, err := s.uc.GetAllowDetail(ctx, req.Id)
	if err != nil {
		slog.ErrorContext(ctx, "get allowlist error : ", err)
		return nil, plugin.ServerErr()
	}
	return &pb.GetAllowListReply{
		Name:        allowDetail.Name,
		Description: allowDetail.Description,
		Key:         allowDetail.Key,
		Value:       allowDetail.Value,
		CreatedAt:   time.Unix(allowDetail.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Unix(allowDetail.UpdatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
	}, nil
}
func (s *AllowListService) ListAllowList(ctx context.Context, req *pb.ListAllowListRequest) (*pb.ListAllowListReply, error) {
	limit := req.PageSize
	offset := req.PageSize * (req.PageNow - 1)
	total, listAllow, err := s.uc.ListAllow(ctx, limit, offset, req.Name)
	if err != nil {
		slog.ErrorContext(ctx, "list allowlist error : ", err)
		return nil, plugin.ServerErr()
	}
	listAllowReply := make([]*pb.AllowListInfo, 0)
	for _, allowInfo := range listAllow {
		allowReply := &pb.AllowListInfo{
			Id:        int64(allowInfo.ID),
			Name:      allowInfo.Name,
			Key:       allowInfo.Key,
			Value:     allowInfo.Value,
			CreatedAt: time.Unix(allowInfo.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Unix(allowInfo.UpdatedAt.Unix(), 0).Format("2006-01-02 15:04:05"),
		}
		listAllowReply = append(listAllowReply, allowReply)
	}
	return &pb.ListAllowListReply{
		Total:         total,
		ListAllowList: listAllowReply,
	}, nil
}
