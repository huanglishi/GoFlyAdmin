package gf

import (
	"reflect"
)

func InArray(array []interface{}, element interface{}) bool {
	// 实现查找整形、string类型和bool类型是否在数组中
	if element == nil || array == nil {
		return false
	}
	for _, value := range array {
		// fmt.Println("比较类型", reflect.TypeOf(value).Kind(), reflect.TypeOf(element).Kind())
		if reflect.TypeOf(value).Kind() == reflect.TypeOf(element).Kind() { // 首先判断类型是否一致
			if value == element { // 比较值是否一致
				return true
			}
		}
	}
	return false
}
