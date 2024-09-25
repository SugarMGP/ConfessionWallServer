package userService

import (
	"ConfessionWall/app/models"
	"ConfessionWall/config/database"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		zap.L().Error("密码验证失败", zap.Error(err))
		return err
	}
	return nil
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		zap.L().Error("获取用户失败", zap.String("username", username), zap.Error(result.Error))
		return models.User{}, result.Error
	}
	return user, nil
}

func Register(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("密码哈希失败", zap.Error(err))
		return err
	}
	user.Password = string(hashedPassword)

	result := database.DB.Create(&user)
	if result.Error != nil {
		zap.L().Error("注册用户失败", zap.String("username", user.Username), zap.Error(result.Error))
		return result.Error
	}
	return nil
}

func GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := database.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		zap.L().Error("获取用户失败", zap.Uint("user_id", id), zap.Error(result.Error))
		return models.User{}, result.Error
	}
	return user, nil
}
