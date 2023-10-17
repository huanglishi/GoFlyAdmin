package member

import (
	"encoding/json"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils"
	"gofly/utils/results"
	"io"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 用于自动注册路由
type Index struct {
}

func init() {
	fpath := Index{}
	utils.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取成员列表
func (api Index) Get_list(c *gin.Context) {
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	cid := c.DefaultQuery("cid", "0")
	name := c.DefaultQuery("name", "")
	mobile := c.DefaultQuery("cimobiled", "")
	page := c.DefaultQuery("page", "1")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(_pageSize)
	MDB := model.DB().Table("business_member").
		Where("businessID", user.BusinessID)
	if cid != "0" {
		MDB.Where("dept_id", cid)
	}
	if name != "" {
		MDB.Where("name", "like", "%"+name+"%")
	}
	if mobile != "" {
		MDB.Where("mobile", mobile)
	}
	list, err := MDB.Limit(pageSize).Page(pageNo).Order("id asc").Get()
	if err != nil {
		results.Failed(c, err.Error(), nil)
	} else {
		rooturl, _ := model.DB().Table("common_config").Where("keyname", "rooturl").Value("keyvalue")
		for _, val := range list {
			//头像
			if val["avatar"] == nil {
				val["avatar"] = rooturl.(string) + "resource/staticfile/avatar.png"
			} else if !strings.Contains(val["avatar"].(string), "http") && rooturl != nil {
				val["avatar"] = rooturl.(string) + val["avatar"].(string)
			}
		}
		var totalCount int64
		totalCount, _ = model.DB().Table("business_member").Count()
		results.Success(c, "获取全部列表", map[string]interface{}{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}, nil)
	}
}

// 保存、编辑
func (api Index) Save(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	//当前用户
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)

	var f_id float64 = 0
	if parameter["id"] != nil {
		f_id = parameter["id"].(float64)
	}
	if parameter["password"] != nil && parameter["password"] != "" {
		rnd := rand.New(rand.NewSource(6))
		salt := strconv.Itoa(rnd.Int())
		mdpass := parameter["password"].(string) + salt
		parameter["password"] = utils.Md5(mdpass)
		parameter["salt"] = salt
	}
	if parameter["avatar"] == "" {
		parameter["avatar"] = "resource/staticfile/avatar.png"
	}
	if f_id == 0 {
		delete(parameter, "id")
		parameter["createtime"] = time.Now().Unix()
		parameter["accountID"] = user.Accountid
		parameter["businessID"] = user.BusinessID
		addId, err := model.DB().Table("business_member").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		delete(parameter, "rolename")
		delete(parameter, "depname")
		parameter["updatetime"] = time.Now().Unix()
		res, err := model.DB().Table("business_member").
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
func (api Index) UpStatus(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	res2, err := model.DB().Table("business_member").Where("id", parameter["id"]).Data(map[string]interface{}{"status": parameter["status"]}).Update()
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
func (api Index) Del(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	res2, err := model.DB().Table("business_member").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		results.Success(c, "删除成功！", res2, nil)
	}
}
