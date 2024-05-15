package bootstrap

import (
	"path"
	"template-project-name/internal/types"
	"template-project-name/internal/utils"
)

var (
	AppLog      *types.AppLogT
	_logConfs   LogConfsT
	_appLogConf AppLogConfT
)

type LogConfsT struct {
	AppLogConf AppLogConfT `toml:"app_log"`
}

type AppLogConfT struct {
	OutputFile    string `toml:"output_file"`
	OutputConsole bool   `toml:"coutput_console"`
	Level         string `toml:"level"`
	Format        string `toml:"format"` // json or text
}

func loadlogConfigs() LogConfsT {
	confPath := path.Join(CONFIG_DIR, "log.toml")
	config := utils.ReadTomlConfig[LogConfsT](confPath)

	return config
}

func InitLog() {
	_logConfs = loadlogConfigs()
	// fmt.Printf("log conf all item: %#v\n", _logConfs)
	_appLogConf = _logConfs.AppLogConf

	AppLog = types.NewAppLog(_appLogConf.Level,
		_appLogConf.OutputFile, _appLogConf.OutputConsole, _appLogConf.Format)
}

func GetLogListConfs() LogConfsT {
	return _logConfs
}

func GetAppLogConf() AppLogConfT {
	return _appLogConf
}
