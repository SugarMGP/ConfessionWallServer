package models

type User struct {
	ID       uint   `json:"user_id"`
	Username string `json:"-"`
	Nickname string `json:"-"`
	Password string `json:"-"`
}
