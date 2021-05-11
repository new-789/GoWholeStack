package main

import (
	"fmt"
	"github.com/GoWholeStack/oldboyGo/book/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var r *gin.Engine

func main() {
	// 初始化数据库
	err := db.InitDB()
	if err != nil {
		fmt.Println("初始化数据库失败~~~~", err)
		return
	}
	r = gin.Default()
	// 加载页面
	r.LoadHTMLGlob("./templates/*")
	// 查询所有图书
	rGroup := r.Group("/book")
	{
		rGroup.GET("/list", bookListHandler)
		rGroup.GET("/delete", deleteBookHandler)
		rGroup.GET("/new", bookInsertHandler)
		rGroup.POST("/new", insertBookHandler)
	}
	_ = r.Run(":8000")
}

func bookListHandler(c *gin.Context) {
	bookList, err := db.QueryAllBook()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": err,
		})
		return
	}
	// 返回数据
	c.HTML(http.StatusOK, "book_list.html", gin.H{
		"code": 0,
		"data": bookList,
	})
}

func bookInsertHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "new_book.html", gin.H{
		"code": 1,
	})
}

func insertBookHandler(c *gin.Context) {
	title := c.PostForm("title")
	price := c.PostForm("price")
	pri, err := strconv.Atoi(price)
	if err != nil {
		fmt.Println("string to int error", err)
		return
	}
	err = db.InsertBook(title, int64(pri))
	if err != nil {
		fmt.Println("insert book failed, err", err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, "http://localhost:8000/book/list")
}

func deleteBookHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Println("string to int failed, err", err)
		return
	}
	bookId := int64(id)
	err = db.DeleteBook(bookId)
	if err != nil {
		fmt.Println("delete book failed, err:", err)
		return
	}
	// 重定向到指定路由
	c.Request.URL.Path = "/book/list"
	r.HandleContext(c)
}