package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// cookie 练习

func MyMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		name, err := c.Cookie("name")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "please login!",
			})
		}
		if name == "root" {
			c.Next()
		}
		// 返回错误
		c.JSON(http.StatusNotModified, gin.H{
			"err": "error",
		})
		// 如果验证不通过，不在调用后续的函数处理
		c.Abort()
		return
	}
}

func main() {
	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		// 设置 Cookie
		c.SetCookie("name", "root", 30, "/", "localhost", false, true)
		// 返回信息
		c.JSON(http.StatusOK, gin.H{
			"msg": "login success!",
		})
	})
	r.GET("/home",MyMiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "Welcome~！",
		})
	})

	r.Run(":8080")
}
