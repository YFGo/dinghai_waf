package allow

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log/slog"
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"
)

type ListAllowRepo interface {
	iface.BaseRepo[model.AllowList]
}

type ListAllowUsecase struct {
	repo ListAllowRepo
}

func NewListAllowUsecase(repo ListAllowRepo) *ListAllowUsecase {
	return &ListAllowUsecase{repo: repo}
}

// checkAllowName 检查白名单昵称是否存在
func (uc *ListAllowUsecase) checkAllowName(ctx context.Context, id int64, name string) bool {
	_, err := uc.repo.GetByNameAndID(ctx, name, id)
	if err == nil { //查询到数据 , 昵称重复
		return true
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false //没有查询到数据
	}
	return true
}

// CreateAllow 新增白名单
func (uc *ListAllowUsecase) CreateAllow(ctx context.Context, allowInfo model.AllowList) error {
	// 检查昵称是否重复
	if uc.checkAllowName(ctx, 0, allowInfo.Name) {
		return status.Error(codes.AlreadyExists, "昵称重复")
	}
	_, err := uc.repo.Create(ctx, allowInfo)
	if err != nil {
		slog.ErrorContext(ctx, "create_allow is failed", err)
		return err
	}
	return nil
}

// UpdateAllow 修改白名单
func (uc *ListAllowUsecase) UpdateAllow(ctx context.Context, id int64, allowInfo model.AllowList) error {
	// 检查昵称是否重复
	if uc.checkAllowName(ctx, id, allowInfo.Name) {
		return status.Error(codes.AlreadyExists, "昵称重复")
	}
	err := uc.repo.Update(ctx, id, allowInfo)
	if err != nil {
		slog.ErrorContext(ctx, "update_allow is failed", err)
		return err
	}
	return nil
}

// DeleteAllow 删除白名单
func (uc *ListAllowUsecase) DeleteAllow(ctx context.Context, ids []int64) error {
	_, err := uc.repo.Delete(ctx, ids)
	if err != nil {
		slog.ErrorContext(ctx, "delete_allow is failed", err)
		return err
	}
	return nil
}

// GetAllowDetail 获取白名单详情
func (uc *ListAllowUsecase) GetAllowDetail(ctx context.Context, id int64) (model.AllowList, error) {
	allowDetailInfo, err := uc.repo.Get(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "get_allow_detail is failed", err)
		return model.AllowList{}, err
	}
	return allowDetailInfo, nil
}

// ListAllow 获取白名单列表
func (uc *ListAllowUsecase) ListAllow(ctx context.Context, limit, offset int64, name string) (int64, []model.AllowList, error) {
	withOptions := make([]iface.WhereOptionWithReturn, 0)
	if len(name) != 0 {
		withOptions = append(withOptions, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", "%"+name+"%")
		})
	}
	total, err := uc.repo.Count(ctx, withOptions...) //获取符合条件的总数据
	if err != nil {
		slog.ErrorContext(ctx, "list_allow_count is failed", err)
		return total, nil, err
	}
	listAllow, err := uc.repo.ListByWhere(ctx, limit, offset, withOptions...)
	if err != nil {
		slog.ErrorContext(ctx, "list_allow_info is failed", err)
		return total, nil, err
	}
	return total, listAllow, nil
}
