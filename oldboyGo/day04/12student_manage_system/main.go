package main

import (
	"fmt"
	"os"
	"time"
)

// 学生管理系统完整版
type administrator struct {
	name     string
	password string
}

type student struct {
	id    int64
	age   int
	score int
	name  string
}

var (
	admin      map[string]*administrator
	allStudent map[int64]*student
)

// administrator 结构体类型构造函数
func newAdmin(name, password string) *administrator {
	return &administrator{
		name:     name,
		password: password,
	}
}

// student 结构体类型构造函数
func newStudent(id int64, name string, age, score int) *student {
	return &student{
		id:    id,
		age:   age,
		score: score,
		name:  name,
	}
}

// 管理员函数设计是否合理？存在的问题，学生不能登录！！
func (a *administrator) register() {
	// 获取用户输入的注册信息
	var (
		name     string
		password string
		pwd      string
	)
	for i := 0; i <= 3; {
		fmt.Print("请输入用户名：")
		fmt.Scan(&name)
		fmt.Print("请输入密码：")
		fmt.Scan(&password)
		fmt.Print("再次输入密码：")
		fmt.Scan(&pwd)
		if password != pwd {
			fmt.Println("两次输入密码不一致，请重新输入")
			i++
			if i == 3 {
				fmt.Println("尝试次数太多，程序将退出")
				time.Sleep(time.Millisecond * 3)
				os.Exit(1)
			}
		} else {
			// 注册：往 admin 的 map 中添加数据
			newAdmin := newAdmin(name, password)
			admin[name] = newAdmin
			fmt.Println()
			fmt.Print("注册成功，请登录：")
			a.login()
		}
	}
}

// 管理员登录函数
func (a *administrator) login() {
	// 获取用户输入的信息，和存储的内容进行比较
	var (
		uname string
		pwd   string
	)
	for k := 0; k <= 3; {
		k++
		fmt.Print("请输入用户名：")
		fmt.Scan(&uname)
		fmt.Print("请输入密码：")
		fmt.Scan(&pwd)
		// 从 map 中获取管理员用户名及密码
		if _, ok := admin[uname]; ok {
			username := admin[uname].name
			upwd := admin[uname].password
			if uname == username && pwd == upwd {
				a.printMenu()
			}
		} else {
			fmt.Println("您尚未注册哦，请先注册然后在尝试登录~！")
			fmt.Println("===================== 请注册 ===================")
			a.register()
		}
	}
}

// 管理员登录成功之后打印菜单函数
func (a *administrator) printMenu() {
	// 打印学生操作菜单
	fmt.Println("===================== 登录成功 ===================")
	for j := 0; j <= 3; {
		fmt.Println(`
	1、查看学生信息
	2、添加学生信息
	3、修改学生信息
	4、删除学生信息
	5、退出程序
	`)
		fmt.Print("选择对应能序列号进行相应的操作：")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			a.showAllStu()
		case 2:
			a.addStu()
		case 3:
			a.editStu()
		case 4:
			a.deleteStu()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("\n")
			j++
			if j < 3 {
				fmt.Println("输入无效，请重新选择!")
			} else if j == 3 {
				fmt.Println("尝试次数过多，程序以退出~！")
				os.Exit(1) // 退出程序
			}
		}
	}
}

// 学生操作相关方法的
func (a *administrator) showAllStu() {
	// 遍历 map 打印所有学生
	fmt.Println("===========================\n")
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s 年龄：%d 分数：%d\n", k, v.name, v.age, v.score)
	}
	fmt.Println("===========================\n")
	time.Sleep(time.Millisecond * 8000)
}

func (a *administrator) addStu() {
	// 获取用户输入信息
	var (
		id    int64
		age   int
		score int
		name  string
	)
	fmt.Print("请输入学生的 ID 号：")
	fmt.Scan(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scan(&name)
	fmt.Print("请输入学生年龄：")
	fmt.Scan(&age)
	fmt.Print("请输入学生分数：")
	fmt.Scan(&score)
	// 通过构造函数获取学生对象
	newStu := newStudent(id, name, age, score)
	// 追加到学生 map 中
	allStudent[id] = newStu
}

func (a *administrator) editStu() {
	// 修改学生之前先打印一遍所有学生内容，以便用户选择需要修改的学生信息
	// 遍历 map 打印所有学生
	fmt.Println("===========================\n")
	for k, v := range allStudent {
		fmt.Printf("学号：%d 姓名：%s 年龄：%d 分数：%d\n", k, v.name, v.age, v.score)
	}
	fmt.Println("===========================\n")
	time.Sleep(time.Millisecond * 8000)
	if a.name != "" {
		// 获取用户输入的数据
		var (
			sid              int64
			sname            string
			sage, score, cid int
		)
		fmt.Print("请输入想要修改的学生 id：")
		fmt.Scan(&sid)
		fmt.Println("待修改的学生信息内容如下：")
		fmt.Printf("学号：%d 姓名：%s 年龄：%d 分数：%d\n", sid, allStudent[sid].name, allStudent[sid].age, allStudent[sid].score)
		fmt.Println(`
	1、姓名
	2、年龄
	3、分数
	`)
		fmt.Print("请选择需要修改的内容，输入相应序号：")
		fmt.Scan(&cid)
		// 获取用户输入的修改数据
		switch cid {
		case 1:
			fmt.Print("请输入新的姓名：")
			fmt.Scan(&sname)
			if sname != " " { // 如果输入为空则不进行修改操作
				allStudent[sid].name = sname
			}
		case 2:
			fmt.Print("请输入新的年龄：")
			fmt.Scan(&sage)
			if sage != 0 {
				allStudent[sid].age = sage
			}
		case 3:
			fmt.Print("请输入新的分数是：")
			fmt.Scan(&score)
			allStudent[sid].score = score
		}
	} else {
		fmt.Println("暂时没有学生信息可供修改，即将退回到登录状态，请选择添加学生功能添加学生信息~~~")
		time.Sleep(time.Millisecond * 5000)
		a.printMenu()
	}
}

func (a *administrator) deleteStu() {
	var sid int64
	fmt.Print("请输入要删除的学生学号：")
	fmt.Scan(&sid)
	delete(allStudent, sid)
}

func main() {
	// 1、程序运行时初始化 map
	a := &administrator{}
	admin = make(map[string]*administrator, 5)
	allStudent = make(map[int64]*student, 20)
	// 2、打印开始欢迎信息及菜单
	fmt.Println("欢迎来到 XX 学生信息管理系统")
	for i := 0; i <= 3; {
		fmt.Println(`
	1、管理员注册
	2、管理员登录
	3、查看介绍
	4、退出程序
		`)
		fmt.Print("请输入相应序号，选择对应功能进行操作：")
		var ipt int
		fmt.Scan(&ipt)
		switch ipt {
		case 1:
			a.register()
		case 2:
			a.login()
		case 3:
			fmt.Println("===================================================================")
			fmt.Println(`
欢迎来到 XX 学生信息管理系统，在系统中您可以查看学生的相关信息以及
对学生信息进行相应操作等。此系统的开发耗费了我们团队大量精力，
【点击链接】赞助我们，若存在疑问请拨打教务处联系电话：xxxxxxxxxxx
			`)
			fmt.Println("===================================================================\n")
			time.Sleep(time.Millisecond * 20000)
		case 4:
			os.Exit(1)
		}
	}
}
