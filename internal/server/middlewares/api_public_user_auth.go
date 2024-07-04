package middlewares

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"template-project-name/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

/*
接口公共参数的传参方式说明:
1. 在原参数外添加：
	issuedAt（请求签发时间，10位，unix时间戳）；
	expireTime（过期时间，10位，unix时间戳）；
	key(随机key，建议随机字符串(可以是数字，也可以是是字母或字母加数字)（8-20位, 非'+'、'='）)；
	keySign(签名key，生成keySign：md5(key + secret))；

2. header头直接传参，变为公共参数（除sign）base64加密；
header名：x-app-token。
如base64(accountId=1+version=125+...)，每个key与value用"="连接，每个keyVal组用"+"连接；

3. header头sign：x-app-sign。

4. 生成签名：md5(urlEncode(sort(key1=value1&key2=value2&key3=value3)):secret)

注：
1. 只传两个header，x-app-token和x-app-sign；
2. 服务器会验证过期时间，当没有过期时间时，会验证请求签发时间（10分钟）
*/

type apiPublicParamsT struct {
	// app version app版本号
	Version string `form:"version" json:"version" binding:"required"`
	// api use platform 台平: android, ios, pc, mobile_website
	Platform string `form:"platform" json:"platform" binding:"required"`
	// device info 设备型号信息
	Device string `form:"device" json:"device,omitempty"`
	// brand info 设备品牌信息
	Brand string `form:"brand" json:"brand,omitempty"`
	// user account id
	AccountId string `form:"accountId" json:"accountId" binding:"required"`
	// sign
	Sign string `form:"sign" json:"sign" binding:"required"`
	// expire time 过期时间，10位，unix时间戳
	ExpireTime string `form:"expireTime" json:"expireTime,omitempty"`
	// issued at 请求签发时间，10位，unix时间戳
	IssuedAt string `form:"issuedAt" json:"issuedAt,omitempty"`
	// key - random string 随机key，建议随机字符串(可以是数字，也可以是是字母或字母加数字)（8-20位, 非'+'、'='）
	Key string `form:"key" json:"key"`
	// key sign - md5(key + Secret)
	KeySign string `form:"keySign" json:"keySign"`
}

// toEncodeMap apiPublicParamsT to map
// Sign field is not required.
func (ua *apiPublicParamsT) toEncodeMap() map[string]string {
	return map[string]string{
		"accountId":  ua.AccountId,
		"brand":      ua.Brand,
		"device":     ua.Device,
		"expireTime": ua.ExpireTime,
		"issuedAt":   ua.IssuedAt,
		"key":        ua.Key,
		"keySign":    ua.KeySign,
		"version":    ua.Version,
		"platform":   ua.Platform}
}

// 验证签名
func (ua *apiPublicParamsT) validate() bool {
	// 获取签名参数
	paramsMap := ua.toEncodeMap()
	// 排序签名参数
	paramsSortMapPair := utils.SortMapByKey(paramsMap, utils.SortOrderAsc)

	// 拼接参数，通过 url.Values.Encode() 方法转码
	v := url.Values{}
	for _, p := range paramsSortMapPair {
		v.Add(p.Key, fmt.Sprintf("%s", p.Val))
	}

	// 生成签名：md5(urlEncode(key1=value1&key2=value2&key3=value3):secret)
	signStr := utils.MD5(v.Encode() + ":" + utils.GetEnvSecretValue())

	// fmt.Println("v.Encode():", v.Encode())
	// fmt.Println("signStr:", signStr)
	// fmt.Println("ua.Sign: ", ua.Sign)

	return ua.Sign == signStr
}

// 验证key
func (ua *apiPublicParamsT) keyValidate() bool {
	// key: 建议随机字符串(可以是数字，也可以是是字母或字母加数字)（8-20位, 除'+'、'='）
	// 生成keysign：md5(key + secret)
	keySignStr := utils.MD5(ua.Key + utils.GetEnvSecretValue())

	return ua.KeySign == keySignStr
}

// 验证过期时间
func (ua *apiPublicParamsT) expireTimeValidate() int {
	// unix时间戳（10位）
	if len(ua.ExpireTime) == 10 { // 如果有过期时间，查看过期时间是否
		expireUnix, err := strconv.Atoi(ua.ExpireTime)
		if err != nil {
			return 0
		}
		expireTime := time.Unix(int64(expireUnix), 0)
		if time.Now().After(expireTime) {
			return -1
		}
	} else { // 否则如果创建时间距离现在已超过10分钟，则返回过期
		issuedAtUnix, err := strconv.Atoi(ua.IssuedAt)
		if err != nil {
			return 0
		}
		issuedAtTime := time.Unix(int64(issuedAtUnix), 0)
		if time.Now().After(issuedAtTime.Add(time.Minute * 10)) {
			return -1
		}
	}

	return 1
}

// unToken 解析token
func (ua *apiPublicParamsT) unToken(token string) error {
	decoded, err := base64.StdEncoding.DecodeString(token)
	// log.Printf("%s\n", decoded)
	// log.Printf("%v\n", ua.Sign)

	if err != nil {
		return errors.New("参数错误")
	}

	if len(decoded) < 1 {
		return errors.New("参数错误")
	}

	for _, keyVal := range strings.Split(string(decoded), "+") {
		if len(keyVal) < 1 {
			continue
		}
		keyValList := strings.Split(keyVal, "=")
		switch keyValList[0] {
		case "expireTime":
			ua.ExpireTime = keyValList[1]
		case "issuedAt":
			ua.IssuedAt = keyValList[1]
		case "key":
			ua.Key = keyValList[1]
		case "keySign":
			ua.KeySign = keyValList[1]
		case "accountId":
			ua.AccountId = keyValList[1]
		case "version":
			ua.Version = keyValList[1]
		case "platform":
			ua.Platform = keyValList[1]
		case "device":
			ua.Device = keyValList[1]
		case "brand":
			ua.Brand = keyValList[1]
		}
	}

	return nil
}

// 验证api公共参数中间件方法
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ua := &apiPublicParamsT{
			Sign: c.Request.Header.Get("x-app-sign"),
		}

		token := c.Request.Header.Get("x-app-token")
		if len(token) < 1 || len(ua.Sign) < 1 {
			c.Abort()
			utils.ErrorJSON(c, http.StatusUnauthorized, "参数错误")
			return
		}

		// 解析token
		ua.unToken(token)
		// fmt.Printf("%#v\n", ua)

		if len(ua.Version) < 1 || len(ua.Platform) < 1 || len(ua.Key) < 1 ||
			len(ua.KeySign) < 1 || len(ua.IssuedAt) != 10 {

			c.Abort()
			utils.ErrorJSON(c, http.StatusBadRequest, "参数错误")
			return
		}

		if !ua.keyValidate() {
			c.Abort()
			utils.ErrorJSON(c, http.StatusUnauthorized, "无权限使用")
			return
		}

		expVal := ua.expireTimeValidate()
		if expVal == 0 {
			c.Abort()
			utils.ErrorJSON(c, http.StatusBadRequest, "参数错误")
			return
		} else if expVal == -1 {
			c.Abort()
			utils.ErrorJSON(c, http.StatusUnauthorized, "已失效")
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
