package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1. 创建路由
	r := gin.Default()
	/*
	r.POST("/form", func(c *gin.Context) {
		// 接受 form 表单参数, 表单参数设置默认值
		typeStr := c.DefaultPostForm("type", "alter")
		// 接收其他 form 表单参数
		username := c.PostForm("username")
		password := c.PostForm("password")
		// 接收多选框参数
		hobbys := c.PostFormArray("hobby")
		 // 返回数据
		 c.String(http.StatusOK,
		 	fmt.Sprintf("type is %s username is %s password is %s hoppys is %v", typeStr, username, password, hobbys))
	})
	 */

	// 表单取文件，获取单个文件
	/*
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			fmt.Printf("获取文件失败，错误：%v\n", err)
			return
		}
		log.Println(file.Filename)
		// 将文件保存到项目的根目录
		err = c.SaveUploadedFile(file, file.Filename)
		if err != nil {
			fmt.Printf("save file failed, err:%v\n", err)
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("%s upload success", file.Filename))
	})
	 */
	
	// 表单取文件，获取多个文件
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("get err: %v\n", err))
		}
		// 获取所有图片
		files := form.File["file"]
		for _, file := range files {
			err := c.SaveUploadedFile(file, file.Filename)
			if err != nil {
				c.JSON(http.StatusBadRequest, fmt.Sprintf("get file form files err:%v\n", err))
				return
			}
		}
		c.String(http.StatusOK, "upload ok %d files", len(files))
	})
	r.Run(":8000")
}
