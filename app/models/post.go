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
	PostTime  time.Time      `json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	Private   bool           `json:"private"`
}
