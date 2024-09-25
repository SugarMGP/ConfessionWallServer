package userController

import (
	"ConfessionWall/app/services/userService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("参数绑定失败", zap.Error(err))
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	// 判断用户是否存在并获取用户信息
	user, err := userService.GetUserByUsername(data.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.L().Debug("用户不存在", zap.String("username", data.Username))
			utils.JsonErrorResponse(c, 200501, "用户不存在")
		} else {
			zap.L().Error("获取用户信息失败", zap.Error(err))
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	// 判断密码是否正确
	err = userService.VerifyPassword(data.Password, user.Password)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			zap.L().Debug("密码错误", zap.String("username", data.Username))
			utils.JsonErrorResponse(c, 200501, "密码错误")
		} else {
			zap.L().Error("密码验证失败", zap.Error(err))
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		zap.L().Error("生成 Token 失败", zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	response := LoginResponse{
		Token: token,
	}

	// 返回用户信息
	utils.JsonSuccessResponse(c, response)
	zap.L().Info("用户登录成功", zap.String("username", data.Username))
}
