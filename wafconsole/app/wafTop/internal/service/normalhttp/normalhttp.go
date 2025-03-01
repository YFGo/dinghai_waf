package service

import (
	"context"
	"wafconsole/app/wafTop/internal/biz/normalhttp"

	pb "wafconsole/api/wafTop/v1"
)

type NormalHttpService struct {
	uc *normalhttp.UsecaseNormalHttp

	pb.UnimplementedNormalHttpServer
}

func NewNormalHttpService(uc *normalhttp.UsecaseNormalHttp) *NormalHttpService {
	return &NormalHttpService{
		uc: uc,
	}
}

func (s *NormalHttpService) CreateNormalHttp(ctx context.Context, req *pb.CreateNormalHttpRequest) (*pb.CreateNormalHttpReply, error) {
	return &pb.CreateNormalHttpReply{}, nil
}
func (s *NormalHttpService) UpdateNormalHttp(ctx context.Context, req *pb.UpdateNormalHttpRequest) (*pb.UpdateNormalHttpReply, error) {
	return &pb.UpdateNormalHttpReply{}, nil
}
func (s *NormalHttpService) DeleteNormalHttp(ctx context.Context, req *pb.DeleteNormalHttpRequest) (*pb.DeleteNormalHttpReply, error) {
	return &pb.DeleteNormalHttpReply{}, nil
}
func (s *NormalHttpService) GetNormalHttp(ctx context.Context, req *pb.GetNormalHttpRequest) (*pb.GetNormalHttpReply, error) {
	return &pb.GetNormalHttpReply{}, nil
}
func (s *NormalHttpService) ListNormalHttp(ctx context.Context, req *pb.ListNormalHttpRequest) (*pb.ListNormalHttpReply, error) {
	return &pb.ListNormalHttpReply{}, nil
}
