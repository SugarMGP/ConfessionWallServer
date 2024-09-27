package uploadController

import (
	"ConfessionWall/app/services/uploadService"
	"ConfessionWall/app/utils"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UploadResponse struct {
	Url string `json:"url"`
}

func PictureUpload(c *gin.Context) {
	id := c.GetUint("user_id")

	file, err := c.FormFile("picture")
	if err != nil {
		zap.L().Error("图片获取失败", zap.Error(err))
		utils.JsonErrorResponse(c, 200506, "参数错误")
		return
	}

	if file.Size > 8*1024*1024 {
		utils.JsonErrorResponse(c, 200511, "文件过大")
		return
	}

	ext := filepath.Ext(file.Filename)

	// 打开文件
	src, err := file.Open()
	if err != nil {
		zap.L().Error("图片打开失败", zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	defer src.Close()

	hashString, err := uploadService.SumMD5(src)
	if err != nil {
		zap.L().Error("图片哈希计算失败", zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	filePath := "./static/" + hashString + ext
	url := "http://" + c.Request.Host + "/static/" + hashString + ext

	// 检查文件是否已存在
	if _, err := os.Stat(filePath); err == nil {
		utils.JsonSuccessResponse(c, UploadResponse{
			Url: url,
		})
		return
	}

	c.SaveUploadedFile(file, filePath)

	// 成功上传图片
	zap.L().Info("图片上传成功", zap.Uint("user_id", id), zap.String("path", filePath))
	utils.JsonSuccessResponse(c, UploadResponse{
		Url: url,
	})
}
