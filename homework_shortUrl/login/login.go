package login

import (
	"database/sql"
	"fmt"
	"gocode/homework_shortUrl/model"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.User
	//用户名和密码登录
	user.Name = c.PostForm("name")
	user.Passwd = c.PostForm("passwd")
	row := model.DB.QueryRow("select id from users where name=? and passwd=?",
		user.Name, user.Passwd)
	if row != nil {
		err := row.Scan(&user.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(401, gin.H{"error": "用户名或密码错误"})
			} else {
				c.JSON(500, gin.H{"error": "服务器内部错误"})
			}
			return
		}
	}

	val, err := model.RDB.Get(c, "name").Result()
	if err != nil {
		fmt.Println("纯登录model.RDB.Get err=", err)
	}
	fmt.Println("name:", val)
	val, err = model.RDB.Get(c, "passwd").Result()
	if err != nil {
		fmt.Println("纯登录model.RDB.Get err=", err)
	}
	fmt.Println("passwd:", val)

	c.JSON(200, gin.H{
		"status": 10000,
		"info":   "登录成功",
	})
}
