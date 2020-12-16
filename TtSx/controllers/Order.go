package controllers

import (
	"TtSx/models"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
	alipay "github.com/smartwalle/alipay"
)

// 订单页还差分页显示功能
type OrderController struct {
	beego.Controller
}

// 订单页面展示控制器
func (o *OrderController) ShowOrder() {
	// 获取数据
	username := o.GetSession("username")
	skuids := o.GetStrings("skuid")
	if len(skuids) == 0 {
		beego.Info("获取订单请求数据错误")
		o.Redirect("/user/Mycart", 302)
		return
	}
	// 处理数据
	ormObj := orm.NewOrm()
	// 获取用数据
	user := models.User{
		Name: username.(string),
	}
	ormObj.Read(&user, "Name")
	// 链接 redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		beego.Info("订单展示页 redis 链接错误", err)
		o.Redirect("/user/Mycart", 302)
		return
	}
	defer conn.Close()
	goodsBuffer := make([]map[string]interface{}, len(skuids))

	totalPrice := 0
	totalCount := 0
	for i, skuid := range skuids {
		temp := make(map[string]interface{})
		id, _ := strconv.Atoi(skuid)
		// 查询商品数据
		goodsSku := models.GoodsSKU{
			Id: id,
		}
		ormObj.Read(&goodsSku)
		temp["goods"] = goodsSku
		// 从 redis 中获取商品数量
		count, _ := redis.Int(conn.Do("hget", "cart_"+strconv.Itoa(user.Id), id))
		temp["count"] = count
		// 计算小计
		amount := count * goodsSku.Price
		temp["amount"] = amount
		// 计算总金额和总件数
		totalCount += count
		totalPrice += amount
		goodsBuffer[i] = temp
	}
	// 获取收货地址数据
	var addrs []models.Address
	ormObj.QueryTable("Address").RelatedSel("User").Filter("User__Id", user.Id).All(&addrs)
	transferPrice := 10 // 定义运费价格
	// 返回数据视图
	o.Data["username"] = username
	o.Data["goodsBuffer"] = goodsBuffer
	o.Data["addrs"] = addrs
	o.Data["totalCount"] = totalCount
	o.Data["totalPrice"] = totalPrice
	o.Data["transferPrice"] = transferPrice
	o.Data["relyPrice"] = totalPrice + transferPrice

	// 传递所有商品的 id
	o.Data["skuids"] = skuids
	o.TplName = "place_order.html"
}

// 添加订单
func (o *OrderController) AddOrder() {
	// 准备返回给 ajax 请求数据结构
	resp := make(map[string]interface{})
	defer o.ServeJSON()
	// 获取数据
	addrId, _ := o.GetInt("addrId")
	payId, _ := o.GetInt("payId")
	skuid := o.GetString("skuids")
	ids := skuid[1 : len(skuid)-1]
	skuids := strings.Split(ids, " ")
	// totalPrice, _ := o.GetInt("totalPrice")
	totclCount, _ := o.GetInt("totalCount")
	transferPrice, _ := o.GetInt("transferPrice")
	relyPrice, _ := o.GetInt("relyPrice")
	username := o.GetSession("username")
	if len(skuids) == 0 {
		beego.Info("提交订单，获取数据错误")
		resp["code"] = 1
		resp["errmsg"] = "提交订单，获取数据错误"
		o.Data["json"] = resp
		return
	}
	// 处理数据
	ormObj := orm.NewOrm()

	// orm 开启事务操作
	ormObj.Begin()
	// 查询用户信息
	user := models.User{
		Name: username.(string),
	}
	ormObj.Read(&user, "Name")
	// 查询地址信息
	addr := models.Address{
		Id: addrId,
	}
	ormObj.Read(&addr, "Id")
	// 向订单表中插入数据
	order := models.OrderInfo{
		OrderId:      time.Now().Format("20060102150405") + strconv.Itoa(user.Id),
		User:         &user,
		Orderstatus:  1,
		PayMethod:    payId,
		TotalCount:   totclCount,
		TotalPrice:   relyPrice,
		TransitPrice: transferPrice,
		Address:      &addr,
	}
	ormObj.Insert(&order)
	// 向订单商品表插入数据
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		beego.Info("提交订单链接 redis 数据库错误", err)
		resp["code"] = 2
		resp["errmsg"] = "提交订单链接 redis 数据库错误"
		o.Data["json"] = resp
		ormObj.Rollback() // orm 事务回滚操作
		return
	}
	defer conn.Close()
	for _, skuid := range skuids {
		id, _ := strconv.Atoi(skuid)
		// 获取商品信息
		goods := models.GoodsSKU{
			Id: id,
		}
		i := 3
		for i > 0 {
			ormObj.Read(&goods)
			// 从 redis 数据库中获取商品数量
			count, err := redis.Int(conn.Do("hget", "cart_"+strconv.Itoa(user.Id), id))
			if err != nil {
				beego.Info("提交订单获取 redis 数据错误", err)
				resp["code"] = 3
				resp["errmsg"] = "提交订单获取 redis 数据错误"
				o.Data["json"] = resp
				ormObj.Rollback() // orm 事务回滚操作
				return
			}
			if count > goods.Stock {
				beego.Info("提交订单获取 redis 数据错误", err)
				resp["code"] = 4
				resp["errmsg"] = "提交订单错误，商品库存不足"
				o.Data["json"] = resp
				ormObj.Rollback() // orm 事务回滚操作
				return
			}
			preCount := goods.Stock
			// 向商品表和订单表多对多关系表中插入数据
			orderGoods := models.OrderGoods{
				GoodsSKU:  &goods,
				OrderInfo: &order,
				Count:     count,
				Price:     count * goods.Price,
			}
			ormObj.Insert(&orderGoods)

			// 更新库存
			goods.Stock -= count
			goods.Sales += count
			// orm 高级更新数据库语法，指定更新条件并更新指定字段
			updateCount, err := ormObj.QueryTable("GoodsSKU").Filter("Id", goods.Id).Filter("Stock", preCount).Update(orm.Params{"Stock": goods.Stock, "Sales": goods.Sales})
			if updateCount == 0 {
				if i > 0 {
					i -= 1
					continue
				}
				beego.Info("商品库存改变，订单提交失败", err)
				resp["code"] = 6
				resp["errmsg"] = "商品库存改变，订单提交失败"
				o.Data["json"] = resp
				ormObj.Rollback() // orm 事务回滚操作
				return
			} else {
				conn.Do("hdel", "cart_"+strconv.Itoa(user.Id), goods.Id)
				break
			}
		}
	}
	ormObj.Commit() // orm 事务提交方法
	// 返回数据
	resp["code"] = 5
	resp["errmsg"] = "ok"
	o.Data["json"] = resp
}

