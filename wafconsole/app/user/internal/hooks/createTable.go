package hooks

import (
	"gorm.io/gorm"
	"wafconsole/app/user/internal/data/model"
)

func CreateTable(db *gorm.DB) {
	if !db.Migrator().HasTable(&model.UserInfo{}) {
		db.AutoMigrate(model.UserInfo{})
	}
}
