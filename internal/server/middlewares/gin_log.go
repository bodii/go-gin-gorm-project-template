package middlewares

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

// GinLogFormat 中间件会将日志写入
func GinLogFormat() gin.HandlerFunc {

	// Disable Console Color, you don't need console color when writing the logs to file.
	// gin.DisableConsoleColor()

	// Logging to a file.
	if os.Getenv("LogToFile") == "true" {
		logFilename := os.Getenv("LogFilename")
		f, _ := os.Create(path.Join("logs", logFilename))
		gin.DefaultWriter = io.MultiWriter(f)
	}

	// 自定义日志格式
	logFormatterFunc := func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}

	return gin.LoggerWithFormatter(logFormatterFunc)
}
