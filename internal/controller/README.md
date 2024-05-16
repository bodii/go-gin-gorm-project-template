## example

```go
package controller

import (
	"template-project-name/internal/bootstrap"
	"template-project-name/internal/types"
	"template-project-name/internal/utils"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type exampleController struct {
	DB    *types.MysqlDBT
	Redis *types.RedisDBT
	Log   *types.AppLogT
}

func NewExampleController() *exampleController {
	// log config
	logger := bootstrap.AppLog

	// db config in 'internal/config/dabasese.toml' file on [[msyql]]
	db, err := bootstrap.GetMysqlDB("dbname")
	if err != nil {
		logger.Fatal(err.Error())
	}

	// redis config in 'internal/config/redis.toml' file on [[redis-server]]
	redis, err := bootstrap.GetRedis("server1")
	if err != nil {
		logger.Fatal(err.Error())
	}

	return &exampleController{
		DB:    db,
		Redis: redis,
		Log:   logger,
	}
}

func (example *exampleController) HelloWorld(c *gin.Context) {
	// 禁止外部访问
	if c.ClientIP() != c.RemoteIP() {
		c.JSON(http.StatusNotFound, nil)
	}

	e.Log.Info("Hello World")
	e.Log.Error("Hello World")
	// e.Log.Fatal("Hello World")
	e.Log.Warn("Hello World")

	resp := make(map[string]string)
	resp["message"] = "Hello World"

	// 将api公共参数写入到响应体
	apiPublicParams := utils.GetApiPublicParamsAllMap(c)
	for k, v := range apiPublicParams {
		resp[k] = v
	}

	// 获取api公共参数的userid
	// userid := utils.GetApiPublicParamsUserid(c)

	c.JSON(http.StatusOK, resp)
}

func (example *exampleController) Health(c *gin.Context) {
	// 禁止外部访问
	if c.ClientIP() != c.RemoteIP() {
		c.JSON(http.StatusNotFound, nil)
	}

	c.JSON(http.StatusOK, map[string]string{"health": "ok"})
}

```