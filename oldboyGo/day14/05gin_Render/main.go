package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

// gin 渲染 Demo

func main() {
	// 创建路由
	r := gin.Default()
	// 1. json 响应
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message":"JSON", "status":"200"})
	})
	// 2. 结构体响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name string
			Message string
			Number int
		}
		msg.Name = "root"
		msg.Message ="STRUCT"
		msg.Number = 88
		c.JSON(http.StatusOK, msg)
	})
	// 3. xml 响应
	r.GET("someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message":"XML"})
	})
	// 4. YAML 响应
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message":"YAML"})
	})
	// 5. protobuf 响应, 谷歌开发的高效读取的工具
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "label"
		data := &protoexample.Test{
			Label: &label,
			Reps: reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})

	r.Run(":8080")
}
