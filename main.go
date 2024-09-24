package main

import (
	"ConfessionWall/app/midwares"
	"ConfessionWall/config/database"
	"ConfessionWall/config/router"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	// 创建 zap 配置
	cfg := zap.NewDevelopmentConfig()

	// 确保 logs 目录存在，如果不存在则创建
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0755)
		if err != nil {
			panic("Failed to create logs directory: " + err.Error())
		}
	}

	path := "logs/" + time.Now().Format("2006-01-02") + ".log"
	cfg.OutputPaths = []string{path}
	cfg.ErrorOutputPaths = []string{path}

	// 创建 zap.Logger 实例
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	// 设置全局日志记录器
	zap.ReplaceGlobals(logger)
}

func main() {
	database.Init()
	r := gin.Default()
	r.NoMethod(midwares.HandleNotFound) // 中间件统一处理404
	r.NoRoute(midwares.HandleNotFound)
	router.Init(r)

	err := r.Run()
	if err != nil {
		zap.L().Fatal("Server start failed", zap.Error(err))
	}
}
