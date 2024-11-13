package allow

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log/slog"
	"strconv"
	"wafconsole/app/wafTop/internal/biz/iface"
	"wafconsole/app/wafTop/internal/data/model"
)

const (
	allowPrefix = "allow_"
)

type ListAllowRepo interface {
	iface.BaseRepo[model.AllowList]
	SaveAllowToEtcd(ctx context.Context, key, value string) error
	DeleteAllowFromEtcd(ctx context.Context, key string) error
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

// saveAllowToEtcd 保存白名单到etcd
func (uc *ListAllowUsecase) saveAllowToEtcd(ctx context.Context, id int64, allowInfo model.AllowList) error {
	key := allowPrefix + strconv.Itoa(int(id))
	etcdAllow := struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{
		Key:   allowInfo.Key,
		Value: allowInfo.Value,
	}
	value, err := json.Marshal(&etcdAllow)
	if err != nil {
		slog.ErrorContext(ctx, "save_allow_to_etcd json marshal is failed: ", err)
		return err
	}
	err = uc.repo.SaveAllowToEtcd(ctx, key, string(value))
	if err != nil {
		slog.ErrorContext(ctx, "save_allow_to_etcd is failed", err)
		return err
	}
	return nil
}

// CreateAllow 新增白名单
func (uc *ListAllowUsecase) CreateAllow(ctx context.Context, allowInfo model.AllowList) error {
	// 检查昵称是否重复
	if uc.checkAllowName(ctx, 0, allowInfo.Name) {
		return status.Error(codes.AlreadyExists, "昵称重复")
	}
	id, err := uc.repo.Create(ctx, allowInfo)
	if err != nil {
		slog.ErrorContext(ctx, "create_allow is failed", err)
		return err
	}
	// 将信息保存到etcd
	err = uc.saveAllowToEtcd(ctx, id, allowInfo)
	if err != nil {
		slog.ErrorContext(ctx, "CreateAllow save_allow_to_etcd is failed", err)
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
	// 将信息保存到etcd
	err = uc.saveAllowToEtcd(ctx, id, allowInfo)
	if err != nil {
		slog.ErrorContext(ctx, "UpdateAllow save_allow_to_etcd is failed", err)
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
	// 删除etcd中的白名单信息
	for _, id := range ids {
		key := allowPrefix + strconv.Itoa(int(id))
		err = uc.repo.DeleteAllowFromEtcd(ctx, key)
		if err != nil {
			slog.ErrorContext(ctx, "delete_allow_from_etcd is failed", err)
			return err
		}
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
