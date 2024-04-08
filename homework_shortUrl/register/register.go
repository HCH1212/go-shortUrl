package register

import (
	"fmt"
	"gocode/homework_shortUrl/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User

	//注册
	user.Name = c.PostForm("name")
	user.Passwd = c.PostForm("passwd")
	if user.Name == "" || user.Passwd == "" {
		c.JSON(200, gin.H{"error": "注册信息不完整，无法继续进行"})
		return
	}
	if model.DB != nil {
		result, err := model.DB.Exec("insert into users (name, passwd) value (?,?)",
			user.Name, user.Passwd)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"error": "注册错误"})
			return
		}
		fmt.Println(result.LastInsertId())

		model.RDB.Set(c, "name", user.Name, 0)
		model.RDB.Set(c, "passwd", user.Passwd, 0)

		c.JSON(200, gin.H{
			"status": 10000,
			"info":   "注册成功",
		})
	} else {
		fmt.Println("空指针错误..")
	}
}
