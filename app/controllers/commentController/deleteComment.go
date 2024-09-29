package commentController

import (
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
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	comment, err := commentService.GetCommentByID(data.CommentID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Debug("评论不存在", zap.Uint("comment_id", data.CommentID))
			utils.JsonErrorResponse(c, 200508, "评论不存在")
		} else {
			zap.L().Error("获取评论失败", zap.Uint("comment_id", data.CommentID), zap.Error(err))
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	// 检查用户是否有权限删除评论
	if comment.UserID != id {
		zap.L().Debug("请求的用户与发评论人不符", zap.Uint("user_id", id), zap.Uint("comment_user_id", comment.UserID))
		utils.JsonErrorResponse(c, 200509, "请求的用户与发评论人不符")
		return
	}

	err = commentService.DeleteComment(data.CommentID)
	if err != nil {
		zap.L().Error("删除评论失败", zap.Uint("user_id", id), zap.Uint("comment_id", data.CommentID), zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	zap.L().Info("评论删除成功", zap.Uint("user_id", id), zap.Uint("comment_id", data.CommentID))
	utils.JsonSuccessResponse(c, nil)
}
