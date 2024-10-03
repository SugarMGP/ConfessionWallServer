package postController

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/services/likeService"
	"ConfessionWall/app/utils"
	"ConfessionWall/config/rds"
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LikePostData struct {
	PostID uint `json:"post_id" binding:"required"`
}

func LikePost(c *gin.Context) {
	id := c.GetUint("user_id")
	user := int64(id)

	var data LikePostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		c.AbortWithError(200, apiException.ParamsError)
		return
	}

	postKey := likeService.GetPostKey(data.PostID)
	redis := rds.GetRedis()
	ctx := context.Background()

	res, err := redis.GetBit(ctx, postKey, user-1).Result()
	if err != nil {
		zap.L().Error("从Redis获取点赞状态失败", zap.Uint("post_id", data.PostID), zap.Uint("user_id", id))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	if res == 1 { // 用户已经点赞过
		_, err = redis.SetBit(ctx, postKey, user-1, 0).Result()
		if err != nil {
			zap.L().Error("取消点赞失败", zap.Uint("post_id", data.PostID), zap.Uint("user_id", id), zap.Error(err))
			c.AbortWithError(200, apiException.InternalServerError)
			return
		}

		err = likeService.PostLikeCount(data.PostID)
		if err != nil {
			zap.L().Error("累计点赞数失败", zap.Uint("post_id", data.PostID), zap.Uint("user_id", id), zap.Error(err))
			c.AbortWithError(200, apiException.InternalServerError)
			return
		}

		utils.JsonSuccessResponse(c, nil)
		return
	}

	// 用户未点赞过
	_, err = redis.SetBit(ctx, postKey, user-1, 1).Result()
	if err != nil {
		zap.L().Error("设置点赞状态失败", zap.Uint("post_id", data.PostID), zap.Uint("user_id", id), zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	err = likeService.PostLikeCount(data.PostID)
	if err != nil {
		zap.L().Error("累计点赞数失败", zap.Uint("post_id", data.PostID), zap.Uint("user_id", id), zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
