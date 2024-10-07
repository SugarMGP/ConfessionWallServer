package activityServive

import (
	"ConfessionWall/app/models"
	"ConfessionWall/config/database"
)

func ChangeActivity(id uint, point int) error {
	var user models.User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	user.Activity += point

	result := database.DB.Save(&user)
	return result.Error
}

func GetActivityUser() (users []models.User, err error) {
	result := database.DB.Where("activity > ?", 0).Order("activity desc").Find(&users)
	err = result.Error
	return
}
