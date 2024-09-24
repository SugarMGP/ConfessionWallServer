package userService

import (
	"ConfessionWall/app/models"
	"ConfessionWall/config/database"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(password, hashedPassword string) error {
	// 将密码与hash密码进行比较
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetUserByUsername(username string) (user models.User, err error) {
	result := database.DB.Where("username = ?", username).First(&user)
	err = result.Error
	return
}

func Register(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	result := database.DB.Create(&user)
	return result.Error
}

func GetUserByID(id uint) (user models.User, err error) {
	result := database.DB.Where("id = ?", id).First(&user)
	err = result.Error
	return
}
