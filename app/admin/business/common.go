package business

import (
	"strings"

	"gofly/utils/gform"
)

// 获取菜单子树结构
func GetMenuChildrenArray(pdata []gform.Data, parent_id int64) []gform.Data {
	var returnList []gform.Data
	for _, v := range pdata {
		if v["pid"].(int64) == parent_id {
			children := GetMenuChildrenArray(pdata, v["id"].(int64))
			if children != nil {
				v["children"] = children
			}
			returnList = append(returnList, v)
		}
	}
	return returnList
}

// tool-获取树状数组
func GetTreeArray(num []gform.Data, pid int64, itemprefix string) []gform.Data {
	childs := ToolFar(num, pid) //获取pid下的所有数据
	var chridnum []gform.Data
	if childs != nil {
		var number int = 1
		var total int = len(childs)
		for _, v := range childs {
			j := ""
			k := ""
			if number == total {
				j += "└"
				k = ""
				if itemprefix != "" {
					k = "&nbsp;"
				}

			} else {
				j += "├"
				k = ""
				if itemprefix != "" {
					k = "│"
				}
			}
			spacer := ""
			if itemprefix != "" {
				spacer = itemprefix + j
			}
			v["spacer"] = spacer
			v["childlist"] = GetTreeArray(num, v["id"].(int64), itemprefix+k+"&nbsp;")
			chridnum = append(chridnum, v)
			number++
		}
	}
	return chridnum
}

// 2.将getTreeArray的结果返回为二维数组
func getTreeList_txt(data []gform.Data, field string) []gform.Data {
	var midleArr []gform.Data
	for _, v := range data {
		var childlist []gform.Data
		if _, ok := v["childlist"]; ok {
			childlist = v["childlist"].([]gform.Data)
		} else {
			childlist = make([]gform.Data, 0)
		}
		delete(v, "childlist")
		v[field+"_txt"] = v["spacer"].(string) + " " + v[field+""].(string)
		if len(childlist) > 0 {
			v["haschild"] = 1
		} else {
			v["haschild"] = 0
		}
		if _, ok := v["id"]; ok {
			midleArr = append(midleArr, v)
		}
		if len(childlist) > 0 {
			newarr := getTreeList_txt(childlist, field)
			midleArr = ArrayMerge(midleArr, newarr)
		}
	}
	return midleArr
}

// base_tool-获取pid下所有数组
func ToolFar(data []gform.Data, pid int64) []gform.Data {
	var mapString []gform.Data
	for _, v := range data {
		if v["pid"].(int64) == pid {
			mapString = append(mapString, v)
		}
	}
	return mapString
}

// 数组拼接
func ArrayMerge(ss ...[]gform.Data) []gform.Data {
	n := 0
	for _, v := range ss {
		n += len(v)
	}
	s := make([]gform.Data, 0, n)
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
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

// 把字符串打散为数组
func Axplode(data interface{}) []interface{} {
	var rule_ids_arr []interface{}
	ids_arr := strings.Split(data.(string), `,`)
	for _, intv := range ids_arr {
		rule_ids_arr = append(rule_ids_arr, intv)
	}
	return rule_ids_arr
}
