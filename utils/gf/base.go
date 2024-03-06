package gf

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"gofly/model"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

// 获取post传过来的data
func PostParam(c *gin.Context) (map[string]interface{}, error) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	err := json.Unmarshal(body, &parameter)
	if err != nil {
		return nil, err
	}
	return parameter, nil
}

// 批量获取请求参数-通用
func RequestParam(c *gin.Context) (dataMap map[string]interface{}, err error) {
	c.Request.ParseForm()
	dataMap = make(map[string]interface{})
	if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
		if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json") {
			body, _ := io.ReadAll(c.Request.Body)
			var parameter map[string]interface{}
			err := json.Unmarshal(body, &parameter)
			if err != nil {
				return nil, err
			}
			dataMap = parameter
		} else {
			//说明:须post方法,加: 'Content-Type': 'application/x-www-form-urlencoded'
			for key, valueArray := range c.Request.PostForm {
				if len(valueArray) > 1 {
					errMsg := fmt.Sprintf("#ERROR#[%s]参数设置了[%d]次,只能设置一次.", key, len(valueArray))
					return nil, errors.New(errMsg)
				}
				dataMap[key] = c.PostForm(key)
			}
		}
	}
	for key, _ := range c.Request.URL.Query() {
		dataMap[key] = c.Query(key)
	}
	return
}

// 判断路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 获取ip函数
func GetIp(c *gin.Context) string {
	reqIP := c.Request.Header.Get("X-Forwarded-For")
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}

// 判断元素是否存在数组中
func IsContain(items []interface{}, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 判断元素是否存在数组中-字符串类型
func IsContainStrin(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 多维数组合并
func ArrayMerge(data []interface{}) []interface{} {
	var rule_ids_arr []interface{}
	for _, mainv := range data {
		ids_arr := strings.Split(mainv.(string), `,`)
		for _, intv := range ids_arr {
			rule_ids_arr = append(rule_ids_arr, intv)
		}
	}
	return rule_ids_arr
}

// 多维数组合并
func ArraymoreMerge(data []interface{}) []interface{} {
	var rule_ids_arr []interface{}
	for _, mainv := range data {
		ids_arr := strings.Split(mainv.(string), `,`)
		for _, intv := range ids_arr {
			rule_ids_arr = append(rule_ids_arr, intv)
		}
	}
	return rule_ids_arr
}

// 合并数组-两个数组合并为一个数组
func MergeArr(a, b []interface{}) []interface{} {
	var arr []interface{}
	for _, i := range a {
		arr = append(arr, i)
	}
	for _, j := range b {
		arr = append(arr, j)
	}
	return arr
}

// 转JSON编码为字符串
func JSONMarshalToString(v interface{}) string {
	s, err := jsoniter.MarshalToString(v)
	if err != nil {
		return ""
	}
	return s
}

// 字符串转JSON编码
func StringToJSON(val interface{}) interface{} {
	str := val.(string)
	if strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}") {
		var parameter interface{}
		_ = json.Unmarshal([]byte(str), &parameter)
		return parameter
	} else {
		var parameter []interface{}
		_ = json.Unmarshal([]byte(str), &parameter)
		return parameter
	}
}

// 获取子菜单包含的父级ID-返回全部ID
func GetRulesID(tablename string, field string, menus interface{}) interface{} {
	menus_rang := menus.([]interface{})
	var fnemuid []interface{}
	for _, v := range menus_rang {
		fid := getParentID(tablename, field, v)
		if fid != nil {
			fnemuid = MergeArr(fnemuid, fid)
		}
	}
	r_nemu := MergeArr(menus_rang, fnemuid)
	uni_fnemuid := UniqueArr(r_nemu) //去重
	return uni_fnemuid
}

// 获取所有父级ID
func getParentID(tablename string, field string, id interface{}) []interface{} {
	var pids []interface{}
	pid, _ := model.DB().Table(tablename).Where("id", id).Value(field)
	if pid != nil {
		a_pid := pid.(int64)
		var zr_pid int64 = 0
		if a_pid != zr_pid {
			pids = append(pids, a_pid)
			getParentID(tablename, field, pid)
		}
	}
	return pids
}

// 去重
func UniqueArr(m []interface{}) []interface{} {
	d := make([]interface{}, 0)
	tempMap := make(map[int]bool, len(m))
	for _, v := range m { // 以值作为键名
		keyv := InterfaceToInt(v)
		if tempMap[keyv] == false {
			tempMap[keyv] = true
			d = append(d, v)
		}
	}
	return d
}

// interface{}转int
func InterfaceToInt(data interface{}) int {
	var t2 int
	switch data.(type) {
	case uint:
		t2 = int(data.(uint))
		break
	case int8:
		t2 = int(data.(int8))
		break
	case uint8:
		t2 = int(data.(uint8))
		break
	case int16:
		t2 = int(data.(int16))
		break
	case uint16:
		t2 = int(data.(uint16))
		break
	case int32:
		t2 = int(data.(int32))
		break
	case uint32:
		t2 = int(data.(uint32))
		break
	case int64:
		t2 = int(data.(int64))
		break
	case uint64:
		t2 = int(data.(uint64))
		break
	case float32:
		t2 = int(data.(float32))
		break
	case float64:
		t2 = int(data.(float64))
		break
	case string:
		t2, _ = strconv.Atoi(data.(string))
		break
	default:
		t2 = data.(int)
		break
	}
	return t2
}

// interface{}转int64
func InterfaceToInt64(data interface{}) int64 {
	var t2 int64
	switch data.(type) {
	case uint:
		t2 = int64(data.(uint))
		break
	case int8:
		t2 = int64(data.(int8))
		break
	case uint8:
		t2 = int64(data.(uint8))
		break
	case int16:
		t2 = int64(data.(int16))
		break
	case uint16:
		t2 = int64(data.(uint16))
		break
	case int32:
		t2 = int64(data.(int32))
		break
	case uint32:
		t2 = int64(data.(uint32))
		break
	case int:
		t2 = int64(data.(int))
		break
	case uint64:
		t2 = int64(data.(uint64))
		break
	case float32:
		t2 = int64(data.(float32))
		break
	case float64:
		t2 = int64(data.(float64))
		break
	case string:
		t2, _ = strconv.ParseInt(data.(string), 10, 64)
		break
	default:
		t2 = data.(int64)
		break
	}
	return t2
}

// interface{}float64
func InterfaceFloat64(data interface{}) float64 {
	var f2 float64
	switch data.(type) {
	case string:
		f2, _ = strconv.ParseFloat(data.(string), 64)
		break
	case int:
		f2 = float64(data.(int))
		break
	case float64:
		f2 = data.(float64)
		break
	}
	return f2
}

// interface{}转string
func InterfaceTostring(i interface{}) string {
	if i == nil {
		return ""
	} else {
		return fmt.Sprintf("%v", i)
	}
}

// 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// 字符串首字母小写
func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// 获取随机数
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// md5加密
func Md5Str(origin string) string {
	m := md5.New()
	m.Write([]byte(origin))
	return hex.EncodeToString(m.Sum(nil))
}

// 删除本地附件
func Del_file(file_list []interface{}) {
	for _, val := range file_list {
		dir := fmt.Sprintf("./%v", val)
		os.Remove(dir)
	}
}

// 截取指定字符串中间字符串的方法
func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	} else {
		n = n + 1
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}
