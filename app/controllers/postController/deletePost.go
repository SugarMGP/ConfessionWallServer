package postController

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DeleteData struct {
	PostID uint `json:"post_id" binding:"required"`
}

func DeletePost(c *gin.Context) {
	id := c.GetUint("user_id")

	// 绑定请求数据
	var data DeleteData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		c.AbortWithError(200, apiException.ParamsError)
		return
	}

	// 获取帖子信息
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

	// 检查用户是否有权限删除帖子
	if post.User != id {
		zap.L().Debug("请求的用户与发帖人不符", zap.Uint("user_id", id), zap.Uint("post_user_id", post.User))
		c.AbortWithError(200, apiException.NoOperatePermission)
		return
	}

	// 删除帖子
	err = postService.DeletePost(data.PostID)
	if err != nil {
		zap.L().Error("删除帖子失败", zap.Uint("post_id", data.PostID), zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	// 成功删除帖子
	zap.L().Info("帖子删除成功", zap.Uint("user_id", id), zap.Uint("post_id", data.PostID))
	utils.JsonSuccessResponse(c, nil)
}
