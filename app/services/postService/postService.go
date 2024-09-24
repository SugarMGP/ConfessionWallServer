package postService

import (
	"ConfessionWall/app/models"
	"ConfessionWall/config/database"
)

func NewPost(post models.Post) error {
	result := database.DB.Create(&post)
	return result.Error
}
