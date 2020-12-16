package main

import (
	_ "BeegoOne/models"
	_ "BeegoOne/routers"

	"github.com/astaxie/beego"
)

func main() {
	// 设置静态资源路径，很少用
	// beego.SetStaticPath("/tests/", "/tests")

	// 在运行前映射视图函数
	beego.AddFuncMap("UpPage", UpPage)
	beego.AddFuncMap("DownPage", DownPage)
	beego.Run()
}

// 与视图函数上一页对应的函数名
func UpPage(in int) (pageIndex int) {
	pageIndex = in - 1
	if pageIndex < 1 {
		pageIndex = 1
	}
	return
}

// 与视图函数下一页对应的函数名
func DownPage(inIndex int) (pageIndex int) {
	pageIndex = inIndex + 1
	return
}
