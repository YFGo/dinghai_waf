package data

import (
	"context"
	"gorm.io/gorm"
	"wafconsole/app/wafTop/internal/biz/iface"
	siteBiz "wafconsole/app/wafTop/internal/biz/site"
	"wafconsole/app/wafTop/internal/data/model"
)

type serverPrincipalRepo struct {
	data *Data
}

func NewServerPrincipalRepo(data *Data) siteBiz.ServerPrincipalRepo {
	return &serverPrincipalRepo{
		data: data,
	}
}

// Count 实现siteBiz.ServerPrincipalRepo接口
func (s *serverPrincipalRepo) Count(ctx context.Context, opts ...iface.WhereOptionWithReturn) (int64, error) {
	// 实现Count方法逻辑
	return 0, nil
}

func (s *serverPrincipalRepo) Get(ctx context.Context, i int64) (model.PrincipalInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serverPrincipalRepo) GetByNameAndID(ctx context.Context, name string,
	i int64) (model.PrincipalInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serverPrincipalRepo) Create(ctx context.Context, principalInfo model.PrincipalInfo) (
	int64, error) {
	panic("implement me")
}

func (s *serverPrincipalRepo) Update(ctx context.Context, i int64, t model.PrincipalInfo) error {
	//TODO implement me
	panic("implement me")
}

func (s *serverPrincipalRepo) Delete(ctx context.Context, int64s []int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *serverPrincipalRepo) ListByWhere(ctx context.Context, limit, offset int64,
	opts ...iface.WhereOptionWithReturn) ([]model.PrincipalInfo, error) {
	principalInfoList := make([]model.PrincipalInfo, 0)
	db := s.data.db
	for _, opt := range opts {
		db.Where(opt)
	}
	err := db.Find(&principalInfoList).Error
	if err != nil {
		return nil, err
	}
	return principalInfoList, nil
}

// GetPrincipalIdsByServerId 获取网站负责人id
func (s *serverPrincipalRepo) GetPrincipalIdsByServerId(ctx context.Context,
	serverId int64) ([]int64, error) {
	principalIDs := make([]int64, 0)
	err := s.data.db.Model(&model.ServerPrincipal{}).Where("server_id = ?", serverId).Find(&principalIDs).Error
	if err != nil {
		return nil, err
	}
	return principalIDs, nil
}

func (s *serverPrincipalRepo) CreatePrincipalBatch(ctx context.Context, serverId int64,
	principalList []model.PrincipalInfo) error {

	err := s.data.db.Transaction(func(tx *gorm.DB) error {
		// 1. 先插入主表数据
		if err := tx.Create(principalList).Error; err != nil {
			return err
		}
		serverPrincipalList := make([]model.ServerPrincipal, 0)
		for _, principal := range principalList {

			serverPrincipal := model.ServerPrincipal{
				ServerID:    serverId,
				PrincipalID: principal.ID,
			}
			serverPrincipalList = append(serverPrincipalList, serverPrincipal)
		}
		// 2. 插入中间表
		if err := tx.Create(&serverPrincipalList).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
