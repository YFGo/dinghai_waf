package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"wafconsole/app/user/internal/biz"
	"wafconsole/app/user/internal/biz/iface"
	"wafconsole/app/user/internal/data/types"
)

type wafUserRepo struct {
	data *Data
	log  *log.Helper
}

func NewWafUserRepo(data *Data, logger log.Logger) biz.WafUserRepo {
	return &wafUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (w wafUserRepo) Get(ctx context.Context, id int64) (types.UserInfo, error) {
	var userInfo types.UserInfo
	err := w.data.db.Where("id = ?", id).First(&userInfo).Error
	return userInfo, err
}

func (w wafUserRepo) GetByNameAndID(ctx context.Context, s string, i int64) (types.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (w wafUserRepo) Create(ctx context.Context, userInfo types.UserInfo) (int64, error) {
	err := w.data.db.Create(&userInfo).Error
	return int64(userInfo.ID), err
}

func (w wafUserRepo) Update(ctx context.Context, i int64, t types.UserInfo) error {
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

func (w wafUserRepo) ListByWhere(ctx context.Context, limit, offset int64, opts ...iface.WhereOptionWithReturn) ([]types.UserInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (w wafUserRepo) LoginByEmailPassword(ctx context.Context, user types.UserInfo) (types.UserInfo, error) {
	err := w.data.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	return user, err
}

func (w wafUserRepo) GetRedisValueByKey(ctx context.Context, key string) (string, error) {
	code, err := w.data.rdb.Get(ctx, key).Result()
	if err != nil {
		w.log.WithContext(ctx).Error(err)
		return "", err
	}
	return code, nil
}

func (w wafUserRepo) GetUserInfoByEmail(ctx context.Context, email string) (types.UserInfo, error) {
	var res types.UserInfo
	err := w.data.db.Where("email = ?", email).First(&res).Error
	return res, err
}
