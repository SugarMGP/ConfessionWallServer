package main

import (
	"ConfessionWall/app/midwares"
	"ConfessionWall/config/database"
	"ConfessionWall/config/router"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	// 创建 zap 配置
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{"server.log"}
	cfg.ErrorOutputPaths = []string{"server.log"}

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
		log.Fatal("Server start failed: ", err)
	}
}
