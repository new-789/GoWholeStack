package controllers

import (
	"BeegoOne/models"
	"bytes"
	"encoding/gob"
	"math"
	"path"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
	"github.com/weilaihui/fdfs_client"
)

// 文章功能控制器
type ArticleController struct {
	beego.Controller
}

// 上传文件函数封装
func UploadFile(this *beego.Controller, filePath string) string {
	//处理文件上传
	file, head, err := this.GetFile(filePath)
	if head.Filename == "" {
		return "NoImg"
	}

	if err != nil {
		this.Data["errmsg"] = "文件上传失败"
		this.TplName = "add.html"
		return ""
	}
	defer file.Close()

	//1.文件大小
	if head.Size > 5000000 {
		this.Data["errmsg"] = "文件太大，请重新上传"
		this.TplName = "add.html"
		return ""
	}

	//2.文件格式
	//a.jpg
	ext := path.Ext(head.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		this.Data["errmsg"] = "文件格式错误。请重新上传"
		this.TplName = "add.html"
		return ""
	}
	/*
		//3.防止重名
		fileName := time.Now().Format("2006-01-02-15:04:05") + ext
		//存储
		this.SaveToFile(filePath, "./static/img/"+fileName)
		// return "/static/img/" + fileName
	*/
	// 获取字节数组，大小和文件大小相等
	fileBuffer := make([]byte, head.Size)
	client, err := fdfs_client.NewFdfsClient("/etc/fdfs/client.conf")
	if err != nil {
		beego.Info("fdfs 链接错误", err)
		return ""
	}
	// 将文件字节流写入到字节数组中
	file.Read(fileBuffer)
	fdfsresponse, err := client.UploadByBuffer(fileBuffer, ext[1:])
	if err != nil {
		beego.Info("上传文件到 fdfs 失败", err)
		return ""
	}
	return fdfsresponse.RemoteFileId
}

/*处理下拉框改变发的请求*/
func (this *ArticleController) HandleSelect() {
	/*
		// 获取前端数据
		typeName := this.GetString("select")
		if typeName == "" {
			beego.Info("获取下拉框数据失败")
			return
		}
		// 根据获取到的下拉框数据，从数据库中查找相对应的文章信息
		ormObject := orm.NewOrm()
		Articles := make([]models.Goods, 0)
		//
			Filter 方法介绍：该方法类似于 sql 语句中的  where 语句，该方法接收两个参数
				- 参数一：用来过滤的条件字段 string 类型，如果是多表查询可使用 `__` 双下划线进行跨表查询，双下划线两边连接的是 `表名__另一张表中的字段`,前提是两张表建立了表关系
				- 参数二：条件字段匹配的值
				- 返回值：返回一个 QuerySeter 对象
			另外需要注意一点的是：如果是多表查询必须要在 QueryTable() 后面使用 `RelateSel()` 方法指定关联的表，否则会默认使用 orm 的惰性查询特性
		//
		ormObject.QueryTable("Goods").RelatedSel("GoodsType").Filter("GoodsType__Name", typeName).All(&Articles)
	*/
}

