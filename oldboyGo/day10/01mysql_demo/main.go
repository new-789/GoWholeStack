package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
)

// Go连接 mysql 示例

func main() {
	// 数据库连接需要提供的信息
	dsn := "root:root@(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=true"
	// 连接数据库
	db, err := sql.Open("mysql", dsn) // Open 不会校验数据库用户名和密码是否正确
	if err != nil {
		fmt.Printf("open %s failed, err:%v\n",dsn, err)
		return
	}
	defer db.Close()
	err = db.Ping() // 尝试连接数据库,dsn 格式不正确时以及用户名和密码不正确时报错
	if err != nil {
		fmt.Println("connect database failed, err:", err)
		return
	}
	fmt.Println("connect database success")
}
