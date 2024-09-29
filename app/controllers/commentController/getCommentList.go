package commentController

import (
	"ConfessionWall/app/apiException"
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
	Post uint `form:"post"`
}

func GetCommentList(c *gin.Context) {
	id := c.GetUint("user_id")

	// 绑定请求数据
	var data GetCommentListData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		c.AbortWithError(200, apiException.ParamsError)
		return
	}

	preCommentList, err := commentService.GetCommentsByPostID(data.Post)
	if err != nil {
		zap.L().Error("获取评论列表失败", zap.Uint("post_id", data.Post), zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	blocks, err := blockService.GetBlocksByUserID(id)
	if err != nil {
		zap.L().Error("获取拉黑列表失败", zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
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
