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

// CreateServer 创建服务器站点
func (s *ServerService) CreateServer(ctx context.Context, req *pb.CreateServerRequest) (*pb.CreateServerReply, error) {
	serverInfo := model.ServerWaf{
		Name:         req.Name,
		Host:         req.Host,
		IP:           req.Ip,
		Port:         int(req.Port),
		StrategiesID: req.StrategyIds,
	}
	err := s.uc.CreateServerSite(ctx, serverInfo)
	if err != nil {
		slog.ErrorContext(ctx, "create server_waf service error: ", err)
		return nil, err
	}
	return &pb.CreateServerReply{}, nil
}

// UpdateServer 更新服务器站点
func (s *ServerService) UpdateServer(ctx context.Context, req *pb.UpdateServerRequest) (*pb.UpdateServerReply, error) {
	serverInfo := model.ServerWaf{
		Name:         req.Name,
		Host:         req.Host,
		IP:           req.Ip,
		Port:         int(req.Port),
		StrategiesID: req.StrategyIds,
	}
	err := s.uc.UpdateServerSite(ctx, req.Id, serverInfo)
	if err != nil {
		slog.ErrorContext(ctx, "update server_waf service error: ", err)
		return nil, err
	}
	return &pb.UpdateServerReply{}, nil
}

// DeleteServer 删除服务器站点
func (s *ServerService) DeleteServer(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	err := s.uc.DeleteServerSite(ctx, req.Ids)
	if err != nil {
		slog.ErrorContext(ctx, "delete server_waf service error: ", err)
		return nil, err
	}
	return &pb.DeleteReply{}, nil
}

// GetServer 获取服务器站点详情
func (s *ServerService) GetServer(ctx context.Context, req *pb.GetServerRequest) (*pb.GetServerReply, error) {
	serverInfo, appInfo, err := s.uc.GetServerSite(ctx, req.Id)
	if err != nil {
		slog.ErrorContext(ctx, "get server_waf service error: ", err)
		return nil, err
	}
	serverReply := &pb.GetServerReply{
		Name:         serverInfo.Name,
		Ip:           serverInfo.IP,
		Host:         serverInfo.Host,
		Port:         int64(serverInfo.Port),
		StrategiesId: serverInfo.StrategiesID,
	}
	if appInfo != nil {
		wafAppInfo := &pb.WafAppInfo{
			Id:   int64(appInfo.ID),
			Name: appInfo.Name,
			Url:  appInfo.Url,
		}
		serverReply.WafApps = wafAppInfo
	}
	return serverReply, nil
}

// ListServer 获取服务器站点列表
func (s *ServerService) ListServer(ctx context.Context, req *pb.ListServerRequest) (*pb.ListServerReply, error) {
	limit := req.PageSize
	offset := req.PageSize * (req.PageNow - 1)
	serverList, total, err := s.uc.GetServerSiteList(ctx, limit, offset, req.Name)
	if err != nil {
		slog.ErrorContext(ctx, "list server_waf service error: ", err)
		return nil, err
	}
	var serverInfoList []*pb.ServerInfo
	for _, server := range serverList {
		serverInfo := &pb.ServerInfo{
			Id:   int64(server.ID),
			Name: server.Name,
			Ip:   server.IP,
			Host: server.Host,
			Port: int64(server.Port),
		}
		serverInfoList = append(serverInfoList, serverInfo)
	}
	return &pb.ListServerReply{
		Total:       total,
		ListServers: serverInfoList,
	}, nil
}
