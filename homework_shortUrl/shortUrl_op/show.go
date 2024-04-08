package shorturlop

import (
	"database/sql"
	"gocode/homework_shortUrl/model"

	"github.com/gin-gonic/gin"
)

func Show(c *gin.Context) {
	var user model.User
	user.Name = c.PostForm("name")
	row := model.DB.QueryRow("select short_url from users where name=?",
		user.Name)
	if row != nil {
		err := row.Scan(&user.Short_url)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(401, gin.H{"error": "用户名不存在"})
			} else {
				c.JSON(500, gin.H{"error": "服务器内部错误"})
			}
			return
		}
	}
	c.JSON(200, gin.H{
		"info": "查询成功",
		"用户名":  user.Name,
		"短链":   user.Short_url,
	})
}
