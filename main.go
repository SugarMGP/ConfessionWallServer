package main

import (
	"ConfessionWall/app/midwares"
	"ConfessionWall/config/config"
	"ConfessionWall/config/database"
	"ConfessionWall/config/logger"
	"ConfessionWall/config/router"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	debug := config.Config.GetBool("debug")
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	logger.Init(debug)
	database.Init()
	r := gin.Default()
	r.NoMethod(midwares.HandleNotFound) // 中间件统一处理404
	r.NoRoute(midwares.HandleNotFound)
	router.Init(r)

	err := r.Run()
	if err != nil {
		zap.L().Fatal("Server start failed", zap.Error(err))
	} else {
		zap.L().Info("Server start success")
	}
}
