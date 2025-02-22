package data

import (
	"context"
	"wafconsole/app/dashBorad/internal/biz/iface"
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

// GetAttackCountByTime 获取指定日期内攻击次数的变化
func (d *dataViewRepo) GetAttackCountByTime(ctx context.Context, startTime, endTime string) ([]dto.AttackCountByTime, error) {
	var ans []dto.AttackCountByTime
	// 确保startTime和endTime是正确的日期格式，例如："YYYY-MM-DD"
	err := d.data.clickhouseDB.Model(&model.SecLog{}).
		Select("count(*) as attack_count, count(distinct(client_ip)) as attack_ip_count, toStartOfDay(ctime) as time").
		Where("toStartOfDay(ctime) >= ? and toStartOfDay(ctime) <= ?", startTime, endTime).
		Order("toStartOfDay(ctime) ASC").Group("toStartOfDay(ctime)").
		Scan(&ans).Error
	return ans, err
}

// ListByWhere 根据条件获取攻击信息
func (d *dataViewRepo) ListByWhere(ctx context.Context, limit, offset int64, opts ...iface.WhereOptionWithReturn) ([]model.SecLog, error) {
	var ans []model.SecLog
	db := d.data.clickhouseDB.Select("log_id , uri , ctime , rule_name")
	for _, opt := range opts {
		db = opt(db)
	}
	err := db.Limit(int(limit)).Offset(int(offset)).Find(&ans).Error
	return ans, err
}

// Count 根据条件 获取次数
func (d *dataViewRepo) Count(ctx context.Context, withReturn ...iface.WhereOptionWithReturn) (int64, error) {
	var total int64
	db := d.data.clickhouseDB.Table(model.SecLogTableName)
	for _, opt := range withReturn {
		db = opt(db)
	}
	err := db.Count(&total).Error
	return total, err
}

// GetSecLog 获取日志文件
func (d *dataViewRepo) GetSecLog(ctx context.Context, logId string) (model.SecLog, error) {
	var ans model.SecLog
	err := d.data.clickhouseDB.Where("log_id = ?", logId).First(&ans).Error
	return ans, err
}

// ListIpAddr 获取指定时间内的攻击IP
func (d *dataViewRepo) ListIpAddr(ctx context.Context, startTime, endTime string) ([]dto.AttackIp, error) {
	var ans []dto.AttackIp
	err := d.data.clickhouseDB.Model(&model.SecLog{}).Select("client_ip , count(client_ip) as count").Where("date(ctime) >= ? and date(ctime) <= ?", startTime, endTime).Group("client_ip").Find(&ans).Error
	return ans, err
}

// GetIPToAddress 获取IP对应的地址
func (d *dataViewRepo) GetIPToAddress(ctx context.Context, ip string) (string, error) {
	region, err := d.data.ipDB.SearchByStr(ip)
	return region, err
}
