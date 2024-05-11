package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"template-project-name/internal/bootstrap"
	"template-project-name/internal/routes"
)

func NewServer() *http.Server {
	// 引导
	bootstrap.Init()

	// get dev mode in env file
	mode := os.Getenv("MODE")
	// set dev mode
	switch mode {
	case gin.TestMode:
		gin.SetMode(gin.TestMode)
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	// Disable Console Color, you don't need console color when writing the logs to file.
	// gin.DisableConsoleColor()

	// // Logging to a file.
	if os.Getenv("LogToFile") == "true" {
		logFilename := os.Getenv("LogFilename")
		f, _ := os.Create(fmt.Sprintf("logs/%s", logFilename))
		gin.DefaultWriter = io.MultiWriter(f)
	}

	app := gin.New()

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

	// LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	app.Use(gin.LoggerWithFormatter(logFormatterFunc))

	app.Use(gin.Recovery())

	// get port in env file
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      routes.RegisterRoutes(app),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
