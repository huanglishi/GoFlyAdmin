package createcode

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
type Codecate struct{}

func init() {
	fpath := Codecate{}
	utils.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取列表
func (api *Codecate) Get_list(c *gin.Context) {
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	name := c.DefaultQuery("name", "")
	page := c.DefaultQuery("page", "1")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(_pageSize)
	MDB := model.DB().Table("business_createcode_cate").Where("businessID", user.BusinessID)
	MDBC := model.DB().Table("business_createcode_cate").Where("businessID", user.BusinessID)
	if name != "" {
		MDB.Where("name", "like", "%"+name+"%")
		MDBC.Where("name", "like", "%"+name+"%")
	}
	list, err := MDB.Limit(pageSize).Page(pageNo).Order("id desc").Get()
	if err != nil {
		results.Failed(c, err.Error(), nil)
	} else {
		rooturl, _ := model.DB().Table("common_config").Where("keyname", "rooturl").Value("keyvalue")
		for _, val := range list {
			if _, ok := val["image"]; ok && val["image"] != "" && !strings.Contains(val["image"].(string), "http") && rooturl != nil {
				val["image"] = rooturl.(string) + val["image"].(string)
			}
		}
		var totalCount int64
		totalCount, _ = MDBC.Count("*")
		results.Success(c, "获取全部列表", map[string]interface{}{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}, nil)
	}
}

// 获取分类列表
func (api *Codecate) Get_cate(c *gin.Context) {
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	list, err := model.DB().Table("business_createcode_cate").Where("businessID", user.BusinessID).Fields("id as value,name as label").Order("weigh desc,id desc").Get()
	if err != nil {
		results.Failed(c, err.Error(), nil)
	} else {
		results.Success(c, "获取选择列表", list, nil)
	}
}

// 保存
func (api *Codecate) Save(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	var f_id float64 = 0
	if parameter["id"] != nil {
		f_id = parameter["id"].(float64)
	}
	if f_id == 0 {
		delete(parameter, "id")
		parameter["createtime"] = time.Now().Unix()
		parameter["businessID"] = user.BusinessID
		addId, err := model.DB().Table("business_createcode_cate").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			if addId != 0 {
				model.DB().Table("business_createcode_cate").
					Data(map[string]interface{}{"weigh": addId}).
					Where("id", addId).
					Update()
			}
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		delete(parameter, "catename")
		res, err := model.DB().Table("business_createcode_cate").
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
func (api *Codecate) UpStatus(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	res2, err := model.DB().Table("business_createcode_cate").Where("id", parameter["id"]).Data(map[string]interface{}{"status": parameter["status"]}).Update()
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
func (api *Codecate) Del(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	res2, err := model.DB().Table("business_createcode_cate").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		results.Success(c, "删除成功！", res2, nil)
	}
}
