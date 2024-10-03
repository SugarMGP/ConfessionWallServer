package likeService

import (
	"ConfessionWall/app/models"
	"ConfessionWall/config/database"
	"ConfessionWall/config/rds"
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func GetPostKey(PostID uint) string {
	return fmt.Sprintf("post:likes:%d", PostID)
}

func GetCommentKey(CommentID uint) string {
	return fmt.Sprintf("comment:likes:%d", CommentID)
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

func CommentLikeCount(CommentID uint) error {
	var keys string = GetCommentKey(CommentID)
	redisClient := rds.GetRedis()
	defer redisClient.Close()

	commentCount := redis.BitCount{Start: 0, End: -1}
	err := redisClient.BitCount(context.Background(), keys, &commentCount).Err()
	if err != nil {
		return err
	}
	result := database.DB.Where("id = ?", CommentID).First(&models.Comment{}).Update("likes", commentCount)
	return result.Error
}
