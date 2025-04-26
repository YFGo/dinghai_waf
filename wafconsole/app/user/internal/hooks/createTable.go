package hooks

import (
	"gorm.io/gorm"
	"wafconsole/app/user/internal/data/types"
)

func CreateTable(db *gorm.DB) {
	if !db.Migrator().HasTable(&types.UserInfo{}) {
		db.AutoMigrate(types.UserInfo{})
	}
}
