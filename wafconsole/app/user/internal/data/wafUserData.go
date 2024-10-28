package data

import (
	"context"
	"wafconsole/app/user/internal/biz"
	"wafconsole/app/user/internal/biz/iface"
	"wafconsole/app/user/internal/data/model"
)

type wafUserRepo struct {
	data *Data
}

func NewWafUserRepo(data *Data) biz.WafUserRepo {
	return &wafUserRepo{data: data}
}

func (w wafUserRepo) Get(ctx context.Context, id int64) (model.UserInfo, error) {
	var userInfo model.UserInfo
	err := w.data.db.Where("id = ?", id).First(&userInfo).Error
	return userInfo, err
}

func (w wafUserRepo) GetByNameAndID(ctx context.Context, s string, i int64) (model.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (w wafUserRepo) Create(ctx context.Context, userInfo model.UserInfo) (int64, error) {
	err := w.data.db.Create(&userInfo).Error
	return int64(userInfo.ID), err
}

func (w wafUserRepo) Update(ctx context.Context, i int64, t model.UserInfo) error {
	//TODO implement me
	panic("implement me")
}

func (w wafUserRepo) Delete(ctx context.Context, int64s []int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (w wafUserRepo) Count(ctx context.Context, withReturn ...iface.WhereOptionWithReturn) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (w wafUserRepo) ListByWhere(ctx context.Context, limit, offset int64, opts ...iface.WhereOptionWithReturn) ([]model.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (w wafUserRepo) LoginByEmailPassword(ctx context.Context, user model.UserInfo) (model.UserInfo, error) {
	err := w.data.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	return user, err
}
