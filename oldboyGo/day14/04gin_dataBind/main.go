package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin 数据解析和绑定

// Login ...
type Login struct {
	// binding:"required" 修饰字段，若接收值为空则报错，是必须字段
	User string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	// JSON 绑定数据与解析

	r.POST("/loginJSON", func(c *gin.Context) {
		var jsonData Login
		// 将 request 的 body 中的数据，自动按照 json 格式解析到结构体
		if err := c.ShouldBind(&jsonData); err != nil {
			// 返回错误信息
			// gin.H 封装了生产 json 数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 判断用户名密码是否正确
		if jsonData.User != "root" || jsonData.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	// form 表单数据绑定与解析
	/*
	r.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// c.ShouldBind(&form) 等同于 c.Bind(&form)
		if err := c.ShouldBind(&form);err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status":"200"})
	})
	 */

	// URL 数据绑定与解析
	/*
	r.GET("/loginURL/:user/:password", func(c *gin.Context) {
			var loginUrl Login
			if err := c.ShouldBindUri(&loginUrl);err == nil {
				if loginUrl.User != "root" || loginUrl.Password != "admin" {
					c.JSON(http.StatusBadRequest, gin.H{"status":"304"})
					return
				}
				c.JSON(http.StatusOK, gin.H{"name":loginUrl.User,"password":loginUrl.Password})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		}
	})
	 */
	r.Run(":8080")
}
