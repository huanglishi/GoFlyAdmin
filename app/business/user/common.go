package user

import (
	"gofly/model"

	"gofly/utils/gform"
)

// 获取权限菜单
func GetMenuArray(pdata []gform.Data, parent_id int64, roles []interface{}) []gform.Data {
	var returnList []gform.Data
	var one int64 = 1
	for _, v := range pdata {
		if v["pid"].(int64) == parent_id {
			mid_item := map[string]interface{}{
				"path":      v["routePath"],
				"name":      v["routeName"],
				"component": v["component"],
			}
			children := GetMenuArray(pdata, v["id"].(int64), roles)
			if children != nil {
				mid_item["children"] = children
			}
			//1.标题
			var Menu_title interface{}
			if v["locale"] != nil && v["locale"] != "" {
				Menu_title = v["locale"]
			} else {
				Menu_title = v["title"]
			}
			meta := map[string]interface{}{
				"locale": Menu_title,
				"id":     v["id"],
			}
			//2.重定向
			if v["redirect"] != nil && v["redirect"] != "" {
				mid_item["redirect"] = v["redirect"]
			}
			//3.隐藏子菜单
			if v["hideChildrenInMenu"] != nil && v["hideChildrenInMenu"].(int64) == one {
				meta["hideChildrenInMenu"] = true
			}
			//3.图标
			if v["icon"] != nil && v["icon"] != "" {
				meta["icon"] = v["icon"]
			}
			//4.缓存
			if v["keepalive"] != nil && v["keepalive"].(int64) == one {
				// meta["ignoreCache"] = false
			} else {
				meta["ignoreCache"] = true
			}
			//5.隐藏菜单
			if v["hideInMenu"] != nil && v["hideInMenu"].(int64) == one {
				meta["hideInMenu"] = true
			}
			//6.在标签隐藏
			if v["noAffix"] != nil && v["noAffix"].(int64) == one {
				meta["noAffix"] = true
			}
			//7.详情页在本业打开-用于配置详情页时左侧激活的菜单路径
			if v["activeMenu"] != nil && v["activeMenu"] == one {
				meta["activeMenu"] = true
			}
			//8.是否需要登录鉴权
			if v["requiresAuth"] != nil && v["requiresAuth"] == one {
				meta["requiresAuth"] = true
			}
			//9.权限
			if len(roles) == 0 {
				permission, _ := model.DB().Table("business_auth_rule").Where("type", 2).Where("pid", v["id"]).Pluck("permission")
				if permission != nil && len(permission.([]interface{})) > 0 {
					meta["roles"] = permission
				}
			} else {
				permission, _ := model.DB().Table("business_auth_rule").Where("type", 2).Where("pid", v["id"]).WhereIn("id", roles).Pluck("permission")
				if permission != nil && len(permission.([]interface{})) > 0 {
					meta["roles"] = permission
				} else {
					meta["roles"] = [1]string{"*"}
				}
			}
			//赋值
			mid_item["meta"] = meta
			returnList = append(returnList, mid_item)
		}
	}
	return returnList
}
