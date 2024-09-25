package models

// Comment 评论模型

import "time"

type Comment struct {
	PostID    int       
	UserID    int       
	Content   string    
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
