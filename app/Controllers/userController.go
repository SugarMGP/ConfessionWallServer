
package Controllers

import (
	"easy-forum/"
	"easy-forum/"

	"github.com/gin-gonic/gin"
)
type RegDate struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	
}
func Reg(c *gin.Context){
    var data RegDate
    err := c.ShouldBind(&data)
    if err != nil{
	utils.JsonErrorResponse(c, 200506, "参数错误")   //格式错误
		return
    }
    // 用户名校验
    if !isUsernameValid(data.Username) {
        utils.JsonErrorResponse(c, 200501, "太过简但，用户名必须包含字母和特殊符号")
        return
    }
    // 密码长度校验
    if len(data.Password) < 16 || len(data.Password) > 32 {
        utils.JsonErrorResponse(c, 200501, "你的密码不够安全，密码长度必须在16-32位")
        return
    }
    if !isPasswordValid(data.Password) {
        utils.JsonErrorResponse(c, 200501, "你的密码不够安全，密码必须包含字母和数字")
        return
    }
     c.JSON(http.StatusOK, gin.H{
        "code":    200,
        "message": "用户注册成功",
    })
    //函数部分
    // isUsernameValid 检查用户名是否包含字母和特殊符号
func isUsernameValid(username string) bool {
    regex := regexp.MustCompile(`^(?=.*[a-zA-Z])(?=.*[\W]).+$`)
    return regex.MatchString(username)
}
    // isPasswordValid 检查密码是否包含字母和数字
func isPasswordValid(password string) bool {
    regex := regexp.MustCompile(`^(?=.*[a-zA-Z])(?=.*\d).+$`)
    return regex.MatchString(password)
}
func main() {
    r := gin.Default()
    r.POST("/register", RegisterUser)
    r.Run(":8080")
}

