package gf

import (
	"math/rand"
)

var (
	chars = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// 获取随机数
func RandString(l int) string {
	bs := []byte{}
	for i := 0; i < l; i++ {
		bs = append(bs, chars[rand.Intn(len(chars))])
	}
	return string(bs)
}
