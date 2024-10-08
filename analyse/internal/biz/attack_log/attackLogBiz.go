package attack_log

import (
	"analyse/internal/biz/iface"
	"analyse/internal/data/model"
)

type AttackLogRepo interface {
	iface.BaseRepo[model.SecLog]
}

type AttackLogUsercase struct {
	repo AttackLogRepo
}

func NewAttackLogUsercase(repo AttackLogRepo) *AttackLogUsercase {
	return &AttackLogUsercase{repo: repo}
}

// ConsumerAttackEvents kafka消费攻击事件
func (a *AttackLogUsercase) ConsumerAttackEvents() func() {
	return a.repo.Consumer()
}
