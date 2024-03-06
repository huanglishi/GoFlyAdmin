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
type Configuration struct{}

func init() {
	fpath := Configuration{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取邮箱
func (api *Configuration) Get_email(c *gin.Context) {
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	data, _ := model.DB().Table("common_email").Where("businessID", user.BusinessID).First()
	results.Success(c, "获取邮箱", data, nil)
}

// 保存邮箱
func (api *Configuration) SaveEmail(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	GetID, _ := model.DB().Table("common_email").Where("businessID", user.BusinessID).Value("id")
	if GetID == nil {
		parameter["businessID"] = user.BusinessID
		parameter["data_from"] = "business"
		addId, err := model.DB().Table("common_email").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		res, err := model.DB().Table("common_email").
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
