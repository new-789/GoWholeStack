package db

import (
	"fmt"
	"github.com/GoWholeStack/oldboyGo/book/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// InitDB 初始化数据库连接
func InitDB() (err error) {
	addr := "root:root@tcp(127.0.0.1:3306)/book"
	db, err = sqlx.Connect("mysql", addr)
	if err != nil {
		return err
	}
	// 设置最大连接
	db.SetMaxOpenConns(100)
	// 设置最大空闲
	db.SetMaxIdleConns(16)
	return
}

// QueryAllBook 查询数据库操作
func QueryAllBook() (bookList []*model.Book, err error) {
	sqlStr := `select id, title, price from book`
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println("查询失败", err)
		return
	}
	return
}

// InsertBook 插入数据操作
func InsertBook(title string, price int64) (err error) {
	sqlStr := "insert into book(title, price) values(?,?)"
	_, err = db.Exec(sqlStr,title, price)
	if err != nil {
		fmt.Println("插入数据失败", err)
		return
	}
	return
}

// DeleteBook 删除书籍操作
func DeleteBook(id int64) (err error) {
	sqlStr := "delete from book where id=?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("删除数据失败",err)
		return
	}
	return
}