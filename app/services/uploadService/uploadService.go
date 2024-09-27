package uploadService

import (
	"crypto/md5"
	"fmt"
	"io"
)

func SumMD5(file io.Reader) (string, error) {
	// 读取文件内容
	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	// 计算MD5
	hash := md5.Sum(data)
	hashStr := fmt.Sprintf("%x", hash)
	return hashStr, nil
}
