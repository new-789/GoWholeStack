package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// sqlx_demo
var db *sqlx.DB  // 注意标准库中的 sql.DB 的区别，这里的 sqlx.DB
type user struct {
	// 使用 sqlx 包，定义的结构体用来存储获取到的数据的字段必须大写开头
	Id int
	Name string
	Age int
}

// 连接数据了
func initDB() (err error) {
	dsn := `root:root@tcp(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=true`
	// 使用 sqlx 连接数据库, Connect 会帮助我们连接数据库，无需像标准库中的 sql 还需要调用以下 Ping 方法；
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	// 设置最大连接数
	db.SetMaxOpenConns(10)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(5)
	return
}

// 使用 sqlx 查询单条数据，使用 Get 方法
func queryOne() {
	sqlStr := `select id, name, age from user where id=1`
	var u user
	// 执行查询操作，注：如没有使用参数，第三个参数不用传参，不要自作聪明的传个 0 或者 nil 会导致查询不到数据
	db.Get(&u, sqlStr)
	fmt.Printf("u:%#v\n", u)
}

// 使用 sqlx 库查询多条数据，查询多条使用 Select 方法
func queryMany() {
	sqlStr := `select id, name, age from user`
	var userList =make([]user, 1) // 定义一个 slice 变量用来保存查询到的多条数据
	// 查询多条记录使用 Select，注意：查询多条数据时，第一个参数传递的切片因为必须的指针类型
	db.Select(&userList, sqlStr)
	fmt.Printf("userlist:%#v\n", userList)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("connect database failed, err:%v\n", err)
		return
	}
	fmt.Println("connect database success...!")
	queryOne()
	queryMany()
}
