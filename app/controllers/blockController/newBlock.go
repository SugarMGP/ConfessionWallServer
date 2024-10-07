package blockController

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/models"
	"ConfessionWall/app/services/blockService"
	"ConfessionWall/app/services/commentService"
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NewBlockData struct {
	PostID    uint `json:"post_id"`
	CommentID uint `json:"comment_id"`
}

func NewBlock(c *gin.Context) {
	id := c.GetUint("user_id")

	var data NewBlockData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		c.AbortWithError(200, apiException.ParamsError)
		return
	}

	if (data.PostID == 0 && data.CommentID == 0) || (data.PostID != 0 && data.CommentID != 0) {
		c.AbortWithError(200, apiException.ParamsError)
		return
	}

	var user uint         // 被拉黑者的ID
	if data.PostID != 0 { // 帖子被拉黑
		post, err := postService.GetPostByID(data.PostID)
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
		user = post.UserID
	} else { // 评论被拉黑
		comment, err := commentService.GetCommentByID(data.CommentID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				zap.L().Debug("评论不存在", zap.Uint("comment_id", data.CommentID))
				c.AbortWithError(200, apiException.CommentNotFound)
			} else {
				zap.L().Error("获取评论信息失败", zap.Uint("comment_id", data.CommentID), zap.Error(err))
				c.AbortWithError(200, apiException.InternalServerError)
			}
			return
		}
		user = comment.UserID
	}
	if user == id {
		c.AbortWithError(200, apiException.AttemptToBlockSelf)
		return
	}

	_, err = blockService.GetBlockByID(id, user)
	if err == nil {
		zap.L().Debug("拉黑关系已存在", zap.Uint("user_id", id), zap.Uint("target_id", user))
		c.AbortWithError(200, apiException.HasBlocked)
		return
	} else if err != gorm.ErrRecordNotFound {
		zap.L().Error("查询拉黑信息失败", zap.Uint("user_id", id), zap.Uint("target_id", user), zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	err = blockService.NewBlock(models.Block{
		UserID:   id,
		TargetID: user,
	})
	if err != nil {
		zap.L().Error("新建拉黑失败", zap.Uint("user_id", id), zap.Uint("target_id", user), zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	// 成功创建拉黑
	zap.L().Info("创建拉黑成功", zap.Uint("user_id", id), zap.Uint("target_id", user))
	utils.JsonSuccessResponse(c, nil)
}
