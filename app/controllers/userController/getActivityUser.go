package userController

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/services/activityServive"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetActivityUserResponse struct {
	ActivityUserList []ActivityUserElement `json:"activity_user_list"`
}
type ActivityUserElement struct {
	ID       uint
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Activity int    `json:"activity"`
}

func GetActivityUserList(c *gin.Context) {
	activityUserList, err := activityServive.GetActivityUser()
	if err != nil {
		zap.L().Error("获取活跃用户列表失败", zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}
	List := make([]ActivityUserElement, 0)
	for _, user := range activityUserList {
		userList := ActivityUserElement{
			ID:       user.ID,
			Username: user.Username,
			Avatar:   user.Avatar,
			Activity: user.Activity,
		}
		List = append(List, userList)
	}

	// 成功获取帖子列表
	zap.L().Info("获取活跃用户列表成功", zap.Int("count", len(List)))
	utils.JsonSuccessResponse(c, GetActivityUserResponse{
		ActivityUserList: List,
	})

}
