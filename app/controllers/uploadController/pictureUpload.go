package uploadController

import (
	"ConfessionWall/app/services/uploadService"
	"ConfessionWall/app/utils"
	"image"
	"image/jpeg"
	"os"

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

	filePath := "./static/" + hashString + ".webp"

	// 检查文件是否已存在
	if _, err := os.Stat(filePath); err == nil {
		return
	}
	output, err := os.Create(filePath)
	if err != nil {
		zap.L().Error("文件创建失败", zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	defer output.Close()

	// 解码图片
	img, _, err := image.Decode(src)
	if err != nil {
		zap.L().Error("图片解码失败", zap.Error(err))
		utils.JsonErrorResponse(c, 200512, "文件无法被解码为图片")
		return
	}

	err = jpeg.Encode(output, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		zap.L().Error("图片编码失败", zap.Error(err))
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 成功上传图片
	zap.L().Info("图片上传成功", zap.Uint("user_id", id), zap.String("path", filePath))
	utils.JsonSuccessResponse(c, UploadResponse{
		Url: "/static/" + hashString + ".webp",
	})
}
