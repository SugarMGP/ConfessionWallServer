package activityServive

import (
	"ConfessionWall/app/models"
	"ConfessionWall/config/database"
)

func IncreaseActivity(id uint, actType int) error {
	var user models.User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	switch actType {
	case 1:
		user.Activity++ //被解除拉黑
	case 2:
		user.Activity += 2 // 评论 签到
	case 3:
		user.Activity += 5 //发帖
	}
	result := database.DB.Save(&user)
	return result.Error
}

func DecreaseActivity(id uint, actType int) error {
	var user models.User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	switch actType {
	case 1:
		user.Activity-- //被拉黑
	case 2:
		user.Activity -= 2 //删评论
	case 3:
		user.Activity -= 5 //删帖
	}
	result := database.DB.Save(&user)
	return result.Error
}
func GetActivityUser() (users []models.User, err error) {
	result := database.DB.Where("activity != ?", 0).Order("activity desc").Find(&users)
	err = result.Error
	return
}
