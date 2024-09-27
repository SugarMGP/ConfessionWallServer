package uploadService

import (
	"crypto/md5"
	"fmt"
)

func SumMD5(data []byte) (string, error) {
	hash := md5.Sum(data)
	hashStr := fmt.Sprintf("%x", hash)
	return hashStr, nil
}
