package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

// Init 初始化数据库连接
func Init(dns string) (err error) {
	DB, err = sqlx.Open("mysql", dns)
	if err != nil {
		return err
	}
	// 测试查看是否连接成功
	err = DB.Ping()
	if err != nil {
		return err
	}
	// 设置最大连接数和空闲连接数
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}