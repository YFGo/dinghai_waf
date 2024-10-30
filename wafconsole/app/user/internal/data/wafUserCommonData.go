package data

import "wafconsole/app/user/internal/biz"

type wafUserCommonRepo struct {
	data *Data
}

func NewWafUserCommonRepo(data *Data) biz.WafUserCommonRepo {
	return &wafUserCommonRepo{data: data}
}
