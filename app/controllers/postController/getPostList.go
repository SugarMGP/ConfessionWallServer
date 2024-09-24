package postController

import (
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/services/userService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
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
	// TODO: 黑名单
	// id := c.GetUint("user_id")

	postList, err := postService.GetPostList()
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 创建一个Confession数组
	// 遍历postList，将信息填入数组中
	// 最后将数组序列化为json响应
	var confessionList []Confession = make([]Confession, 0)
	for _, post := range postList {
		nickname := ""
		if !post.Unnamed {
			user, err := userService.GetUserByID(post.User)
			if err == nil { // 如果能获取到用户
				nickname = user.Nickname
			}
		}
		confession := Confession{
			ID:       post.ID,
			Nickname: nickname,
			Content:  post.Content,
		}
		confessionList = append(confessionList, confession)
	}

	utils.JsonSuccessResponse(c, GetListResponse{
		ConfessionList: confessionList,
	})
}
