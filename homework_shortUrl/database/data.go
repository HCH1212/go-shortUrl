// 连接mysql数据库和redis用于缓存
package database

import (
	"database/sql"
	"fmt"
	"gocode/homework_shortUrl/model"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

func Database() {
	sqlStr := "root:041212@tcp(127.0.0.1:3306)/homework_shortUrl?parseTime=true"
	model.DB, model.ERR = sql.Open("mysql", sqlStr)
	if model.ERR != nil {
		fmt.Println("连接数据库失败..", model.ERR.Error())
		return
	}

	model.RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       6,
	})
}
