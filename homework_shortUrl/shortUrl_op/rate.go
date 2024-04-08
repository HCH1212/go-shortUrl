package shorturlop

import (
	"fmt"
	"gocode/homework_shortUrl/model"

	"github.com/gin-gonic/gin"
)

func Rate(c *gin.Context) {
	var user model.User
	// 查询同类所有数据
	rows, err := model.DB.Query("select short_url from users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// 遍历结果集
	var data []string
	for rows.Next() {
		err := rows.Scan(&user.Short_url)
		if err != nil {
			panic(err.Error())
		}
		data = append(data, user.Short_url)
	}

	// 输出结果
	for _, d := range data {
		fmt.Printf("%s\n", d)
	}

	c.JSON(200, gin.H{
		"info": "按id排列所有短链接",
		"data": data,
	})
}
