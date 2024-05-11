package utils

import "github.com/gin-gonic/gin"

// Response struct
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

// ServerJSON : json error server response function
func ErrorJSON(c *gin.Context, statusCode int, msg string) {
	c.JSON(statusCode, gin.H{"msg": msg})
}

// ResponseJSON : json response function
func ResponseJSON(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(200, gin.H{"code": code, "msg": msg, "data": data})
}

func OkJSON(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"code": 1, "msg": "ok", "data": data})
}
