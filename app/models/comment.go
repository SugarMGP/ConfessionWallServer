package models

// Comment 评论模型

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint
	PostID    uint
	UserID    uint
	Content   string `gorm:"type:text"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
