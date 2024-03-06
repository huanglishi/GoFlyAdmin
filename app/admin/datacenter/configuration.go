package datacenter

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
type Configuration struct{}

func init() {
	fpath := Configuration{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取邮箱
func (api *Configuration) Get_email(c *gin.Context) {
	data, _ := model.DB().Table("common_email").Where("data_from", "common").First()
	results.Success(c, "获取邮箱", data, nil)
}

// 保存邮箱
func (api *Configuration) SaveEmail(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	GetID, _ := model.DB().Table("common_email").Where("data_from", "common").Value("id")
	if GetID == nil {
		parameter["data_from"] = "common"
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
