package database

import (
	"ConfessionWall/app/models"

	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Block{},
		&models.Comment{},
	)
	return err
}
