package logger

import (
	"os"
	"time"

	"go.uber.org/zap"
)

func Init(debug bool) {
	// 创建 zap 配置
	cfg := zap.NewDevelopmentConfig()

	// 确保 logs 目录存在，如果不存在则创建
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0755)
		if err != nil {
			panic("Failed to create logs directory: " + err.Error())
		}
	}

	// 设置日志级别
	if debug {
		cfg.Level.SetLevel(zap.DebugLevel)
	} else {
		cfg.Level.SetLevel(zap.InfoLevel)
	}

	// 设置日志文件路径
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
