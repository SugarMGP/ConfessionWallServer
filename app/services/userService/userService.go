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

	zap.L().Info("密码验证成功")
	return nil
}

func GetUserByUsername(username string) (user models.User, err error) {
	result := database.DB.Where("username = ?", username).First(&user)
	err = result.Error
	if err != nil {
		zap.L().Error("获取用户失败", zap.String("username", username), zap.Error(err))
		return models.User{}, err
	}

	zap.L().Info("获取用户成功", zap.String("username", username))
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

	zap.L().Info("注册用户成功", zap.String("username", user.Username))
	return nil
}

func GetUserByID(id uint) (user models.User, err error) {
	result := database.DB.Where("id = ?", id).First(&user)
	err = result.Error
	if err != nil {
		zap.L().Error("获取用户失败", zap.Uint("user_id", id), zap.Error(err))
		return models.User{}, err
	}

	zap.L().Info("获取用户成功", zap.Uint("user_id", id))
	return user, nil
}
