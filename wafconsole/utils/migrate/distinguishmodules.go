package migrate

import (
	"bufio"
	"fmt"
	"os"
	"wafconsole/utils/const/common"
	"wafconsole/utils/const/waftop"
)

// ModuleMigrateInfo 模块信息
type ModuleMigrateInfo struct {
	ModuleName  string
	MigrateType string
}

type ModuleMigrator interface {
	GetModuleMigratePath(moduleMigrateInfo ModuleMigrateInfo) string
}

// WafTopMigratePath wafTop模块迁移路径
type WafTopMigratePath struct{}

func (w WafTopMigratePath) GetModuleMigratePath(moduleMigrateInfo ModuleMigrateInfo) string {
	if moduleMigrateInfo.MigrateType == common.MySqlMigrate {
		return waftop.MigratePathMysql
	}

	return waftop.MigratePathCk
}

func GetModuleMigratePath(moduleType string) ModuleMigrator {
	switch moduleType {
	case waftop.WafTop:
		return WafTopMigratePath{}
	default:
		return nil
	}
}

// DistinguishModules 区分不同模块的迁移 , 返回指定模块的目录
func DistinguishModules(migrateTxtPath string) (map[string]string, error) {
	// 读取txt文件内容
	migrateTxtFile, err := os.Open(migrateTxtPath)
	if err != nil {
		return nil, err
	}
	defer migrateTxtFile.Close()
	// 按行解析txt文件内容，获取模块名称
	sacnner := bufio.NewScanner(migrateTxtFile)
	for sacnner.Scan() {
		line := sacnner.Text()
		fmt.Println(line)
		GetModuleMigratePath(line)
	}
	// 根据模块名称，返回指定模块的目录
	return nil, nil
}
