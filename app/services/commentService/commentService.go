package commentService

import (
	"ConfessionWall/app/models"
	"ConfessionWall/config/database"
	"time"
)

type CommentService struct {
	Comment *models.Comment
}

// 在数据库中创建新的评论
func NewComment(comment models.Comment) error {
	result := database.DB.Create(&comment)
	return result.Error
}

// 创建一条新的评论
func CreateComment(postID uint, userID uint, content string) error {
	comment := models.Comment{
		PostID:    postID,
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return NewComment(comment)
}

func GetCommentList() (comments []models.Comment, err error) {
	//检索所有评论
	result := database.DB.Find(&comments)
	err = result.Error
	return
}

func GetCommentByID(id uint) (comment models.Comment, err error) {
	//根据ID检索评论
	result := database.DB.Where("id =?", id).First(&comment)
	err = result.Error
	return
}

// GetCommentsByPostID 获取某篇帖子的所有评论
func GetCommentsByPostID(postID uint) ([]models.Comment, error) {
	var comments []models.Comment
	result := database.DB.Where("post_id =?", postID).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

// UpdateComment 更新评论
func UpdateComment(comment models.Comment) error {
	result := database.DB.Save(&comment)
	return result.Error
}

// DeleteComment 删除评论
func DeleteComment(id uint) error {
	result := database.DB.Delete(&models.Comment{}, id)
	return result.Error
}
