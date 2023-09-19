package article

import (
	"gofly/model"
	"gofly/utils"
	"gofly/utils/results"
	"reflect"

	"github.com/gin-gonic/gin"
)

// 用于自动注册路由
type Api struct {
}

func init() {
	fpath := Api{}
	utils.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 测试获取数据接口
func (api *Api) Get_data(c *gin.Context) {
	results.Success(c, "测试获取数据接口", "张三的数据", "扩展数据")
}

// 测试获取列表数据接口
func (api *Api) Get_list(c *gin.Context) {
	list, err := model.DB().Table("business_auth_rule").Fields("id,pid,title,component").Order("id desc").Get()
	results.Success(c, "测试获取列表数据接口", list, err)
}
