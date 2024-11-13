package siteBiz

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"wafconsole/app/wafTop/internal/biz/allow"
	strategyBiz "wafconsole/app/wafTop/internal/biz/strategy"

	"gorm.io/gorm"
	"log/slog"
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"
)

const (
	serverAddrKey  = "_real"
	serverStrategy = "_strategy"
	serverAllow    = "_allow"
	cutOff         = "_"
)

// ServerRepo 服务器上层实现
type ServerRepo interface {
	iface.BaseRepo[model.ServerWaf]
	GetServerStrategiesID(ctx context.Context, id int64) ([]int64, error)
	SaveServerToEtcd(ctx context.Context, serverStrategiesKey, serverRealAddrKey, serverStrategies, serverRealAddr, serverAllowKey, serverAllowValue string) error
	DeleteServerToEtcd(ctx context.Context, serverStrategiesKey, serverRealAddrKey, serverAllowKey string) error
	ListHostByIds(ctx context.Context, ids []int64) ([]string, error)
	GetServerAllowIDList(ctx context.Context, id int64) ([]int64, error)
}

type ServerUsecase struct {
	repo         ServerRepo
	allowRepo    allow.ListAllowRepo
	strategyRepo strategyBiz.WafStrategyRepo
	appRepo      WafAppRepo
}

func NewServerUsecase(repo ServerRepo, appRepo WafAppRepo, allowRepo allow.ListAllowRepo, strategyRepo strategyBiz.WafStrategyRepo) *ServerUsecase {
	return &ServerUsecase{
		repo:         repo,
		appRepo:      appRepo,
		allowRepo:    allowRepo,
		strategyRepo: strategyRepo,
	}
}

// checkAllowExists 检查白名单是否存在
func (s *ServerUsecase) checkAllowExists(ctx context.Context, ids []int64) bool {
	for _, id := range ids {
		_, err := s.allowRepo.Get(ctx, id)
		if errors.Is(err, gorm.ErrRecordNotFound) { //白名单不存在
			return false
		}
		if err != nil {
			slog.ErrorContext(ctx, "get allow failed: ", err, "id", id)
			return false
		}
	}
	return true //所有选择的白名单均存在
}

// checkStrategyExists 检查策略是否存在
func (s *ServerUsecase) checkStrategyExists(ctx context.Context, ids []int64) bool {
	for _, id := range ids {
		_, err := s.strategyRepo.Get(ctx, id)
		if errors.Is(err, gorm.ErrRecordNotFound) { //策略不存在
			return false
		}
		if err != nil {
			slog.ErrorContext(ctx, "get strategy failed: ", err, "id", id)
			return false
		}
	}
	return true //所有选择的策略均存在
}

// checkServerExist 检查服务器是否存在
func (s *ServerUsecase) checkServerExist(ctx context.Context, name string, id int64) bool {
	_, err := s.repo.GetByNameAndID(ctx, name, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false //服务器不存在
	}
	if err != nil {
		slog.ErrorContext(ctx, "get server failed: ", err, "name", name)
		return true
	}
	return true
}

// updateServerInfoEtcd 将服务器信息整理存入etcd
func (s *ServerUsecase) updateServerInfoEtcd(ctx context.Context, serverInfo model.ServerWaf) error {
	serverStrategiesKey := serverInfo.Host + serverStrategy //站点和策略的key
	serverAllowKey := serverInfo.Host + serverAllow         //站点和白名单的key
	var (
		serverStrategies string
		serverAllowValue string
	)
	for i := 0; i < len(serverInfo.StrategiesID); i++ { //拼接策略
		if i == len(serverInfo.StrategiesID)-1 {
			serverStrategies += strconv.Itoa(int(serverInfo.StrategiesID[i]))
		} else {
			serverStrategies += strconv.Itoa(int(serverInfo.StrategiesID[i])) + cutOff
		}
	}
	for i := 0; i < len(serverInfo.AllowListID); i++ {
		if i == len(serverInfo.AllowListID)-1 {
			serverAllowValue += strconv.Itoa(int(serverInfo.AllowListID[i]))
		} else {
			serverAllowValue += strconv.Itoa(int(serverInfo.AllowListID[i])) + cutOff
		}
	}
	serverRealAddrKey := serverStrategiesKey + serverAddrKey // 站点真实地址
	serverRealAddrValue := serverInfo.IP + ":" + strconv.Itoa(serverInfo.Port)
	if err := s.repo.SaveServerToEtcd(ctx, serverStrategiesKey, serverRealAddrKey, serverStrategies, serverRealAddrValue, serverAllowKey, serverAllowValue); err != nil { //存储站点对应的关系
		slog.ErrorContext(ctx, "save server to etcd failed: ", err, "server_info", serverInfo)
		return err
	}
	return nil
}

