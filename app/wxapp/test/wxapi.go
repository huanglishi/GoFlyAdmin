package article

import (
	"gofly/utils/gf"
	"gofly/utils/results"
	"reflect"

	"github.com/gin-gonic/gin"
)

// 用于自动注册路由
type wxapi struct {
}

func init() {
	fpath := wxapi{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 测试获取数据接口
func (api *wxapi) Get_data(c *gin.Context) {
	results.Success(c, "测试获取数据接口", "张三的数据", "扩展数据")
}
