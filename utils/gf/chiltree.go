package gf

import "gofly/model"

/**
* 1.批量获取子节点id
* @tablename 数据表名称
  @ids 要获取的id
*/
func GetAllChilIds(tablename string, ids []interface{}) []interface{} {
	var allsubids []interface{}
	for _, id := range ids {
		sub_ids := GetAllChilId(tablename, id)
		allsubids = append(allsubids, sub_ids...)
	}
	return allsubids
}

// 1.2获取所有子级ID
func GetAllChilId(tablename string, id interface{}) []interface{} {
	var subids []interface{}
	sub_ids, _ := model.DB().Table(tablename).Where("pid", id).Pluck("id")
	if len(sub_ids.([]interface{})) > 0 {
		for _, sid := range sub_ids.([]interface{}) {
			subids = append(subids, sid)
			subids = append(subids, GetAllChilId(tablename, sid)...)
		}
	}
	return subids
}
