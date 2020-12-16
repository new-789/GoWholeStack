package main

import "github.com/kataras/iris/v12"

func main3004() {
	app := iris.New()
	// 用户模块
	// xxx/user/register  注册
	// xxx/user/login  登录
	// xxx/user/info  获取用户信息

	/*
		路由组请求
	*/
	userParty := app.Party("/users", func(ctx iris.Context) {
		// 处理下一级请求
		ctx.Next()
	})
	/*
		 路由组下面的下一级请求
			xxx/users/register
	*/
	userParty.Get("/register", func(ctx iris.Context) {
		app.Logger().Info("用户注册功能")
		ctx.HTML("<h1>用户注册功能</h1>")
	})

	/*
		路由组下面的下一级请求
		xxx/users/login
	*/
	userParty.Get("/login", func(ctx iris.Context) {
		app.Logger().Info("用户登录功能")
		ctx.HTML("<h1>用户登录功能</h1>")
	})

	userRouter := app.Party("/admin", userMiddleware)

	// Done 方法，表示只要在下一级路由的逻辑处理函数中执行了 ctx.Next() 方法则会自动执行 Done 方法，表示该请求所有操作都完成了
	userRouter.Done(func(ctx iris.Context) {
		ctx.Application().Logger().Info("reponse sent to" + ctx.Path())
	})
	userRouter.Get("/info", func(ctx iris.Context) {
		ctx.HTML("<h1>用户信息</h1>")
		ctx.Next() // 执行到该语句表示该请求所有逻辑已执行完毕接下来会自动执行该路由组下的 Done 方法，属于手动显示调用
	})
	userRouter.Get("/query", func(ctx iris.Context) {
		ctx.HTML("<h1>查询信息</h1>")
	})

	app.Run(iris.Addr(":8080"))
}

func userMiddleware(ctx iris.Context) {
	ctx.Next()
}
