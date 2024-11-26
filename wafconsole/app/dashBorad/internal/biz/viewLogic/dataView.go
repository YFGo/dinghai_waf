package viewLogic

import (
	"context"
	"log/slog"
	"wafconsole/app/dashBorad/internal/data/dto"
)

type DataViewRepo interface {
	GetDayAttackCount(ctx context.Context, today string, yesterday string) (dto.AttackDayCount, error)
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