// 首页控制器函数，展示文章列表内容
func (this *ArticleController) ShowArticleList() {
	// 获取 session 中保存的用户名
	username := this.GetSession("username")
	typeName := this.GetString("select")
	// 1. 查询操作
	ormObject := orm.NewOrm()
	qs := ormObject.QueryTable("GoodsSKU")

	// 分页实现：获取总页数
	PageSize := 3 // 指定每页显示的内容条数
	// 获取后台传递的当前页码
	pageIndex, err := strconv.Atoi(this.GetString("pageIndex"))
	if err != nil {
		pageIndex = 1 // 指定默认显示的页码(即首页的页码)
	}
	// 定义切面保存根据文章类型获取到的文章数据集
	var articleWithType []models.GoodsSKU
	// 上一页下一页标志位
	FirstPage := false
	EndPage := false
	var CountPage float64
	var count int64
	// 当没有传递文章类型时,表示进入了首页,获取所有文章数据的数量
	if typeName == "" {
		// 获取数据库中文章的总条数
		count, err = qs.Count()
		if err != nil {
			beego.Info("查询错误")
			return
		}
		// 用浮点数相除为了避免出现无法整除半页的情况
		CountPage = float64(count) / float64(PageSize)
		// 判断上一页下一页数据改变按钮状态处理
		if pageIndex == 1 {
			FirstPage = true
		}
		if int(math.Ceil(CountPage)) <= pageIndex {
			EndPage = true
		}
		// 当没有传递起始页时表示进入了首页,此时查询所有数据进行显示
		qs.Limit(PageSize, PageSize*(pageIndex-1)).RelatedSel("GoodsType").All(&articleWithType)
	} else {
		// 根据前段传递的文章类型从数据库获文章表取当前文章类型的总条数
		count, err = qs.RelatedSel("GoodsType").Filter("GoodsType__Name", typeName).Count()
		// 用浮点数相除为了避免出现无法整除半页的情况
		CountPage = float64(count) / float64(PageSize)
		if err != nil {
			beego.Info("查询错误")
			return
		}
		// 判断上一页下一页数据改变按钮状态处理
		if pageIndex == 1 {
			FirstPage = true
		}
		if pageIndex > int(math.Ceil(CountPage)) {
			EndPage = true
		}
		// 用浮点数相除为了避免出现无法整除半页的情况
		CountPage = float64(count) / float64(PageSize)
		qs.Limit(PageSize, PageSize*(pageIndex-1)).RelatedSel("GoodsType").Filter("GoodsType__Name", typeName).All(&articleWithType)
	}

	// 获取类型数据
	articleTypes := make([]models.GoodsType, 0) // 存储所有 type 对象的数组
	conn, err := redis.Dial("tcp", ":6379")     // 链接 redis
	if err != nil {
		beego.Info("链接 redis 数据库错误", err)
	}
	// 从 redis 中读取数据
	rel, _ := redis.Bytes(conn.Do("get", "articleTypes"))
	// if err != nil {
	// 	beego.Info("获取 redis 数据失败", err)
	// 	return
	// }
	// 由于存入 redis 的是序列化为字节类型数据，所以在读取出来之后需要进行解码操作
	dec := gob.NewDecoder(bytes.NewReader(rel))
	dec.Decode(&articleTypes)

	// 如果从 redis 中获取到的数据为空则从 mysql 数据库获取文章类型数据，并存入 redis
	if len(articleTypes) == 0 {
		_, err = ormObject.QueryTable("GoodsType").All(&articleTypes)
		if err != nil {
			beego.Info("查询文章分类信息错误", err)
			return
		}
		// 执行 redis 存储操作,由于 redis 不支持自定义对象，所以需要现将自定义对象编码成字节流然后存入 redis
		var buffer bytes.Buffer
		enc := gob.NewEncoder(&buffer)
		enc.Encode(&articleTypes)
		_, err = conn.Do("set", "types", buffer.Bytes())
		if err != nil {
			beego.Info("Do 操作 redis 数据库错误", err)
			return
		}
	}

	this.Data["articleTypes"] = articleTypes
	// 将首页数据存入到 redis 可加快网页访问的速度

	// 传递数据给前端页面
	this.Data["count"] = count
	this.Data["FirstPage"] = FirstPage
	this.Data["endPage"] = EndPage
	this.Data["typeName"] = typeName
	this.Data["articles"] = articleWithType
	this.Data["username"] = username
	this.Data["pageIndex"] = pageIndex
	// math.Ceil() 向上取整，离需要取整的数最近，但比我大的数，与之相反的还有 math.Floor() 向下取整
	this.Data["countPage"] = math.Ceil(CountPage)
	// 指定 layout
	this.Layout = "layout.html"
	this.TplName = "index.html"
}

