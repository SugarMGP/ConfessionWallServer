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

		api.GET("/my_confession", midwares.JWTAuth, postController.GetMyPostList)
		api.GET("/confession", midwares.JWTAuth, postController.GetPostList)
		api.POST("/confession", midwares.JWTAuth, postController.NewPost)
		api.PUT("/confession", midwares.JWTAuth, postController.UpdatePost)
		api.DELETE("/confession", midwares.JWTAuth, postController.DeletePost)

		api.POST("/blacklist", midwares.JWTAuth, blockController.NewBlock)
		api.POST("/confession/comment", midwares.JWTAuth, commentController.NewComment)
		api.GET("/confession/comment", midwares.JWTAuth, commentController.GetCommentList)
		api.DELETE("/confession/comment", midwares.JWTAuth, commentController.DeleteComment)

		api.POST("/upload", midwares.JWTAuth, uploadController.PictureUpload)
	}
}
