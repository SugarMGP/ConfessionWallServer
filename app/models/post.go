package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct { //json tag for "GetMyPostList"
	ID        uint           `json:"post_id"`
	Content   string         `json:"content" gorm:"type:text"`
	UserID    uint           `json:"-"`
	Unnamed   bool           `json:"unnamed"`
	Likes     int64          `json:"likes" gorm:"default:0"`
	PostTime  time.Time      `json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Private   bool           `json:"private"`
}
