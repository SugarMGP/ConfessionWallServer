package models

// Comment 评论模型

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint
	PostID    uint `json:"post_id"`
	UserID    uint
	Content   string `gorm:"type:text"`
	Likes     int64  `json:"likes" gorm:"default:0"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
