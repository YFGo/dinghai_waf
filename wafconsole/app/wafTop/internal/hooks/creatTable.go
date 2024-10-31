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
		if err := db.Exec("INSERT INTO waf_buildin_rule (id , created_at, updated_at, name, description, group_id, seclang  , risk_level) VALUES (90001 , NOW() , NOW() , 'coraza基础配置', 'waf内核的基础配置文件' , 1 , 'Include wafCoraza/ruleset/coraza.conf'  , 4),(90002 , NOW() , NOW() , 'coraza的内置防护规则1' , 'coraza内置防护规则的文件' , 1 , 'Include wafCoraza/ruleset/coreruleset/crs-setup.conf.example' , 4),(90003 , NOW() , NOW() , 'coraza的内置防护规则2' , 'coraza内置防护规则文件' , 1 , 'Include wafCoraza/ruleset/coreruleset/rules/*.conf'  , 4);").Error; err != nil {
			panic(err)
		}
	}
	if !db.Migrator().HasTable(&model.RuleGroup{}) {
		db.AutoMigrate(model.RuleGroup{})
		if err := db.Exec("INSERT INTO waf_rule_group (id , created_at, updated_at, name, description, is_buildin) VALUES (1 , NOW() , NOW() , 'coraza的内置防护规则' , 'coraza内置防护规则' , 1)").Error; err != nil {
			panic(err)
		}
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
