package service

import (
	"context"
	siteBiz "wafconsole/app/wafTop/internal/biz/site"

	pb "wafconsole/api/wafTop/v1"
)

type ServerService struct {
	uc *siteBiz.ServerUsecase
	pb.UnimplementedServerServer
}

func NewServerService(uc *siteBiz.ServerUsecase) *ServerService {
	return &ServerService{
		uc: uc,
	}
}

func (s *ServerService) CreateServer(ctx context.Context, req *pb.ChangeServerRequest) (*pb.CreateServerReply, error) {
	return &pb.CreateServerReply{}, nil
}
func (s *ServerService) UpdateServer(ctx context.Context, req *pb.ChangeServerRequest) (*pb.UpdateServerReply, error) {
	return &pb.UpdateServerReply{}, nil
}
func (s *ServerService) DeleteServer(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	return &pb.DeleteReply{}, nil
}
func (s *ServerService) GetServer(ctx context.Context, req *pb.GetServerRequest) (*pb.GetServerReply, error) {
	return &pb.GetServerReply{}, nil
}
func (s *ServerService) ListServer(ctx context.Context, req *pb.ListServerRequest) (*pb.ListServerReply, error) {
	return &pb.ListServerReply{}, nil
}
