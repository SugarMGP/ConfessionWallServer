package postService

import (
	"ConfessionWall/app/models"
	"ConfessionWall/config/database"
)

func NewPost(post models.Post) error {
	result := database.DB.Create(&post)
	return result.Error
}

func GetPostList() (posts []models.Post, err error) {
	result := database.DB.Find(&posts)
	err = result.Error
	return
}

func GetPostByID(id uint) (post models.Post, err error) {
	result := database.DB.Where("id = ?", id).First(&post)
	err = result.Error
	return
}

func DeletePost(id uint) error {
	result := database.DB.Where("id = ?", id).Delete(&models.Post{})
	return result.Error
}

func UpdatePost(id uint, content string) error {
	result := database.DB.Where("id = ?", id).First(&models.Post{}).Update("content", content)
	return result.Error
}
