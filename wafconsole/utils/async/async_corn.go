package utils

import "context"

// AsyncCorn 异步执行的定时任务
type AsyncCorn struct {
	ctx context.Context
}

// AsyncCornSaveAttackLog 将kafka中的攻击日志异步保存到clickhouse中
func (a *AsyncCorn) AsyncCornSaveAttackLog() {

}
