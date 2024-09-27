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
		confession := api.Group("/confession")
		{
			confession.POST("/comment", midwares.JWTAuth, commentController.NewComment)
			confession.GET("/comment", midwares.JWTAuth, commentController.GetCommentsByPostID)
			confession.DELETE("/comment", midwares.JWTAuth, commentController.DeleteComment)
		}

		api.POST("/upload", midwares.JWTAuth, uploadController.PictureUpload)
	}
}
