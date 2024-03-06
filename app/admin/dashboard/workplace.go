package dashboard

import (
	"fmt"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils/gf"
	"gofly/utils/results"
	"reflect"
	"time"

	"gofly/utils/gform"

	"github.com/gin-gonic/gin"
)

// 用于自动注册路由
type Workplace struct {
}

// 初始化生成路由
func init() {
	fpath := Workplace{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 统计总数据
func (api *Workplace) Get_statistical(c *gin.Context) {
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	allBook, _ := model.DB().Table("doc_book").Where("status", 0).Where("admin_id", user.ID).Count()
	allFile, err := model.DB().Table("doc_file").Where("status", 0).Where("admin_id", user.ID).Count()
	visitNum, _ := model.DB().Table("doc_file").Where("status", 0).Where("admin_id", user.ID).Sum("visit")
	if err != nil {
		results.Failed(c, "统计访问数据失败", err)
	} else {
		results.Success(c, "统计总数据", map[string]interface{}{"totalData": map[string]interface{}{"allBook": allBook, "allFile": allFile, "visitNum": gf.InterfaceToInt(visitNum)}}, err)
	}
}

// 统计访问数据
func (api *Workplace) Get_visit(c *gin.Context) {
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	list, err := model.DB().Table("doc_file").Where("status", 0).Where("admin_id", user.ID).Fields("createtime as x,visit as y").Order("createtime asc").Get()
	if err != nil {
		results.Failed(c, "统计访问数据失败", err)
	} else {
		if list == nil {
			list = make([]gform.Data, 0) //赋空值
		}
		nlist := []map[string]interface{}{}
		for _, val := range list {
			val["x"] = time.Unix(val["x"].(int64), 0).Format("2006-01-02")
			ishase := true
			for _, list := range nlist {
				if list["x"].(string) == val["x"].(string) {
					list["y"] = (list["y"].(int64)) + (val["y"].(int64))
					ishase = false
				}
			}
			if ishase {
				nlist = append(nlist, val)
			}
		}
		results.Success(c, "统计访问数据", nlist, err)
	}
}

// 统计热门文档
func (api *Workplace) Get_popular(c *gin.Context) {
	cid := c.DefaultQuery("cid", "0")
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	MDB := model.DB().Table("doc_file").Where("status", 0).Where("admin_id", user.ID)
	if cid != "0" {
		bookdata, _ := model.DB().Table("doc_book").Where("folder_id", cid).Pluck("id")
		MDB.WhereIn("book_id", bookdata.([]interface{}))
	}
	list, err := MDB.Fields("id,title,visit,createtime").Order("visit desc").Limit(15).Get()
	if err != nil {
		results.Failed(c, "统计热门文档失败", err)
	} else {
		if list == nil {
			list = make([]gform.Data, 0) //赋空值
		}
		countnum, _ := model.DB().Table("doc_file").Where("status", 0).Where("admin_id", user.ID).Sum("visit")
		countnum_int := gf.InterfaceToInt(countnum)
		for key, val := range list {
			visit_int := gf.InterfaceToInt(val["visit"])
			val["increases"] = fmt.Sprintf("%.2f%%", float32(visit_int)/float32(countnum_int)*100)
			val["keys"] = key + 1
		}
		catedata, _ := model.DB().Table("doc_folder").Where("status", 0).Where("admin_id", user.ID).Fields("id,name").Get()
		results.Success(c, "统计热门文档", map[string]interface{}{"list": list, "cate": catedata}, nil)
	}
}
