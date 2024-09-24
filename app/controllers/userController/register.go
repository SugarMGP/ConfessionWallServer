package userController

import (
	"ConfessionWall/app/models"
	"ConfessionWall/app/services/userService"
	"ConfessionWall/app/utils"
	"regexp"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegDate struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var data RegDate
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	// 用户名校验
	if !isUsernameValid(data.Username) {
		utils.JsonErrorResponse(c, 200505, "用户名不符合规范")
		return
	}

	// 密码校验
	if !isPasswordValid(data.Password) {
		utils.JsonErrorResponse(c, 200505, "密码不符合规范")
		return
	}

	// 判断用户是否已经注册
	_, err = userService.GetUserByUsername(data.Username)
	if err == nil {
		utils.JsonErrorResponse(c, 200507, "用户名已存在")
		return
	} else if err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 注册用户
	err = userService.Register(models.User{
		Username: data.Username,
		Password: data.Password,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

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
