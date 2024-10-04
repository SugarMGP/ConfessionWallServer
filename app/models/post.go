package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint
	Content   string `gorm:"type:text"`
	UserID    uint
	Unnamed   bool
	Likes     int64 `gorm:"default:0"`
	PostTime  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Private   bool
}
