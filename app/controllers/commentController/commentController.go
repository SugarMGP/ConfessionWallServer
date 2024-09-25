package commentController

import (
	"ConfessionWall/app/services/commentService"
	"ConfessionWall/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *commentService.CommentService
}

func NewCommentController(commentService *commentService.CommentService) *CommentController {
	return &CommentController{
		commentService: commentService,
	}
}

// CreateCommentHandler 创建评论
func (cc *CommentController) CreateComment(ctx *gin.Context) {
	postID, err := strconv.Atoi(ctx.Param("postID"))
	if err != nil {
		utils.JsonErrorResponse(ctx, 200511, "无效的帖子ID")
		return
	}
	var req struct {
		UserID  int    `json:"user_id"`
		Content string `json:"content"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.JsonErrorResponse(ctx, 200511, "无效的请求参数")
		return
	}

	comment, err := cc.commentService.CreateComment(uint(postID), uint(req.UserID), req.Content)
	if err != nil {
		utils.JsonErrorResponse(ctx, 200511, "创建评论失败")
		return
	}

	utils.JsonSuccessResponse(ctx, gin.H{
		"comment_id": comment.ID,
	})
}
