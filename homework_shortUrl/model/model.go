// 存放对象数据和全局变量
package model

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var DB *sql.DB
var ERR error

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Passwd    string `json:"passwd"`
	Long_url  string `json:"long_url"`
	Short_url string `json:"short_url"`
}
