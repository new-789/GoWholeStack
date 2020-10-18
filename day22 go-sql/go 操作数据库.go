package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 设置连接数据库需要的参数
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/itcast")
	if err != nil {
		fmt.Println("sql.Open failed:", err)
		return
	}
	defer db.Close() // 延迟关闭数据库

	// 连接数据库操作
	er := db.Ping()
	if er != nil {
		fmt.Println("db.Ping failed:", er)
		return
	}

	// 操作一：执行 sql 数据操作语句
	/*
		sql := "insert into stu values(2,'berry');"
		result, err1 := db.Exec(sql)  // 执行 sql 语句
		if err1 != nil {
			fmt.Println("db.Exec failed:", err1)
			return
		}
		n, err2 := result.RowsAffected()  // sql 语句执行成功后获取受影响的记录数
		if err2 != nil {
			fmt.Println("result.RowsAffected failed:", err2)
			return
		}
		fmt.Printf("受影响的记录数是: %d\n", n)
	*/

	// 操作二：执行预处理
	/*
		stu := [][]string{{"5", "张良"}, {"6", "刘伯温"}}             // 定义一个二维切片，用来保存需要往数据库中写入的数据
		stmt, err3 := db.Prepare("insert into stu values(?,?);") // 获取预处理对象 stmt
		if err3 != nil {
			fmt.Println("db.Prepare failed:", err3)
			return
		}
		for _, v := range stu {
			stmt.Exec(v[0], v[1]) // 调用执行预处理语句
		}
	*/

	// 操作三：单行查询
	/*
		var id,name string
		row := db.QueryRow("select * from stu where id=4;")  // 获取一行数据
		row.Scan(&id,&name)  // 将 rows 中的数据，存放到 id 和 name 变量中
		fmt.Println(id, name)
	*/

	// 操作四：多行查询
	rows, err4 := db.Query("select * from stu;") // 获取多行数据
	if err4 != nil {
		fmt.Println("err", err4)
		return
	}
	var id, name string
	/*
		由于 Query 执行查询语句时默认会将光标指向数据表的字段行，rows.Next() 将光标移到数据表中的下一行，
		所以第一步就需要将移动一次光标到字段行下的第一条记录，该方法返回 bool 类型数据，
		所以只要没有取完一致循环
	*/
	for rows.Next() { // 循环显示所有的数据
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}
}
