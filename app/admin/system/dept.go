package system

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
type Dept struct {
}

func init() {
	fpath := Dept{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取部门列表
func (api *Dept) Get_list(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	status := c.DefaultQuery("status", "")
	MDB := model.DB().Table("admin_auth_dept")
	if name != "" {
		MDB.Where("name", "like", "%"+name+"%")
	}
	if status != "" {
		MDB.Where("status", status)
	}
	list, _ := MDB.Order("weigh asc").Get()
	if list == nil {
		list = make([]gform.Data, 0)
	} else {
		list = gf.GetTreeArray(list, 0, "")
	}
	// list = gf.GetMenuChildrenArray(list, 0, "pid")
	results.Success(c, "获取部门列表", list, nil)
}

// 获取部门列表-表单
func (api *Dept) Get_parent(c *gin.Context) {
	list, _ := model.DB().Table("admin_auth_dept").Fields("id,pid,name").Order("weigh asc").Get()
	list = gf.GetMenuChildrenArray(list, 0, "pid")
	if list == nil {
		list = make([]gform.Data, 0)
	}
	results.Success(c, "获取部门列表", list, nil)
}

// 保存
func (api *Dept) Save(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	//当前用户
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	parameter["uid"] = user.ID
	var f_id float64 = 0
	if parameter["id"] != nil {
		f_id = parameter["id"].(float64)
	}
	parameter["createtime"] = time.Now().Unix()
	if f_id == 0 {
		delete(parameter, "id")
		addId, err := model.DB().Table("admin_auth_dept").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			if addId != 0 {
				model.DB().Table("admin_auth_dept").
					Data(map[string]interface{}{"weigh": addId}).
					Where("id", addId).
					Update()
			}
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		res, err := model.DB().Table("admin_auth_dept").
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
func (api *Dept) UpStatus(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	res2, err := model.DB().Table("admin_auth_dept").Where("id", parameter["id"]).Data(map[string]interface{}{"status": parameter["status"]}).Update()
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

// 更新父级-拖拽更新父id
func (api *Dept) Upgrouppid(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	b_ids, _ := json.Marshal(parameter["ids"])
	var ids_arr []interface{}
	json.Unmarshal([]byte(b_ids), &ids_arr)
	res2, err := model.DB().Table("admin_auth_dept").WhereIn("id", ids_arr).Data(map[string]interface{}{"pid": parameter["pid"]}).Update()
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
func (api *Dept) Del(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	res2, err := model.DB().Table("admin_auth_dept").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		results.Success(c, "删除成功！", res2, nil)
	}
}
