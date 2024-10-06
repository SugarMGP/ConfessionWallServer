package userController

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/services/activityServive"
	"ConfessionWall/app/services/signService"
	"ConfessionWall/app/utils"
	"ConfessionWall/config/rds"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Sign(c *gin.Context) {
	id := c.GetUint("user_id")
	var offset int = time.Now().Local().Day() - 1
	var keys string = signService.GetIdKey(id)
	resid := rds.GetRedis()
	ctx := context.Background()

	// 检查是否已经签到
	isSignedIn, err := resid.GetBit(ctx, keys, int64(offset)).Result()
	if err != nil {
		zap.L().Error("获取签到状态失败", zap.Uint("user_id", id))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	if isSignedIn == 1 {
		zap.L().Info("用户今日已签到", zap.Uint("user_id", id))
		return
	}

	// 如果没有签到，则进行签到操作
	_, err = resid.SetBit(ctx, keys, int64(offset), 1).Result()
	if err != nil {
		zap.L().Error("签到失败", zap.Uint("user_id", id))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}
	zap.L().Info("签到成功", zap.Uint("user_id", id))

	// 增加活跃度
	err = activityServive.IncreaseActivity(id, 2)
	if err != nil {
		zap.L().Error("增加活跃度失败", zap.Uint("user_id", id), zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	// 返回成功响应
	utils.JsonSuccessResponse(c, nil)
}