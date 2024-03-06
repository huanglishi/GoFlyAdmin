package goredis

import (
	"context"
	"fmt"
	"time"
)

/**********2.常用操作工具类**********/

var client = GetRedisClient()
var ctx = context.Background()

/*------------------------------------ 字符 操作 ------------------------------------*/

// Set 设置 key的值
func Set(key, value string) bool {
	result, err := client.Set(ctx, key, value, 0).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return result == "OK"
}

// SetEX 设置 key的值并指定过期时间
func SetEX(key, value string, ex time.Duration) bool {
	result, err := client.Set(ctx, key, value, ex).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return result == "OK"
}

// Get 获取 key的值
func Get(key string) (bool, string) {
	result, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	return true, result
}

// GetSet 设置新值获取旧值
func GetSet(key, value string) (bool, string) {
	oldValue, err := client.GetSet(ctx, key, value).Result()
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	return true, oldValue
}

// Incr key值每次加一 并返回新值
func Incr(key string) int64 {
	val, err := client.Incr(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val
}

// IncrBy key值每次加指定数值 并返回新值
func IncrBy(key string, incr int64) int64 {
	val, err := client.IncrBy(ctx, key, incr).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val
}

// IncrByFloat key值每次加指定浮点型数值 并返回新值
func IncrByFloat(key string, incrFloat float64) float64 {
	val, err := client.IncrByFloat(ctx, key, incrFloat).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val
}

// Decr key值每次递减 1 并返回新值
func Decr(key string) int64 {
	val, err := client.Decr(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val
}

// DecrBy key值每次递减指定数值 并返回新值
func DecrBy(key string, incr int64) int64 {
	val, err := client.DecrBy(ctx, key, incr).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val
}

// Del 删除 key
func Del(key string) bool {
	result, err := client.Del(ctx, key).Result()
	if err != nil {
		return false
	}
	return result == 1
}

// Expire 设置 key的过期时间
func Expire(key string, ex time.Duration) bool {
	result, err := client.Expire(ctx, key, ex).Result()
	if err != nil {
		return false
	}
	return result
}

/*------------------------------------ list 操作 ------------------------------------*/

// LPush 从列表左边插入数据，并返回列表长度
func LPush(key string, date ...interface{}) int64 {
	result, err := client.LPush(ctx, key, date).Result()
	if err != nil {
		fmt.Println(err)
	}
	return result
}

// RPush 从列表右边插入数据，并返回列表长度
func RPush(key string, date ...interface{}) int64 {
	result, err := client.RPush(ctx, key, date).Result()
	if err != nil {
		fmt.Println(err)
	}
	return result
}

// LPop 从列表左边删除第一个数据，并返回删除的数据
func LPop(key string) (bool, string) {
	val, err := client.LPop(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	return true, val
}

// RPop 从列表右边删除第一个数据，并返回删除的数据
func RPop(key string) (bool, string) {
	val, err := client.RPop(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	return true, val
}

// LIndex 根据索引坐标，查询列表中的数据
func LIndex(key string, index int64) (bool, string) {
	val, err := client.LIndex(ctx, key, index).Result()
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	return true, val
}

// LLen 返回列表长度
func LLen(key string) int64 {
	val, err := client.LLen(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val
}

// LRange 返回列表的一个范围内的数据，也可以返回全部数据
func LRange(key string, start, stop int64) []string {
	vales, err := client.LRange(ctx, key, start, stop).Result()
	if err != nil {
		fmt.Println(err)
	}
	return vales
}

// LRem 从列表左边开始，删除元素data， 如果出现重复元素，仅删除 count次
func LRem(key string, count int64, data interface{}) bool {
	_, err := client.LRem(ctx, key, count, data).Result()
	if err != nil {
		fmt.Println(err)
	}
	return true
}

// LInsert 在列表中 pivot 元素的后面插入 data
func LInsert(key string, pivot int64, data interface{}) bool {
	err := client.LInsert(ctx, key, "after", pivot, data).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

/*------------------------------------ set 操作 ------------------------------------*/

// SAdd 添加元素到集合中
func SAdd(key string, data ...interface{}) bool {
	err := client.SAdd(ctx, key, data).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// SCard 获取集合元素个数
func SCard(key string) int64 {
	size, err := client.SCard(ctx, "key").Result()
	if err != nil {
		fmt.Println(err)
	}
	return size
}

// SIsMember 判断元素是否在集合中
func SIsMember(key string, data interface{}) bool {
	ok, err := client.SIsMember(ctx, key, data).Result()
	if err != nil {
		fmt.Println(err)
	}
	return ok
}

// SMembers 获取集合所有元素
func SMembers(key string) []string {
	es, err := client.SMembers(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return es
}

// SRem 删除 key集合中的 data元素
func SRem(key string, data ...interface{}) bool {
	_, err := client.SRem(ctx, key, data).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// SPopN 随机返回集合中的 count个元素，并且删除这些元素
func SPopN(key string, count int64) []string {
	vales, err := client.SPopN(ctx, key, count).Result()
	if err != nil {
		fmt.Println(err)
	}
	return vales
}

/*------------------------------------ hash 操作 ------------------------------------*/

// HSet 根据 key和 field字段设置，field字段的值
func HSet(key, field, value string) bool {
	err := client.HSet(ctx, key, field, value).Err()
	if err != nil {
		return false
	}
	return true
}

// HGet 根据 key和 field字段，查询field字段的值
func HGet(key, field string) string {
	val, err := client.HGet(ctx, key, field).Result()
	if err != nil {
		fmt.Println(err)
	}
	return val
}

// HMGet 根据key和多个字段名，批量查询多个 hash字段值
func HMGet(key string, fields ...string) []interface{} {
	vales, err := client.HMGet(ctx, key, fields...).Result()
	if err != nil {
		panic(err)
	}
	return vales
}

// HGetAll 根据 key查询所有字段和值
func HGetAll(key string) map[string]string {
	data, err := client.HGetAll(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return data
}

// HKeys 根据 key返回所有字段名
func HKeys(key string) []string {
	fields, err := client.HKeys(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return fields
}

// HLen 根据 key，查询hash的字段数量
func HLen(key string) int64 {
	size, err := client.HLen(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}
	return size
}

// HMSet 根据 key和多个字段名和字段值，批量设置 hash字段值
func HMSet(key string, data map[string]interface{}) bool {
	result, err := client.HMSet(ctx, key, data).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return result
}

// HSetNX 如果 field字段不存在，则设置 hash字段值
func HSetNX(key, field string, value interface{}) bool {
	result, err := client.HSetNX(ctx, key, field, value).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return result
}

// HDel 根据 key和字段名，删除 hash字段，支持批量删除
func HDel(key string, fields ...string) bool {
	_, err := client.HDel(ctx, key, fields...).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// HExists 检测 hash字段名是否存在
func HExists(key, field string) bool {
	result, err := client.HExists(ctx, key, field).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return result
}