// CreateServerSite 新增服务器站点
func (s *ServerUsecase) CreateServerSite(ctx context.Context, serverInfo model.ServerWaf) error {
	// 1. 检测
	if s.checkServerExist(ctx, serverInfo.Name, 0) {
		return status.Error(codes.AlreadyExists, "服务器已存在")
	}
	if !s.checkAllowExists(ctx, serverInfo.AllowListID) {
		return status.Error(codes.NotFound, "白名单不存在")
	}
	if !s.checkStrategyExists(ctx, serverInfo.StrategiesID) {
		return status.Error(codes.NotFound, "策略不存在")
	}
	// 2. 新增服务器
	if _, err := s.repo.Create(ctx, serverInfo); err != nil {
		slog.ErrorContext(ctx, "create server failed: ", err, "server_info", serverInfo)
		return err
	}
	// 3. 保存服务器到etcd
	if err := s.updateServerInfoEtcd(ctx, serverInfo); err != nil {
		slog.ErrorContext(ctx, "update server to etcd failed: ", err, "server_info", serverInfo)
		return err
	}
	return nil
}

// UpdateServerSite 修改服务器站点
func (s *ServerUsecase) UpdateServerSite(ctx context.Context, id int64, serverInfo model.ServerWaf) error {
	// 1. 检测服务器名称是否重复
	if s.checkServerExist(ctx, serverInfo.Name, id) {
		return status.Error(codes.AlreadyExists, "服务器已存在")
	}
	if !s.checkAllowExists(ctx, serverInfo.AllowListID) {
		return status.Error(codes.NotFound, "白名单不存在")
	}
	if !s.checkStrategyExists(ctx, serverInfo.StrategiesID) {
		return status.Error(codes.NotFound, "策略不存在")
	}
	if err := s.repo.Update(ctx, id, serverInfo); err != nil {
		slog.ErrorContext(ctx, "update server failed: ", err, "server_info", serverInfo)
		return err
	}
	// 3. 保存服务器到etcd
	if err := s.updateServerInfoEtcd(ctx, serverInfo); err != nil {
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
	hosts, err := s.repo.ListHostByIds(ctx, ids)
	if err != nil {
		slog.ErrorContext(ctx, "get server host failed: ", err, "ids", ids)
		return err
	}
	for _, host := range hosts {
		if err = s.repo.DeleteServerToEtcd(ctx, host+serverStrategy, host+serverAddrKey, host+serverStrategy); err != nil { // 删除etcd中 站点应用的策略 , 真实地址 , 白名单
			slog.ErrorContext(ctx, "delete server to etcd failed: ", err, "host", host)
			return err
		}
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
	strategiesID, err := s.repo.GetServerStrategiesID(ctx, id) // 获取应用的策略id
	if err != nil {
		slog.ErrorContext(ctx, "get server strategies failed: ", err, "id", id)
		return model.ServerWaf{}, nil, err
	}
	serverWaf.StrategiesID = strategiesID
	allowListIds, err := s.repo.GetServerAllowIDList(ctx, id) // 获取应用的白名单id
	if err != nil {
		slog.ErrorContext(ctx, "get server allow failed: ", err, "id", id)
		return model.ServerWaf{}, nil, err
	}
	serverWaf.AllowListID = allowListIds
	// 获取对应的app信息
	appInfo, err := s.appRepo.GetAppWafByServerId(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.ErrorContext(ctx, "get app failed: ", err, "id", id)
		return model.ServerWaf{}, nil, err
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
