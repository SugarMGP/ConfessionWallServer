package router

import (
	"ConfessionWall/app/controllers/postController"
	"ConfessionWall/app/controllers/userController"
	"ConfessionWall/app/midwares"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/user/reg", userController.Register)
		api.POST("/user/login", userController.Login)

		api.POST("/confession", midwares.JWTAuth, postController.NewPost)
	}
}
