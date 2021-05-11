package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

// es insert data Demo
type student struct {
	Name string  `json:"name"`
    Age int 	 `json:"age"`
	Married bool `json:"married"`
}

func main() {
	// 初始化连接 elasticsearch 创建一个 client 客户端
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Printf("connect to elasticsearch failed, err:%v\n", err)
		return
	}

	data := student{Name: "guanyin", Age: 2000, Married: false}
	// 执行链式操作，开始执行插入数据操作
	put, err := client.Index().Index("student").BodyJson(data).Do(context.Background())
	if err != nil {
		fmt.Printf("set data to elastic failed, err:%v\n", err)
		return
	}
	fmt.Printf("index student:%s to index:%s, type:%v\n", put.Id, put.Index, put.Type)
}
