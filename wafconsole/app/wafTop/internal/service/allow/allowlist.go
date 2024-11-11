package service

import (
	"context"
	"wafconsole/app/wafTop/internal/biz/allow"

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
	return &pb.CreateAllowListReply{}, nil
}
func (s *AllowListService) UpdateAllowList(ctx context.Context, req *pb.UpdateAllowListRequest) (*pb.UpdateAllowListReply, error) {
	return &pb.UpdateAllowListReply{}, nil
}
func (s *AllowListService) DeleteAllowList(ctx context.Context, req *pb.DeleteAllowListRequest) (*pb.DeleteAllowListReply, error) {
	return &pb.DeleteAllowListReply{}, nil
}
func (s *AllowListService) GetAllowList(ctx context.Context, req *pb.GetAllowListRequest) (*pb.GetAllowListReply, error) {
	return &pb.GetAllowListReply{}, nil
}
func (s *AllowListService) ListAllowList(ctx context.Context, req *pb.ListAllowListRequest) (*pb.ListAllowListReply, error) {
	return &pb.ListAllowListReply{}, nil
}
