package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// gin 中间件

// MiddleWare 定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了...........")
		// 设置变量到 context 的 key 中，可以通过 Get() 获取
		c.Set("request", "中间件中Set设置的值")
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time", t2)
	}
}

func main() {
	r := gin.Default()
	// 注册全局中间件
	r.Use(MiddleWare())
	// 大括号是为了代码规范
	{
		r.GET("/middleware", func(c *gin.Context) {
			// 取中间件设置中的值
			resp, _ := c.Get("request")
			fmt.Println(resp)
			// 页面返回
			c.JSON(http.StatusOK, gin.H{"request": resp})
		})

		// 定义局部中间件: 路由后面紧跟着的第二个参数 MiddleWare() 即为定义的局部中间件
		r.GET("/middleware2", MiddleWare(), func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Println("局部中间件：request", req)
			c.JSON(http.StatusOK, gin.H{"局部中间件结果:request": req})
		})
	}

	r.Run(":8080")
}
