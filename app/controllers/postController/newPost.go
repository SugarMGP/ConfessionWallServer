package postController

import (
	"ConfessionWall/app/models"
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/services/userService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NewPostData struct {
	Content string `json:"content" binding:"required"`
	Unnamed bool   `json:"unnamed" binding:"required"`
}

func NewPost(c *gin.Context) {
	id := c.GetUint("user_id")
	var data NewPostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	_, err = userService.GetUserByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200508, "用户不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	// 新建帖子
	err = postService.NewPost(models.Post{
		Content: data.Content,
		User:    id,
		Unnamed: data.Unnamed,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
