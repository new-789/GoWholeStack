package main

import "github.com/kataras/iris/v12"

func main3003() {
	app := iris.New()
	/*
		Handle 方法说明
			该方法和前面所说的 GET、POST、PUT 等方法一样同样是用来处理　url 请求用的，它的源码实现如下：
			```go
			func (api *APIBuilder) Handle(method string, relativePath string, handlers ...context.Handler) *Route {}
			```
			- 参数说明：
				- 参数1：表示请求方法，string 类型，请求方法即前面所说的 HTTP 请求你的八种方式都可
				- 参数2：请求的 url 地址 string 类型
				- 参数3：请求成功后处理相关逻辑的匿名函数
			- 返回值说明：该方法返回一个路由对象
	*/
	app.Handle("GET", "/userinfo", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info(path)
	})

	// Handle 处理 post 请求
	app.Handle("POST", "/postcommit", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info("this post request url is:" + path)
		ctx.WriteString(path)
	})
	// Get 正则表达式路由,请求地址：http://localhost:8080/hello/2020-12-01/beijing, {data} 和 {city} 用来接收 url 后面的 data 和 city
	app.Get("/hello/{date}/{city}", func(ctx iris.Context) {
		path := ctx.Path()
		app.Logger().Info(path)

		// 获取正则路由变量的值方法
		data := ctx.Params().Get("date")
		city := ctx.Params().Get("city")
		ctx.WriteString(path + " , " + data + " , " + city)
	})

	// 正则表达式路由 {name},请求地址 http://localhost:8080/hello/1000
	app.Get("/hello/{name}", func(ctx iris.Context) {
		path := ctx.Path()
		name := ctx.Params().Get("name")
		ctx.HTML("<p>" + path + "<span>" + name + "</span></p>")
	})

	// 自定义正则表达式路由变量请求，并使用  uint64 类型来限制变量的类型,，请求地址:http://localhost:8080/api/users/10000
	app.Get("/api/users/{userid: uint64}", func(ctx iris.Context) {
		userid, err := ctx.Params().GetUint64("userid")
		if err != nil {
			app.Logger().Info(err)
			// 设置请求状态码方法
			ctx.JSON(map[string]interface{}{
				"requestcode": 201,
				"message":     "bad request",
			})
			return
		}
		ctx.JSON(map[string]interface{}{
			"requestcode": 200,
			"User_id":     userid,
		})
	})

	// 自定义正则表达式路由请求，限制正则表达变量类型为 bool，请求地址:http://localhost:8080/api/users/false
	app.Get("/api/users/{isLogin:bool}", func(ctx iris.Context) {
		isLogin, err := ctx.Params().GetBool("isLogin")
		if err != nil {
			// iris.StatusNonAuthoritativeInfo
			ctx.StatusCode(iris.StatusNonAuthoritativeInfo)
			return
		}
		if isLogin {
			ctx.WriteString("已登录")
		} else {
			ctx.WriteString("未登录")
		}
	})
	app.Run(iris.Addr(":8080"))
}
