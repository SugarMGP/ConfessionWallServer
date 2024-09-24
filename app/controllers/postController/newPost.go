package postController

import (
	"ConfessionWall/app/models"
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/services/userService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NewPostData struct {
	Content string `json:"content" binding:"required"`
	Unnamed bool   `json:"unnamed"`
}

func NewPost(c *gin.Context) {
	id := c.GetUint("user_id")

	// 绑定请求数据
	var data NewPostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	// 检查用户是否存在
	user, err := userService.GetUserByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Error("用户不存在", zap.Uint("user_id", id))
			utils.JsonErrorResponse(c, 200508, "用户不存在")
		} else {
			zap.L().Error("查询用户信息失败", zap.Uint("user_id", id), zap.Error(err))
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	// 新建帖子
	post := models.Post{
		Content: data.Content,
		User:    user.ID,
		Unnamed: data.Unnamed,
	}
	err = postService.NewPost(post)
	if err != nil {
		zap.L().Error("新建帖子失败", zap.Uint("user_id", id), zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 成功创建帖子
	zap.L().Info("帖子创建成功", zap.Uint("user_id", id), zap.String("content", data.Content))
	utils.JsonSuccessResponse(c, nil)
}
