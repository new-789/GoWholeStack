package main

type Student1 struct {
	name   string
	age    int
	sex    string
	cscore int
	mscore int
	escore int
}

// 扩展，为对象赋值操作
func (s *Student1) InitInfo(name, sex string, age, cscore, mscore, escore int) {
	// 使用该方法为对象(结构体)中的字段进行赋值操作
	s.name = name
	s.age = age
	s.sex = sex
	s.cscore = cscore
	s.mscore = mscore
	s.escore = escore
}

func main() {
	//var s Student1 = Student1{"杨广", 15, "男", 80, 60, 90}
	var s Student1
	// 初始化对象信息
	s.InitInfo("薛宝钗", "女", 16, 90, 100, 88)
	s.SayHello()
	s.PrintScore()
}
