package service

import (
	"context"

	pb "wafconsole/api/wafTop/v1"
	"wafconsole/app/wafTop/internal/biz/normalhttp"
	"wafconsole/app/wafTop/internal/data/model"
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

func (s *NormalHttpService) CreateNormalHttpBatch(ctx context.Context, req *pb.CreateNormalHttpRequest) (*pb.CreateNormalHttpReply, error) {
	var normalHttpInfoList []model.NormalHttpModel
	for _, normalHttpInfo := range req.NormalHttpInfos {
		normalHttpInfoList = append(normalHttpInfoList, model.NormalHttpModel{
			ID:            normalHttpInfo.Id,
			IP:            normalHttpInfo.Ip,
			RequestURI:    normalHttpInfo.RequestUri,
			RequestMethod: normalHttpInfo.RequestMethod,
			RequestTime:   normalHttpInfo.RequestTime.AsTime(),
			Protocol:      normalHttpInfo.Protocol,
			RequestBody:   normalHttpInfo.RequestBody,
		})
	}
	err := s.uc.SaveNormalHttpInfo(ctx, normalHttpInfoList)
	if err != nil {
		return nil, err
	}
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
