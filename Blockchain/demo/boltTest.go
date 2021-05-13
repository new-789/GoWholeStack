package main

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	// 1. 连接数据库
	db, err := bolt.Open("./demo/test.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	// 2. 操作数据库(更新）
	db.Update(func(tx *bolt.Tx) error {
		// 打开 bucket(抽屉）, 没有则创建
		bucket := tx.Bucket([]byte("b1"))
		// bucket 等于空则说明没有数据库，则创建
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("b1"))
			if err != nil {
				log.Panic("创建 bucket(b1) 数据库失败")
			}
		}
		// 3. 写数据操作
		bucket.Put([]byte("1111"), []byte("hello"))
		bucket.Put([]byte("2222"), []byte("world"))
		return nil
	})

	// 4. 操作数据库，读数据
	db.View(func(tx *bolt.Tx) error {
		// 打开 bucket (抽屉），如果报错会返回 nil，不为 nil 则表示打开 bucket 成功
		bucket := tx.Bucket([]byte("b1"))
		if bucket != nil {
			// 根据 key 获取对应的数据
			v1 := bucket.Get([]byte("1111"))
			v2 := bucket.Get([]byte("2222"))
			fmt.Printf("v1:%s\n", v1)
			fmt.Printf("v2:%s\n", v2)
		}
		return errors.New("bucket b1 不应该为空，请检查！！！")
	})
}
