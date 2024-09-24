package midwares

import (
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JWTAuth(c *gin.Context) {
	// 通过 header 中的 Authorization 来认证
	token := c.Request.Header.Get("Authorization")
	if token == "" { // 没有携带 token
		zap.L().Error("无权限访问", zap.String("Authorization", token))
		utils.JsonErrorResponse(c, 200502, "无权限访问")
		c.Abort() // 中止当前请求
		return
	}

	id, err := utils.ExtractToken(token)
	if err != nil {
		if err == utils.ErrTokenHandlingFailed { // token 处理失败
			zap.L().Error("token 处理失败", zap.String("Authorization", token), zap.Error(err))
			utils.JsonInternalServerErrorResponse(c)
		} else { // token 无效
			zap.L().Error("密钥无效", zap.String("Authorization", token), zap.Error(err))
			utils.JsonErrorResponse(c, 200502, "密钥无效")
		}
		c.Abort() // 中止当前请求
		return
	}

	// 将 user_id 重新写入 gin.Context 对象中
	c.Set("user_id", id)
	zap.L().Info("JWT认证成功", zap.Uint("user_id", id))
}
