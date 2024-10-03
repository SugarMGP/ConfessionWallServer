package userController

import (
	"ConfessionWall/app/apiException"
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
		c.AbortWithError(200, apiException.ParamsError)
		return
	}

	if data.Nickname == "" && data.Avatar == "" {
		c.AbortWithError(200, apiException.ParamsError)
		return
	}

	if data.Nickname != "" {
		if data.Nickname == "匿名用户" {
			c.AbortWithError(200, apiException.NicknameOccupied)
			return
		}

		if len(data.Nickname) > 16 {
			zap.L().Debug("用户昵称设置过长", zap.Uint("user_id", id), zap.Error(err))
			c.AbortWithError(200, apiException.NicknameTooLong)
			return
		}

		_, err := userService.GetUserByNickname(data.Nickname)
		if err == nil {
			zap.L().Debug("昵称已被占用", zap.Uint("user_id", id), zap.String("nickname", data.Nickname))
			c.AbortWithError(200, apiException.NicknameOccupied)
			return
		} else if err != gorm.ErrRecordNotFound {
			zap.L().Error("查询用户信息失败", zap.Error(err))
			c.AbortWithError(200, apiException.InternalServerError)
			return
		}

		err = userService.SetNickname(id, data.Nickname)
		if err != nil {
			zap.L().Error("用户昵称设置失败", zap.Uint("user_id", id), zap.Error(err))
			c.AbortWithError(200, apiException.InternalServerError)
			return
		}
	}

	if data.Avatar != "" {
		err = userService.SetAvatar(id, data.Avatar)
		if err != nil {
			zap.L().Error("用户头像设置失败", zap.Uint("user_id", id), zap.Error(err))
			c.AbortWithError(200, apiException.InternalServerError)
			return
		}
	}

	utils.JsonSuccessResponse(c, nil)
	zap.L().Info("用户个人信息设置成功", zap.Uint("user_id", id), zap.String("nickname", data.Nickname), zap.String("avatar", data.Avatar))
}
