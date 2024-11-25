package data

import "wafconsole/app/dashBorad/internal/biz/viewLogic"

type dataViewRepo struct {
	data *Data
}

func NewDataViewRepo(data *Data) viewLogic.DataViewRepo {
	return &dataViewRepo{
		data: data,
	}
}
