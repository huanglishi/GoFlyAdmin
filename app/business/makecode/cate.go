package makecode

import (
	"encoding/json"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils/gf"
	"gofly/utils/results"
	"io"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 用于自动注册路由
type Cate struct{}

func init() {
	fpath := Cate{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取列表
func (api *Cate) Get_list(c *gin.Context) {
	//请求参数
	params, _ := gf.RequestParam(c)
	MDB := model.DB().Table("createcode_product_cate").
		Fields("createtime,id,name")
	MDBC := model.DB().Table("createcode_product_cate")
	if gf.IsHaseField("createcode_product_cate", "businessID") {
		getuser, _ := c.Get("user")
		user := getuser.(*middleware.UserClaims)
		MDB.Where("businessID", user.BusinessID)
		MDBC.Where("businessID", user.BusinessID)
	}
	if valname, ok := params["name"]; ok && valname != "" {
		MDB.Where("name", "like", "%"+gf.InterfaceTostring(valname)+"%")
		MDBC.Where("name", "like", "%"+gf.InterfaceTostring(valname)+"%")
	}
	if valstatus, ok := params["status"]; ok && valstatus != "" {
		MDB.Where("status", valstatus)
		MDBC.Where("status", valstatus)
	}
	if valTime, ok := params["createdTime"]; ok && valTime != "" {
		datetime_arr := strings.Split(gf.InterfaceTostring(valTime), ",")
		star_time := gf.StringTimestamp(datetime_arr[0]+" 00:00", "datetime")
		end_time := gf.StringTimestamp(datetime_arr[1]+" 23:59", "datetime")
		MDB.WhereBetween("createtime", []interface{}{star_time, end_time})
		MDBC.WhereBetween("createtime", []interface{}{star_time, end_time})
	}
	list, err := MDB.Limit(gf.InterfaceToInt(params["pageSize"])).Page(gf.InterfaceToInt(params["page"])).Order("id desc").Get()
	if err != nil {
		results.Failed(c, err.Error(), nil)
	} else {
		rooturl, _ := model.DB().Table("common_config").Where("keyname", "rooturl").Value("keyvalue")
		for _, val := range list {
			if _, ok := val["image"]; ok && val["image"] != "" && !strings.Contains(val["image"].(string), "http") && rooturl != nil {
				val["image"] = rooturl.(string) + val["image"].(string)
			}
			if _, ok := val["cid"]; ok {
				catename, _ := model.DB().Table("createcode_product_cate").Where("id", val["cid"]).Value("name")
				val["catename"] = catename
			}
		}
		var totalCount int64
		totalCount, _ = MDBC.Count("*")
		results.Success(c, "获取全部列表", map[string]interface{}{
			"page":     params["page"],
			"pageSize": params["pageSize"],
			"total":    totalCount,
			"items":    list}, nil)
	}
}

// 保存
func (api *Cate) Save(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
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
		if gf.IsHaseField("createcode_product_cate", "createtime") {
			parameter["createtime"] = time.Now().Unix()
		}
		if gf.IsHaseField("createcode_product_cate", "businessID") {
			parameter["businessID"] = user.BusinessID
		}
		addId, err := model.DB().Table("createcode_product_cate").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			if addId != 0 {
				model.DB().Table("createcode_product_cate").
					Data(map[string]interface{}{"weigh": addId}).
					Where("id", addId).
					Update()
			}
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		delete(parameter, "catename")
		res, err := model.DB().Table("createcode_product_cate").
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
func (api *Cate) UpStatus(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	res2, err := model.DB().Table("createcode_product_cate").Where("id", parameter["id"]).Data(map[string]interface{}{"status": parameter["status"]}).Update()
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
func (api *Cate) Del(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	res2, err := model.DB().Table("createcode_product_cate").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		results.Success(c, "删除成功！", res2, nil)
	}
}
