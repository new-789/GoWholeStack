package controllers

import (
	"TtSx/models"
	"math"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
)

// 获取 session 中保存的用户名方法
func GetSessionUser(ctl *beego.Controller) string {
	username := ctl.GetSession("username")
	if username == nil {
		ctl.Data["username"] = ""
		// 如果没有获取到用户名则跳转到登录页面
		// g.Redirect("/login", 302)
	} else {
		ctl.Data["username"] = username.(string)
		return username.(string)
	}
	return ""
}

// 获取商品类型方法
func ShowLayout(this *beego.Controller) {
	// 查询类型
	ormObj := orm.NewOrm()
	var types []models.GoodsType
	ormObj.QueryTable("GoodsType").All(&types)
	this.Data["types"] = types
	// 获取用户信息
	GetSessionUser(this)
	// 指定 Layout
	this.Layout = "goodsLayout.html"
}

// 分页控制助手函数
func PageTool(pageCount, pageIndex int) (pages []int) {
	if pageCount <= 5 {
		pages = make([]int, pageCount)
		for i, _ := range pages {
			pages[i] = i + 1
		}
	} else if pageIndex <= 3 {
		pages = []int{1, 2, 3, 4, 5}
	} else if pageIndex > pageCount-3 {
		pages = []int{pageCount - 4, pageCount - 3, pageCount - 2, pageCount - 1, pageCount}
	} else {
		pages = []int{pageIndex - 2, pageIndex - 1, pageIndex, pageIndex + 1, pageIndex + 2}
	}
	return
}

type GoodsController struct {
	beego.Controller
}

// 首页展示控制器
func (g *GoodsController) ShowIndex() {
	GetSessionUser(&g.Controller)
	// 获取数据库数据
	ormObj := orm.NewOrm()
	// 获取类型数据
	goodsTypes := make([]models.GoodsType, 0)
	ormObj.QueryTable("GoodsType").All(&goodsTypes)
	g.Data["goodsTypes"] = goodsTypes

	// 获取轮播图数据
	var IndexGoodsBanner []models.IndexGoodsBanner
	// 查询所有轮播图顺序，并按照索引进行排倒序
	ormObj.QueryTable("IndexGoodsBanner").OrderBy("Index").All(&IndexGoodsBanner)
	g.Data["IndexGoodsBanner"] = IndexGoodsBanner

	// 获取促销品数据
	var promotionGoods []models.IndexPromotionBanner
	ormObj.QueryTable("IndexPromotionBanner").OrderBy("Index").All(&promotionGoods)
	g.Data["promotionGoods"] = promotionGoods

	/* 获取首页展示商品数据
	数据结构:
		[ -->  goods
			{"type": -->  temp1
				"类型名称",
			},
			{"textGoods":  --> 文字商品
				[
						"商品名称",
				]
			},
			{"imageGoods"  --> 图片商品
				"商品名称",
			}
		]
	*/
	goods := make([]map[string]interface{}, len(goodsTypes))
	// 向 interface 切面中插入类型数据
	for inx, val := range goodsTypes {
		// 获取对应类型的首页展示商品
		temp := make(map[string]interface{})
		temp["type"] = val
		goods[inx] = temp
		// goods = append(goods, temp)
	}
	// 商品数据
	for _, val := range goods {
		var textGoods []models.IndexTypeGoodsBanner
		var imgGoods []models.IndexTypeGoodsBanner
		// 获取文字商品数据
		ormObj.QueryTable("IndexTypeGoodsBanner").RelatedSel("GoodsType", "GoodsSKU").OrderBy("Index").Filter("GoodsType", val["type"]).Filter("Display_Type", 0).All(&textGoods)
		// 获取图片商品数据
		ormObj.QueryTable("IndexTypeGoodsBanner").RelatedSel("GoodsType", "GoodsSKU").OrderBy("Index").Filter("GoodsType", val["type"]).Filter("Display_Type", 1).All(&imgGoods)
		val["textGoods"] = textGoods
		val["imgGoods"] = imgGoods
	}
	cartCount := GetCartCount(&g.Controller)
	g.Data["cartCount"] = cartCount
	// 返回视图
	g.Data["goods"] = goods
	g.TplName = "index.html"
}

