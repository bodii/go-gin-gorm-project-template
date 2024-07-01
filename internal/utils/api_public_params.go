package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetApiPublicParamsAllMap 获取所有公共参数的map形式
func GetApiPublicParamsAllMap(c *gin.Context) map[string]any {
	resp := make(map[string]any)
	if apiPublicParams, ok := c.Get("api_public_params"); ok {
		for k, v := range apiPublicParams.(map[string]string) {
			if k == "accountId" {
				convVal, _ := strconv.ParseUint(v, 10, 64)
				resp[k] = convVal
			}
		}
	}

	return resp
}

// GetApiPublicParamsToUserid 获取公共参数中的accountId
func GetApiPublicParamsToUserid(c *gin.Context) uint64 {
	params := GetApiPublicParamsAllMap(c)

	return params["accountId"].(uint64)
}
