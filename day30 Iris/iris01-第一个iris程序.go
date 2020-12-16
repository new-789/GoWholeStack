package main

import "github.com/kataras/iris/v12"

func main3001() {
	// 实例化 app 对象
	app := iris.New()
	// 使用 app 对象运行创建服务，参数一为指定的 ip 和端，参数二为：当 iris 没有发现服务时的错误信息输出，参数二可省略
	// app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	// 启动服务方式二
	app.Run(iris.Addr(":8080"))
}
