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

	zap.L().Info("创建帖子成功", zap.Uint("post_id", post.ID))
	return nil
}

func GetPostList() (posts []models.Post, err error) {
	result := database.DB.Find(&posts)
	if result.Error != nil {
		zap.L().Error("获取帖子列表失败", zap.Error(result.Error))
		return nil, result.Error
	}

	zap.L().Info("获取帖子列表成功", zap.Int("count", len(posts)))
	return posts, nil
}

func GetPostByID(id uint) (post models.Post, err error) {
	result := database.DB.Where("id = ?", id).First(&post)
	if result.Error != nil {
		zap.L().Error("获取帖子失败", zap.Uint("post_id", id), zap.Error(result.Error))
		return models.Post{}, result.Error
	}

	zap.L().Info("获取帖子成功", zap.Uint("post_id", post.ID))
	return post, nil
}

func DeletePost(id uint) error {
	result := database.DB.Where("id = ?", id).Delete(&models.Post{})
	if result.Error != nil {
		zap.L().Error("删除帖子失败", zap.Uint("post_id", id), zap.Error(result.Error))
		return result.Error
	}

	zap.L().Info("删除帖子成功", zap.Uint("post_id", id))
	return nil
}

func UpdatePost(id uint, content string) error {
	result := database.DB.Where("id = ?", id).First(&models.Post{}).Update("content", content)
	if result.Error != nil {
		zap.L().Error("更新帖子失败", zap.Uint("post_id", id), zap.Error(result.Error))
		return result.Error
	}

	zap.L().Info("更新帖子成功", zap.Uint("post_id", id))
	return nil
}
