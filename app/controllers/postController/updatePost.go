package postController

import (
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/services/userService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpdatePostData struct {
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
}

func UpdatePost(c *gin.Context) {
	id := c.GetUint("user_id")

	var data UpdatePostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	// 验证用户是否存在
	_, err = userService.GetUserByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200508, "用户不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	// 检查请求用户是否为发帖人
	post, err := postService.GetPostByID(data.PostID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200508, "帖子不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}
	if post.User != id {
		utils.JsonErrorResponse(c, 200509, "请求的用户与发帖人不符")
		return
	}

	// 编辑帖子
	err = postService.UpdatePost(data.PostID, data.Content)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