// alipay method
func (o *OrderController) HanldlePay() {
	// 获取前端传过来的商品id和总价
	orderid := o.GetString("orderid")
	totalPrice := o.GetString("totalPrice")
	if orderid == "" || totalPrice == "" {
		beego.Info("付款请求数据错误")
		o.Redirect("/user/userCenterOrder", 302)
		return
	}
	// 个人私钥
	var privateKey = "MIIEpAIBAAKCAQEAvVPlgbzVBMOxzVCYiTeVD/0C/HVcZuUCEUYYAl0g8fqCMdo1" +
		"3Av+iBtNtqumAkUWTIKUhDNr73a5WK+08qvuWdLtNqlYbrdjgHb9wiWj5GA0e/3i" +
		"L9M4W89wTLilLBgGblLHpLmUUB8ZKGkE6ngFy3iBuXp3cKTkTyr+OXaZHrXRqZxq" +
		"P3dM5ybOaNJB2YV12BNWC7+j+KfDg/7vEQVwiL4lf2Qxagi7sjr8hifI89eALerc" +
		"LZM3ZPjoIN7Ai5j1d+ERztZxhHOp/5oDbHFsiumAPXyUqLn/AS54gx1JwyAzFuit" +
		"jdbxiiaKZsJIFWSMBFyRlnKGgLAQVnH+TPQ2gwIDAQABAoIBAQCjBY2Wc/F8g8Ei" +
		"yCY/6QNKXBLpLWkeFEz+cLuVtXj7wze+E2XTDuThb5kwFIcCZ1i3Eq0tVFW5QSKB" +
		"cfI+DrtSzuOA8Lp7wx9Da7I/LUf9CrXGmircAWUC8RCCj29yE1HeRxsvBLjncI+q" +
		"ExuvhR+AP0h8XC5iaSlJ3fw/oJZiMLN329AJnnfYns++05Q2cvL3KbKJ1SAjLgfT" +
		"Ce9mOT05HHvh3ErTyk22qSNc+69g8J+8sDCOBMIfWzAMCRJDMVXlWxVpMwGu7mGN" +
		"rcWLkDhLwxn/BGSCBXwo0tcztTWRJP48Syhxdbou0XGOMalIRffOaA1m9CA3G62m" +
		"UpqHve75AoGBAOPIdIYYlopCukhgAZXJWQvOz/obYrECJ6l2a5mHV9MHbDhLn0f8" +
		"gIBkehQuFV7pZpZy12Zhy2+2Vuu+kyPq23CxOBdz5HI7lpt/+RAhUnlhCN7M00/V" +
		"xY0FbnFpOXCWqakPEh0G8iyYerVh6dWk423ea5RU0pH/pqk8iXq6Lze/AoGBANTH" +
		"7nkY5g+3nWvGMehW9w3JIcf94iek38gHK8pvGInyWBP/Nn+b2Eax8EcbUZe1ABdk" +
		"1lRbPwfcFSqHQ0uIL5yP9qKSivy0IQsM72K7O5plz+RZPMfAH7M4utxL8IR/gFq+" +
		"x8SEYS9QUEwTq3qcKsvvkT4UHLQz02RhGT7wV5I9AoGBAKFhhCz9fqq2XASrAHR/" +
		"Eveqru8kN+fw/uixXcqIeVCqEaJQ8aiu3NwaU+O4Qqvb8TLmanW/E8MyDKM5gkYQ" +
		"Pyi4ZmU3VsMYNLw3Dt6+vRACD6bKt2v4qN4g2v6+7lM2az27z7mBElNGNKoNlP4o" +
		"tHlC/DU7Y3TNC+cbXMv23T1BAoGAN38R1qnEz1Kfj015FNUhyImt8uBnzNk2uJAO" +
		"PtUs/EWl7UlCKJQ4G3ArcBBS/pNpu7BEKmpGDUG8j4QJ7DhP2rC9rfc0ouZEPAKa" +
		"qRCWYhWt3CveK2cTnYXtTqnLyHMMwh5CMiIoytNSrhTvQ6JZkykfo6ZROjrOJr8g" +
		"e4bqAC0CgYAKtowS8+FYKcDYd+0xcmLKii610vw/mMVYlBANDXXDqGxUIxkTQkaq" +
		"sVAO2ffM1M6bNKMVFTiR/uG2FcR1kjyQjn+QM34hsKJpGfZUeCVX6UKiKTaeNvFW" +
		"O9Gytkk2mTMtVhsW2DEea6SqpDN9IMfr2qTS9mhTEMmMWYT+kMX/9Q=="

	// 支付宝公钥
	var aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1ceK6Z8jSWy5HLDKSrs" +
		"TExhP70wBW9tAe7XBpjsw1O83Inq+lCxA3sQN45VANIN1WzVWkdRhI4gDB59LtL" +
		"xsuc5yigozk/HVcESeXdj6RyShB1Y/rW2jKQBXYkm+mR/9gJHaWXeqqa1XS47kQ" +
		"66N1aU8DRNdRyah4auHFrbK/2RYlceWyJ0TMN8rgSJ5su9RufNGrQz7CTaXxeR3" +
		"Bh19NNstQaOQW4IZiysvfY9mLLtYrANfSyHTDElwFOUiwt3RaJ2AxVwdZDw3h0l" +
		"/e9D5PrA42vZ7TMP+FjNd08mG+M8SwvDjElTRgZ6q9R3q2TEwZU045W0cSecuAk" +
		"VLabpONQIDAQAB"

	// 阿里开放平台创建的应用 Id
	appid := "2016092600599040"
	var client = alipay.New(appid, aliPublicKey, privateKey, false)
	// 指定支付页面为电脑端
	var p = alipay.AliPayTradePagePay{}
	// 异步请求地址
	p.NotifyURL = "http://xxx"
	// 同步地址
	p.ReturnURL = "http://192.168.8.100:8080/user/Payok"
	// titel 标题
	p.Subject = "珠江鲜啤-购物平台"
	// 唯一订单号
	p.OutTradeNo = orderid
	// 价格
	p.TotalAmount = totalPrice
	p.ProductCode = "QUICK_WAP_WAY"

	var url, err = client.TradePagePay(p)
	if err != nil {
		fmt.Println(err)
	}
	var payURL = url.String()
	o.Redirect(payURL, 302)
}

// 支付成功，修改订单状态方法
func (o *OrderController) Payok() {
	// 获取数据,获取的内容均为支付成功后支付宝返回 url 中所带的参数
	orderid := o.GetString("out_trade_no")
	if orderid == "" {
		beego.Info("支付返回数据错误")
		o.Redirect("/user/UserOrderCenter", 302)
		return
	}
	// 操作数据
	ormObj := orm.NewOrm()
	count, _ := ormObj.QueryTable("OrderInfo").Filter("OrderId", orderid).Update(orm.Params{"OrderId": 2})
	if count == 0 {
		beego.Info("支付返回更新数据错误")
		o.Redirect("/user/UserOrderCenter", 302)
		return
	}
	// 返回视图
	o.Redirect("/user/UserOrderCenter", 302)
}
