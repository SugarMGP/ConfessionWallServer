package blockController

import (
	"ConfessionWall/app/services/blockService"
	"ConfessionWall/app/services/userService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetBlacklistResponse struct {
	Blacklist []BlacklistElement `json:"blacklist"`
}

type BlacklistElement struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	TargetID uint   `json:"target_id"`
}

func GetBlacklist(c *gin.Context) {
	id := c.GetUint("user_id")

	// 获取拉黑列表
	blocks, err := blockService.GetBlocksByUserID(id)
	if err != nil {
		zap.L().Error("获取拉黑列表失败", zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	blacklist := make([]BlacklistElement, 0)
	for _, block := range blocks {
		target, err := userService.GetUserByID(block.TargetID)
		if err != nil {
			zap.L().Error("获取用户信息失败", zap.Uint("user_id", block.TargetID), zap.Error(err))
			continue
		}

		blacklist = append(blacklist, BlacklistElement{
			Nickname: target.Nickname,
			Avatar:   target.Avatar,
			TargetID: block.TargetID,
		})
	}

	// 成功获取拉黑列表
	zap.L().Info("获取拉黑列表成功", zap.Int("count", len(blacklist)))
	utils.JsonSuccessResponse(c, GetBlacklistResponse{
		Blacklist: blacklist,
	})
}
