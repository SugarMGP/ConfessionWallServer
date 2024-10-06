package commentController

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/models"
	"ConfessionWall/app/services/activityServive"
	"ConfessionWall/app/services/commentService"
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NewCommentData struct {
	PostID  uint   `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func NewComment(c *gin.Context) {
	id := c.GetUint("user_id")

	var data NewCommentData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		c.AbortWithError(200, apiException.ParamsError)
		return
	}

	// 检查帖子是否存在
	_, err = postService.GetPostByID(data.PostID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Debug("帖子不存在", zap.Uint("post_id", data.PostID))
			c.AbortWithError(200, apiException.PostNotFound)
		} else {
			zap.L().Error("获取帖子信息失败", zap.Uint("post_id", data.PostID), zap.Error(err))
			c.AbortWithError(200, apiException.InternalServerError)
		}
		return
	}

	if len(data.Content) > 1000 {
		zap.L().Debug("评论内容过长", zap.Uint("user_id", id), zap.Int("length", len(data.Content)))
		c.AbortWithError(200, apiException.ContentTooLong)
		return
	}

	err = commentService.NewComment(models.Comment{
		PostID:  data.PostID,
		UserID:  id,
		Content: data.Content,
	})
	if err != nil {
		zap.L().Error("发布评论失败", zap.Uint("user_id", id), zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	err = activityServive.IncreaseActivity(id, 2)
	if err != nil {
		zap.L().Error("增加活跃度失败", zap.Uint("user_id", id), zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	// 成功创建评论
	zap.L().Info("发布评论成功", zap.Uint("user_id", id), zap.Uint("post_id", data.PostID))
	utils.JsonSuccessResponse(c, nil)
}
