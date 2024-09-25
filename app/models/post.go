package models

import "time"

type Post struct {
	ID        uint      `json:"post_id"`
	Content   string    `json:"content"`
	User      uint      `json:"-"`
	Unnamed   bool      `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
