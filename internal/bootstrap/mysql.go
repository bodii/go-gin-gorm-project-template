package bootstrap

import (
	"fmt"
	"log"
	"os"
	"time"

	"template-project-name/internal/types"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	mysqlDBs map[string]*types.MysqlDBT = make(map[string]*types.MysqlDBT)
)

// db mysql config struct type
type mysqlConfT struct {
	DriverName  string `toml:"driver_name"`
	Host        string `toml:"host"`
	Port        int    `toml:"port"`
	User        string `toml:"user"`
	Password    string `toml:"password"`
	Database    string `toml:"database"`
	Charset     string `toml:"charset"`
	MaxLineNums int    `toml:"max_line_nums"`
	ShowSQL     bool   `toml:"show_sql_log"`
	LogLevel    int    `toml:"log_level"`
	LogColor    bool   `toml:"log_color"`
}

func initMysqlDBConnect(allMysqlConf []mysqlConfT) {

	for _, conf := range allMysqlConf {
		dsn := fmt.Sprintf(
			"%s:%s@(%s:%d)/%s?charset=%s",
			conf.User,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.Database,
			conf.Charset,
		)

		gormConfig := &gorm.Config{}

		// 是否输出日志
		if conf.ShowSQL {
			var logLevel logger.LogLevel
			switch conf.LogLevel {
			case 1:
				logLevel = logger.Silent
			case 2:
				logLevel = logger.Error
			case 3:
				logLevel = logger.Warn
			case 4:
				logLevel = logger.Info
			default:
				logLevel = logger.Silent
			}

			gormConfig.Logger = logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold: time.Second, // Slow SQL threshold
					LogLevel:      logLevel,    // Log level
					// IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
					// ParameterizedQueries:      true,        // Don't include params in the SQL log
					Colorful: conf.LogColor, // use color
				},
			)
		}

		db, err := gorm.Open(
			mysql.New(mysql.Config{DSN: dsn}),
			gormConfig)

		if err != nil {
			log.Fatalf("Failed to connect to %s database!", conf.Database)
		}

		// auto sync
		// DB_Engine.Sync2()
		mysqlDBs[conf.Database] = &types.MysqlDBT{DB: db}

	}

	log.Println("init databases success!")
}

func GetMysqlDB(dbName string) (*types.MysqlDBT, error) {
	db, ok := mysqlDBs[dbName]
	if !ok {
		return nil, fmt.Errorf(
			"the currently specified database name '%s' does not exist",
			dbName)
	}

	return db, nil
}

func GetMysqlAllDBNames() []string {
	all := []string{}
	for dbname := range mysqlDBs {
		all = append(all, dbname)
	}

	return all
}

func GetMysqlAllDB() map[string]*types.RedisDBT {
	return redisDBs
}
