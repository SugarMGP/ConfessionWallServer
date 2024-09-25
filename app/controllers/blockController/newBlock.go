package blockController

import (
	"ConfessionWall/app/models"
	"ConfessionWall/app/services/blockService"
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type NewBlockData struct {
	PostID uint `json:"post_id" binding:"required"`
}

func NewBlock(c *gin.Context) {
	id := c.GetUint("user_id")

	var data NewBlockData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	post, err := postService.GetPostByID(data.PostID)
	if err != nil {
		zap.L().Error("查询帖子信息失败", zap.Uint("post_id", data.PostID), zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	if post.User == id {
		utils.JsonErrorResponse(c, 200510, "不能屏蔽自己的帖子")
		return
	}

	_, err = blockService.GetBlockByID(id, post.User)
	if err == nil {
		zap.L().Debug("拉黑关系已存在", zap.Uint("user_id", id), zap.Uint("target_id", post.User))
		utils.JsonErrorResponse(c, 200503, "拉黑关系已存在")
		return
	} else if err != gorm.ErrRecordNotFound {
		zap.L().Error("查询拉黑信息失败", zap.Uint("user_id", id), zap.Uint("target_id", post.User), zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	err = blockService.NewBlock(models.Block{
		UserID:   id,
		TargetID: post.User,
	})
	if err != nil {
		zap.L().Error("新建拉黑失败", zap.Uint("user_id", id), zap.Uint("target_id", post.User), zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 成功创建拉黑
	zap.L().Info("创建拉黑成功", zap.Uint("user_id", id), zap.Uint("target_id", post.User))
	utils.JsonSuccessResponse(c, nil)
}
