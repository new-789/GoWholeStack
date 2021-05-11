package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
// gin 框架的 helloWorld 程序

func main() {
	// 1. 创建路由,Default 方法默认使用了两个中间件 Logger()、Recover()
	r := gin.Default()
	/*
	r := gin.New()  // 创建不带中间件的路由
	// 2. 绑定路由规则,执行的函数
	// gin.Context，封装了 request 和 Response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.POST("/xxxPost", geting)
	r.PUT("/xxxPut")
	 */

	// 3. API 参数
	/* 用来处理请求地址 http://localhost:8000/user/wukong/test
		:name 用来获取 /user/ 后面的 wukong
		*action 用来获取 /wukong 后面的所有内容
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name") // 获取 API 中的 name 参数
		action := c.Param("action")
		c.String(http.StatusOK, name + " is " + action)
	})
	 */

	// 4. url 参数：用来获取 url 地址 ? 号后面所带的参数
	// url 为：http://localhost:8000/welcome?name=wukong 即获取参数 name 所带的参数 wukong
	r.GET("/welcome", func(c *gin.Context) {
		name := c.DefaultQuery("name", "jack")
		c.String(http.StatusOK, fmt.Sprintf("Hello %s", name))
	})

	// 3. 设置监听的端口, 默认在 8080 端口
	err := r.Run(":8000")
	if err != nil {
		fmt.Printf("gin run failed, err:%v\n", err)
		return
	}
}

//func geting(c *gin.Context) {
//
//}
