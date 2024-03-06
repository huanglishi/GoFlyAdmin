package gf

import (
	"encoding/json"
	"gofly/global"
	"gofly/model"
)

// 判断某个数据表是否存在指定字段
// tablename=表名 field=字段
func IsHaseField(tablename, fields string) bool {
	//获取数据库名
	dielddata, _ := model.DB().Query("select COLUMN_NAME from information_schema.columns where TABLE_SCHEMA='" + global.App.Config.DBconf.Database + "' AND TABLE_NAME='" + tablename + "'")
	var tablefields []interface{}
	for _, val := range dielddata {
		var valjson map[string]interface{}
		mdata, _ := json.Marshal(val)
		json.Unmarshal(mdata, &valjson)
		tablefields = append(tablefields, valjson["COLUMN_NAME"].(string))
	}
	return IsContain(tablefields, fields)
}
