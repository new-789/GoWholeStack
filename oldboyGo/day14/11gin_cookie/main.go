package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// gin cookie Demo

func main() {
	r := gin.Default()
	// 服务端要给客户端 cookie
	r.GET("/cookie", func(c *gin.Context) {
		// 获取客户单是否携带了 cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NoSet"
			// 给客户端设置 cookie
			/*
			参数一和参数二：name, value string;  设置的 key 和 value
			参数三：maxAge int; 超时时间，单位为秒
			参数四和参数五：path, domain string； path 为 cookie 所在目录，domain 为 域名
			参数六和参数七：secure, httpOnly bool; secure 是否只能通过 https 访问，httpOnly 是否允许别人通过 js 获取自己的 cookie
			 */
			c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
		}
		fmt.Printf("cookie value is :%v\n", cookie)
	})
	r.Run(":8000")
}
