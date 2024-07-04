package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"template-project-name/internal/routes/api"
	"template-project-name/internal/server/middlewares"
)

// RegisterRoutes 注册总路由
func RegisterRoutes(router *gin.Engine) http.Handler {

	// example
	// exampleContro := controller.NewExampleController()
	// router.GET("/", exampleContro.HelloWorld)

	// 注册api路由组及其子路由, 并使用jwt验证中间件
	api.RegisterRoutes(router.Group("/api", middlewares.JWTAuth()))

	return router
}
