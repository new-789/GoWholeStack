package main

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/boltdb"
)

func main3007() {
	sessionId := "MySession"

	app := iris.New()
	// 1. 创建 session 对象
	sess := sessions.New(sessions.Config{
		Cookie: sessionId,
		// 设置有效期
		Expires: time.Second * 600,
	})

	// 下面通过案例来说明 session 操作方法
	// 用户登录功能
	app.Post("/login", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info("请求 path:", path)
		// 获取 post 请求携带的值
		userName := ctx.PostValue("name")
		pwd := ctx.PostValue("pwd")
		if userName == "tom" && pwd == "000000" {
			// 2. 通过获取到的 session 对象调用 Start 方法开启 session 并将 context 当做参数传入，返回一个 session
			session := sess.Start(ctx)
			// 3. 往  session 中保存数据
			session.Set("username", userName)
			session.Set("isLogin", true) // 设置登录状态
			ctx.WriteString("用户登录成功")
		} else {
			session := sess.Start(ctx)
			// 保存数据到 session
			session.Set("IsLogin", false)
			ctx.WriteString("登录失败，请重新尝试")
		}
	})

	// 用户退出功能
	app.Get("/logout", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info("退出登录:", path)
		session := sess.Start(ctx)
		// 删除  session 中保存的数据操作
		session.Delete("isLogin")
		session.Delete("username")
		ctx.WriteString("退出登录成功")
	})

	// 查询功能
	app.Get("/query", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info("查询信息 path", path)
		session := sess.Start(ctx)
		// 获取 bool 值方法
		isLogin, err := session.GetBoolean("isLogin")
		if err != nil {
			ctx.WriteString("账户还未登录，请先登录")
			return
		}
		if isLogin {
			app.Logger().Info("用户以登录")
			ctx.WriteString("用户以登录")
		} else {
			app.Logger().Info("用户未登录")
			ctx.WriteString("用户未登录")
		}
	})

	// session 和数据库绑定使用方法,
	/*
		boltdb 是iris框架中自带的一种数据库
			New 方法可以用来生成一个数据库文件，
				参数一：生成的数据库文件名
				参数二：数据库文件的操作权限
	*/
	db, err := boltdb.New("session.db", 0600)
	if err != nil {
		panic(err.Error)
	}
	// 程序中断时将数据库关闭
	iris.RegisterOnInterrupt(func() {
		defer db.Close()
	})
	// session 和 db 绑定，绑定之后，session 中的数据就能够存储到数据库文件中
	sess.UseDatabase(db)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
