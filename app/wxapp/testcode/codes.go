package testcode

import (
	"encoding/json"
	"fmt"
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
type Codes struct {
}

func init() {
	fpath := Codes{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取成员列表
func (api *Codes) Get_list(c *gin.Context) {
	businessID := c.GetHeader("Businessid")
	title := c.DefaultQuery("title", "")
	lastid := c.DefaultQuery("lastid", "")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageSize, _ := strconv.Atoi(_pageSize)
	MDB := model.DB().Table("business_createcode_api").
		Fields("accountID,businessID,id,des").Where("businessID", businessID)
	MDBC := model.DB().Table("business_createcode_api").
		Where("businessID", businessID)
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

// 获取详情
func (api *Codes) Get_detail(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		results.Failed(c, "请传参数id", nil)
	} else {
		data, err := model.DB().Table("business_createcode_api").Where("id", id).
			Fields("accountID,businessID,id,des").First()
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
func (api *Codes) Save(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	//当前用户
	var f_id float64 = 0
	if parameter["id"] != nil {
		f_id = parameter["id"].(float64)
	}
	parameter["createtime"] = time.Now().Unix()
	if f_id == 0 {
		delete(parameter, "id")
		businessID := c.GetHeader("Businessid")
		parameter["businessID"] = businessID
		addId, err := model.DB().Table("business_createcode_api").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			if addId != 0 {
				model.DB().Table("business_createcode_api").
					Data(map[string]interface{}{"weigh": addId}).
					Where("id", addId).
					Update()
			}
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		res, err := model.DB().Table("business_createcode_api").
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

// 删除
func (api *Codes) Del(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	res2, err := model.DB().Table("business_createcode_api").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		results.Success(c, "删除成功！", res2, nil)
	}
}
