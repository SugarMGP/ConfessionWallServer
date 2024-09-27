package models

// Comment 评论模型

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint           `json:"comment_id"`
	PostID    uint           `json:"-"`
	UserID    uint           `json:"user_id"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
