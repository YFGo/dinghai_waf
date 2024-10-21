package siteBiz

import (
	"context"
	"errors"
	"strconv"

	"gorm.io/gorm"
	"log/slog"
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"
)

const (
	serverAddrKey = "_real"
)

// ServerRepo 服务器上层实现
type ServerRepo interface {
	iface.BaseRepo[model.ServerWaf]
	GetServerStrategiesID(ctx context.Context, id int64) ([]int64, error)
	SaveServerToEtcd(ctx context.Context, serverStrategiesKey, serverRealAddrKey, serverStrategies, serverRealAddr string) error
	DeleteServerToEtcd(ctx context.Context, serverStrategiesKey, serverRealAddrKey string) error
}

type ServerUsecase struct {
	repo    ServerRepo
	appRepo WafAppRepo
}

func NewServerUsecase(repo ServerRepo, appRepo WafAppRepo) *ServerUsecase {
	return &ServerUsecase{
		repo:    repo,
		appRepo: appRepo,
	}
}

func (s *ServerUsecase) GetServerInfoByName(ctx context.Context, name string, id int64) (model.ServerWaf, error) {
	return s.repo.GetByNameAndID(ctx, name, id)
}

// UpdateServerInfoEtcd 将服务器信息整理存入etcd
func (s *ServerUsecase) UpdateServerInfoEtcd(ctx context.Context, serverInfo model.ServerWaf) error {
	serverStrategiesKey := serverInfo.Host
	var serverStrategies string
	for i := 0; i < len(serverInfo.StrategiesID); i++ {
		if i == len(serverInfo.StrategiesID)-1 {
			serverStrategies += strconv.Itoa(int(serverInfo.StrategiesID[i]))
		} else {
			serverStrategies += strconv.Itoa(int(serverInfo.StrategiesID[i])) + "_"
		}
	}
	serverRealAddrKey := serverStrategiesKey + serverAddrKey
	serverRealAddrValue := serverInfo.IP + ":" + strconv.Itoa(serverInfo.Port)
	if err := s.repo.SaveServerToEtcd(ctx, serverStrategiesKey, serverRealAddrKey, serverStrategies, serverRealAddrValue); err != nil {
		slog.ErrorContext(ctx, "save server to etcd failed: ", err, "server_info", serverInfo)
		return err
	}
	return nil
}

// CreateServerSite 新增服务器站点
func (s *ServerUsecase) CreateServerSite(ctx context.Context, serverInfo model.ServerWaf) error {
	// 1. 检测服务器名称是否重复
	_, err := s.GetServerInfoByName(ctx, serverInfo.Name, 0)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.ErrorContext(ctx, "get server failed: ", err, "server_info", serverInfo)
		return err
	}
	if err == nil { //说明存在昵称重复的情况 , 禁止插入
		return errors.New("server name is already exist")
	}
	// 2. 新增服务器
	if _, err = s.repo.Create(ctx, serverInfo); err != nil {
		slog.ErrorContext(ctx, "create server failed: ", err, "server_info", serverInfo)
		return err
	}
	// 3. 保存服务器到etcd
	if err = s.UpdateServerInfoEtcd(ctx, serverInfo); err != nil {
		slog.ErrorContext(ctx, "update server to etcd failed: ", err, "server_info", serverInfo)
		return err
	}
	return nil
}

// UpdateServerSite 修改服务器站点
func (s *ServerUsecase) UpdateServerSite(ctx context.Context, id int64, serverInfo model.ServerWaf) error {
	// 1. 检测服务器名称是否重复
	_, err := s.GetServerInfoByName(ctx, serverInfo.Name, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.ErrorContext(ctx, "get server failed: ", err, "server_info", serverInfo)
		return err
	}
	if err == nil { //说明存在昵称重复的情况 , 禁止插入
		return errors.New("server name is already exist")
	}
	if err = s.repo.Update(ctx, id, serverInfo); err != nil {
		slog.ErrorContext(ctx, "update server failed: ", err, "server_info", serverInfo)
		return err
	}
	// 3. 保存服务器到etcd
	if err = s.UpdateServerInfoEtcd(ctx, serverInfo); err != nil {
		slog.ErrorContext(ctx, "update server to etcd failed: ", err, "server_info", serverInfo)
		return err
	}
	return nil
}

// DeleteServerSite 删除服务器站点
func (s *ServerUsecase) DeleteServerSite(ctx context.Context, ids []int64) error {
	if _, err := s.repo.Delete(ctx, ids); err != nil {
		slog.ErrorContext(ctx, "delete server failed: ", err, "ids", ids)
		return err
	}
	return nil
}

// GetServerSite 获取服务器站点详情
func (s *ServerUsecase) GetServerSite(ctx context.Context, id int64) (model.ServerWaf, *model.AppWaf, error) {
	serverWaf, err := s.repo.Get(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "get server failed: ", err, "id", id)
		return model.ServerWaf{}, nil, err
	}
	strategiesID, err := s.repo.GetServerStrategiesID(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "get server strategies failed: ", err, "id", id)
		return model.ServerWaf{}, nil, err
	}
	serverWaf.StrategiesID = strategiesID
	// 获取对应的app信息
	appInfo, err := s.appRepo.GetAppWafByServerId(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.ErrorContext(ctx, "get app failed: ", err, "id", id)
		return model.ServerWaf{}, nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return serverWaf, nil, nil
	}
	return serverWaf, &appInfo, nil
}

// GetServerSiteList 获取服务器列表
func (s *ServerUsecase) GetServerSiteList(ctx context.Context, limit, offset int64, name string) ([]model.ServerWaf, int64, error) {
	whereOptions := make([]iface.WhereOptionWithReturn, 0)
	if len(name) == 0 {
		whereOptions = append(whereOptions, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", "%"+name+"%")
		})
	}
	total, err := s.repo.Count(ctx, whereOptions...)
	if err != nil {
		slog.ErrorContext(ctx, "get server count failed: ", err)
		return nil, 0, err
	}
	serverWafList, err := s.repo.ListByWhere(ctx, limit, offset, whereOptions...)
	if err != nil {
		slog.ErrorContext(ctx, "get server list failed: ", err)
		return nil, 0, err
	}
	return serverWafList, total, nil
}
