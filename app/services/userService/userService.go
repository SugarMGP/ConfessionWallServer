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
	user.Nickname = "用户" + hashedUsername[:12]

	result := database.DB.Create(&user)
	return result.Error
}

func GetUserByID(id uint) (user models.User, err error) {
	result := database.DB.Where("id = ?", id).First(&user)
	err = result.Error
	return
}

func SetNickname(id uint, nickname string) error {
	result := database.DB.Where("id = ?", id).First(&models.User{}).Update("nickname", nickname)
	return result.Error
}

func SetAvatar(id uint, avatar string) error {
	result := database.DB.Where("id = ?", id).First(&models.User{}).Update("avatar", avatar)
	return result.Error
}

func GetUserByNickname(nickname string) (user models.User, err error) {
	result := database.DB.Where("nickname = ?", nickname).First(&user)
	err = result.Error
	return
}
