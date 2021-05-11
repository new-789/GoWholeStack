package main

import (
	"fmt"
	"os"
	"time"
)

/*
 学生管理系统函数版
	写一个系统能够查看、新增、删除学生
*/
type student struct {
	id   int64
	name string
}

var (
	allStudent map[int64]*student // 变量声明
)

// newStudent 是 student 结构体类型的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showALLStudent() {
	// 遍历 map 打印所有学生
	fmt.Println("===========================\n")
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", k, v.name)
	}
	fmt.Println("===========================\n")
	time.Sleep(time.Millisecond * 3000)
}

func addStudent() {
	// 向 allStudent 中添加新的学生
	// 1. 创建一个新学生
	var (
		id   int64
		name string
	)
	// 1.1：获取用户输入
	fmt.Print("请输入需要添加的学生学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生的姓名：")
	fmt.Scanln(&name)
	// 1.2：构造学生（调用构造函数）
	newStu := newStudent(id, name)
	// 2. 追加到 allStudent map 中
	allStudent[id] = newStu
}

func deleteStudent() {
	// 1、请用户输入需要删除的学生学号
	fmt.Print("请输入要删除的学生学号：")
	var stuNum int64
	fmt.Scanln(&stuNum)
	// 从 map 中删除数据
	delete(allStudent, stuNum)
}

func main() {
	allStudent = make(map[int64]*student, 50) // 初始化开辟内存空间
	// 1. 打印菜单
	fmt.Println("欢迎光临学生管理系统~!")
	for i := 0; i <= 3; {
		fmt.Println(`
	1、查看所有学生
	2、新增学生
	3、删除学生
	4、退出程序
		`)
		fmt.Print("请输入相应的序号，选择对应的操作：")
		// 2. 等待用户选择功能
		var choice int
		fmt.Scanln(&choice)
		// 3. 执行对应的函数
		switch choice {
		case 1:
			showALLStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("\n")
			i++
			if i < 3 {
				fmt.Println("输入无效，请重新选择!")
			} else if i == 3 {
				fmt.Println("尝试次数过多，程序以退出~！")
				os.Exit(1) // 退出程序
			}
		}
	}
}
