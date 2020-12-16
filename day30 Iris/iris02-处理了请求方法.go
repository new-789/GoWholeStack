package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main3002() {
	app := iris.New()
	/*
		Get 方法说明：
			参数一：用来用来定义请求的地址
			参数二：匿名函数，在该匿名函数中可书写前面请求地址要处理了的逻辑代码
	*/
	app.Get("/getRequest", func(ctx context.Context) {
		path := ctx.Path() // 获取请求的地址
		// 打印输出
		app.Logger().Info(path)
	})
	app.Get("/userpath", func(ctx context.Context) {
		path := ctx.Path()
		app.Logger().Info(path)
		// 写返回数据给请求端 string 类型
		ctx.WriteString("请求路径：" + path)
	})
	// 请求地址：http://localhost:8080/userinfo?username=Tom&pwd=Hello
	app.Get("/userinfo", func(ctx context.Context) {
		path := ctx.Path()
		app.Logger().Info(path)
		// ctx.URLParam 该方法用来获取前段 Get 请求 URL中传过来的数据
		username := ctx.URLParam("username")
		pwd := ctx.URLParam("pwd")
		app.Logger().Info(username, pwd)
		// 将 html 格式的数据返回给前端
		ctx.HTML("<h2>" + username + "," + pwd + "</h2>")

	})

	// 处理 post 请求，form 表单的字段获取
	app.Post("/postPath", func(ctx context.Context) {
		path := ctx.Path()
		app.Logger().Info(path)
		// ctx.PostValue 获取 POST 请求传过来的参数方法
		name := ctx.PostValue("name")
		pwd := ctx.PostValue("pwd")
		app.Logger().Info(name, pwd)
		ctx.HTML("<p> Hello," + name + "</p>")
	})

	// 处理 post 请求，Json 格式数据
	app.Post("/postJson", func(ctx context.Context) {
		/*
			处理 post 请求，JSON 格式数据
				方法说明：`ctx.ReadJSON`
					 用来读取解析 JSON 格式的数据，接收一个内存地址类型的结构体变量为参数，用来存放读取到的 json 数据内容，语法如下：
					 	ctx.ReadJSON(&变量名)
		*/
		var person Person
		if err := ctx.ReadJSON(&person); err != nil {
			panic(err.Error())
		}
		// 格式化返回数据到前端页面
		ctx.Writef("ReadJson data %v \n", person)
	})

	/*
		处理 post 请求， Xml 格式数据
			方法说明：`ctx.ReadXML`
				该方法专门用来读取 XML 格式的数据，接收一个参数，用来保存读到的数据类型，参数类型为地址内容的结构体变量语法如下：
					ctx.ReadXML(&变量名)
			请求数据
				<student>
					<stu_name>Tom</stu_name>
					<stu_age>28</stu_age>
				</student>
	*/
	app.Post("/postXml", func(ctx iris.Context) {
		var student Student
		if err := ctx.ReadXML(&student); err != nil {
			panic(err.Error())
		}
		ctx.Writef("read xml data:%v\n", student)
	})
	app.Run(iris.Addr(":8080"))
}

// 该结构体用来存放 POST 请求中获取到的json 格式的数据
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 该结构体用来存放 POST 请求中获取到的 XML 格式的数据
type Student struct {
	StuName string `xml:"stu_name"`
	StuAge  int    `xml:"stu_age"`
}
