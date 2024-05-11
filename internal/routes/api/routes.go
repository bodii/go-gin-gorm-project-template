package api

import (
	v1 "template-project-name/internal/routes/api/v1"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册api路由组及其子路由
func RegisterRoutes(api *gin.RouterGroup) {

	// 注册api下v1子路由组及其v1子路由
	v1.RegisterRoutes(api.Group("/v1"))
}
