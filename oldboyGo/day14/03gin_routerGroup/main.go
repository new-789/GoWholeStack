package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由组

func main() {
	r := gin.Default()
	// 路由组1，处理 GET 请求
	v1 := r.Group("/v1")
	// {} 是书写规范，不写也是可以的
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit",submit)
	}
	r.Run(":8080")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(http.StatusOK, "Hello " + name + "\n")
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lili")
	c.String(http.StatusOK, "Hello " + name + "\n")
}