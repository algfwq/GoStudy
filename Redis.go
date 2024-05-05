package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址和端口
		Password: "",               // Redis 访问密码，如果没有可以为空字符串
		DB:       0,                // 使用的 Redis 数据库编号，默认为 0
	})

	//设置键值对，为0是永久
	err := rdb.Set(ctx, "name", "John Doe", 10*time.Second).Err()
	if err != nil {
		fmt.Println("Failed to set key:", err)
		return
	}

	// 获取值
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println("Failed to get value:", err)
		return
	}
	fmt.Println("Name:", val)

	// 删除键
	_, err = rdb.Del(ctx, "name").Result()
	if err != nil {
		fmt.Println("Failed to delete key:", err)
		return
	}
}
