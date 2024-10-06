package postController

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/models"
	"ConfessionWall/app/services/activityServive"
	"ConfessionWall/app/services/postService"
	"ConfessionWall/app/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NewPostData struct {
	Content  string `json:"content" binding:"required"`
	Unnamed  bool   `json:"unnamed"`
	PostUnix int64  `json:"post_unix"`
	Private  bool   `json:"private"`
}

func NewPost(c *gin.Context) {
	id := c.GetUint("user_id")

	// 绑定请求数据
	var data NewPostData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		zap.L().Error("请求数据绑定失败", zap.Error(err))
		c.AbortWithError(200, apiException.ParamsError)
		return
	}

	if len(data.Content) > 5000 {
		zap.L().Debug("帖子内容过长", zap.Uint("user_id", id), zap.Int("length", len(data.Content)))
		c.AbortWithError(200, apiException.ContentTooLong)
		return
	}

	postTime := time.Now()
	if data.PostUnix != 0 {
		postTime = time.Unix(data.PostUnix, 0)
	}

	// 新建帖子
	err = postService.NewPost(models.Post{
		Content:  data.Content,
		UserID:   id,
		Unnamed:  data.Unnamed,
		PostTime: postTime,
		Private:  data.Private,
	})
	if err != nil {
		zap.L().Error("新建帖子失败", zap.Uint("user_id", id), zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}
	err = activityServive.IncreaseActivity(id, 3)
	if err != nil {
		zap.L().Error("增加活跃度失败", zap.Uint("user_id", id), zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}
	// 成功创建帖子
	zap.L().Info("帖子创建成功", zap.Uint("user_id", id), zap.String("content", data.Content))
	utils.JsonSuccessResponse(c, nil)
}
