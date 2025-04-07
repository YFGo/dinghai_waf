package model

import (
	"time"
)

// ServerPrincipal 服务器负责人关联表
type ServerPrincipal struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ServerID    int64     `gorm:"column:server_id;not null" json:"server_id"`
	PrincipalID int64     `gorm:"column:principal_id;not null" json:"principal_id"`
	CreateTime  time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime  time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"update_time"`
}

// TableName 设置表名
func (s *ServerPrincipal) TableName() string {
	return "server_principal"
}

// PrincipalInfo 负责人信息表
type PrincipalInfo struct {
	ID                int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PrincipalType     string    `gorm:"column:principal_type;not null;type:varchar(25)" json:"principal_type"`         // 通知网站负责人类型
	PrincipalName     string    `gorm:"column:principal_name;not null;type:varchar(35)" json:"principal_name"`         // 通知网站负责人姓名
	PrincipalPosition string    `gorm:"column:principal_position;not null;type:varchar(35)" json:"principal_position"` // 通知网站负责人职位
	NoticeInfo        string    `gorm:"column:notice_info;not null;type:varchar(35)" json:"notice_info"`               // 通知网站负责人信息
	CreateTime        time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime        time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"update_time"`

	ServerIDs []int64 `gorm:"-"`
}

// TableName 设置表名
func (p *PrincipalInfo) TableName() string {
	return "principal_info"
}
