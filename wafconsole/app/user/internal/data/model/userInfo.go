package model

import "gorm.io/gorm"

const TableNameUserInfo = "waf_user_info"

type UserInfo struct {
	*gorm.Model
	Email      string `json:"email" gorm:"type:varchar(50);not null;unique;comment:'用户邮箱'"`
	UserName   string `json:"user_name" gorm:"type:varchar(20);"`
	Password   string `json:"password" gorm:"type:varchar(20);not null"`
	AvatarAddr string `json:"avatar_addr" gorm:"type:varchar(200);comment:'头像地址'"`
	Phone      string `json:"phone"  gorm:"type:varchar(20);comment:'手机号'"`
}

func (u *UserInfo) TableName() string {
	return TableNameUserInfo
}
