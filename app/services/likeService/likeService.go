package likeService

import (
	"ConfessionWall/app/models"
	"ConfessionWall/config/database"
	"ConfessionWall/config/rds"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func GetPostKey(PostID uint) string {
	return fmt.Sprintf("post:likes:%d", PostID)
}

func PostLikeCount(PostID uint) error {
	var keys string = GetPostKey(PostID)
	redisClient := rds.GetRedis()
	defer redisClient.Close()
	postCount := redis.BitCount{Start: 0, End: -1}
	err := redisClient.BitCount(context.Background(), keys, &postCount).Err()
	if err != nil {
		return err
	}
	result := database.DB.Where("id = ?", PostID).First(&models.Post{}).Update("likes", postCount)
	return result.Error
}
func GetCommentKey(CommentID uint) string {
	return fmt.Sprintf("comment:likes:%d", CommentID)
}
func CommentLikeCount(PostID uint, CommentID uint) error {
	var keys string = GetCommentKey(CommentID)
	redisClient := rds.GetRedis()
	defer redisClient.Close()
	commentCount := redis.BitCount{Start: 0, End: -1}
	err := redisClient.BitCount(context.Background(), keys, &commentCount).Err()
	if err != nil {
		return err
	}
	result := database.DB.Where("id = ? AND post_id = ?", CommentID, PostID).First(&models.Comment{}).Update("likes", commentCount)
	return result.Error
}
