package postController

import (
	"ConfessionWall/app/models"
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NewPostData struct {
	Content  string `json:"content" binding:"required"`
	Unnamed  bool   `json:"unnamed"`
	PostUnix string `json:"post_unix"`
}

func NewPost(c *gin.Context) {
	id := c.GetUint("user_id")

	// 绑定请求数据
	var data NewPostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	postTime := time.Time{}
	if data.PostUnix != "" {
		unix, err := strconv.ParseInt(data.PostUnix, 10, 64)
		if err != nil {
			zap.L().Debug("string转换int64失败", zap.Error(err))
			utils.JsonErrorResponse(c, 200506, "参数错误")
			return
		}
		postTime = time.Unix(unix, 0)
	}

	// 新建帖子
	err = postService.NewPost(models.Post{
		Content:  data.Content,
		User:     id,
		Unnamed:  data.Unnamed,
		PostTime: postTime,
	})
	if err != nil {
		zap.L().Error("新建帖子失败", zap.Uint("user_id", id), zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 成功创建帖子
	zap.L().Info("帖子创建成功", zap.Uint("user_id", id), zap.String("content", data.Content))
	utils.JsonSuccessResponse(c, nil)
}
