package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin HTML 模板渲染 demo

func main() {
	r := gin.Default()
	// 1. 加载模板文件
	r.LoadHTMLGlob("templates/*") // 匹配 templates 目录下的所有文件
	//r.LoadHTMLFiles("templates/index.html")
	r.GET("/index", func(c *gin.Context) {
		// 根据文件名渲染，第一个参数为状态码，第二个参数为渲染的文件名，第三个参数为 html 中需要渲染的内容
		c.HTML(http.StatusOK, "index.html", gin.H{"title":"我的标题"})
	})
	r.Run(":8080")
}