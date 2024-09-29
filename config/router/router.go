package router

import (
	"ConfessionWall/app/controllers/blockController"
	"ConfessionWall/app/controllers/commentController"
	"ConfessionWall/app/controllers/postController"
	"ConfessionWall/app/controllers/uploadController"
	"ConfessionWall/app/controllers/userController"
	"ConfessionWall/app/midwares"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.POST("/api/user/reg", userController.Register)
	r.POST("/api/user/login", userController.Login)

	api := r.Group("/api")
	api.Use(midwares.JWTAuth)
	{
		api.GET("/user", userController.GetProfile)
		api.PUT("/user", userController.SetProfile)

		api.GET("/my_confession", postController.GetMyPostList)
		api.GET("/confession", postController.GetPostList)
		api.POST("/confession", postController.NewPost)
		api.PUT("/confession", postController.UpdatePost)
		api.DELETE("/confession", postController.DeletePost)

		api.POST("/blacklist", blockController.NewBlock)
		api.DELETE("/blacklist", blockController.DeleteBlock)
		api.GET("/blacklist", blockController.GetBlacklist)
		api.POST("/confession/comment", commentController.NewComment)
		api.GET("/confession/comment", commentController.GetCommentList)
		api.DELETE("/confession/comment", commentController.DeleteComment)

		api.POST("/upload", uploadController.PictureUpload)
	}
}
