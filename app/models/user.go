package models

type User struct {
	ID       uint
	Username string
	Nickname string
	Password string
	Avatar   string
	Activity int `gorm:"default:0"`
}
