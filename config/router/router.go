package router

import (
	"ConfessionWall/app/controllers/blockController"
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
		api.GET("/confession", midwares.JWTAuth, postController.GetPostList)
		api.GET("/my_confession", midwares.JWTAuth, postController.GetMyPostList)
		api.PUT("/confession", midwares.JWTAuth, postController.UpdatePost)
		api.DELETE("/confession", midwares.JWTAuth, postController.DeletePost)

		api.POST("/blacklist", midwares.JWTAuth, blockController.NewBlock)
	}
}
