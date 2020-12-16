package routers

import (
	"BeegoOne/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	// 注册过滤器函数
	beego.InsertFilter("/Article/*", beego.BeforeExec, filterFunc)

	// 注册路由
	beego.Router("/register", &controllers.RegController{}, "get:ShowReg;post:HandlerReg")
	// 登录路由
	beego.Router("/", &controllers.LoginController{}, "get:ShowLogin;post:HandlerLogin")
	beego.Router("/logout", &controllers.LoginController{}, "get:HandlerLogout")
	// 首页路由,展示文件列表页
	beego.Router("/Article/ShowArticle", &controllers.ArticleController{}, "get:ShowArticleList;post:HandleSelect")
	// 插入文章路由
	beego.Router("/Article/AddArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")
	// 查看文章详情路由
	beego.Router("/Article/ArticleContent", &controllers.ArticleController{}, "get:ShowContent")
	// 删除文章路由
	beego.Router("/Article/DeleteArticle", &controllers.ArticleController{}, "get:HandleDelete")
	// 更新文章路由
	beego.Router("/Article/EditArticle", &controllers.ArticleController{}, "get:ShowUpdateArticle;post:UpdateArticle")
	// 添加文章类型路由
	beego.Router("/Article/AddArticleType", &controllers.ArticleController{}, "get:ShowAddType;post:HandleAddType")
	// 删除文章类型路由
	beego.Router("/Article//DelType", &controllers.ArticleController{}, "get:DelArticleType")
	// 发送邮件功能路由
	beego.Router("/sendEmail", &controllers.ArticleController{}, "get:SendMail")
}

// 过滤器函数执行的自定义函数
var filterFunc = func(ctx *context.Context) {
	username := ctx.Input.Session("username")
	if username == nil {
		ctx.Redirect(302, "/")
		return
	}
}
