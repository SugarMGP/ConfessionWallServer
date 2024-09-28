package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `json:"post_id"`
	Content   string         `json:"content" gorm:"type:text"`
	User      uint           `json:"-"`
	Unnamed   bool           `json:"unnamed"`
	PostTime  time.Time      `json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
