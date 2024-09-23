package database

import (
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
	// &models.User{},
	)
	return err
}
