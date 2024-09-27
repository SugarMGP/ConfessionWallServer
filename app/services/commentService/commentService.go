package commentService

import (
	"ConfessionWall/app/models"
	"ConfessionWall/config/database"
)

// 在数据库中创建新的评论
func NewComment(comment models.Comment) error {
	result := database.DB.Create(&comment)
	return result.Error
}

// GetCommentsByPostID 获取某篇帖子的所有评论
func GetCommentsByPostID(postID uint) (comments []models.Comment, err error) {
	result := database.DB.Where("post_id =?", postID).Find(&comments)
	err = result.Error
	return
}

// GetCommentByID 根据ID获取评论
func GetCommentByID(id uint) (comment models.Comment, err error) {
	result := database.DB.Where("id =?", id).Find(&comment)
	err = result.Error
	return
}

// DeleteComment 删除评论
func DeleteComment(id uint) error {
	result := database.DB.Where("id = ?", id).Delete(&models.Comment{})
	return result.Error
}
