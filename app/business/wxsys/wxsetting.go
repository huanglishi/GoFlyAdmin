package wxsys

import (
	"encoding/json"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils"
	"gofly/utils/results"
	"io/ioutil"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

// 用于自动注册路由
type Wxsetting struct {
}

func init() {
	fpath := Wxsetting{}
	utils.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取内容
func (api *Wxsetting) Get_account(c *gin.Context) {
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	data, err := model.DB().Table("business_wxsys_officonfig").Where("businessID", user.BusinessID).First()
	if err != nil {
		results.Failed(c, "获取内容失败", err)
	} else {
		results.Success(c, "获取内容成功！", data, nil)
	}
}

// 保存
func (api *Wxsetting) Save_account(c *gin.Context) {
	//获取post传过来的data
	body, _ := ioutil.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	//当前用户
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	var f_id float64 = 0
	if parameter["id"] != nil {
		f_id = parameter["id"].(float64)
	}
	parameter["createtime"] = time.Now().Unix()
	if f_id == 0 {
		delete(parameter, "id")
		parameter["accountID"] = user.Accountid
		parameter["businessID"] = user.BusinessID
		addId, err := model.DB().Table("business_wxsys_officonfig").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		res, err := model.DB().Table("business_wxsys_officonfig").
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

/**********************小程序***********************/
// 获取内容
func (api *Wxsetting) Get_wxapp(c *gin.Context) {
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	data, err := model.DB().Table("business_wxsys_wxappconfig").Where("businessID", user.BusinessID).First()
	if err != nil {
		results.Failed(c, "获取内容失败", err)
	} else {
		results.Success(c, "获取内容成功！", data, nil)
	}
}

// 保存
func (api *Wxsetting) Save_wxapp(c *gin.Context) {
	//获取post传过来的data
	body, _ := ioutil.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	//当前用户
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	var f_id float64 = 0
	if parameter["id"] != nil {
		f_id = parameter["id"].(float64)
	}
	parameter["createtime"] = time.Now().Unix()
	if f_id == 0 {
		delete(parameter, "id")
		parameter["accountID"] = user.Accountid
		parameter["businessID"] = user.BusinessID
		addId, err := model.DB().Table("business_wxsys_wxappconfig").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		res, err := model.DB().Table("business_wxsys_wxappconfig").
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
