package goredis

import (
	"time"

	"github.com/go-redis/redis/v8"
)

/**********1.配置**********/
var rclient *redis.Client

func init() {
	InitRedisClient()
}
func InitRedisClient() {
	rclient = redis.NewClient(&redis.Options{
		Addr:        "localhost:6379", // 连接地址
		Password:    "",               // 密码
		DB:          0,                // 数据库编号
		DialTimeout: 1 * time.Second,  // 链接超时
	})
}
func GetRedisClient() *redis.Client {
	return rclient
}
