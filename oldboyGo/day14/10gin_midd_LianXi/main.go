package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// gin 中间件练习

// MyTime 定义中间件，注意此种写法定义中间件没有返回 gin.HandlerFunc
func MyTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	// 统计时间
	since := time.Since(start)
	fmt.Println("程序用时时间：", since)
}

func main() {
	r := gin.New()
	// 注册全局中间件，对于没有返回gin.HandlerFunc 函数的中间件注册时无需对中间件函数加括号
	r.Use(MyTime)
	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		shoppingGroup.GET("/home", shopHomeHandler)
	}
	r.Run(":8080")
}

func shopIndexHandler(c *gin.Context) {
	time.Sleep(time.Second * 5)
}

func shopHomeHandler(c *gin.Context) {
	time.Sleep(time.Second * 3)
}