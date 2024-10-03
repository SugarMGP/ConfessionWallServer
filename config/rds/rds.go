package rds

import (
	"ConfessionWall/config/config"
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         config.Config.GetString("redis.addr"),
		Password:     config.Config.GetString("redis.password"),
		DB:           config.Config.GetInt("redis.DB"),
		PoolSize:     config.Config.GetInt("redis.PoolSize"),
		MinIdleConns: config.Config.GetInt("redis.MinIdleConns"),
	})

	// 测试连接是否成功
	ctx := context.Background()
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		zap.L().Fatal("Could not connect to Redis", zap.Error(err))
	} else if pong != "PONG" {
		zap.L().Fatal("Unexpected response from Redis", zap.String("response", pong))
	}

	zap.L().Info("Connected to Redis successfully.")
}

func GetRedis() *redis.Client {
	return redisClient
}
