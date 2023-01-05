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
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}

func TestConnection() {

	fmt.Println(rdb)

	//client.Set(ctx, "number", "10086", 10000)
	//number := client.Get(ctx, "number")

	//只在键 key 不存在的情况下， 将键 key 的值设置为 value 。	//若键 key 已经存在， 则 SETNX 命令不做任何动作。
	//rdb.SetNX(ctx, "db", "redis", 10 * time.Minute)

	//将键 key 的值设置为 value ， 并将键 key 的生存时间设置为 seconds 秒钟。//如果键 key 已经存在， 那么 SETEX 命令将覆盖已有的值。
	//rdb.SetEX(ctx, "db", "mysql", 10 * time.Minute)

	//rdb.MSet(ctx, "db", "mongodb", "oracle")

	//number := rdb.GetSet(ctx, "number", "1024")
	//fmt.Println(number)
	fmt.Println(rdb.MGet(ctx, "db"))

	_ = rdb.Close()
}

