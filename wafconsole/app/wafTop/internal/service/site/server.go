package service

import (
	"context"
	"log/slog"
	siteBiz "wafconsole/app/wafTop/internal/biz/site"
	"wafconsole/app/wafTop/internal/data/model"

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

func (s *ServerService) CreateServer(ctx context.Context, req *pb.CreateServerRequest) (*pb.CreateServerReply, error) {
	if req == nil {
		slog.ErrorContext(ctx, "create server http request is null")
		return nil, nil
	}
	serverInfo := model.ServerWaf{
		Name: req.Name,
		Host: req.Host,
		IP:   req.Ip,
		Port: int(req.Port),
	}
	err := s.uc.CreateServerSite(ctx, serverInfo)
	if err != nil {
		slog.ErrorContext(ctx, "create server_waf service error: ", err)
		return nil, err
	}
	return &pb.CreateServerReply{}, nil
}
func (s *ServerService) UpdateServer(ctx context.Context, req *pb.UpdateServerRequest) (*pb.UpdateServerReply, error) {
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
