package controllers

import (
	"TtSx/models"
	"encoding/base64"
	"regexp"
	"strconv"

	"github.com/astaxie/beego/utils"
	"github.com/gomodule/redigo/redis"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

// 显示注册页面
func (u *UserController) ShowReg() {
	u.TplName = "register.html"
}

// 处理注册数据
func (u *UserController) HandleReg() {
	// 1. 获取数据
	userName := u.GetString("user_name")
	pwd := u.GetString("pwd")
	cpwd := u.GetString("cpwd")
	email := u.GetString("email")
	// 2. 校验数据
	if userName == "" || pwd == "" || cpwd == "" || email == "" {
		u.Data["errmsg"] = "提交是数据不完整，请重新注册"
		u.TplName = "register.html"
		return
	}
	if pwd != cpwd {
		u.Data["errmsg"] = "两次输入的密码不一致"
		u.TplName = "register.html"
		return
	}
	// 正则匹配邮箱
	reg, _ := regexp.Compile("^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$")
	if reg.FindString(email) == "" {
		u.Data["errmsg"] = "邮箱格式不正确"
		u.TplName = "register.html"
		return
	}
	// 3. 处理数据
	user := models.User{
		Name:     userName,
		PassWord: pwd,
		Email:    email,
	}
	ormObj := orm.NewOrm()
	_, err := ormObj.Insert(&user)
	if err != nil {
		u.Data["errmsg"] = "注册失败，请更换数据注册"
		u.TplName = "register.html"
		return
	}
	// 发送邮件激活用户
	emailConfig := `{"username":"2314574867@qq.com",
				"password":"oeqqcmqggieleaig",
				"host":"smtp.qq.com",
				"port":587
				}`
	// 获取邮箱实体
	emailObj := utils.NewEMail(emailConfig)
	// 设置发送邮件的地址
	emailObj.From = "2314574867@qq.com"
	// 设置邮件接收地址,值为字符串切片格式，切片中可以写多个接收地址用来群发
	emailObj.To = []string{email}
	// 设置邮件标题
	emailObj.Subject = "珠江鲜啤，鲜活每一天·系统注册服务"
	// 设置邮件内容
	// emailObj.Text = "192.168.8.100:8080/active?id=" + strconv.Itoa(user.Id)
	emailObj.HTML = "<h1>珠江鲜啤用户注册,请点击下面的链接激活账号</h1><p><a href=192.168.8.100:8080/active?id=" + strconv.Itoa(user.Id) + ">点击链接激活账号：192.168.8.100:8080/active?id=" + strconv.Itoa(user.Id) + "</a></p>"

	emailObj.Send()

	// 4.返回视图
	u.Data["active"] = "注册成功，请前往邮箱激活您的账号"
	u.TplName = "register.html"
}

// 用户激活方法
func (u *UserController) ActiveUser() {
	// 获取发送过来的用户 id
	id, err := u.GetInt("id")
	if err != nil {
		u.Data["errmsg"] = "要激活的用户不存在"
		u.TplName = "register.html"
		return
	}
	// 处理数据，做数据库查询操作及更改数据库中用户的激活状态信息
	ormObj := orm.NewOrm()
	user := models.User{
		Id: id,
	}
	err = ormObj.Read(&user)
	if err != nil {
		u.Data["errmsg"] = "要激活的用户不存在"
		u.TplName = "register.html"
		return
	}
	user.Active = true
	ormObj.Update(&user)
	// 返回登录视图
	u.Redirect("/login", 302)
}

// 展示用户登录
func (u *UserController) ShowLogin() {
	username := u.Ctx.GetCookie("username")
	// 解码：将存在 cookie 中 byte 类型的用户名解码
	temp, _ := base64.StdEncoding.DecodeString(username)
	if string(temp) != "" {
		u.Data["username"] = string(temp)
		u.Data["check"] = "checked"
	}
	u.TplName = "login.html"
}

// 用户登录功能业务处理
func (u *UserController) HandleLogin() {
	// 1. 获取数据
	username := u.GetString("username")
	pwd := u.GetString("pwd")
	// 2.校验数据
	if username == "" || pwd == "" {
		u.Data["errmsg"] = "登录数据不完整，登录失败"
		u.TplName = "login.html"
		return
	}
	// 3. 处理数据,根据获取到的用户名和密码从数据库查询数据操作，并做各种判断
	ormObj := orm.NewOrm()
	user := models.User{
		Name: username,
	}
	err := ormObj.Read(&user, "Name")
	if err != nil {
		u.Data["errmsg"] = "用户名或密码错误，登录失败"
		u.TplName = "login.html"
		return
	}
	if user.PassWord != pwd {
		u.Data["errmsg"] = "用户名或密码错误，登录失败"
		u.TplName = "login.html"
		return
	}
	if user.Active != true {
		u.Data["errmsg"] = "用户尚未激活，请先前往邮箱激活用户"
		u.TplName = "login.html"
		return
	}
	// 记住用户名操作
	remember := u.GetString("remember")
	if remember == "on" {
		// cookie 中不能存储中文，所以需先将获取到的用户名转换为 byte 类型，以防用户名存在中文的情况
		ByteUser := base64.StdEncoding.EncodeToString([]byte(username))
		u.Ctx.SetCookie("username", ByteUser, 24*3600*30)
	} else {
		// 删除 cookie
		u.Ctx.SetCookie("username", username, -1)
	}
	// 4. 登录成功设置 session 返回视图,跳转到首页
	u.SetSession("username", username)
	u.Redirect("/", 302)
}

// 退出登录,退出登录后跳转到登录页面
func (u *UserController) Logout() {
	u.DelSession("username")
	u.Redirect("/login", 302)
}

// 展示用户中心页功能控制器
func (u *UserController) ShowUserCenterInfo() {
	// 获取用户相关信息
	username := GetSessionUser(&u.Controller)
	// 思考，不登录能访问到这个函数吗？
	// 从数据库查询用户收货地址等信息
	ormObj := orm.NewOrm()
	// 高级查询，表关联
	var addr models.Address
	ormObj.QueryTable("Address").RelatedSel("User").Filter("User__Name", username).Filter("Isdefault", true).One(&addr)

	// 返回数据和视图
	if addr.Id == 0 {
		u.Data["addr"] = ""
	} else {
		u.Data["addr"] = addr
	}
	// 从 redis 中获取用户浏览历史记录
	if username != "" {
		user := models.User{
			Name: username,
		}
		ormObj.Read(&user, "Name")
		// 获取用户浏览历史记录
		conn, err := redis.Dial("tcp", "127.0.0.1:6379")
		defer conn.Close()
		if err != nil {
			beego.Info("user 模块连接 redis 数据库错误")
		}
		rep, err := conn.Do("lrange", "history_"+strconv.Itoa(user.Id), 0, 4)
		// 回复助手函数
		goodsIds, _ := redis.Ints(rep, err)
		var goodsSKUs []models.GoodsSKU
		for _, val := range goodsIds {
			var goods models.GoodsSKU
			goods.Id = val
			ormObj.Read(&goods)
			goodsSKUs = append(goodsSKUs, goods)
		}
		u.Data["goodsSKUs"] = goodsSKUs
	}
	u.Layout = "userCenterLayout.html"
	u.TplName = "user_center_info.html"
}

// 展示用户中心订单页
func (u *UserController) ShowUserCenterOrder() {
	username := GetSessionUser(&u.Controller)
	ormObj := orm.NewOrm()
	user := models.User{
		Name: username,
	}
	ormObj.Read(&user, "Name")
	// 获取订单表数据
	var orderInfos []models.OrderInfo
	ormObj.QueryTable("OrderInfo").RelatedSel("User").Filter("User__Id", user.Id).All(&orderInfos)
	// 获取订单商品表数据
	goodsBuffer := make([]map[string]interface{}, len(orderInfos))
	for idx, orderInfo := range orderInfos {
		var orderGoods []models.OrderGoods
		ormObj.QueryTable("OrderGoods").RelatedSel("OrderInfo", "GoodsSKU").Filter("OrderInfo__Id", orderInfo.Id).All(&orderGoods)
		temp := make(map[string]interface{})
		temp["orderInfo"] = orderInfo
		temp["orderGoods"] = orderGoods
		goodsBuffer[idx] = temp
	}
	u.Data["goodsBuffer"] = goodsBuffer
	u.Layout = "userCenterLayout.html"
	u.TplName = "user_center_order.html"
}

// 展示用户中心地址页
func (u *UserController) ShowUserCenterSite() {
	username := GetSessionUser(&u.Controller)
	// u.Data["username"] = username
	// 查询数据库数据获取地址信息
	ormObj := orm.NewOrm()
	var addr models.Address
	ormObj.QueryTable("Address").RelatedSel("User").Filter("User__Name", username).Filter("Isdefault", true).One(&addr)

	u.Data["addr"] = addr

	u.Layout = "userCenterLayout.html"
	u.TplName = "user_center_site.html"
}

// 处理用户中心地址数据
func (u *UserController) HandleUserCenterSite() {
	// 获取页面数据
	receive := u.GetString("receive")
	addr := u.GetString("addr")
	zipcode := u.GetString("zipcode")
	phone := u.GetString("phone")
	// 校验数据
	if receive == "" || addr == "" || zipcode == "" || phone == "" {
		beego.Info("添加数据不完整")
		u.Redirect("/user/userCenterSite", 302)
		return
	}
	// 处理数据
	ormObj := orm.NewOrm()
	var addrUser models.Address
	addrUser.Isdefault = true
	// 查询默认地址
	err := ormObj.Read(&addrUser, "Isdefault")
	// 错误等于 nil 说明有默认地址，因为是添加所以修改之前的地址为不是默认地址
	if err == nil {
		addrUser.Isdefault = false
		ormObj.Update(&addrUser)
	}
	// 此处有坑注意：更新地址时，如果使用原来的 addrUser 对象则说明是给原地址对象的 ID 赋值，如果插入则 id 相同插入数据哭会失败，所以需要新建一个地址对象
	// 做数据表关联
	username := u.GetSession("username")
	user := models.User{
		Name: username.(string),
	}
	// 注意此处必须按照 Name 进新查找，否则后面插入数据会错误
	ormObj.Read(&user, "Name")
	addrUserNew := models.Address{
		Receiver:  receive,
		Zipcode:   zipcode,
		Addr:      addr,
		Phone:     phone,
		Isdefault: true,
		User:      &user,
	}
	ormObj.Insert(&addrUserNew)
	// 返回视图
	u.Redirect("/user/userCenterSite", 302)
}
