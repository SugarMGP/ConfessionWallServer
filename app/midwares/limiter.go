package midwares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func Limiter() gin.HandlerFunc {
	// 每秒允许3个请求
	rate := limiter.Rate{
		Period: 1 * time.Second,
		Limit:  3,
	}

	// 内存存储
	store := memory.NewStore()

	// 创建限流中间件
	return mgin.NewMiddleware(limiter.New(store, rate), mgin.WithKeyGetter(func(c *gin.Context) string {
		token := c.Request.Header.Get("Authorization")

		// 如果没有Token，则按IP进行限流
		if token == "" {
			return c.ClientIP()
		}

		// 按Token进行限流
		return token
	}))
}
