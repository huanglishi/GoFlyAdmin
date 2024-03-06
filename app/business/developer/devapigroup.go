package developer

import (
	"encoding/json"
	"gofly/model"
	"gofly/utils/gf"
	"gofly/utils/results"
	"io"
	"reflect"

	"gofly/utils/gform"

	"github.com/gin-gonic/gin"
)

/**
* 接口分类
 */
// 用于自动注册路由
type Devapigroup struct{}

func init() {
	fpath := Devapigroup{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取列表
func (api *Devapigroup) Get_list(c *gin.Context) {
	list, err := model.DB().Table("common_apidoc_group").Order("id asc").Get()
	if err != nil {
		results.Failed(c, err.Error(), nil)
	} else {
		for _, val := range list {
			typename, _ := model.DB().Table("common_apidoc_type").Where("id", val["type_id"]).Value("name")
			val["typename"] = typename
		}
		results.Success(c, "获取数据列表", list, nil)
	}
}

// 获取父级数据
func (api *Devapigroup) Get_parent(c *gin.Context) {
	list, _ := model.DB().Table("common_apidoc_group").Fields("id,pid,name").Order("id asc").Get()
	list = gf.GetMenuChildrenArray(list, 0, "pid")
	if list == nil {
		list = make([]gform.Data, 0)
	}
	results.Success(c, "获取分组列表", list, nil)
}

// 保存
func (api *Devapigroup) Save(c *gin.Context) {
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
		addId, err := model.DB().Table("common_apidoc_group").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		res, err := model.DB().Table("common_apidoc_group").
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

// 更新状态
func (api *Devapigroup) UpStatus(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	res2, err := model.DB().Table("common_apidoc_group").Where("id", parameter["id"]).Data(map[string]interface{}{"status": parameter["status"]}).Update()
	if err != nil {
		results.Failed(c, "更新失败！", err)
	} else {
		msg := "更新成功！"
		if res2 == 0 {
			msg = "暂无数据更新"
		}
		results.Success(c, msg, res2, nil)
	}
}

// 删除
func (api *Devapigroup) Del(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	res2, err := model.DB().Table("common_apidoc_group").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		results.Success(c, "删除成功！", res2, nil)
	}
}
