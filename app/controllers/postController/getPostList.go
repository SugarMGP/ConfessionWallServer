package postController

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/services/blockService"
	"ConfessionWall/app/services/likeService"
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/services/userService"
	"ConfessionWall/app/utils"
	"ConfessionWall/config/rds"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetListResponse struct {
	ConfessionList []Confession `json:"confession_list"`
}

type Confession struct {
	ID       uint   `json:"post_id"`
	Nickname string `json:"nickname"`
	Content  string `json:"content"`
	Likes    int64  `json:"likes"`
	Avatar   string `json:"avatar"`
	IsLiked  bool   `json:"is_liked"`
}

func GetPostList(c *gin.Context) {
	id := c.GetUint("user_id")

	// 获取帖子列表
	postList, err := postService.GetPostList()
	if err != nil {
		zap.L().Error("获取帖子列表失败", zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	blocks, err := blockService.GetBlocksByUserID(id)
	if err != nil {
		zap.L().Error("获取拉黑列表失败", zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	// 遍历postList，将信息填入Confession数组中
	confessionList := make([]Confession, 0)
	for _, post := range postList {
		// 判断是否被屏蔽
		blocked := false
		for _, block := range blocks {
			if block.TargetID == post.UserID {
				blocked = true
				break
			}
		}
		if blocked {
			continue
		}

		// 判断是否到达发布时间
		if post.PostTime.After(time.Now()) {
			continue
		}

		// 获取用户信息
		nickname := ""
		avatar := ""
		if !post.Unnamed {
			user, err := userService.GetUserByID(post.UserID)
			if err == nil { // 如果能获取到用户
				nickname = user.Nickname
				avatar = user.Avatar
			} else {
				zap.L().Error("获取用户信息失败", zap.Uint("post_id", post.ID), zap.Uint("user_id", post.UserID), zap.Error(err))
			}
		}

		// 获取点赞状态
		postKey := likeService.GetPostKey(post.ID)
		redis := rds.GetRedis()
		ctx := context.Background()

		res, err := redis.GetBit(ctx, postKey, int64(id)-1).Result()
		if err != nil {
			zap.L().Error("从Redis获取点赞状态失败", zap.Uint("post_id", post.ID), zap.Uint("user_id", id))
		}

		confession := Confession{
			ID:       post.ID,
			Nickname: nickname,
			Content:  post.Content,
			Avatar:   avatar,
			Likes:    post.Likes,
			IsLiked:  res == 1,
		}
		confessionList = append(confessionList, confession)
	}

	// 成功获取帖子列表
	zap.L().Info("获取帖子列表成功", zap.Int("count", len(postList)))
	utils.JsonSuccessResponse(c, GetListResponse{
		ConfessionList: confessionList,
	})
}
