package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// redis
var rdb *redis.Client

// v8 版本连接 redis 示例
func initRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "", // no password
		DB: 0, // use default DB
		PoolSize: 100,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

// Zset 示例
func redisExample() {
	key := "rank"
	items := []*redis.Z{
		&redis.Z{Score: 90, Member: "php"},
		&redis.Z{Score: 96, Member: "Goland"},
		&redis.Z{Score: 97, Member: "Python"},
		&redis.Z{Score: 99, Member: "Java"},
	}
	// 将元素都追加到 key
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	rdb.ZAdd(ctx, key, items...)

	// 将 Golang 分数加 10
	rdb.ZIncrBy(ctx, key, 10.0, "Goland")

	// 查询分数最高了3个
	zSlice := rdb.ZRevRangeWithScores(ctx, key,0, 2)
	for _, v := range zSlice.Val() {
		fmt.Println(v.Member, v.Score)
	}
	fmt.Println("========================")
	// 查询 95~100分的数据
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret := rdb.ZRangeByScoreWithScores(ctx, key, &op)
	for _, v := range ret.Val() {
		fmt.Println(v.Member, v.Score)
	}
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Println("conect redis failed, err:",err)
		return
	}
	fmt.Println("connect redis success")
	redisExample()
}
