package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 1. 放结构体，即表设计的结构体
// 2. 放表的初始化操作方法

// 定义用来对应数据库表的结构体,注结构体字段首字母必须为大写
type User struct {
	Id       int
	UserName string
	Password string
	Articles []*Article `orm:"rel(m2m)"`
}

// 定义文章表在数据库中的表解构
type Article struct {
	// 文章标题
	Id int `orm:"pk;auto"`
	// 文章标题
	Title string `orm:"size(100)"`
	// 文章内容
	Content string `orm:"size(500)"`
	// 图片，存放是路径
	Img string `orm:"size(50);null"`
	// 文章发布时间, 时间类型
	Time time.Time `orm:"type(datetime);auto_now_add"`
	// 阅读量
	Count int `orm:"default(0)"`
	// 建立与文章类型的一对多关系，一个类型可以对应多篇文章，所以此处为 fk 关系
	ArticleType *ArticleType `orm:"rel(fk)"`
	// 与用户表建立多对多关系
	Users []*User `orm:"reverse(many)"`
}

// 文章类型表
type ArticleType struct {
	Id       int
	TypeName string     `orm:"size(20)"`
	Article  []*Article `orm:"reverse(many)"`
}

func init() {
	// 连接数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/newsWeb?charset=utf8")
	// 注册表
	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	// 生成表
	orm.RunSyncdb("default", false, true)
}
