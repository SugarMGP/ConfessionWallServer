package signService

import "fmt"

func GetIdKey(id uint) string {
	return fmt.Sprintf("id:%d", id)
}
