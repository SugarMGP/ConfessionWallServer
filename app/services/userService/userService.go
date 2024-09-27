package userService

import (
	"ConfessionWall/app/models"
	"ConfessionWall/config/database"
	"crypto/md5"
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

func GetUserByUsername(username string) (user models.User, err error) {
	result := database.DB.Where("username = ?", username).First(&user)
	err = result.Error
	return
}

func Register(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("密码哈希失败", zap.Error(err))
		return err
	}
	user.Password = string(hashedPassword)

	// 昵称默认为用户名的哈希值
	hashedUsername := fmt.Sprintf("%x", md5.Sum([]byte(user.Username)))
	user.Nickname = "用户" + hashedUsername

	result := database.DB.Create(&user)
	return result.Error
}

func GetUserByID(id uint) (user models.User, err error) {
	result := database.DB.Where("id = ?", id).First(&user)
	err = result.Error
	return
}
