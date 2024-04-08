package main

import (
	"gocode/homework_shortUrl/database"
	"gocode/homework_shortUrl/login"
	longtoshort "gocode/homework_shortUrl/longToShort"
	"gocode/homework_shortUrl/redirect"
	"gocode/homework_shortUrl/register"
	shorturlop "gocode/homework_shortUrl/shortUrl_op"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	// 设置日志级别为info
	log.SetLevel(log.InfoLevel)

	// 设置日志输出格式为json
	log.SetFormatter(&log.JSONFormatter{})

	//连接mysql和redis
	database.Database()

	//启动gin
	router := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	router.POST("/register", register.Register)                    //注册
	router.GET("/login", login.Login)                              //登录
	router.POST("/shortUrl", longtoshort.Write)                    //写入长链并生成短链
	router.GET("/redirect/:short_url", redirect.Redirect)          //根据短链接重定向到长链接
	router.DELETE("shortUrl/delete/:short_url", shorturlop.Delete) //删除指定用户短链接
	router.POST("shortUrl/change", shorturlop.Change)              //重新自定义短链接
	router.GET("shortUrl/show", shorturlop.Show)                   //查询指定用户短链接
	router.GET("shortUrl/rate", shorturlop.Rate)                   //排列所有短链接

	router.Run(":8080")

}
