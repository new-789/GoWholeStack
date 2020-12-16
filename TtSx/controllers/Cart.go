package controllers

import (
	"TtSx/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
)

type CartController struct {
	beego.Controller
}

// 获取购物车数量
func GetCartCount(this *beego.Controller) int {
	// 从 redis 中获取数据
	username := this.GetSession("username")
	if username == nil {
		return 0
	}
	ormObj := orm.NewOrm()
	user := models.User{
		Name: username.(string),
	}
	ormObj.Read(&user, "Name")
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return 0
	}
	defer conn.Close()
	cartCount, err := redis.Int(conn.Do("hlen", "cart_"+strconv.Itoa(user.Id)))
	if err != nil {
		return 0
	}
	return cartCount
}

// 添加购物车实现方法
func (c *CartController) HandleAddCart() {
	// 获取校验数据
	id, err1 := c.GetInt("skuid")
	num, err2 := c.GetInt("num")
	typeid := c.GetString("typeid")
	resp := make(map[string]interface{})
	defer c.ServeJSON()
	if err1 != nil || err2 != nil {
		c.Data["code"] = 1
		c.Data["msg"] = "传递的数据不正确"
		c.Data["json"] = resp
		return
	}
	username := c.GetSession("username")
	if username == "" {
		c.Data["code"] = 2
		c.Data["msg"] = "您当前尚未登录，请先登录"
		c.Data["json"] = resp
		return
	}
	ormObj := orm.NewOrm()
	user := models.User{
		Name: username.(string),
	}
	ormObj.Read(&user, "Name")
	// 处理数据
	// 将购物车数据存在 redis 中，使用 hash 类型
	conn, err3 := redis.Dial("tcp", "127.0.0.1:6379")
	if err3 != nil {
		beego.Info("添加购物车模块链接 redis 错误")
		return
	}
	defer conn.Close()
	// 先获取原来的数量,然后和购物车中已有商品的数量加起来
	preCount, _ := redis.Int(conn.Do("hget", "cart_"+strconv.Itoa(user.Id), id))
	conn.Do("hset", "cart_"+strconv.Itoa(user.Id), id, num+preCount)
	// 使用回复助手函数，将从 redis 中获取到的值转换为需要的类型
	cartCount, _ := redis.Int(conn.Do("hlen", "cart_"+strconv.Itoa(user.Id)))

	resp["cartCount"] = cartCount
	// 返回 json 数据
	resp["code"] = 5
	resp["msg"] = "ok"
	// 指定返回的数据格式
	c.Data["json"] = resp
	// 将 json 数据格式返回给前段
	if typeid != "" {
		if typeid == "5" {
			c.Redirect("/user/userCenterInfo", 302)
		} else {
			c.Redirect("/goodsList?typeId="+typeid, 302)
		}
	}
}

// 展示购物车页面
func (c *CartController) ShowCart() {
	username := GetSessionUser(&c.Controller)
	ormObj := orm.NewOrm()
	user := models.User{
		Name: username,
	}
	ormObj.Read(&user, "Name")
	// 从 redis 中获取数据
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		beego.Info("购物车模块链接redis 错误", err)
		return
	}
	defer conn.Close()
	goodsMap, err := redis.IntMap(conn.Do("hgetall", "cart_"+strconv.Itoa(user.Id)))
	if err != nil {
		beego.Info("获取redis 数据错误", err)
		return
	}
	goods := make([]map[string]interface{}, len(goodsMap))
	var i = 0
	var totalPrice = 0
	var totalNum = 0
	for idx, val := range goodsMap {
		skuid, _ := strconv.Atoi(idx)
		goodsSku := models.GoodsSKU{
			Id: skuid,
		}
		ormObj.Read(&goodsSku)
		temp := make(map[string]interface{})
		temp["goodsSku"] = goodsSku
		temp["num"] = val
		temp["addPrice"] = goodsSku.Price * val
		totalPrice += goodsSku.Price * val
		totalNum += val
		goods[i] = temp
		i += 1
	}
	c.Data["goods"] = goods
	c.Data["totalPrice"] = totalPrice
	c.Data["totalNum"] = totalNum
	c.TplName = "cart.html"
}

// 购物车中更新物品数量
func (c *CartController) HandleUpdateCart() {
	// 获取数据
	skuid, err1 := c.GetInt("skuid")
	count, err2 := c.GetInt("count")
	resp := make(map[string]interface{})
	// 获取数据错误返回 json 数据
	defer c.ServeJSON()
	if err1 != nil || err2 != nil {
		resp["code"] = 1
		resp["errmsg"] = "请求数据不正确"
		c.Data["json"] = resp
	}
	username := c.GetSession("username")
	if username == nil {
		resp["code"] = 3
		resp["errmsg"] = "当前用户未登录"
		c.Data["json"] = resp
	}
	ormObj := orm.NewOrm()
	user := models.User{
		Name: username.(string),
	}
	ormObj.Read(&user, "Name")
	// 处理数据根据用户 ID 将数据更新到 redis 中
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		resp["code"] = 2
		resp["errmsg"] = "链接 redis 数据库失败"
		c.Data["json"] = resp
	}
	defer conn.Close()
	conn.Do("hset", "cart_"+strconv.Itoa(user.Id), skuid, count)
	resp["code"] = 5
	resp["errmsg"] = "ok"
	c.Data["json"] = resp
}

// 删除购物车中商品数据
func (c *CartController) HandleDeleteCart() {
	username := c.GetSession("username")
	resp := make(map[string]interface{})
	defer c.ServeJSON()
	skuid, err := c.GetInt("skuid")
	if err != nil {
		resp["code"] = 1
		resp["errmsg"] = "请求数据不正确"
		c.Data["json"] = resp
		return
	}
	// 处理数据,根据用户 id 从 redis 中删除指定内容
	ormObj := orm.NewOrm()
	user := models.User{
		Name: username.(string),
	}
	ormObj.Read(&user, "Name")
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		resp["code"] = 2
		resp["errmsg"] = "购物车删除商品功能后台链接 redis 数据库错误"
		c.Data["json"] = resp
		return
	}
	defer conn.Close()
	conn.Do("hdel", "cart_"+strconv.Itoa(user.Id), skuid)
	// if err != nil {
	// 	resp["code"] = 3
	// 	resp["errmsg"] = "后台操作 redis 数据库错误，请稍后在尝试"
	// 	c.Data["json"] = resp
	// 	return
	// }
	resp["code"] = 5
	resp["errmsg"] = "ok"
	c.Data["json"] = resp
}
