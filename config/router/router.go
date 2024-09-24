package router

import (
	"ConfessionWall/app/controllers/userController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/user/reg", userController.Register)
		api.POST("/user/login", userController.Login)
		// api.POST("/user/reg", userController.Register)
		// api.GET("/student/post", postController.GetAllPosts)

		// api.POST("/student/post", midwares.JWTAuth, postController.NewPost)
		// api.DELETE("/student/post", midwares.JWTAuth, postController.DeletePost)
		// api.PUT("/student/post", midwares.JWTAuth, postController.EditPost)
	}
}