// 显示添加文章页面控制器
func (this *ArticleController) ShowAddArticle() {
	/*
		// 查询文章类型数据
		ormObject := orm.NewOrm()
		articleTypes := make([]models.ArticleType, 0)
		_, err := ormObject.QueryTable("ArticleType").All(&articleTypes)
		if err != nil {
			beego.Info("查询文件类型错误", err)
			return
		}
		username := this.GetSession("username")
		this.Data["username"] = username
		this.Data["articleTypes"] = articleTypes
		this.Layout = "layout.html"
		this.TplName = "add.html"
	*/
}

// 添加文章功能实现控制器实现(包含文件上传功能)
func (this *ArticleController) HandleAddArticle() {
	/*
		// 一.获取前端传递的数据
		title := this.GetString("articleName")
		content := this.GetString("content")
		// 获取前端上传的文件
		file, fileHeader, err := this.GetFile("uploadname")
		if err != nil {
			beego.Info("上传文件失败", err)
			return
		}
		defer file.Close()
		filename := fileHeader.Filename // 获取文件名
		fileSize := fileHeader.Size     // 获取文件大小
		ext := path.Ext(filename)       // 获取文件格式方法,直接返回文件后缀
		// 二、上传文件处理
		// 2.1. 判断文件格式
		if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
			beego.Info("上传的文件格式不正确")
			return
		}
		// 2.2 判断文件大小,字节类型
		if fileSize >= 500000 {
			beego.Info("上传的文件太大，不允许上传")
			return
		}
		// 2.3. 保证上传的文件不能重名
		saveFileName := time.Now().Format("2006-01-02 15:04:05")
		// 2.4. 保存文件
		err = this.SaveToFile("uploadname", "./static/img/"+saveFileName+ext)
		if err != nil {
			beego.Info("保存文件失败", err)
			return
		}

		// 三、插入内容到数据库
		// 3.1：获取 orm 对象
		ormObject := orm.NewOrm()
		// 3.2 创建表的插入对象
		article := models.Article{}
		// 3.3 将插入的内容赋值给结构体字段名
		article.Title = title
		article.Content = content
		article.Img = "/static/img/" + saveFileName + ext
		article.Time = time.Now()

		/* 给 Article 对象赋值分析,
		1. 首先将获取到前端传过来的类型数据，并检查
		2. 拿到 articleType 对象
		3. 然后将数据赋值给articleType表中的TypeName字段，
		4. 通过获取到的数据查找类型表中的数据检查前端传过来的内容在类型表中是否存在，因为类型表与文章表建立了关系，如果不存在则运行插入文章表中的关联字段,存在则会拿到对应的类型 Id
		5. 给文章表中的关联字段赋值(因为在关联表示关联字段使用的是指针类型，所以需要赋值 articleType 的地址)。
		6. 最后插入数据到文章表中的关联字段
		//
		// 获取下拉框传递过来的数据
		TypeName := this.GetString("select")
		if TypeName == "" {
			beego.Info("下拉框数据获取失败")
			return
		}
		// 获取 articleType 对象
		var articleType models.ArticleType
		articleType.TypeName = TypeName
		err = ormObject.Read(&articleType, "TypeName")
		if err != nil {
			beego.Info("获取类型错误")
			return
		}
		article.ArticleType = &articleType

		// 3.4：存入数据库操作
		_, err = ormObject.Insert(&article)
		if err != nil {
			beego.Info("添加内容失败", err)
			return
		}

		// 插入完成重定向到列表页进行展示
		this.Redirect("/Article/ShowArticle", 302)
	*/
}

