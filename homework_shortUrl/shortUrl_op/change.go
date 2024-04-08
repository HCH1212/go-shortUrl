package shorturlop

import (
	"database/sql"
	"fmt"
	"gocode/homework_shortUrl/model"

	"github.com/gin-gonic/gin"
)

// 相当于自定义短链接
func Change(c *gin.Context) {
	var user model.User
	user.Name = c.PostForm("name")
	user.Passwd = c.PostForm("passwd")
	new_short_url := c.PostForm("new_short_url")
	row := model.DB.QueryRow("select long_url, short_url from users where name=? and passwd=?",
		user.Name, user.Passwd)
	if row != nil {
		err := row.Scan(&user.Long_url, &user.Short_url)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(401, gin.H{"error": "用户名或密码错误"})
			} else {
				fmt.Println(err)
				c.JSON(500, gin.H{"error": "服务器内部错误"})
			}
			return
		}
		if user.Long_url == "" || user.Short_url == "" {
			c.JSON(401, gin.H{"error": "长链接或短连接不存在"})
			return
		}
	}
	//更新短链接
	//mysql
	if model.DB != nil {
		result, err := model.DB.Exec("update users set short_url=? where name=?",
			new_short_url, user.Name)
		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{"error": "mysql更新自定义短链接错误"})
			return
		}
		fmt.Println(result.LastInsertId())
	} else {
		fmt.Println("空指针错误..")
	}
	//redis
	// 更新key
	err := model.RDB.Rename(c, user.Short_url, new_short_url).Err()
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"info":  "自定义短链接成功",
		"新短链接为": new_short_url,
	})
}
