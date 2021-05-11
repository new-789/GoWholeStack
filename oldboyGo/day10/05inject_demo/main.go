package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
// sql 注入
var db *sqlx.DB
type user struct {
	Id int
	Age int
	Name string
}
func initDB() (err error) {
	dsn := `root:root@tcp(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=true`
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	return
}

// SQL 注入问题演示 database/sql 库会自动检测 sql 注入问题，发现 sql 注入则跑出 panic 错误
//func sqlInjectDemo(name string) {
//	// 自己拼接 sql 语句的字符串，注意在实际生产环境中永远不要做这样的拼接
//	sqlStr := fmt.Sprintf("select id, name, age from user from name=%s", name)
//	fmt.Printf("sqlStr:%s\n", sqlStr)
//
//	var u user
//	err := db.QueryRow(sqlStr).Scan(&u.Id, &u.Name, &u.Age)
//	if err != nil {
//		fmt.Println("row.Scan failed, err:", err)
//		return
//	}
//	fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
//}

// SQL 注入示例
func sqlInjectDemo(name string) {
	// SQL 注入写法，切记勿用
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL:%s\n", sqlStr)
	var u []user
	err := db.Select(&u, sqlStr)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	for _, v := range u {
		fmt.Printf("%#v\n", v)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("connect database failed, err:",err)
		return
	}
	//sqlInjectDemo("精怪")
	//sqlInjectDemo("xxx' or 1=1 #")
	sqlInjectDemo("xxx' union select * from user #")
}
