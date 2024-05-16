package middlewares

import (
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"template-project-name/internal/utils"

	"github.com/gin-gonic/gin"
)

type apiPublicParamsT struct {
	// app version app版本号
	Version string `form:"version" json:"version" binding:"required"`
	// api use platform 台平: android, ios, pc, mobile_website
	Platform string `form:"platform" json:"platform" binding:"required"`
	// device info 设备型号信息
	Device string `form:"device" json:"device" binding:"required"`
	// brand info 设备品牌信息
	Brand string `form:"brand" json:"brand"`
	// user account id
	AccountId string `form:"accountId" json:"accountId" binding:"required"`
	// sign
	Sign string `form:"sign" json:"sign" binding:"required"`
}

// toEncodeMap apiPublicParamsT to map
// Sign field is not required.
func (ua *apiPublicParamsT) toEncodeMap() map[string]any {
	return map[string]any{
		"device":    ua.Device,
		"accountId": ua.AccountId,
		"version":   ua.Version,
		"platform":  ua.Platform,
		"brand":     ua.Brand}
}

type pair struct {
	key string
	val any
}

// sortMap sort map
func sortMap(m map[string]any) []pair {
	pairs := []pair{}
	for k, v := range m {
		pairs = append(pairs, pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].key < pairs[j].key
	})

	return pairs
}

// 验证签名
func (ua *apiPublicParamsT) validate() bool {
	// 获取签名参数
	paramsMap := ua.toEncodeMap()
	// 排序签名参数
	paramsSortMap := sortMap(paramsMap)

	// 拼接参数，通过 url.Values.Encode() 方法转码
	v := url.Values{}
	for _, p := range paramsSortMap {
		v.Add(p.key, fmt.Sprintf("%s", p.val))
	}

	// 生成签名：md5(urlEncode(key1=value1&key2=value2&key3=value3):secret)
	sign_str := utils.MD5(v.Encode() + ":" + utils.GetEnvSecretValue())

	// fmt.Println("v.Encode():", v.Encode())
	// fmt.Println("sign_str:", sign_str)
	// fmt.Println("ua.Sign: ", ua.Sign)

	return ua.Sign == sign_str
}

// 验证api公共参数中间件方法
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ua := &apiPublicParamsT{}
		if err := c.ShouldBindJSON(ua); err != nil {
			c.Abort()
			utils.ErrorJSON(c, http.StatusBadRequest, "参数错误")
			return
		}

		if !ua.validate() {
			c.Abort()
			utils.ErrorJSON(c, http.StatusUnauthorized, "无权限使用")
			return
		}

		// 将参数放入上下文中
		c.Set("api_public_params", ua.toEncodeMap())

		c.Next()
	}
}
