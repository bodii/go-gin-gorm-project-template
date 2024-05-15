package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"template-project-name/internal/routes/api"
)

// RegisterRoutes 注册总路由
func RegisterRoutes(router *gin.Engine) http.Handler {

	// example
	// exampleContro := controller.NewExampleController()
	// router.GET("/", exampleContro.HelloWorld)

	// 注册api路由组及其子路由
	api.RegisterRoutes(router.Group("/api"))

	return router
}
