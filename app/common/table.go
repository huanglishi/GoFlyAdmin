package common

import (
	"encoding/json"
	"gofly/model"
	"gofly/utils/gf"
	"gofly/utils/results"
	"io"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Table struct {
}

func init() {
	fpath := Table{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 数据通用排序
func (api *Table) Weigh(context *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(context.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	//排序的数组
	ids := parameter["ids"]
	bids, _ := json.Marshal(&ids)
	var ids_arr []interface{}
	var ids_arr_int []int64
	// 将字符串反解析为数组
	json.Unmarshal([]byte(bids), &ids_arr)
	json.Unmarshal([]byte(bids), &ids_arr_int)
	// ids_arr := strings.Split(ids, `,`)
	//拖动的记录ID
	_changeid := parameter["changeid"].(float64)
	changeid := int64(_changeid)
	// //操作字段
	field := parameter["field"].(string)
	// //操作的数据表
	tablename := parameter["table"].(string)
	//父级id
	_pid := parameter["pid"].(float64)
	pid := int64(_pid)
	// //排序方式
	orderway := parameter["orderway"].(string)
	// //主键id
	prikey := parameter["prikey"].(string)
	// 1.如果设定了pid的值,此时只匹配满足条件的ID,其它忽略
	if _, ok := parameter["pid"]; ok {
		// var hasids []map[string]interface{}
		list_id, _ := model.DB().Table(tablename).WhereIn(prikey, ids_arr).Where("pid", pid).Pluck("id")
		list_int := list_id.([]interface{})
		list_intb := make([]int64, len(list_int))
		for i := range list_int {
			list_intb[i] = list_int[i].(int64)
		}
		ids_arr_int = intersect(list_intb, ids_arr_int)
	}
	winids_base, _ := json.Marshal(&ids_arr_int)
	var winids []interface{}
	_ = json.Unmarshal(winids_base, &winids)
	list, _ := model.DB().Table(tablename).WhereIn(prikey, winids).Fields(prikey + "," + field).Order(field + " " + orderway).Get()
	var sour []int64
	weighdata := make(map[int64]int64)
	for _, v := range list {
		sour = append(sour, v[prikey].(int64))
		weighdata[v[prikey].(int64)] = v[field].(int64)
	}
	position := array_search(changeid, ids_arr_int)
	desc_id := sour[position] //移动到目标的ID值,取出所处改变前位置的值
	// change_id, _ := strconv.ParseInt(changeid, 8, 64) //强转int64
	change_id := changeid
	temp := difference(ids_arr_int, sour)
	for k, v := range temp {
		var offset int64
		if v == change_id {
			offset = desc_id
		} else {
			if change_id == temp[0] {
				nk := k + 1
				if len(temp) > nk {
					offset = temp[nk]
				} else {
					offset = change_id
				}
			} else {
				nk := k - 1
				if nk >= 0 {
					offset = temp[nk]
				} else {
					offset = change_id
				}
			}
		}
		model.DB().Table(tablename).Where(prikey, v).Data(map[string]interface{}{field: weighdata[offset]}).Update()
	}
	results.Success(context, "排序成功！", sour, desc_id)
}

// 函数在数组中搜索某个键值，并返回对应的键名
func array_search(changeid int64, arr []int64) int {
	// changeidint, _ := strconv.ParseInt(changeid, 8, 64)
	for k, v := range arr {
		if v == changeid {
			return k
		}
	}
	return -1
}

// 1.比较两个数组的值，并返回交集;2.返回数组中所有的值（不保留键名）：
func intersect(nums1 []int64, nums2 []int64) []int64 {
	m := make(map[int64]int64)
	var arr []int64
	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {
		times, ok := m[v] //v是nums2中的值,m[v]是map中的值.m[v]==times
		if ok && times > 0 {
			arr = append(arr, v)
			m[v]-- //所有出现的数字都+1,最后要减掉1
		}
	}
	return arr
}

// 2.求差集
func difference(slice1, slice2 []int64) []int64 {
	var arr []int64
	for k, v := range slice1 {
		for key, value := range slice2 {
			if k == key && v != value {
				arr = append(arr, v)
			}
		}
	}
	if len(slice1) > len(slice2) {
		sn := len(slice2)
		n_arr := slice1[sn:]
		arr = ArrayMerge(arr, n_arr)
	}
	return arr
}

// 数组拼接
func ArrayMerge(ss ...[]int64) []int64 {
	n := 0
	for _, v := range ss {
		n += len(v)
	}
	s := make([]int64, 0, n)
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
}
