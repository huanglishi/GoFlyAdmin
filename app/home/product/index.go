package product

import (
	"gofly/app/model"
	"gofly/utils"
	"gofly/utils/results"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

/**
*使用 Index 是省略路径中的index
*本路径为： /admin/user/login -省去了index
 */
func init() {
	utils.Register(&Index{}, reflect.TypeOf(Index{}).PkgPath())
}

type Index struct {
}

// 产品
func (api *Index) Get_list(c *gin.Context) {
	cid := c.DefaultQuery("cid", "0")
	name := c.DefaultQuery("name", "")
	page := c.DefaultQuery("page", "1")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(_pageSize)
	MDB := model.DB().Table("business_store_product").Fields("id,cid,name,image,des,showHome,showMarket,weigh,status,visit,star,createtime")
	if cid != "0" {
		MDB.Where("cid", cid)
	}
	if name != "" {
		MDB.Where("name", "like", "%"+name+"%")
	}
	list, err := MDB.Limit(pageSize).Page(pageNo).Order("id asc").Get()
	if err != nil {
		results.Failed(c, err.Error(), nil)
	} else {
		rooturl, _ := model.DB().Table("common_config").Where("keyname", "rooturl").Value("keyvalue")
		for _, val := range list {
			catename, _ := model.DB().Table("business_store_cate").Where("id", val["cid"]).Value("name")
			val["catename"] = catename
			//头像
			if val["image"] != "" && !strings.Contains(val["image"].(string), "http") && rooturl != nil {
				val["image"] = rooturl.(string) + val["image"].(string)
			}
		}
		var totalCount int64
		totalCount, _ = model.DB().Table("business_user").Count()
		results.Success(c, "获取产品列表", map[string]interface{}{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}, nil)
	}
}
