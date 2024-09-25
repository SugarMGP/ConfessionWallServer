package models

import "time"

type Comment struct {
	Username  string    `json:"-"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"-"`
}
