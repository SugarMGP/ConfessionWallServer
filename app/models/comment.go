package models

// Comment 评论模型

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	PostID    uint
	UserID    uint
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
