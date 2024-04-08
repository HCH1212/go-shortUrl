// 写入长链并返回生成的短链
package longtoshort

import (
	"database/sql"
	"fmt"
	"gocode/homework_shortUrl/model"

	"github.com/gin-gonic/gin"
)

func Write(c *gin.Context) {
	var user model.User
	user.Name = c.PostForm("name")
	user.Passwd = c.PostForm("passwd")
	user.Long_url = c.PostForm("long_url")
	//登录
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

	//更新long_url到数据库
	result, err := model.DB.Exec("update users set long_url=? where id=?",
		user.Long_url, user.Id)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": "长连接注入错误"})
		return
	}
	fmt.Println(result.LastInsertId())

	user.Short_url = To(user.Id, user.Long_url)

	//更新短链接
	result, err = model.DB.Exec("update users set short_url=? where id=?",
		user.Short_url, user.Id)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": "短连接注入错误"})
		return
	}
	fmt.Println(result.LastInsertId())

	model.RDB.Set(c, user.Short_url, user.Long_url, 0) //以短链接为key,长链接为val，存入缓存

	c.JSON(200, gin.H{
		"status":    10000,
		"info":      "长短链接插入成功",
		"long_url":  user.Long_url,
		"short_url": user.Short_url,
	})

}
