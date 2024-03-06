package datacenter

import (
	"encoding/json"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils/gf"
	"gofly/utils/results"
	"io"
	"reflect"

	"github.com/gin-gonic/gin"
)

// 用于自动注册路由
type Common_config struct{}

func init() {
	fpath := Common_config{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取配置
func (api *Common_config) Get_config(c *gin.Context) {
	keyname := c.DefaultQuery("keyname", "")
	data, _ := model.DB().Table("common_config").Where("keyname", keyname).First()
	results.Success(c, "获取配置", data, nil)
}

// 保存邮箱
func (api *Common_config) SaveConfig(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	GetID, _ := model.DB().Table("common_config").Where("keyname", parameter["keyname"]).Value("id")
	if GetID == nil {
		parameter["data_from"] = "business"
		parameter["businessID"] = user.BusinessID
		addId, err := model.DB().Table("common_config").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		res, err := model.DB().Table("common_config").
			Data(parameter).
			Where("id", GetID).
			Update()
		if err != nil {
			results.Failed(c, "更新失败", err)
		} else {
			results.Success(c, "更新成功！", res, nil)
		}
	}
}