// 商品详情页展示
func (g *GoodsController) ShowGoodsDetail() {
	// 获取前端数据
	id, err := g.GetInt("id")
	if err != nil {
		beego.Error("浏览器请求错误")
		g.Redirect("/", 302)
		return
	}
	// 数据库查询数据
	ormObj := orm.NewOrm()
	goodsSku := models.GoodsSKU{
		Id: id,
	}
	// ormObj.Read(&goodsSku, "Id")
	ormObj.QueryTable("GoodsSKU").RelatedSel("GoodsType", "Goods").Filter("Id", id).One(&goodsSku)
	// 获取同类型时间靠前的两条商品数据
	var goodsNew []models.GoodsSKU
	ormObj.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType", goodsSku.GoodsType).OrderBy("Time").Limit(2, 0).All(&goodsNew)
	// 添加历史浏览记录
	// 判断用户是否登录
	username := g.GetSession("username")
	if username != nil {
		// 查询用户信息
		user := models.User{
			Name: username.(string),
		}
		ormObj.Read(&user, "Name")
		// 添加历史记录,使用 redis 存储
		conn, err := redis.Dial("tcp", "127.0.0.1:6379")
		defer conn.Close()
		if err != nil {
			beego.Info("redis 链接错误:", err)
		}
		// 添加浏览记录前先将以前相同商品的浏览历史记录删除
		conn.Do("lrem", "history_"+strconv.Itoa(user.Id), 0, id)
		// 添加新的商品浏览记录
		conn.Do("lpush", "history_"+strconv.Itoa(user.Id), id)
	}

	// 返回数据视图
	g.Data["goodsSku"] = goodsSku
	g.Data["goodsNew"] = goodsNew
	ShowLayout(&g.Controller)
	cartCount := GetCartCount(&g.Controller)
	g.Data["cartCount"] = cartCount
	g.TplName = "detail.html"
}

// 展示商品列表页
func (g *GoodsController) ShowGoodsList() {
	// 获取前端数据
	typeId, err := g.GetInt("typeId")
	if err != nil {
		beego.Error("浏览器请求错误，没有接收到类型 id 无法获取到对应的内容")
		g.Redirect("/", 302)
	}
	// 处理数据
	ShowLayout(&g.Controller)
	// 获取新品数据
	ormObj := orm.NewOrm()
	var goodsNew []models.GoodsSKU
	ormObj.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id", typeId).OrderBy("Time").Limit(2).All(&goodsNew)
	g.Data["goodsNew"] = goodsNew

	// 获取商品信息
	var goods []models.GoodsSKU
	ormObj.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id", typeId).All(&goods)

	// 分页实现
	// 获取总页码
	count, _ := ormObj.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id", typeId).Count()
	pageSize := 3
	pageCount := math.Ceil(float64(count)) / float64(pageSize)
	pageIndex, err := g.GetInt("pageIndex")
	if err != nil {
		pageIndex = 1
	}
	pages := PageTool(int(pageCount), pageIndex)
	g.Data["pages"] = pages
	g.Data["pageIndex"] = pageIndex
	g.Data["typeId"] = typeId

	start := (pageIndex - 1) * pageSize

	// 获取上一页页码
	prePage := pageIndex - 1
	if prePage <= 1 {
		prePage = 1
	}
	g.Data["prePage"] = prePage
	// 获取下一页页码
	nextPage := pageIndex + 1
	if nextPage >= int(pageCount) {
		nextPage = int(pageCount)
	}
	g.Data["nextPage"] = nextPage

	// 按照一定顺序获取商品
	sort := g.GetString("sort")
	if sort == "" {
		ormObj.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id", typeId).Limit(pageSize, start).All(&goods)
		g.Data["goods"] = goods
		g.Data["sort"] = ""
	} else if sort == "price" {
		ormObj.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id", typeId).OrderBy("Price").Limit(pageSize, start).All(&goods)
		g.Data["goods"] = goods
		g.Data["sort"] = "price"
	} else {
		ormObj.QueryTable("GoodsSKU").RelatedSel("GoodsType").Filter("GoodsType__Id", typeId).OrderBy("Sales").Limit(pageSize, start).All(&goods)
		g.Data["goods"] = goods
		g.Data["sort"] = "sale"
	}
	cartCount := GetCartCount(&g.Controller)
	g.Data["cartCount"] = cartCount
	// 返回视图
	g.TplName = "list.html"
}

// 商品搜索功能实现
func (g *GoodsController) HandleSearch() {
	ormObj := orm.NewOrm()
	var goods []models.GoodsSKU
	// 获取前段数据
	goodsName := g.GetString("goodsName")
	// 如果为空则显示所有数据
	if goodsName == "" {
		ormObj.QueryTable("GoodsSKU").All(&goods)
		g.Data["goods"] = goods
		ShowLayout(&g.Controller)
		g.TplName = "search.html"
	}
	// 处理数据
	ormObj.QueryTable("GoodsSKU").Filter("Name__icontains", goodsName).All(&goods)
	g.Data["goods"] = goods
	ShowLayout(&g.Controller)
	g.TplName = "search.html"
}
