package uploadController

import (
	"ConfessionWall/app/apiException"
	"ConfessionWall/app/utils"
	"crypto/md5"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/h2non/filetype"
	"go.uber.org/zap"
)

type UploadResponse struct {
	Url string `json:"url"`
}

func PictureUpload(c *gin.Context) {
	id := c.GetUint("user_id")

	file, err := c.FormFile("picture")
	if err != nil {
		zap.L().Error("文件获取失败", zap.Error(err))
		c.AbortWithError(200, apiException.ParamsError)
		return
	}

	if file.Size > 4*1024*1024 {
		c.AbortWithError(200, apiException.FileTooLarge)
		return
	}

	// 获取文件扩展名
	ext := filepath.Ext(file.Filename)

	// 打开文件
	src, err := file.Open()
	if err != nil {
		zap.L().Error("文件打开失败", zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}
	defer src.Close()

	// 读取文件内容
	data, err := io.ReadAll(src)
	if err != nil {
		zap.L().Error("文件读取失败", zap.Error(err))
		c.AbortWithError(200, apiException.InternalServerError)
		return
	}

	// 检查文件类型
	if !filetype.IsImage(data) {
		c.AbortWithError(200, apiException.FileTypeError)
		return
	}

	// 计算图片哈希
	hashString := fmt.Sprintf("%x", md5.Sum(data))

	filePath := "./static/" + hashString + ext
	host, _, _ := net.SplitHostPort(c.Request.Host)
	url := "http://" + host + ":8080" + "/static/" + hashString + ext

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
