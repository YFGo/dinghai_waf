package data

import (
	"context"
	"wafconsole/app/dashBorad/internal/biz/viewLogic"
	"wafconsole/app/dashBorad/internal/data/dto"
	"wafconsole/app/dashBorad/internal/data/model"
)

type dataViewRepo struct {
	data *Data
}

func NewDataViewRepo(data *Data) viewLogic.DataViewRepo {
	return &dataViewRepo{
		data: data,
	}
}

// GetDayAttackCount 获取昨日和今日的异常攻击 和 异常IP数量
func (d *dataViewRepo) GetDayAttackCount(ctx context.Context, today string, yesterday string) (dto.AttackDayCount, error) {
	var ans dto.AttackDayCount
	err := d.data.clickhouseDB.Model(&model.SecLog{}).Select("count(*) as attack_count , count(distinct(client_ip)) as attack_ip_count").Where("date(ctime) = ?", today).Find(&ans).Error
	if err != nil {
		return dto.AttackDayCount{}, err
	}
	err = d.data.clickhouseDB.Model(&model.SecLog{}).Select("count(*) as attack_yesterday , count(distinct(client_ip)) as attack_ip_yesterday").Where("date(ctime) = ?", yesterday).Find(&ans).Error
	return ans, err
}
