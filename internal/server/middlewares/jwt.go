package middlewares

import (
	"net/http"
	"template-project-name/internal/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			c.Abort()
			utils.ErrorJSON(c, http.StatusUnauthorized, "无权限登录")
			return
		}
		// 校验token，只要出错直接拒绝请求
		token, err := utils.NewJwtService().ValidateWithClaimsToken(auth)
		if err != nil || !token.Valid {
			c.Abort()
			utils.ErrorJSON(c, http.StatusUnauthorized, "权限登录失败")
			return
		}

		// 解析userid
		userid := token.Claims.(*utils.JWTAuthCustomClaims).Userid

		// 将参数放入上下文中
		c.Set("userid", userid)

		c.Next()
	}
}
