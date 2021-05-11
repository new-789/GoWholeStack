package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // init()
)

// Go连接 mysql 示例
var db *sql.DB // 是一个连接池对象
// User 用来存储从数据库中查询数来的数据
type User struct {
	id int
	age int
	name string
}

func initDB() (err error) {
	// 数据库连接需要提供的信息
	dsn := "root:root@(127.0.0.1:3306)/gotest?charset=utf8mb4&parseTime=true"
	// 连接数据库
	db, err = sql.Open("mysql", dsn) // Open 不会校验数据库用户名和密码是否正确
	if err != nil {
		return
	}
	err = db.Ping() // 尝试连接数据库,dsn 格式不正确时以及用户名和密码不正确时报错
	if err != nil {
		return
	}
	// 设置数据库连接池的最大连接数
	db.SetMaxOpenConns(10)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(5)
	return
}

// 查询单条记录操作
func queryOne(id int) {
	// 查询单条记录
	var u User
	// 1.  // 查询数据库记录的 SQL 语句, ? 号表示占位符
	sqlStr := "select id, name, age from user where id=?"
	// 2. 执行并拿到结果
	// 2.1:从一个连接池里获取一个连接去数据库查询单条记录,
	// 2.2: 调用 Scan 方法拿到数据，必须的，因为 Scan 会释放连接,若不调用则会占用连接数
	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Println("queryRow data failed, err:", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

// 查询多条记录操作
func queryMany(n int) {
	// 1. 编写SQL查询语句
	sqlStr := "select id, name, age from user where id > ?"
	// 执行查询操作
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("db.Query failed, err:%v\n", err)
		return
	}
	// 3. 多条查询记得一定要调用 Close 方法关闭连接
	defer rows.Close()

	// 2. 循环 rows 获取数据
	for rows.Next() {
		var u User
		// 注意：这里的 Scan 与 单条查询中的 Scan 是两个不同的方法, 多条查询中的 Scan 不带自动关闭连接操作
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("rows.Scan failed, err:%v\n", err)
			return
		}
		fmt.Println(u.id, u.name, u.age)
	}
}

// 插入数据
func insert(name string, age int) {
	// 1. 写 Sql 语句
	sqlStr := "insert into user(name, age) values(?,?)"
	// 2. Exec 执行 Sql 语句
	result, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("db.Exec failed, err:%v\n", err)
		return
	}
	// 如果是插入数据操作，LastInsertId 方法可以获取到插入成功的 id 值
	theId, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get LastInsertId failed, err:%v\n", err)
		return
	}
	//result.RowsAffected() // 可以获取到受影响的行数
	fmt.Printf("insert data id is:%d\n", theId)
}

// 更新数据操作
func updateRow(new, id int) {
	sqlStr := "update user set age=? where id=?"
	ret, err := db.Exec(sqlStr, new, id)
	if err != nil {
		fmt.Printf("update data failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update %d data\n", n)
}

// 删除数据操作
func deleteRow(id int) {
	sqlStr := "delete from user where id=?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete date failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("ret.RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete %v number data\n", n)
}

// Go 实现 SQL 预处理,批量插入多条数据
func prepareInsert() {
	sqlStr := `insert into user(name, age) values(?,?)`
	// 将 sql 语句发送给 mysql 服务器进行预处理
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	// 记得一定要关闭连接
	defer stmt.Close()
	// 后续只需要拿到 stmt 调用 Exec 传值去执行即可
	var m = map[string]int{
		"龙马": 5555,
		"白骨": 4444,
		"灰熊": 3333,
		"观音": 2222,
		"如来": 1111,
	}
	var num int64
	for k, v := range m {
		// stmt.Exec 传值并执行相关操作
		if ret, err := stmt.Exec(k, v); err != nil {
			fmt.Printf("stmt.Exec failed, err:%v\n", err)
			return
		} else {
			n, err := ret.RowsAffected()
			if err != nil {
				fmt.Printf("ret.RowsAffected failed, err:%v\n", err)
				return
			}
			num += n
		}
	}
	fmt.Printf("affected %d data\n", num)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connect database success")
	//queryOne(4)
	//queryMany(0)
	//queryMany(0)
	//insert()
	//updateRow(5555, 5)
	//deleteRow(5)
	prepareInsert()
}
