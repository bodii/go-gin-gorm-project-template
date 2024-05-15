package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"template-project-name/internal/bootstrap"
	"template-project-name/internal/routes"
	"template-project-name/internal/server/middlewares"
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

	app := gin.New()

	// gin log format middleware
	app.Use(middlewares.GinLogFormat())

	app.Use(gin.Recovery())

	// use jwt auth middleware
	app.Use(middlewares.JWTAuth())

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
