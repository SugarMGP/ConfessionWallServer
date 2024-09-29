package userController

import (
	"ConfessionWall/app/services/userService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SetProfileData struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

func SetProfile(c *gin.Context) {
	id := c.GetUint("user_id")

	var data SetProfileData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	if data.Nickname != "" {
		if len(data.Nickname) > 16 {
			zap.L().Debug("用户昵称设置过长", zap.Uint("user_id", id), zap.Error(err))
			utils.JsonErrorResponse(c, 200512, "用户昵称过长")
			return
		}

		_, err := userService.GetUserByNickname(data.Nickname)
		if err == nil {
			zap.L().Debug("昵称已被占用", zap.Uint("user_id", id), zap.String("nickname", data.Nickname))
			utils.JsonErrorResponse(c, 200507, "昵称已被占用")
			return
		} else if err != gorm.ErrRecordNotFound {
			zap.L().Error("查询用户信息失败", zap.Error(err))
			utils.JsonInternalServerErrorResponse(c)
			return
		}

		err = userService.SetNickname(id, data.Nickname)
		if err != nil {
			zap.L().Error("用户昵称设置失败", zap.Uint("user_id", id), zap.Error(err))
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}

	if data.Avatar != "" {
		err = userService.SetAvatar(id, data.Avatar)
		if err != nil {
			zap.L().Error("用户头像设置失败", zap.Uint("user_id", id), zap.Error(err))
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}

	utils.JsonSuccessResponse(c, nil)
	zap.L().Info("用户个人信息设置成功", zap.Uint("user_id", id), zap.String("nickname", data.Nickname), zap.String("avatar", data.Avatar))
}
