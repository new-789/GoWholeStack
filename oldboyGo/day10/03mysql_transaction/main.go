package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// mysql 事务操作
var db *sql.DB
type user struct {
	id int
	age int
	name string
}

func initDB() (err error) {
	dsn := `root:root@tcp(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=true`
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

// Go 语言实现 mysql 事务操作
func transactionDemo() {
	// 1. 开启事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("begin transacation failed, err:%v\n", err)
		return
	}
	// 2. 执行多个 sql 操作
	sqlStr1 := `update user set age=age-2 where id=11`
	sqlStr2 := `update xxx set age=age+3 where id=5`
	// 执行 sql1
	_, err = tx.Exec(sqlStr1)
	if err != nil {
		_ = tx.Rollback() // 2. 出错则执行回滚操作
		fmt.Printf("sqlStr1 exec failed, err:%v\n", err)
		return
	}
	// 执行 sql2
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		_ = tx.Rollback() // 2. 出错则执行回滚操作
		fmt.Printf("sqlStr2 exec failed, err:%v\n", err)
		return
	}
	// 3. 提交事务方法
	err = tx.Commit()
	if err != nil {
		fmt.Printf("tx.Comm failed, err:%v\n", err)
		return
	}

	// 查询操作，此小节重点不在这里在上面的事务操作
	var u user
	qSqlStr := `select id, name, age from user`
	result, err := db.Query(qSqlStr)
	if err != nil {
		fmt.Printf("db.Query failed, er:%v\n", err)
		return
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&u.id,&u.name, &u.age)
		if err != nil {
			fmt.Println("result.Scan failed, err:", err)
			return
		}
		fmt.Println(u.id, u.name, u.age)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("Connect database failed, err:%v\n", err)
		return
	}
	fmt.Println("Connect database success")
	transactionDemo()
}