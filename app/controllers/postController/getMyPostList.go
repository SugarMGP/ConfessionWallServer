package postController

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetMyListResponse struct {
	MyConfessionList []MyConfessionElement `json:"my_confession_list"`
}

type MyConfessionElement struct {
	ID       uint   `json:"post_id"`
	Content  string `json:"content"`
	Unnamed  bool   `json:"unnamed"`
	Likes    int64  `json:"likes"`
	PostUnix int64  `json:"post_unix"`
	Private  bool   `json:"private"`
}

func GetMyPostList(c *gin.Context) {
	id := c.GetUint("user_id")

	// 获取帖子列表
	postList, err := postService.GetMyPostList(id)
	if err != nil {
		zap.L().Error("获取帖子列表失败", zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	confessionList := make([]MyConfessionElement, 0)
	for _, post := range postList {
		myConfession := MyConfessionElement{
			ID:       post.ID,
			Content:  post.Content,
			Unnamed:  post.Unnamed,
			Likes:    post.Likes,
			PostUnix: post.PostTime.Unix(),
			Private:  post.Private,
		}
		confessionList = append(confessionList, myConfession)
	}

	// 成功获取帖子列表
	zap.L().Info("获取帖子列表成功", zap.Int("count", len(postList)))
	utils.JsonSuccessResponse(c, GetMyListResponse{
		MyConfessionList: confessionList,
	})
}
