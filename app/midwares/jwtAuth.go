package midwares

import (
	"ConfessionWall/app/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth(c *gin.Context) {
	// 通过 header 中的 token 来认证
	token := c.Request.Header.Get("token")
	if token == "" { // 没有携带 token
		utils.JsonErrorResponse(c, 200502, "无权限访问")
		c.Abort() // 中止当前请求
		return
	}

	id, err := utils.ExtractToken(token)
	if err != nil {
		if err == utils.ErrTokenHandlingFailed { // token 处理失败
			utils.JsonInternalServerErrorResponse(c)
		} else { // token 无效
			utils.JsonErrorResponse(c, 200502, "密钥无效")
		}
		c.Abort() // 中止当前请求
		return
	}

	// 将 user_id 重新写入 gin.Context 对象中
	c.Set("user_id", id)
}
