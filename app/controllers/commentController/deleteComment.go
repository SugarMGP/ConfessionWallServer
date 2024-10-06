package commentController

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/services/activityServive"
	"ConfessionWall/app/services/commentService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DeleteData struct {
	CommentID uint `json:"comment_id" binding:"required"`
}

// 删除评论
func DeleteComment(c *gin.Context) {
	id := c.GetUint("user_id")

	var data DeleteData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		c.AbortWithError(200, apiException.ParamsError)
		return
	}

	comment, err := commentService.GetCommentByID(data.CommentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Debug("评论不存在", zap.Uint("comment_id", data.CommentID))
			c.AbortWithError(200, apiException.CommentNotFound)
		} else {
			zap.L().Error("获取评论失败", zap.Uint("comment_id", data.CommentID), zap.Error(err))
			c.AbortWithError(200, apiException.InternalServerError)
		}
		return
	}

	// 检查用户是否有权限删除评论
	if comment.UserID != id {
		zap.L().Debug("请求的用户与发评论人不符", zap.Uint("user_id", id), zap.Uint("comment_user_id", comment.UserID))
		c.AbortWithError(200, apiException.NoOperatePermission)
		return
	}

	err = commentService.DeleteComment(data.CommentID)
	if err != nil {
		zap.L().Error("删除评论失败", zap.Uint("user_id", id), zap.Uint("comment_id", data.CommentID), zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}
	err = activityServive.DecreaseActivity(id, 2)
	if err != nil {
		zap.L().Error("减少活跃度失败", zap.Uint("user_id", id), zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}
	zap.L().Info("评论删除成功", zap.Uint("user_id", id), zap.Uint("comment_id", data.CommentID))
	utils.JsonSuccessResponse(c, nil)
}
