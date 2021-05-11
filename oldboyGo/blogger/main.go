package main

import (
	"github.com/GoWholeStack/oldboyGo/blogger/controller"
	"github.com/GoWholeStack/oldboyGo/blogger/dao/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 启动程序加载数据库
	dns := "root:root@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	// 加载模板
	router.LoadHTMLGlob("./views/*")
	// 加载静态文件
	router.Static("/static/", "./static")
	// 访问主页
	router.GET("/", controller.IndexHandle)
	router.GET("/category/", controller.CategoryList)
	router.Run(":8000")
}
