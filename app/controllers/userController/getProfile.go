package userController

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/services/userService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetProfileResponse struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Activity int    `json:"activity"`
}

func GetProfile(c *gin.Context) {
	id := c.GetUint("user_id")

	// 获取用户信息
	user, err := userService.GetUserByID(id)
	if err != nil {
		zap.L().Error("获取用户信息失败", zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	// 成功获取用户信息
	zap.L().Info("用户信息获取成功", zap.Uint("user_id", id))
	utils.JsonSuccessResponse(c, GetProfileResponse{
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Activity: user.Activity,
	})
}
