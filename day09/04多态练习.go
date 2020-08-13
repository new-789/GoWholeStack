package main

import "fmt"

// 用多态来实现将移动硬盘或U盘或者 MP3 插到电脑上进行读写数据操作(分析，接口，方法)；
type USBer interface { // 1. 接口的实现
	Read()
	Write()
}

// 2. 创建对象
type UsbDev struct {
	id       int
	name     string
	rspeed   int
	wspeed   int
	capacity string
}

type MobileDisk struct { // 移动硬盘
	UsbDev
}

type Mp3 struct { // MP3
	UsbDev
}

type UDisk struct { // U 盘
	UsbDev
}

// 3. 方法的实现
func (m *MobileDisk) Read() {
	fmt.Printf("有一个 %s 它的容量是 %s, 此时正在读取数据，它的速度为 %d kb/s\n", m.name, m.capacity, m.rspeed)
}

func (m *MobileDisk) Write() {
	fmt.Printf("有一个 %s 它的容量是 %s, 此时正在写入数据，它速度为 %d kb/s\n", m.name, m.capacity, m.wspeed)
}

func (mp *Mp3) Read() {
	fmt.Printf("有一个 %s 它的容量是 %s, 此时正在读取数据，它的速度为 %d kb/s\n", mp.name, mp.capacity, mp.rspeed)
}

func (mp *Mp3) Write() {
	fmt.Printf("有一个 %s 它的容量是 %s, 此时正在写入数据，它速度为 %d kb/s\n", mp.name, mp.capacity, mp.wspeed)
}

func (u *UDisk) Read() {
	fmt.Printf("有一个 %s 它的容量是 %s, 此时正在读取数据，它的速度为 %d kb/s\n", u.name, u.capacity, u.rspeed)
}

func (u *UDisk) Write() {
	fmt.Printf("有一个 %s 它的容量是 %s, 此时正在写入数据，它速度为 %d kb/s\n", u.name, u.capacity, u.wspeed)
}

// 4. 多态的实现， 将接口作为函数参数
func USBDev(usb USBer) {
	usb.Read()
	usb.Write()
}

func main0401() {
	var usb USBer // 定义接口类型变量
	// 初始化移动硬盘结构体，并将其地址赋值给接口类型变量
	usb = &MobileDisk{UsbDev{101, "移动硬盘", 200, 90, "500GB"}}
	USBDev(usb) // 通过调用多态函数，调用它们的方法
	fmt.Println("=============================================================")
	usb = &Mp3{UsbDev{102, "MP3", 45, 20, "4GB"}}
	USBDev(usb)
	fmt.Println("=============================================================")
	usb = &UDisk{UsbDev{103, "U盘", 50, 40, "16GB"}}
	USBDev(usb)
}