// 文章详情页控制器
func (this *ArticleController) ShowContent() {
	/*
		// 获取前端传过来的数据
		id := this.GetString("id")
		// 查询数据库操作
		ormObject := orm.NewOrm()
		id2, _ := strconv.Atoi(id) // 将字符传转换为 int 型
		// 获取查询对象
		article := models.Article{Id: id2}
		err := ormObject.Read(&article, "Id")
		if err != nil {
			beego.Info("查询数据失败")
			return
		}
		// 查看一次更新一次浏览记录的值
		article.Count += 1

		// 多对多插入读者信息
		// 1. 获取操作对象
		// artile := models.Article{Id: id2}
		// 2. 获取多对多操作对象
		m2m := ormObject.QueryM2M(&article, "Users")
		// 3. 获取插入对象,即获取操作表关联字段对应表对象
		username := this.GetSession("username")
		user := models.User{}
		user.UserName = username.(string) //.(string)表式断言
		ormObject.Read(&user, "UserName")
		// 4. 多对多插入操作
		_, err = m2m.Add(&user)
		if err != nil {
			beego.Info("多对多插入失败")
			return
		}

		// 多对多查询
		// 方式一
		// ormObject.LoadRelated(&article, "Users")
		// 方式二: 过滤多对多查询,并去重处理
		var users []models.User
		ormObject.QueryTable("User").Filter("Articles__Article__Id", id2).Distinct().All(&users)
		this.Data["Users"] = users

		// 没有指定更新那一列数据，它会自动进行对比更新
		ormObject.Update(&article)

		this.GetSession("username")
		this.Data["username"] = username
		this.Data["article"] = article
		this.Layout = "layout.html"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["contentHead"] = "head.html"
		this.TplName = "content.html"
	*/
}

// 删除文件控制器
func (this *ArticleController) HandleDelete() {
	/*
		// 获取前端传过来的数据，并将其转换为 int 类型
		id, _ := strconv.Atoi(this.GetString("id"))
		beego.Info(id)
		// 删除数据操作
		ormObject := orm.NewOrm()         // 1.获取 orm 对象
		article := models.Article{Id: id} // 2.获取产出对象
		ormObject.Delete(&article)        // 3. 删除操作
		// 重定向到列表展示页面
		this.Redirect("/Article/ShowArticle", 302)
	*/
}

// 显示更新文章内容界面控制器
func (this *ArticleController) ShowUpdateArticle() {
	/*
		// 1. 获取前端传过来的 id 数据，并将其转换为 int 类型
		id, _ := strconv.Atoi(this.GetString("id"))
		if id == 0 {
			beego.Info("连接错误")
			return
		}
		// 2. 查询数据
		ormObject := orm.NewOrm()
		article := models.Article{Id: id}
		err := ormObject.Read(&article, "Id")
		if err != nil {
			beego.Info("查询数据错误")
			return
		}

		this.Data["username"] = this.GetSession("username")
		this.Data["article"] = article
		this.Layout = "layout.html"
		this.TplName = "update.html"
	*/
}

// 更新文章页面提交更新实现控制器
func (this *ArticleController) UpdateArticle() {
	/*
		// 获取前端传入的数据
		content := this.GetString("content")
		articleName := this.GetString("articleName")
		id, _ := strconv.Atoi(this.GetString("id"))
		// 判断前端传入的数据
		if content == "" || articleName == "" || id == 0 {
			beego.Info("更新数据失败")
			return
		}
		// 获取前端上传的静态文件
		file, header, err := this.GetFile("uploadname")
		if err != nil {
			beego.Info("没有选择上传的图片哦")
			this.Data["errmsg"] = "没有选择上传的图片哦"
			this.Layout = "layout.html"
			this.TplName = "update.html"
			return
		}
		defer file.Close()
		// 文件处理
		filename := header.Filename
		fileSize := header.Size
		ext := path.Ext(filename)
		if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
			this.Data["errmsg"] = "上传的文件格式不正确，请重传"
			this.Layout = "layout.html"
			this.TplName = "update.html"
			return
		}
		if fileSize >= 500000 {
			beego.Info("上传的文件太大，不允许上传", err)
			return
		}
		// 获取原来的文件路径用来替换并拿到文件名，然后将其替换为修改后的文件,防止服务器因文件内容太多撑爆存储空间，因为同名所以再次查看时因为 cookie 的关系可能不会马上改变过来
		ormObject := orm.NewOrm()
		article := models.Article{Id: id}
		// 在更新前一定要有一个读数据库数据操作，否则会更新失败
		err = ormObject.Read(&article, "Id")
		if err != nil {
			beego.Info("要更新的文章不存在")
			return
		}
		filePath := "." + article.Img
		err = this.SaveToFile("uploadname", filePath)
		if err != nil {
			this.Data["errmsg"] = "上传文件失败，更新内容失败"
			this.Layout = "layout.html"
			this.TplName = "update.html"
			return
		}
		// 更新数据库
		article.Title = articleName
		article.Content = content
		article.Img = article.Img
		_, err = ormObject.Update(&article)
		if err != nil {
			this.Data["errmsg"] = "数据库操作错误，更新失败，请稍后重传"
			this.Layout = "layout.html"
			this.TplName = "update.html"
			return
		}
		this.Redirect("/Article/ShowArticle", 302)
	*/
}

