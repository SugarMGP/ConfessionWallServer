package commentController

import (
	"ConfessionWall/app/models"
	"ConfessionWall/app/services/commentService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeleteResponse struct {
	PostID uint `json:"post_id" binding:"required"`
}

// 删除评论
func DeleteComment(c *gin.Context) {
	id := c.GetUint("user_id")
	var data DeleteResponse
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}
	comment, err := commentService.GetCommentsByPostID(data.PostID)
	if err != nil {
		zap.L().Error("获取评论失败", zap.Error(err))
		utils.JsonErrorResponse(c, 200507, "获取评论失败")
		return
	}
	// 检查用户是否有权限删除评论
	// 删除帖子
	var commentToDelete *models.Comment
	for _, comment := range comment {
		if comment.UserID == id {
			commentToDelete = &comment
			break
		}
	}
	if commentToDelete == nil {
		zap.L().Debug("请求的用户与发评论的人不符", zap.Uint("user_id", id))
		utils.JsonErrorResponse(c, 200509, "请求的用户与发评论的人不符")
		return
	}
	err = commentService.DeleteComment(id)
	if err != nil {
		zap.L().Error("删除评论失败", zap.Uint("user_id", id), zap.Uint("post_id", data.PostID), zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	zap.L().Info("评论删除成功", zap.Uint("user_id", id), zap.Uint("post_id", data.PostID))
	utils.JsonSuccessResponse(c, nil)
}
