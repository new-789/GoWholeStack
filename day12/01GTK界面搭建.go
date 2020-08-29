package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"os"
)

func main0101() {
	gtk.Init(&os.Args) // 初始化

	// 通过 GTK 创建界面
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		fmt.Println("创建窗口失败")
	}

	// 设置窗体大小
	win.SetSizeRequest(600, 320)
	// 设置窗体标题
	win.SetTitle("Hello GTK3")

	// 创建一个 label 标签
	l, err1 := gtk.LabelNew("hello___gtk3~!")
	if err1 != nil {
		fmt.Println("创建 label 错误")
	}

	// 创建一个按钮
	// 以水平布局创建一个容器，第二个参数表示其中控件的像素点
	layout, err2 := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 2)
	if err2 != nil {
		fmt.Println("创建容器失败")
	}
	// 创建按钮
	button, err3 := gtk.ButtonNewWithLabel("buttOn")
	if err3 != nil {
		fmt.Println("创建按钮失败")
	}
	// 为按钮设置一个名称,由于创建按钮时已经设置，此步可省略
	button.SetLabel("sss")
	// 将按钮添加到容器中
	layout.Add(button)
	// 将 label 添加到容器中
	layout.Add(l)

	// 创建按钮事件
	button.Connect("clicked", func() { fmt.Println("按钮被按下了》》》》》》》》》》") })
	win.Add(layout) // 将容器添加到 window 中

	win.ShowAll()
	gtk.Main()
}
