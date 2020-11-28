package controllers

import (
	"BeegoOne/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 用户注册控制器
type RegController struct {
	beego.Controller
}

// 用户登录控制器
type LoginController struct {
	beego.Controller
}

// 显示注册页面方法
func (this *RegController) ShowReg() {
	this.TplName = "register.html"
}

// 用户注册实现方法
func (this *RegController) HandlerReg() {
	// 1.获取页面发送过来的数据
	name := this.GetString("userName")
	pwd := this.GetString("password")
	// 2.检查接收到的数据是否为空
	if name != "" || pwd != "" {
		// 3. 插入数据库操作
		ormObject := orm.NewOrm() // 获取 orm 对象
		// 获取插入表对象，并给字段名进行赋值表示插入到数据库表的字段
		user := models.User{}
		user.UserName = name
		user.Password = pwd
		_, err := ormObject.Insert(&user) // 插入操作
		// 4. 插入数据失败则返回注册页面，成功则跳转到登录页面
		if err != nil {
			beego.Info("插入数据错误")
			this.TplName = "register.html"
			return
		}
		// 插入数据成功重定向到到登录页面,
		/*
			- Redirect 重定向,接收两个参数
				第一个参数为跳转的 url 地址 string 类型
				第二个参数为状态码 int 类型
					状态码的类型:
						1xx：继续发送 100
						2xx：请求成功 200
						3xx：资源转移 302 重定向
						4xx：请求错误 404
						5xx：服务器错误 500
			- 重定向和转发(渲染)的区别
				1. 重定向(Redirect)：浏览器端的请求，
					1.1：浏览器再次访问了服务器，url 地址会发生变化；
					1.2：不能传递数据
				2. 转发(TplName)：服务器端的功能，
					2.1：浏览器不会发送请求所以地址栏的 url 地址不会改变用的还是上一次的请求地址；
					2.2：该方法的另一个作用是传递数据给全段页面
		*/
		this.Redirect("/", 302)
	} else {
		beego.Info("用户名或者密码不能为空")
		this.TplName = "register.html"
		return
	}
}

// 显示登录页面方法
func (this *LoginController) ShowLogin() {
	// 读取 cookie
	username := this.Ctx.GetCookie("username")
	// 如果用户没有点击保存用户名按钮，则在登录逻辑中不设置 cookie，此处就不会获取到 cookie 中的值更不会将用户名显示在输入框记住用户名也不会默认选中
	if username != "" { // 不等于空则显示用户名和默认选中记住
		this.Data["username"] = username
		this.Data["check"] = "checked"
	}
	this.TplName = "login.html"
}

// 登录业务方法实现
func (this *LoginController) HandlerLogin() {
	// 1.获取前端输入的用户名和密码
	name := this.GetString("userName")
	pwd := this.GetString("password")
	// 2.数据处理
	if name != "" || pwd != "" {
		// 3.查找数据操作
		ormObject := orm.NewOrm()                // 获取 orm 对象
		user := models.User{UserName: name}      // 获取查找用户表对象
		err := ormObject.Read(&user, "UserName") // 从数据库中查找数据,参数二表示按什么字段进行查找
		if err != nil {
			beego.Info("登录失败,用户名错误")
			this.TplName = "login.html"
			return
		}
		// 4.判断密码是否一致
		if user.Password != pwd {
			beego.Info("登录失败,密码错误")
			this.TplName = "login.html"
			return
		}
		// 设置 cookie,记住用户名
		check := this.GetString("remember")
		if check == "on" {
			this.Ctx.SetCookie("username", name, time.Second*3600)
		} else {
			// 删除 cookie 中的值，即将 cookie 的保存时间更改为负数
			this.Ctx.SetCookie("username", "sss", -1)
		}
		// 登录成功将用户保存在 session 作用
		this.SetSession("username", name)

		// 5.返回视图
		this.Redirect("/Article/ShowArticle", 302)
	} else {
		beego.Info("用户名和密码不能为空")
		this.TplName = "login.html"
		return
	}
}

// 退出登录控制器
func (this *LoginController) HandlerLogout() {
	// 删除登录状态，即删除 session 中保存的信息
	this.DelSession("username")
	this.Redirect("/", 302)
}
