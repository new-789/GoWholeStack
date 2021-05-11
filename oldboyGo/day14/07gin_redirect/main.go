package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin 重定向 Demo

func main() {
	r := gin.Default()
	r.GET("/redirect", func(c *gin.Context) {
		// 支持内部和外部重定向
		c.Redirect(http.StatusMovedPermanently, "https://sogou.com/")
	})
	r.Run(":8080")
}