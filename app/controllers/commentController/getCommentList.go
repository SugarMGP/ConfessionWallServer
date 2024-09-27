package commentController

import (
	"ConfessionWall/app/models"
	"ConfessionWall/app/services/blockService"
	"ConfessionWall/app/services/commentService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetCommentListResponse struct {
	CommentList []models.Comment `json:"comment_list"`
}

type GetCommentListData struct {
	PostID uint `json:"post_id" binding:"required"`
}

func GetCommentList(c *gin.Context) {
	id := c.GetUint("user_id")

	// 绑定请求数据
	var data GetCommentListData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	preCommentList, err := commentService.GetCommentsByPostID(data.PostID)
	if err != nil {
		zap.L().Error("获取评论列表失败", zap.Uint("post_id", data.PostID), zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	blocks, err := blockService.GetBlocksByUserID(id)
	if err != nil {
		zap.L().Error("获取拉黑列表失败", zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	commentList := make([]models.Comment, 0)
	for _, comment := range preCommentList {
		// 判断是否被屏蔽
		blocked := false
		for _, block := range blocks {
			if block.TargetID == comment.UserID {
				blocked = true
				break
			}
		}
		if blocked {
			continue
		}
		commentList = append(commentList, comment)
	}

	zap.L().Info("获取帖子列表成功", zap.Int("count", len(commentList)))
	utils.JsonSuccessResponse(c, GetCommentListResponse{
		CommentList: commentList,
	})
}
