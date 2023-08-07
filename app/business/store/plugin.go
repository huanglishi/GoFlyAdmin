package store

import (
	"encoding/json"
	"gofly/app/model"
	"gofly/route/middleware"
	"gofly/utils"
	"gofly/utils/results"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 用于自动注册路由
type Plugin struct {
}

func init() {
	fpath := Plugin{}
	utils.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取列表
func (api *Plugin) Get_list(c *gin.Context) {
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	cid := c.DefaultQuery("cid", "0")
	name := c.DefaultQuery("name", "")
	page := c.DefaultQuery("page", "1")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(_pageSize)
	MDB := model.DB().Table("business_store_plugin").Fields("id,cid,name,image,des,showHome,showMarket,weigh,status,visit,star,createtime").
		Where("businessID", user.BusinessID)
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
		results.Success(c, "获取全部列表", map[string]interface{}{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}, nil)
	}
}

// 保存
func (api *Plugin) Save(c *gin.Context) {
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
		addId, err := model.DB().Table("business_store_plugin").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			if addId != 0 {
				model.DB().Table("business_store_plugin").
					Data(map[string]interface{}{"weigh": addId}).
					Where("id", addId).
					Update()
			}
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		delete(parameter, "catename")
		res, err := model.DB().Table("business_store_plugin").
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
func (api *Plugin) UpStatus(c *gin.Context) {
	//获取post传过来的data
	body, _ := ioutil.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	res2, err := model.DB().Table("business_store_plugin").Where("id", parameter["id"]).Data(map[string]interface{}{"status": parameter["status"]}).Update()
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
func (api *Plugin) Del(c *gin.Context) {
	//获取post传过来的data
	body, _ := ioutil.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	res2, err := model.DB().Table("business_store_plugin").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		results.Success(c, "删除成功！", res2, nil)
	}
}

// 获取内容
func (api *Plugin) Get_content(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		results.Failed(c, "请传参数id", nil)
	} else {
		res2, err := model.DB().Table("business_store_plugin").Where("id", id).First()
		if err != nil {
			results.Failed(c, "获取内容失败", err)
		} else {
			results.Success(c, "获取内容成功！", res2, nil)
		}
	}

}
