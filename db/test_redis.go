package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DB:           1,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}

func TestConnection() {

	fmt.Println(rdb)
	////////////////////////////////////////// 字符串篇 ////////////////////////////////////////////////////////////////

	//client.Set(ctx, "number", "10086", 10000)
	//number := client.Get(ctx, "number")

	//只在键 key 不存在的情况下， 将键 key 的值设置为 value 。	//若键 key 已经存在， 则 SETNX 命令不做任何动作。
	//rdb.SetNX(ctx, "db", "redis", 10 * time.Minute)

	//将键 key 的值设置为 value ， 并将键 key 的生存时间设置为 seconds 秒钟。//如果键 key 已经存在， 那么 SETEX 命令将覆盖已有的值。
	//rdb.SetEX(ctx, "db", "mysql", 10 * time.Minute)

	//rdb.MSet(ctx, "db", "mongodb", "oracle")

	number := rdb.GetSet(ctx, "number", "1024")
	fmt.Println(number)
	//rdb.SetRange(ctx, "number", 5, "hello world")
	//fmt.Println(rdb.GetRange(ctx, "number", 5, 15))
	//fmt.Println(rdb.StrLen(ctx, "number"))
	//append := rdb.Append(ctx, "number", "world")
	//fmt.Println(append)
	//rdb.Set(ctx,"page_view", "20", 7 * 24 * time.Hour)
	//rdb.Incr(ctx, "page_view")
	//fmt.Println(rdb.Get(ctx, "page_view"))
	//rdb.Decr(ctx, "page_view")
	//fmt.Println(rdb.Get(ctx, "page_view"))
	//rdb.IncrBy(ctx, "page_view", 5)
	//fmt.Println(rdb.Get(ctx, "page_view"))
	//rdb.DecrBy(ctx, "page_view", -3)
	//fmt.Println(rdb.Get(ctx, "page_view"))

	//rdb.IncrByFloat(ctx, "page_view", 3.1)
	//fmt.Println(rdb.Get(ctx, "page_view"))

	rdb.IncrByFloat(ctx, "page_view", -3.31)
	fmt.Println(rdb.Get(ctx, "page_view"))

	////////////////////////////////////////// 字符串篇 ////////////////////////////////////////////////////////////////


	_ = rdb.Close()
}
