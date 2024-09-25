package postService

import (
	"ConfessionWall/app/models"
	"ConfessionWall/config/database"

	"go.uber.org/zap"
)

func NewPost(post models.Post) error {
	result := database.DB.Create(&post)
	if result.Error != nil {
		zap.L().Error("创建帖子失败", zap.Error(result.Error))
		return result.Error
	}
	return nil
}

func GetPostList() ([]models.Post, error) {
	var posts []models.Post
	result := database.DB.Find(&posts)
	if result.Error != nil {
		zap.L().Error("获取帖子列表失败", zap.Error(result.Error))
		return nil, result.Error
	}
	return posts, nil
}

func GetPostByID(id uint) (models.Post, error) {
	var post models.Post
	result := database.DB.Where("id = ?", id).First(&post)
	if result.Error != nil {
		zap.L().Error("获取帖子失败", zap.Uint("post_id", id), zap.Error(result.Error))
		return models.Post{}, result.Error
	}
	return post, nil
}

func DeletePost(id uint) error {
	result := database.DB.Where("id = ?", id).Delete(&models.Post{})
	if result.Error != nil {
		zap.L().Error("删除帖子失败", zap.Uint("post_id", id), zap.Error(result.Error))
		return result.Error
	}
	return nil
}

func UpdatePost(id uint, content string) error {
	result := database.DB.Where("id = ?", id).First(&models.Post{}).Update("content", content)
	if result.Error != nil {
		zap.L().Error("更新帖子失败", zap.Uint("post_id", id), zap.Error(result.Error))
		return result.Error
	}
	return nil
}
