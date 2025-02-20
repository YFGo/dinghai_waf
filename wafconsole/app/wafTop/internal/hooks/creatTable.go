package hooks

import (
	"gorm.io/gorm"
	"wafconsole/app/wafTop/internal/data/model"
)

func CreateTable(db *gorm.DB) {
	if !db.Migrator().HasTable(&model.AppWaf{}) {
		if err := db.AutoMigrate(model.AppWaf{}); err != nil {
			panic(err)
		}
	}
	if !db.Migrator().HasTable(&model.BuildinRule{}) {
		if err := db.AutoMigrate(model.BuildinRule{}); err != nil {
			panic(err)
		}
		if err := db.Exec("INSERT INTO waf_buildin_rule (id , created_at, updated_at, name, description, group_id, seclang  , risk_level) VALUES (90001 , NOW() , NOW() , 'coraza基础配置', 'waf内核的基础配置文件' , 1 , 'Include wafcoraza/ruleset/coraza.conf'  , 4),(90002 , NOW() , NOW() , 'coraza的内置防护规则1' , 'coraza内置防护规则的文件' , 1 , 'Include wafcoraza/ruleset/coreruleset/crs-setup.conf.example' , 4),(90003 , NOW() , NOW() , 'coraza的内置防护规则2' , 'coraza内置防护规则文件' , 1 , 'Include wafcoraza/ruleset/coreruleset/rules/*.conf'  , 4);").Error; err != nil {
			panic(err)
		}
	}
	if !db.Migrator().HasTable(&model.RuleGroup{}) {
		if err := db.AutoMigrate(model.RuleGroup{}); err != nil {
			panic(err)
		}
		if err := db.Exec("INSERT INTO waf_rule_group (id , created_at, updated_at, name, description, is_buildin) VALUES (1 , NOW() , NOW() , 'coraza的内置防护规则' , 'coraza内置防护规则' , 1)").Error; err != nil {
			panic(err)
		}
	}
	if !db.Migrator().HasTable(&model.ServerStrategies{}) {
		if err := db.AutoMigrate(model.ServerStrategies{}); err != nil {
			panic(err)
		}
	}
	if !db.Migrator().HasTable(&model.ServerWaf{}) {
		if err := db.AutoMigrate(model.ServerWaf{}); err != nil {
			panic(err)
		}
	}
	if !db.Migrator().HasTable(&model.Strategy{}) {
		if err := db.AutoMigrate(model.Strategy{}); err != nil {
			panic(err)
		}
	}
	if !db.Migrator().HasTable(&model.StrategyConfig{}) {
		if err := db.AutoMigrate(model.StrategyConfig{}); err != nil {
			panic(err)
		}
	}
	if !db.Migrator().HasTable(&model.UserRule{}) {
		if err := db.AutoMigrate(model.UserRule{}); err != nil {
			panic(err)
		}
	}
	if !db.Migrator().HasTable(&model.AllowList{}) {
		if err := db.AutoMigrate(model.AllowList{}); err != nil {
			panic(err)
		}
	}
	if !db.Migrator().HasTable(&model.ServerAllow{}) {
		if err := db.AutoMigrate(&model.ServerAllow{}); err != nil {
			panic(err)
		}
	}
}
