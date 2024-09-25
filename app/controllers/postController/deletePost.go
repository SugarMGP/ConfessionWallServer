package postController

import (
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DeleteResponse struct {
	PostID uint `json:"post_id" binding:"required"`
}

func DeletePost(c *gin.Context) {
	id := c.GetUint("user_id")

	// 绑定请求数据
	var data DeleteResponse
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	// 获取帖子信息
	post, err := postService.GetPostByID(data.PostID)
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

	// 检查用户是否有权限删除帖子
	if post.User != id {
		zap.L().Debug("请求的用户与发帖人不符", zap.Uint("user_id", id), zap.Uint("post_user_id", post.User))
		utils.JsonErrorResponse(c, 200509, "请求的用户与发帖人不符")
		return
	}

	// 删除帖子
	err = postService.DeletePost(data.PostID)
	if err != nil {
		zap.L().Error("删除帖子失败", zap.Uint("post_id", data.PostID), zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 成功删除帖子
	zap.L().Info("帖子删除成功", zap.Uint("post_id", data.PostID), zap.Uint("user_id", id))
	utils.JsonSuccessResponse(c, nil)
}
