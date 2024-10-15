package siteBiz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"
	"log/slog"
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"
)

// ServerRepo 服务器上层实现
type ServerRepo interface {
	iface.BaseRepo[model.ServerWaf]
}

type ServerUsecase struct {
	repo ServerRepo
}

func NewServerUsecase(repo ServerRepo) *ServerUsecase {
	return &ServerUsecase{repo: repo}
}

func (s *ServerUsecase) GetServerInfoByName(ctx context.Context, name string) (model.ServerWaf, error) {
	return s.repo.GetByName(ctx, name)
}

// CreateServerSite 新增服务器站点
func (s *ServerUsecase) CreateServerSite(ctx context.Context, serverInfo model.ServerWaf) error {
	// 1. 检测服务器名称是否重复
	_, err := s.GetServerInfoByName(ctx, serverInfo.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.ErrorContext(ctx, "get server failed: ", err, "server_info", serverInfo)
		return err
	}
	if err == nil { //说明存在昵称重复的情况 , 禁止插入
		return nil
	}
	// 2. 新增服务器
	if _, err = s.repo.Create(ctx, serverInfo); err != nil {
		slog.ErrorContext(ctx, "create server failed: ", err, "server_info", serverInfo)
		return err
	}
	return nil
}
