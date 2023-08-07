package store

import (
	"encoding/json"
	"gofly/app/model"
	"gofly/route/middleware"
	"gofly/utils"
	"gofly/utils/results"
	"io/ioutil"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose/v2"
)

// 用于自动注册路由
type Cate struct {
}

func init() {
	fpath := Cate{}
	utils.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取部门列表
func (api *Cate) Get_list(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	status := c.DefaultQuery("status", "")
	getuser, _ := c.Get("user") //当前用户
	user := getuser.(*middleware.UserClaims)
	MDB := model.DB().Table("business_store_cate").Where("businessID", user.BusinessID)
	if name != "" {
		MDB.Where("name", "like", "%"+name+"%")
	}
	if status != "" {
		MDB.Where("status", status)
	}
	list, _ := MDB.Order("weigh asc").Get()
	if list == nil {
		list = make([]gorose.Data, 0)
	}
	results.Success(c, "获取列表", list, nil)
}

// 获取列表
func (api *Cate) Get_parent(c *gin.Context) {
	typeid := c.DefaultQuery("type", "1")
	getuser, _ := c.Get("user") //当前用户
	user := getuser.(*middleware.UserClaims)
	list, _ := model.DB().Table("business_store_cate").Where("businessID", user.BusinessID).Where("type", typeid).Fields("id as value,name as label").Order("weigh asc").Get()
	if list == nil {
		list = make([]gorose.Data, 0)
	}
	results.Success(c, "获取列表", list, nil)
}

// 保存
func (api *Cate) Save(c *gin.Context) {
	//获取post传过来的data
	body, _ := ioutil.ReadAll(c.Request.Body)
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
		parameter["businessID"] = user.BusinessID
		addId, err := model.DB().Table("business_store_cate").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			if addId != 0 {
				model.DB().Table("business_store_cate").
					Data(map[string]interface{}{"weigh": addId}).
					Where("id", addId).
					Update()
			}
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		res, err := model.DB().Table("business_store_cate").
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
func (api *Cate) UpStatus(c *gin.Context) {
	//获取post传过来的data
	body, _ := ioutil.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	res2, err := model.DB().Table("business_store_cate").Where("id", parameter["id"]).Data(map[string]interface{}{"status": parameter["status"]}).Update()
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
func (api *Cate) Del(c *gin.Context) {
	//获取post传过来的data
	body, _ := ioutil.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	res2, err := model.DB().Table("business_store_cate").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		results.Success(c, "删除成功！", res2, nil)
	}
}
