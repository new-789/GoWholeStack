package routers

import (
	"TtSx/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	// 注册过滤器
	beego.InsertFilter("/user/*", beego.BeforeExec, filterFunc)

	// 注册页面路由
	beego.Router("/register", &controllers.UserController{}, "get:ShowReg;post:HandleReg")
	// 用户激活路由
	beego.Router("/active", &controllers.UserController{}, "get:ActiveUser")
	// 用户登录
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	// 登录成功之后跳转首页
	beego.Router("/", &controllers.GoodsController{}, "get:ShowIndex")
	// 退出登录功能路由
	beego.Router("/user/logout", &controllers.UserController{}, "get:Logout")
	// 用户中心信息页面路由
	beego.Router("/user/userCenterInfo", &controllers.UserController{}, "get:ShowUserCenterInfo")
	// 用户中心订单页
	beego.Router("/user/userCenterOrder", &controllers.UserController{}, "get:ShowUserCenterOrder")
	// 用户中心地址页
	beego.Router("/user/userCenterSite", &controllers.UserController{}, "get:ShowUserCenterSite;post:HandleUserCenterSite")
	// 商品详情展示页路由
	beego.Router("/goodsDetail", &controllers.GoodsController{}, "get:ShowGoodsDetail")
	// 商品列表页展示路由
	beego.Router("/goodsList", &controllers.GoodsController{}, "get:ShowGoodsList")
	// 商品搜索框路由
	beego.Router("/goodsSerach", &controllers.GoodsController{}, "post:HandleSearch")
	// 添加购物车路由
	beego.Router("/user/addCart", &controllers.CartController{}, "get,post:HandleAddCart")
	// 展示购物车页面
	beego.Router("/user/Mycart", &controllers.CartController{}, "get:ShowCart")
	// 购物车中更新物品数量
	beego.Router("/user/UpdataCart", &controllers.CartController{}, "post:HandleUpdateCart")
	// 删除购物车商品数据
	beego.Router("/user/deleteCart", &controllers.CartController{}, "post:HandleDeleteCart")
	// 订单页面展示
	beego.Router("/user/showOrder", &controllers.OrderController{}, "post:ShowOrder")
	// 提交订单
	beego.Router("/user/addOrder", &controllers.OrderController{}, "post:AddOrder")
	// 阿里支付路由
	beego.Router("/user/Pay", &controllers.OrderController{}, "get:HanldlePay")
	// 支付成功修改订单状态
	beego.Router("/user/Payok", &controllers.OrderController{}, "get:Payok")
}

// 过滤器函数
var filterFunc = func(ctx *context.Context) {
	// 从 session 获取用户数据
	username := ctx.Input.Session("username")
	if username == nil {
		ctx.Redirect(302, "/login")
		return
	}
}
