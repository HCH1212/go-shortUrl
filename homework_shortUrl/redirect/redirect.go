package redirect

import (
	"fmt"
	"gocode/homework_shortUrl/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 由短链接重定向
func Redirect(c *gin.Context) {
	var user model.User
	user.Short_url = c.Param("short_url")                                //从GET请求中获取短链接,注意是Param请求
	user.Long_url, model.ERR = model.RDB.Get(c, user.Short_url).Result() //从Redis缓存中获取长链接
	if model.ERR != nil {
		fmt.Println(model.ERR)
		c.JSON(http.StatusNotFound, gin.H{"error": "Short code not found"}) //如果找不到长链接，返回错误信息
		return
	}
	c.Redirect(http.StatusMovedPermanently, user.Long_url) //根据长链接进行重定向
}
