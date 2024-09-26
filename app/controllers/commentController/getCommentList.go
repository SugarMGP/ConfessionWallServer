package commentController

import (
	"ConfessionWall/app/models"
	"ConfessionWall/app/services/commentService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetCommentListResponse struct {
	ConfessionList []models.Comment `json:"confession_list"`
}

type Confession struct {
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
}

func GetCommentsByPostID(c *gin.Context) {
	id := c.GetUint("post_id")
	commentList, err := commentService.GetCommentsByPostID(id)
	if err != nil {
		zap.L().Error("获取评论列表失败", zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	zap.L().Info("获取帖子列表成功", zap.Int("count", len(commentList)))
	utils.JsonSuccessResponse(c, GetCommentListResponse{
		ConfessionList: commentList,
	})

}
