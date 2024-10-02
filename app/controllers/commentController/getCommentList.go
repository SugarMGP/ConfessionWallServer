package commentController

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/services/blockService"
	"ConfessionWall/app/services/commentService"
	"ConfessionWall/app/services/userService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetCommentListResponse struct {
	CommentList []CommentElement `json:"comment_list"`
}

type CommentElement struct {
	ID       uint   `json:"comment_id"`
	Content  string `json:"content"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Likes    int64  `json:"likes"`
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

	commentList := make([]CommentElement, 0)
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

		nickname := ""
		avatar := ""
		user, err := userService.GetUserByID(comment.UserID)
		if err == nil { // 如果能获取到用户
			nickname = user.Nickname
			avatar = user.Avatar
		} else {
			zap.L().Error("获取用户信息失败", zap.Uint("user_id", comment.UserID), zap.Error(err))
		}

		commentList = append(commentList, CommentElement{
			ID:       comment.ID,
			Content:  comment.Content,
			Nickname: nickname,
			Avatar:   avatar,
			Likes:    comment.Likes,
		})
	}

	zap.L().Info("获取帖子列表成功", zap.Int("count", len(commentList)))
	utils.JsonSuccessResponse(c, GetCommentListResponse{
		CommentList: commentList,
	})
}
