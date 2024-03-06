package common

import (
	"encoding/json"
	"fmt"
	"gofly/global"
	"gofly/model"
	"gofly/utils/gf"
	"gofly/utils/results"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 用于自动注册路由
type Api struct {
}

func init() {
	fpath := Api{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取成员列表
func (api *Api) Get_list(c *gin.Context) {
	businessID := c.GetHeader("Businessid")
	apiID := c.DefaultQuery("apiID", "")
	title := c.DefaultQuery("title", "")
	lastid := c.DefaultQuery("lastid", "")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageSize, _ := strconv.Atoi(_pageSize)
	if apiID == "" {
		results.Failed(c, "请传参数apiID", nil)
	} else {
		//获取接口数据信息
		apitext, _ := model.DB().Table("common_apidoc").Where("id", apiID).Fields("tablename,fields").First()
		MDB := model.DB().Table(apitext["tablename"]).Fields(apitext["fields"].(string)).Where("businessID", businessID)
		MDBC := model.DB().Table(apitext["tablename"]).Where("businessID", businessID)
		if title != "" {
			MDB.Where("title", "like", "%"+title+"%")
			MDBC.Where("title", "like", "%"+title+"%")
		}
		if lastid != "" {
			MDB.Where("id", "<", lastid)
		}
		list, err := MDB.Limit(pageSize).Order("id desc").Get()
		if err != nil {
			results.Failed(c, err.Error(), nil)
		} else {
			var new_lastid interface{}
			rooturl, _ := model.DB().Table("common_config").Where("keyname", "rooturl").Value("keyvalue")
			for key, val := range list {
				if _, ok := val["image"]; ok && !strings.Contains(val["image"].(string), "http") && rooturl != nil {
					val["image"] = rooturl.(string) + val["image"].(string)
				}
				fmt.Println(key)
				if (key + 1) == len(list) {
					new_lastid = val["id"]
				}
			}
			var totalCount int64
			totalCount, _ = MDBC.Count()
			results.Success(c, "获取全部列表", map[string]interface{}{
				"lastid":   new_lastid,
				"pageSize": pageSize,
				"total":    totalCount,
				"items":    list}, nil)
		}
	}
}

// 获取详情
func (api *Api) Get_detail(c *gin.Context) {
	apiID := c.DefaultQuery("apiID", "")
	id := c.DefaultQuery("id", "")
	if apiID == "" {
		results.Failed(c, "请传参数apiID", nil)
	} else if id == "" {
		results.Failed(c, "请传参数id", nil)
	} else {
		//获取接口数据信息
		apitext, _ := model.DB().Table("common_apidoc").Where("id", apiID).Fields("tablename,fields").First()
		data, err := model.DB().Table(apitext["tablename"]).Where("id", id).Fields(apitext["fields"].(string)).First()
		if err != nil {
			results.Failed(c, "获取详情失败", err)
		} else {
			if data != nil {
				rooturl, _ := model.DB().Table("common_config").Where("keyname", "rooturl").Value("keyvalue")
				if _, ok := data["image"]; ok && !strings.Contains(data["image"].(string), "http") && rooturl != nil {
					data["image"] = rooturl.(string) + data["image"].(string)
				}
			}
			results.Success(c, "获取详情成功！", data, nil)
		}
	}
}

// 保存
func (api *Api) Save(c *gin.Context) {
	apiID := c.DefaultQuery("apiID", "")
	if apiID == "" {
		results.Failed(c, "请传参数apiID", nil)
	} else {
		body, _ := io.ReadAll(c.Request.Body)
		var parameter map[string]interface{}
		_ = json.Unmarshal(body, &parameter)
		var f_id float64 = 0
		if parameter["id"] != nil {
			f_id = parameter["id"].(float64)
		}
		apitext, _ := model.DB().Table("common_apidoc").Where("id", apiID).Fields("tablename").First()
		if f_id == 0 {
			delete(parameter, "id")
			if IsHaseField(apitext["tablename"].(string), "createtime") {
				parameter["createtime"] = time.Now().Unix()
			}
			if IsHaseField(apitext["tablename"].(string), "businessID") {
				parameter["businessID"] = c.GetHeader("Businessid")
			}
			addId, err := model.DB().Table(apitext["tablename"]).Data(parameter).InsertGetId()
			if err != nil {
				results.Failed(c, "添加失败", err)
			} else {
				if addId != 0 {
					model.DB().Table(apitext["tablename"]).
						Data(map[string]interface{}{"weigh": addId}).
						Where("id", addId).
						Update()
				}
				results.Success(c, "添加成功！", addId, nil)
			}
		} else {
			res, err := model.DB().Table(apitext["tablename"]).
				Data(parameter).
				Where("id", f_id).
				Update()
			if err != nil {
				results.Failed(c, "更新失败", err)
			} else {
				results.Success(c, "更新成功！", res, nil)
			}
		}
	}
}

// 删除
func (api *Api) Del(c *gin.Context) {
	apiID := c.DefaultQuery("apiID", "")
	if apiID == "" {
		results.Failed(c, "请传参数apiID", nil)
	} else {
		apitext, _ := model.DB().Table("common_apidoc").Where("id", apiID).Fields("tablename").First()
		body, _ := io.ReadAll(c.Request.Body)
		var parameter map[string]interface{}
		_ = json.Unmarshal(body, &parameter)
		ids := parameter["ids"]
		res2, err := model.DB().Table(apitext["tablename"]).WhereIn("id", ids.([]interface{})).Delete()
		if err != nil {
			results.Failed(c, "删除失败", err)
		} else {
			results.Success(c, "删除成功！", res2, nil)
		}
	}
}

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
	return gf.IsContain(tablefields, fields)
}
