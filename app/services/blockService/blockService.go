package blockService

import (
	"ConfessionWall/app/models"
	"ConfessionWall/config/database"
)

func NewBlock(block models.Block) error {
	result := database.DB.Create(&block)
	return result.Error
}

func GetBlockByID(id uint, target uint) (block models.Block, err error) {
	result := database.DB.Where("user_id = ?", id).Where("target_id = ?", target).First(&block)
	err = result.Error
	return
}

func GetBlocksByUserID(id uint) (blocks []models.Block, err error) {
	result := database.DB.Where("user_id = ?", id).Find(&blocks)
	err = result.Error
	return
}

func DeleteBlock(user uint, target uint) error {
	result := database.DB.Where("user_id = ?", user).Where("target_id = ?", target).Delete(&models.Block{})
	return result.Error
}
