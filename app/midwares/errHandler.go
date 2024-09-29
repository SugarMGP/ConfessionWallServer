package midwares

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			if err != nil {
				var Err *apiException.Error

				if e, ok := err.(*apiException.Error); ok {
					Err = e
				} else {
					Err = apiException.InternalServerError
				}

				utils.JsonErrorResponse(c, Err.Code, Err.Msg)
				return
			}
		}
	}
}

// 404处理
func HandleNotFound(c *gin.Context) {
	// 记录未找到的请求信息
	zap.L().Warn("未找到请求路径", zap.String("path", c.Request.URL.Path))

	// 返回404 Not Found响应
	utils.JsonResponse(c, 404, 200404, http.StatusText(http.StatusNotFound), nil)
}
