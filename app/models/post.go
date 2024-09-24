package models

import "time"

type Post struct {
	ID        uint      `json:"post_id"`
	Content   string    `json:"content"`
	User      uint      `json:"user_id"`
	Unnamed   bool      `json:"unnamed"`
	CreatedAt time.Time `json:"time"`
}
