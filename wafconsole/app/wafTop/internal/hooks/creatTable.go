package hooks

import (
	"gorm.io/gorm"
	"wafconsole/app/wafTop/internal/data/model"
)

func CreateTable(db *gorm.DB) {
	if !db.Migrator().HasTable(&model.AppWaf{}) {
		db.AutoMigrate(model.AppWaf{})
	}
	if !db.Migrator().HasTable(&model.BuildinRule{}) {
		db.AutoMigrate(model.BuildinRule{})
	}
	if !db.Migrator().HasTable(&model.RuleGroup{}) {
		db.AutoMigrate(model.RuleGroup{})
	}
	if !db.Migrator().HasTable(&model.ServerStrategies{}) {
		db.AutoMigrate(model.ServerStrategies{})
	}
	if !db.Migrator().HasTable(&model.ServerWaf{}) {
		db.AutoMigrate(model.ServerWaf{})
	}
	if !db.Migrator().HasTable(&model.Strategy{}) {
		db.AutoMigrate(model.Strategy{})
	}
	if !db.Migrator().HasTable(&model.StrategyConfig{}) {
		db.AutoMigrate(model.StrategyConfig{})
	}
	if !db.Migrator().HasTable(&model.UserRule{}) {
		db.AutoMigrate(model.UserRule{})
	}
}
