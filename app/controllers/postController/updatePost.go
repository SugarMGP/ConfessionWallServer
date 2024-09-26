package postController

import (
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UpdatePostData struct {
	PostID  uint   `json:"post_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func UpdatePost(c *gin.Context) {
	id := c.GetUint("user_id")

	// 绑定请求数据
	var data UpdatePostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	// 检查请求用户是否为发帖人
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
	if post.User != id {
		zap.L().Debug("请求的用户与发帖人不符", zap.Uint("user_id", id), zap.Uint("post_user_id", post.User))
		utils.JsonErrorResponse(c, 200509, "请求的用户与发帖人不符")
		return
	}

	// 编辑帖子
	err = postService.UpdatePost(data.PostID, data.Content)
	if err != nil {
		zap.L().Error("编辑帖子失败", zap.Uint("post_id", data.PostID), zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 成功编辑帖子
	zap.L().Info("帖子编辑成功", zap.Uint("post_id", data.PostID), zap.Uint("user_id", id))
	utils.JsonSuccessResponse(c, nil)
}