// 添加类别页面显示控制器
func (this *ArticleController) ShowAddType() {
	// 先从数据库读取到已有类别在页面上进行展示
	ormObj := orm.NewOrm()
	// 定义一个结构体表名类型的切片，用来接收查询到的所有数据
	ArticleTypes := make([]models.GoodsType, 0)
	// 查询表中的所有数据
	_, err := ormObj.QueryTable("GoodsType").All(&ArticleTypes)
	if err != nil {
		beego.Info("查询数据错误", err)
		// return
	}
	this.Data["username"] = this.GetSession("username")
	this.Data["ArticleTypes"] = ArticleTypes
	this.Layout = "layout.html"
	this.TplName = "addType.html"
}

// 添加分类业务实现控制器
func (this *ArticleController) HandleAddType() {
	// 获取前端数据
	typeName := this.GetString("typeName")
	logoPath := UploadFile(&this.Controller, "uploadlogo")
	typeImg := UploadFile(&this.Controller, "uploadTypeImage")
	if typeName == "" || logoPath == "" || typeImg == "" {
		beego.Info("信息不完整，请重新输入")
		return
	}
	ormObj := orm.NewOrm()
	// 获取插入对象，并给字段名赋值
	goodsType := models.GoodsType{
		Name:  typeName,
		Logo:  logoPath,
		Image: typeImg,
	}
	// 插入数据操作
	_, err := ormObj.Insert(&goodsType)
	if err != nil {
		beego.Info("数据库操作错误，添加分类失败", err)
		return
	}
	this.Redirect("/Article/AddArticleType", 302)
}

// 删除类型业务功能实现控制器
func (this *ArticleController) DelArticleType() {
	id, err := this.GetInt("id") // 获取前段传过来的数据
	if err != nil {
		beego.Info("没有获取到需要删除的类型id,删除错误", err)
		return
	}
	// 进行数据库删除操作
	ormObject := orm.NewOrm()
	goodsType := models.GoodsType{Id: id}
	_, err = ormObject.Delete(&goodsType)
	if err != nil {
		beego.Info("需要删除的类型不存在,删除失败", err)
		return
	}
	// 删除成功后跳转到类型页面
	this.Redirect("/Article/AddArticleType", 302)

}

// 发送邮件功能
func (this *ArticleController) SendMail() {
	/*
		// 定义邮件的配置信息
		config := `{"username":"2314574867@qq.com",
					"password":"oeqqcmqggieleaig",
					"host":"smtp.qq.com",
					"port":587
					}`
		// 获取邮箱实体
		email := utils.NewEMail(config)
		// 设置发送邮件的地址
		email.From = "2314574867@qq.com"
		// 设置邮件接收地址,值为字符串切片格式，切片中可以写多个接收地址用来群发
		email.To = []string{"zhuxiujian@foxmail.com"}
		// 设置邮件标题
		email.Subject = "某某操作系统激活账号邮件"
		// 设置邮件内容
		email.Text = "http://192.168.8.100:8086/active?id=1"
		// 设置以 html 渲染的内容,注：设置了 HTML 后会覆盖掉 Text 设置的邮件内容
		email.HTML = `<h1>特别提示</ht><p><a href="192.168.8.100:8086/active?id=1">点击链接激活账号：192.168.8.100:8086/active?id=1</a></p>`
		email.Send() // 发送邮件
		this.Ctx.WriteString("发送邮件成功")
	*/
}
