## example 

file name: routes.go
```go
package v1

import (
	"template-project-name/internal/controller/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册v1路由组及其子路由
func RegisterRoutes(v1 *gin.RouterGroup) {
	/* 用户 */
	userContro := handlers.NewUserController()
	v1.POST("/user/login", userContro.Login)
}

```