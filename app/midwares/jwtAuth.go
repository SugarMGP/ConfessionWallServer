package midwares

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JWTAuth(c *gin.Context) {
	// 通过 header 中的 Authorization 来认证
	token := c.Request.Header.Get("Authorization")
	if token == "" { // 没有携带 token
		zap.L().Debug("无权限访问", zap.String("Authorization", token))
		c.AbortWithError(200, apiException.NoAccessPermission)
		return
	}

	id, err := utils.ExtractToken(token)
	if err != nil {
		if err == utils.ErrTokenHandlingFailed { // Token 处理失败
			zap.L().Error("Token 处理失败", zap.String("Authorization", token), zap.Error(err))
			c.AbortWithError(200, apiException.InternalServerError)
		} else { // Token 无效
			zap.L().Debug("密钥无效", zap.String("Authorization", token), zap.Error(err))
			c.AbortWithError(200, apiException.NoAccessPermission)
		}
		return
	}

	// 将 user_id 重新写入 gin.Context 对象中
	c.Set("user_id", id)
	zap.L().Info("JWT认证成功", zap.Uint("user_id", id))
}
