package developer

import (
	"encoding/json"
	"gofly/model"
	"gofly/utils/gf"
	"gofly/utils/results"
	"io"
	"reflect"

	"github.com/gin-gonic/gin"
)

// 用于自动注册路由
type Devapitype struct{}

func init() {
	fpath := Devapitype{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取列表
func (api *Devapitype) Get_list(c *gin.Context) {
	list, err := model.DB().Table("common_apidoc_type").Order("id asc").Get()
	if err != nil {
		results.Failed(c, err.Error(), nil)
	} else {
		results.Success(c, "获取数据列表", list, nil)
	}
}

// 获取单条数据
func (api *Devapitype) Get_typeinfo(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		results.Failed(c, "参数id不能为空", nil)
	} else {
		data, _ := model.DB().Table("common_apidoc_type").Where("id", id).First()
		results.Success(c, "获取单条接口类型", data, nil)
	}
}

// 保存
func (api *Devapitype) Save(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	//当前用户
	var f_id float64 = 0
	if parameter["id"] != nil {
		f_id = parameter["id"].(float64)
	}
	if f_id == 0 {
		delete(parameter, "id")
		addId, err := model.DB().Table("common_apidoc_type").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		res, err := model.DB().Table("common_apidoc_type").
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
func (api *Devapitype) Del(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	res2, err := model.DB().Table("common_apidoc_type").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		results.Success(c, "删除成功！", res2, nil)
	}
}
