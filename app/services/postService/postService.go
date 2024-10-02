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
	result := database.DB.Order("post_time desc").Where("private = ?", false).Find(&posts)
	err = result.Error
	return

}

func GetMyPostList(user uint) (posts []models.Post, err error) {
	result := database.DB.Where("user_id = ?", user).Find(&posts)
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

func UpdatePost(id uint, content string, unnamed bool, private bool) error {
	result := database.DB.Where("id = ?", id).First(&models.Post{}).Updates(map[string]interface{}{"content": content, "unnamed": unnamed, "private": private})
	return result.Error
}
