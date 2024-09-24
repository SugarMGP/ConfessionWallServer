package userController

import (
	"ConfessionWall/app/models"
	"ConfessionWall/app/services/userService"
	"ConfessionWall/app/utils"
	"regexp"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RegData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var data RegData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	// 用户名校验
	if !isUsernameValid(data.Username) {
		zap.L().Error("用户名不符合规范", zap.String("username", data.Username))
		utils.JsonErrorResponse(c, 200505, "用户名不符合规范")
		return
	}

	// 密码校验
	if !isPasswordValid(data.Password) {
		zap.L().Error("密码不符合规范", zap.String("password", data.Password))
		utils.JsonErrorResponse(c, 200505, "密码不符合规范")
		return
	}

	// 判断用户是否已经注册
	_, err = userService.GetUserByUsername(data.Username)
	if err == nil {
		zap.L().Error("用户名已存在", zap.String("username", data.Username))
		utils.JsonErrorResponse(c, 200507, "用户名已存在")
		return
	} else if err != gorm.ErrRecordNotFound {
		zap.L().Error("查询用户信息失败", zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 注册用户
	err = userService.Register(models.User{
		Username: data.Username,
		Password: data.Password,
	})
	if err != nil {
		zap.L().Error("注册用户失败", zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 成功注册用户
	zap.L().Info("用户注册成功", zap.String("username", data.Username))
	utils.JsonSuccessResponse(c, nil)
}

// 用户名正则，4到16位（字母，数字，下划线，减号）
func isUsernameValid(username string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9_-]{4,16}$`)
	return regex.MatchString(username)
}

// 密码正则，8到32位，至少一个小写字母和一个数字，可包含大小写字母、数字和特殊符号
func isPasswordValid(password string) bool {
	// 检查至少包含一个小写字母
	lowercaseRegex := regexp.MustCompile(`(?i)[a-z]`)
	if !lowercaseRegex.MatchString(password) {
		return false
	}

	// 检查至少包含一个数字
	digitRegex := regexp.MustCompile(`\d`)
	if !digitRegex.MatchString(password) {
		return false
	}

	// 检查长度
	if len(password) < 8 || len(password) > 32 {
		return false
	}

	// 检查只包含字母、数字和特殊符号
	allowedCharsRegex := regexp.MustCompile(`^[A-Za-z\d\W]+$`)
	return allowedCharsRegex.MatchString(password)
}
