package bootstrap

import (
	"template-project-name/internal/utils"
	"path"
	"sync"
)

var (
	_dbInitOnce sync.Once
)

type databasesConfT struct {
	MysqlConfs []mysqlConfT `toml:"mysql"`
}

func DBOnceInit() {
	_dbInitOnce.Do(initDbConnect)
}

func initDbConnect() {
	// 加载所有数据库配置项
	allDBConfigs := loadDatabasesConfigs()

	// 初始化mysql的配置连接
	initMysqlDBConnect(allDBConfigs.MysqlConfs)
}

// read  database.yaml config and set var
func loadDatabasesConfigs() databasesConfT {
	confPath := path.Join(CONFIG_DIR, "database.toml")
	config := utils.ReadTomlConfig[databasesConfT](confPath)

	return config
}
