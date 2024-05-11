package bootstrap

const (
	CONFIG_DIR string = "internal/config"
)

func Init() {
	// 初始化系统配置项 根目录.env file
	LoadEnv()

	// 初始化日志
	InitLog()

	// 创建数据库链接
	RedisOnceInit()

	// 创建数据库链接
	DBOnceInit()
}
