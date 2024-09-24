package postController

import (
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/services/userService"
	"ConfessionWall/app/utils"

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
}

func GetPostList(c *gin.Context) {
	// 获取帖子列表
	postList, err := postService.GetPostList()
	if err != nil {
		zap.L().Error("获取帖子列表失败", zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 创建一个Confession数组
	// 遍历postList，将信息填入数组中
	var confessionList []Confession
	for _, post := range postList {
		nickname := ""
		if !post.Unnamed {
			user, err := userService.GetUserByID(post.User)
			if err == nil { // 如果能获取到用户
				nickname = user.Nickname
			} else {
				zap.L().Error("获取用户信息失败", zap.Uint("user_id", post.User), zap.Error(err))
			}
		}
		confession := Confession{
			ID:       post.ID,
			Nickname: nickname,
			Content:  post.Content,
		}
		confessionList = append(confessionList, confession)
	}

	// 成功获取帖子列表
	zap.L().Info("获取帖子列表成功", zap.Int("count", len(postList)))
	utils.JsonSuccessResponse(c, GetListResponse{
		ConfessionList: confessionList,
	})
}
