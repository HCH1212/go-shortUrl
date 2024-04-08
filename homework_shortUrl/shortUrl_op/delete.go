package shorturlop

import (
	"database/sql"
	"fmt"
	"gocode/homework_shortUrl/model"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	var user model.User
	user.Short_url = c.Param("short_url")
	//查询短链接是否存在
	row := model.DB.QueryRow("select id from users where short_url=?",
		user.Short_url)
	if row != nil {
		err := row.Scan(&user.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(401, gin.H{"error": "短连接不存在"})
			} else {
				c.JSON(500, gin.H{"error": "服务器内部错误"})
			}
			return
		}
	}
	//删除短链接
	//更新mysql
	result, err := model.DB.Exec("update users set short_url=? where id=?",
		"", user.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result.LastInsertId())
	//redis
	err = model.RDB.Del(c, "short_url").Err()
	if err != nil {
		fmt.Println("redis delete short_url err=", err)
	}

	c.JSON(200, gin.H{"data": "短链删除成功"})
}
