package models

// Comment 评论模型

import "time"

type Comment struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	PostID    int       `gorm:"not null"`
	UserID    int       `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
