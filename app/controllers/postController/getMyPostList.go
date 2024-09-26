package postController

import (
	"ConfessionWall/app/models"
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetMyListResponse struct {
	MyConfessionList []models.Post `json:"my_confession_list"`
}

func GetMyPostList(c *gin.Context) {
	id := c.GetUint("user_id")

	// 获取帖子列表
	postList, err := postService.GetMyPostList(id)
	if err != nil {
		zap.L().Error("获取帖子列表失败", zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 成功获取帖子列表
	zap.L().Info("获取帖子列表成功", zap.Int("count", len(postList)))
	utils.JsonSuccessResponse(c, GetMyListResponse{
		MyConfessionList: postList,
	})
}
