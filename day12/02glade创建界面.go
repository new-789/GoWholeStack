package main

import (
	"fmt"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"os"
)

func onActivate(application *gtk.Application) {
	builder, err := gtk.BuilderNewFromFile("E:/CodingFiles/GolangCode/src/github.com/FullStackDevelStudy/day12/UI.glade")
	if err != nil {
		fmt.Println("打开文件错误")
	}

	winObj, er := builder.GetObject("Win1")
	if er != nil {
		fmt.Println("获取 win1 接错失败")
	}
	win := winObj.(*gtk.Window)
	win.SetTitle("hello_gtk3")   // 设置窗口标题
	win.SetSizeRequest(600, 300) // 设置窗口大小
	application.AddWindow(win)

	// 获取按钮
	btnObj, err1 := builder.GetObject("Btn1")
	if err1 != nil {
		fmt.Println("获取按钮错误")
	}
	btn := btnObj.(*gtk.Button)
	btn.Connect("clicked", BtnClick, "点击鼠标事件")
	win.ShowAll()
}

func BtnClick(ctx *glib.AsyncResult) {
	data := ctx.GetUserData()
	fmt.Println(data)
}

func main() {
	const appId = "www.pmlchina.cn"
	// 每个 GTK 程序都需要这一步
	app, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		fmt.Println("创建 app 对象错误")
	}
	// 为 activate 事件绑定函数，activate 会在程序启动时触发，也就是 app.Run() 时
	app.Connect("activate", func() {
		onActivate(app)
	})
	// 运行 gtkApplication
	app.Run(os.Args)
}
