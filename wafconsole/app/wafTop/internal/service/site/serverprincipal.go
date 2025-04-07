package service

import (
	"context"
	pb "wafconsole/api/wafTop/v1"
	siteBiz "wafconsole/app/wafTop/internal/biz/site"
	"wafconsole/app/wafTop/internal/data/model"
)

type ServerPrincipalService struct {
	pb.UnimplementedServerPrincipalServer
	uc *siteBiz.ServerPrincipalUsecase
}

func NewServerPrincipalService(uc *siteBiz.ServerPrincipalUsecase) *ServerPrincipalService {
	return &ServerPrincipalService{
		uc: uc,
	}
}

// GetServerPrincipalByServerId 获取对应网站负责人id
func (s *ServerPrincipalService) GetServerPrincipalByServerId(ctx context.Context, req *pb.GetServerPrincipalRequest) (*pb.GetServerPrincipalReply, error) {
	principalInfoList, err := s.uc.GetServerPrincipalByServerId(ctx, req.ServerId)
	if err != nil {
		return nil, err
	}
	resp := make([]*pb.ServerPrincipalInfo, 0)
	for _, principalInfo := range principalInfoList {
		resp = append(resp, &pb.ServerPrincipalInfo{
			Id:                principalInfo.ID,
			PrincipalType:     principalInfo.PrincipalType,
			PrincipalName:     principalInfo.PrincipalName,
			PrincipalPosition: principalInfo.PrincipalPosition,
			NoticeInfo:        principalInfo.NoticeInfo,
		})
	}
	return &pb.GetServerPrincipalReply{
		ServerPrincipalInfos: resp,
	}, nil
}

func (s *ServerPrincipalService) CreateServerPrincipal(ctx context.Context, req *pb.CreateServerPrincipalRequest) (*pb.CreateServerPrincipalReply, error) {
	principalList := make([]model.PrincipalInfo, 0, len(req.PrincipalInfos))
	for _, principal := range req.PrincipalInfos {
		principalList = append(principalList, model.PrincipalInfo{
			ID:                0,
			PrincipalType:     principal.PrincipalType,
			PrincipalName:     principal.PrincipalName,
			PrincipalPosition: principal.PrincipalPosition,
			NoticeInfo:        principal.NoticeInfo,
		})
	}
	if err := s.uc.CreatePrincipal(ctx, req.ServerId, principalList); err != nil {
		return nil, err
	}
	return &pb.CreateServerPrincipalReply{}, nil
}

func (s *ServerPrincipalService) UpdateServerPrincipal(ctx context.Context, req *pb.UpdateServerPrincipalRequest) (*pb.UpdateServerPrincipalReply, error) {
	return &pb.UpdateServerPrincipalReply{}, nil
}

func (s *ServerPrincipalService) DeleteServerPrincipal(ctx context.Context, req *pb.DeleteServerPrincipalRequest) (*pb.DeleteServerPrincipalReply, error) {
	return &pb.DeleteServerPrincipalReply{}, nil
}
