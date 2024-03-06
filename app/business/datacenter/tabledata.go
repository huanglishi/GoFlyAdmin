package datacenter

import (
	"encoding/json"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils/gf"
	"gofly/utils/results"
	"io"
	"reflect"
	"time"

	"gofly/utils/gform"

	"github.com/gin-gonic/gin"
)

// 用于自动注册路由
type Tabledata struct{}

func init() {
	fpath := Tabledata{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取列表
func (api *Tabledata) Get_list(c *gin.Context) {
	list, _ := model.DB().Table("common_dictionary_table").Fields("id,businessID,title,remark,tablename,status,weigh,data_from").Order("weigh asc").Get()
	if list == nil {
		list = make([]gform.Data, 0)
	}
	results.Success(c, "获取列表", list, nil)
}

// 保存
func (api *Tabledata) Save(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	var f_id float64 = 0
	if parameter["id"] != nil {
		f_id = parameter["id"].(float64)
	}
	if f_id == 0 {
		getuser, _ := c.Get("user")
		user := getuser.(*middleware.UserClaims)
		delete(parameter, "id")
		parameter["data_from"] = "business"
		parameter["businessID"] = user.BusinessID
		parameter["createtime"] = time.Now().Unix()
		addId, err := model.DB().Table("common_dictionary_table").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			if addId != 0 {
				model.DB().Table("common_dictionary_table").
					Data(map[string]interface{}{"weigh": addId}).
					Where("id", addId).
					Update()
			}
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		res, err := model.DB().Table("common_dictionary_table").
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
func (api *Tabledata) Del(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	res2, err := model.DB().Table("common_dictionary_table").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		results.Success(c, "删除成功！", res2, nil)
	}
}
