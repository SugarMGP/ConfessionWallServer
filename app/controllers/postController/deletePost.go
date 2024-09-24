package postController

import (
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteResponse struct {
	PostID uint `json:"post_id"`
}

func DeletePost(c *gin.Context) {
	id := c.GetUint("user_id")

	var data DeleteResponse
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

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

	err = postService.DeletePost(data.PostID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
