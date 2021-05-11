package main

import "fmt"

// 学生管理系统
type student struct {
	id   int64
	name string
}

// 造一个学生的管理者
type studentMgr struct {
	allStudent map[int64]*student
}

// 查看学生
func (s studentMgr) showStudent() {
	// 从 s.allStudent 这个 map 中逐个打印出来
	for _, stu := range s.allStudent { // stu 是具体的每一个学生
		fmt.Printf("学号：%d 姓名：%s\n", stu.id, stu.name)
	}
}

// 增加学生
func (s studentMgr) addStudent() {
	// 1、根据用户输入的内容创建一个新的学生
	var (
		stuID   int64
		stuName string
	)
	// 获取用户输入
	fmt.Print("请输入学号：")
	fmt.Scan(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scan(&stuName)
	// 根据用户输入创建结构体对象
	newStu := &student{
		id:   stuID,
		name: stuName,
	}
	// 2、将新的学生放到 s.allStudent 中
	s.allStudent[newStu.id] = newStu
	fmt.Println("添加成功!")
}

// 修改学生
func (s studentMgr) editStudent() {
	// 获取用户输入的学号
	var stuID int64
	fmt.Print("请输入学号：")
	fmt.Scan(&stuID)
	// 展示该学号对应的学生信息，如果没有提示没有
	stuObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("待修改的学生信息如下：\n === 学号：%d 姓名：%s ===\n", stuObj.id, stuObj.name)
	// 请输入修改后的学生名
	fmt.Print("请输入学生的新姓名:")
	var newName string
	fmt.Scan(&newName)
	// 更新学生信息
	stuObj.name = newName
	fmt.Println("修改成功!")
}

// 删除学生
func (s studentMgr) deleteStudent() {
	// 请用户输入要删除的学生的 id
	var stuID int64
	fmt.Print("请输入要删除的学生 id：")
	fmt.Scan(&stuID)
	// 去 map 中查找是否存在待删除的学生信息,没有则提示，有则直接删除
	_, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	// 存在直接删除
	delete(s.allStudent, stuID)
	fmt.Println("删除成功~!")
}
