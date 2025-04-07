package siteBiz

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"
)

type ServerPrincipalRepo interface {
	iface.BaseRepo[model.PrincipalInfo]
	GetPrincipalIdsByServerId(ctx context.Context, serverId int64) ([]int64, error)
	CreatePrincipalBatch(ctx context.Context, serverId int64, principalList []model.PrincipalInfo) error
}

type ServerPrincipalUsecase struct {
	repo       ServerPrincipalRepo
	serverRepo ServerRepo
	log        *log.Helper
}

func NewServerPrincipalUsecase(repo ServerPrincipalRepo, serverRepo ServerRepo, logger log.Logger) *ServerPrincipalUsecase {
	return &ServerPrincipalUsecase{
		repo:       repo,
		serverRepo: serverRepo,
		log:        log.NewHelper(logger),
	}
}

func (s *ServerPrincipalUsecase) isServerExist(ctx context.Context, serverId int64) bool {
	_, err := s.serverRepo.Get(ctx, serverId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false //服务器站点不存在
		}
		s.log.WithContext(ctx).Error(err)
		return false
	}
	return true
}

func (s *ServerPrincipalUsecase) GetServerPrincipalByServerId(ctx context.Context, serverId int64) ([]model.PrincipalInfo, error) {
	principalIds, err := s.repo.GetPrincipalIdsByServerId(ctx, serverId)
	if err != nil {
		s.log.WithContext(ctx).Errorf("get principal ids is failed:%v", err)
		return nil, err
	}
	options := make([]iface.WhereOptionWithReturn, 0)
	options = append(options, func(d *gorm.DB) *gorm.DB {
		return d.Where("id IN (?)", principalIds)
	})
	principalInfoList, err := s.repo.ListByWhere(ctx, 0, 0, options...)
	if err != nil {
		s.log.WithContext(ctx).Errorf("get list principal is failed:%v", err)
		return nil, err
	}
	return principalInfoList, nil
}

func (s *ServerPrincipalUsecase) CreatePrincipal(ctx context.Context, serverId int64, principalList []model.PrincipalInfo) error {
	if !s.isServerExist(ctx, serverId) {
		return status.Error(codes.NotFound, "serverId is not exist")
	}
	if err := s.repo.CreatePrincipalBatch(ctx, serverId, principalList); err != nil {
		s.log.WithContext(ctx).Errorf("create principal is failed:%v", err)
		return err
	}
	return nil
}
