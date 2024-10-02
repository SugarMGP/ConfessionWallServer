package commentController

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/services/likeService"
	"ConfessionWall/app/utils"
	"ConfessionWall/config/rds"
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LikeCommentData struct {
	CommentID uint `json:"comment_id" binding:"required"`
}

func LikeComment(c *gin.Context) {
	id := c.GetUint("user_id")
	user := int64(id)

	var data LikeCommentData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		c.AbortWithError(200, apiException.ParamsError)
		return
	}

	commentKey := likeService.GetCommentKey(data.CommentID)
	redis := rds.GetRedis()
	defer redis.Close()
	ctx := context.Background()

	res, err := redis.GetBit(ctx, commentKey, user-1).Result()
	if err != nil {
		zap.L().Error("从Redis获取点赞状态失败", zap.Uint("comment_id", data.CommentID), zap.Uint("user_id", id))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	if res == 1 { // 用户已经点赞过
		_, err = redis.SetBit(ctx, commentKey, user-1, 0).Result()
		if err != nil {
			zap.L().Error("取消点赞失败", zap.Uint("comment_id", data.CommentID), zap.Uint("user_id", id), zap.Error(err))
			c.AbortWithError(200, apiException.InternalServerError)
			return
		}

		err := likeService.CommentLikeCount(data.CommentID)
		if err != nil {
			zap.L().Error("累计点赞数失败", zap.Uint("comment_id", data.CommentID), zap.Uint("user_id", id), zap.Error(err))
			c.AbortWithError(200, apiException.InternalServerError)
			return
		}
		return
	} else { // 用户未点赞过
		_, err = redis.SetBit(ctx, commentKey, user-1, 1).Result()
		if err != nil {
			zap.L().Error("设置点赞状态失败", zap.Uint("comment_id", data.CommentID), zap.Uint("user_id", id), zap.Error(err))
			c.AbortWithError(200, apiException.InternalServerError)
			return
		}

		err := likeService.CommentLikeCount(data.CommentID)
		if err != nil {
			zap.L().Error("累计点赞数失败", zap.Uint("comment_id", data.CommentID), zap.Uint("user_id", id), zap.Error(err))
			c.AbortWithError(200, apiException.InternalServerError)
			return
		}
		utils.JsonSuccessResponse(c, nil)
	}
}
