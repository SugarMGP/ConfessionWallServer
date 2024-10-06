package userController

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/services/activityServive"
	"ConfessionWall/app/services/blockService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetActivityUserResponse struct {
	ActivityUserList []ActivityUserElement `json:"activity_user_list"`
}
type ActivityUserElement struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Activity int    `json:"activity"`
}

func GetActivityUserList(c *gin.Context) {
	id := c.GetUint("user_id")

	blocks, err := blockService.GetBlocksByUserID(id)
	if err != nil {
		zap.L().Error("获取拉黑列表失败", zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	activityUserList, err := activityServive.GetActivityUser()
	if err != nil {
		zap.L().Error("获取活跃用户列表失败", zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	list := make([]ActivityUserElement, 0)
	for _, user := range activityUserList {
		// 判断是否被屏蔽
		blocked := false
		for _, block := range blocks {
			if block.TargetID == user.ID {
				blocked = true
				break
			}
		}
		if blocked {
			continue
		}

		userElement := ActivityUserElement{
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
			Activity: user.Activity,
		}
		list = append(list, userElement)
	}

	// 成功获取帖子列表
	zap.L().Info("获取活跃用户列表成功", zap.Int("count", len(list)))
	utils.JsonSuccessResponse(c, GetActivityUserResponse{
		ActivityUserList: list,
	})

}
