package main

import (
	"ConfessionWall/app/midwares"
	"ConfessionWall/config/config"
	"ConfessionWall/config/database"
	"ConfessionWall/config/logger"
	"ConfessionWall/config/router"
	"os"

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
	r.ForwardedByClientIP = true
	r.NoMethod(midwares.HandleNotFound) // 使用404统一处理中间件
	r.NoRoute(midwares.HandleNotFound)
	r.Use(midwares.Limiter())
	r.Use(midwares.ErrHandler()) // 统一错误处理中间件

	// 确保 static 目录存在，如果不存在则创建
	if _, err := os.Stat("static"); os.IsNotExist(err) {
		err := os.Mkdir("static", 0755)
		if err != nil {
			zap.L().Fatal("Failed to create static directory", zap.Error(err))
		}
	}
	r.Static("/static", "./static") // 挂载静态文件目录

	router.Init(r)

	err := r.Run()
	if err != nil {
		zap.L().Fatal("Server start failed", zap.Error(err))
	}
}
