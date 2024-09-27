package commentController

import (
	"ConfessionWall/app/models"
	"ConfessionWall/app/services/commentService"
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CommentData struct {
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
}

// CreateCommentHandler 创建评论
func NewComment(c *gin.Context) {
	id := c.GetUint("user_id")

	var data CommentData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	// 检查帖子是否存在
	_, err = postService.GetPostByID(data.PostID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Debug("帖子不存在", zap.Uint("post_id", data.PostID))
			utils.JsonErrorResponse(c, 200508, "帖子不存在")
		} else {
			zap.L().Error("获取帖子信息失败", zap.Uint("post_id", data.PostID), zap.Error(err))
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	err = commentService.NewComment(models.Comment{
		PostID:  data.PostID,
		UserID:  id,
		Content: data.Content,
	})
	if err != nil {
		zap.L().Error("发布评论失败", zap.Uint("user_id", id), zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 成功创建帖子
	zap.L().Info("发布评论成功", zap.Uint("user_id", id), zap.String("content", data.Content))
	utils.JsonSuccessResponse(c, nil)
}