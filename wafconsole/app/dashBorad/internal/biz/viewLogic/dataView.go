package viewLogic

import (
	"context"
	"gorm.io/gorm"
	"log/slog"
	"wafconsole/app/dashBorad/internal/biz/iface"
	"wafconsole/app/dashBorad/internal/data/dto"
	"wafconsole/app/dashBorad/internal/data/model"
)

type DataViewRepo interface {
	GetDayAttackCount(ctx context.Context, today string, yesterday string) (dto.AttackDayCount, error)
	GetAttackCountByTime(ctx context.Context, startTime, endTime string) ([]dto.AttackCountByTime, error)
	ListIpAddr(ctx context.Context, startTime, endTime string) ([]dto.AttackIp, error)
	GetIPToAddress(ctx context.Context, ip string) (string, error)
	iface.BaseRepo[model.SecLog]
}

type DataViewUsecase struct {
	repo DataViewRepo
}

func NewDataViewUsecase(repo DataViewRepo) *DataViewUsecase {
	return &DataViewUsecase{repo: repo}
}

// GetDayAttack 获取攻击数 , 攻击IP数
func (d *DataViewUsecase) GetDayAttack(ctx context.Context, today, yesterday string) (dto.AttackDayCount, error) {
	attackCount, err := d.repo.GetDayAttackCount(ctx, today, yesterday)
	if err != nil {
		slog.ErrorContext(ctx, "get_day_attack_count is failed", err)
		return dto.AttackDayCount{}, err
	}
	return attackCount, nil
}

// GetAttackByTime 获取攻击数 , 攻击IP数
func (d *DataViewUsecase) GetAttackByTime(ctx context.Context, startTime, endTime string) ([]dto.AttackCountByTime, error) {
	attackCount, err := d.repo.GetAttackCountByTime(ctx, startTime, endTime)
	if err != nil {
		slog.ErrorContext(ctx, "get_attack_count_by_time is failed", err)
		return []dto.AttackCountByTime{}, err
	}
	return attackCount, nil
}

// ListAttackLog 获取攻击日志
func (d *DataViewUsecase) ListAttackLog(ctx context.Context, limit, offset int64, startTime, endTime string) ([]model.SecLog, int64, error) {
	whereOptions := make([]iface.WhereOptionWithReturn, 0)
	whereOptions = append(whereOptions, func(db *gorm.DB) *gorm.DB {
		return db.Where("toStartOfDay(ctime) >= ? and toStartOfDay(ctime) <= ?", startTime, endTime)
	})
	total, err := d.repo.Count(ctx, whereOptions...)
	if err != nil {
		slog.ErrorContext(ctx, "count_attack_log is failed", err)
		return nil, total, err
	}
	attackList, err := d.repo.ListByWhere(ctx, limit, offset, whereOptions...)
	if err != nil {
		slog.ErrorContext(ctx, "list_attack_log is failed", err)
		return nil, total, err
	}
	return attackList, total, nil
}

// GetAttackLogDetail 根据日志id 获取攻击日志详情
func (d *DataViewUsecase) GetAttackLogDetail(ctx context.Context, logId string) (model.SecLog, error) {
	attackLog, err := d.repo.GetSecLog(ctx, logId)
	if err != nil {
		slog.ErrorContext(ctx, "get_attack_log_detail is failed", err)
		return model.SecLog{}, err
	}
	return attackLog, nil
}

// ListIpToAddress 获取IP地址信息
func (d *DataViewUsecase) ListIpToAddress(ctx context.Context, startTime, endTime string) ([]dto.IPToAddress, error) {
	ipList, err := d.repo.ListIpAddr(ctx, startTime, endTime)
	if err != nil {
		slog.ErrorContext(ctx, "list_ip_to_address is failed", err)
		return nil, err
	}
	var ipToAddress []dto.IPToAddress
	for _, ip := range ipList {
		address, err := d.repo.GetIPToAddress(ctx, ip.ClientIp)
		if err != nil {
			slog.ErrorContext(ctx, "get_ip_to_address is failed", err)
			return nil, err
		}
		tempArr := dto.IPToAddress{
			Addr:  address,
			Count: ip.Count,
		}
		ipToAddress = append(ipToAddress, tempArr)
	}
	return ipToAddress, nil
}
