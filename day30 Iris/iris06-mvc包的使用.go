package main

//
import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// 自定义控制器结构体
type UserController struct{}

func main3006() {
	app := iris.New()
	// 设置自定义控制器根路由路径
	mvc.New(app).Handle(new(UserController))

	// 路由组的 mvc 处理方式
	mvc.Configure(app.Party("/user"), func(ctx *mvc.Application) {
		ctx.Handle(new(UserController))
	})

	app.Run(iris.Addr(":8080"))
}

// get url: http://localhost:8080
func (us *UserController) Get() string {
	iris.New().Logger().Info("Get 请求")
	return "hello world"
}

// post url:http://localhost:8080
func (us *UserController) Post() {
	iris.New().Logger().Info("post 请求")
}

// put 请求 url: http://localhost:8080
func (us *UserController) Put() {
	iris.New().Logger().Info("put 请求")
}

/*
第二类请求方式处理方法
	get url:http://localhost:8080/info ,该请求自动匹配到下面的方法
	GetInfo 表示接收一个 Get 的请求方法，且请求地址为 /info
*/
func (us *UserController) GetInfo() mvc.Result {
	iris.New().Logger().Info("get 请求，请求路径为 /info")
	// 返回 json 格式数据到前端
	return mvc.Response{
		Object: map[string]interface{}{
			"code":    1,
			"message": "请求成功",
		},
	}
}

func (us *UserController) PostLogin() mvc.Result {
	iris.New().Logger().Info("post 请求，请求路径为 /login")
	return mvc.Response{
		Object: map[string]interface{}{
			"code":    200,
			"message": "post 方式请求成功",
		},
	}
}

/*
	第三种请求方式：
		BeforeActivation 方法
*/
func (us *UserController) BeforeActivation(b mvc.BeforeActivation) {
	// 注：指定的处理方法首字母必须大写
	b.Handle("GET", "/query", "UserInfo")
}

// 对应处理请求方法
func (us *UserController) UserInfo() mvc.Result {
	iris.New().Logger().Info("/query 请求路径")
	return mvc.Response{
		Object: map[string]interface{}{
			"code":    200,
			"message": "/query url BeforeActivation 方式访问成功",
		},
	}
}
