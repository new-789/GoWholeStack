package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// gin 框架的同步和异步

var (

)
func main() {
	r := gin.Default()
	// 1. 异步
	r.GET("/long_async", func(c *gin.Context) {
		// 需要搞一个副本
		copyContext := c.Copy()
		go func() {
			time.Sleep(time.Second * 3)
			log.Println("异步执行"+copyContext.Request.URL.Path)
		}()
	})
	
	// 2. 同步
	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(time.Second * 3)
		log.Println("同步执行" + c.Request.URL.Path)
	})
	
	r.Run(":8080")
}