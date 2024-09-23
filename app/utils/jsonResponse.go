package utils

import (
	"github.com/gin-gonic/gin"
)

func JsonResponse(c *gin.Context, httpStatusCode int, code int, msg string, data interface{}) {
	c.JSON(httpStatusCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func JsonSuccessResponse(c *gin.Context, data interface{}) {
	JsonResponse(c, 200, 200, "success", data)
}

func JsonErrorResponse(c *gin.Context, code int, msg string) {
	JsonResponse(c, 200, code, msg, nil)
}

func JsonInternalServerErrorResponse(c *gin.Context) {
	JsonResponse(c, 200, 200500, "Internal server error", nil)
}
