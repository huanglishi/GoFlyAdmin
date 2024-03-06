package system

import (
	"encoding/json"
	"fmt"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils/gf"
	"gofly/utils/results"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	fpath := Account{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 用于自动注册路由
type Account struct {
}

// 获取成员列表
func (api *Account) Get_list(c *gin.Context) {
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	cid := c.DefaultQuery("cid", "0")
	name := c.DefaultQuery("name", "")
	mobile := c.DefaultQuery("cimobiled", "")
	page := c.DefaultQuery("page", "1")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(_pageSize)
	MDB := model.DB().Table("admin_account").Fields("id,status,name,username,avatar,tel,mobile,email,remark,dept_id,remark,city,address,company,createtime").
		Where("accountID", user.Accountid)
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
			roleid, _ := model.DB().Table("admin_auth_role_access").Where("uid", val["id"]).Pluck("role_id")
			rolename, _ := model.DB().Table("admin_auth_role").WhereIn("id", roleid.([]interface{})).Pluck("name")
			val["rolename"] = rolename
			val["roleid"] = roleid
			depname, _ := model.DB().Table("admin_auth_dept").Where("id", val["dept_id"]).Value("name")
			val["depname"] = depname
			//头像
			if val["avatar"] == nil {
				val["avatar"] = rooturl.(string) + "resource/staticfile/avatar.png"
			} else if !strings.Contains(val["avatar"].(string), "http") && rooturl != nil {
				val["avatar"] = rooturl.(string) + val["avatar"].(string)
			}
		}
		var totalCount int64
		totalCount, _ = model.DB().Table("admin_account").Count()
		results.Success(c, "获取全部列表", map[string]interface{}{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}, nil)
	}
}

// 保存、编辑
func (api *Account) Save(c *gin.Context) {
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
	var roleid []interface{}
	if parameter["roleid"] != nil {
		roleid = parameter["roleid"].([]interface{})
		delete(parameter, "roleid")
	}
	if parameter["password"] != nil && parameter["password"] != "" {
		salt := time.Now().Unix()
		mdpass := fmt.Sprintf("%v%v", parameter["password"], salt)
		parameter["password"] = gf.Md5(mdpass)
		parameter["salt"] = salt
	}
	if parameter["avatar"] == "" {
		parameter["avatar"] = "resource/staticfile/avatar.png"
	}
	if f_id == 0 {
		delete(parameter, "id")
		parameter["createtime"] = time.Now().Unix()
		parameter["uid"] = user.ID
		parameter["accountID"] = user.Accountid
		addId, err := model.DB().Table("admin_account").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			//添加角色-多个
			appRoleAccess(roleid, addId)
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		delete(parameter, "rolename")
		delete(parameter, "depname")
		parameter["updatetime"] = time.Now().Unix()
		res, err := model.DB().Table("admin_account").
			Data(parameter).
			Where("id", f_id).
			Update()
		if err != nil {
			results.Failed(c, "更新失败", err)
		} else {
			//添加角色-多个
			if roleid != nil {
				appRoleAccess(roleid, f_id)
			}
			results.Success(c, "更新成功！", res, nil)
		}
	}
}

// 更新状态
func (api *Account) UpStatus(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	res2, err := model.DB().Table("admin_account").Where("id", parameter["id"]).Data(map[string]interface{}{"status": parameter["status"]}).Update()
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
func (api *Account) Del(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	res2, err := model.DB().Table("admin_account").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		results.Success(c, "删除成功！", res2, nil)
	}
}

// 添加授权
func appRoleAccess(roleids []interface{}, uid interface{}) {
	//批量提交
	model.DB().Table("admin_auth_role_access").Where("uid", uid).Delete()
	save_arr := []map[string]interface{}{}
	for _, val := range roleids {
		marr := map[string]interface{}{"uid": uid, "role_id": val}
		save_arr = append(save_arr, marr)
	}
	model.DB().Table("admin_auth_role_access").Data(save_arr).Insert()
}

// 获取账号信息
func (api *Account) Get_account(c *gin.Context) {
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	data, _ := model.DB().Table("admin_account").Where("id", user.ID).First()
	results.Success(c, "获取账号信息", data, nil)
}

// 表单-选择角色
func (api *Account) Get_role(c *gin.Context) {
	getuser, _ := c.Get("user") //当前用户
	user := getuser.(*middleware.UserClaims)
	role_id, _ := model.DB().Table("admin_auth_role_access").Where("uid", user.ID).Pluck("role_id")
	role_ids := gf.GetAllChilIds("admin_auth_role", role_id.([]interface{})) //批量获取子节点id
	all_role_id := gf.MergeArr(role_id.([]interface{}), role_ids)
	menuList, _ := model.DB().Table("admin_auth_role").WhereIn("id", all_role_id).Where("status", 0).Fields("id ,pid,name").Order("weigh asc").Get()
	results.Success(c, "表单选择角色多选用数据", menuList, nil)
}

// 获取登录日志
func (api *Account) Get_loginloglist(c *gin.Context) {
	//当前用户
	userID := c.DefaultQuery("uid", "0")
	page := c.DefaultQuery("page", "1")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(_pageSize)
	list, _ := model.DB().Table("login_logs").Where("uid", userID).Where("type", 2).Limit(pageSize).Page(pageNo).Order("createtime desc").Get()
	var totalCount int64
	totalCount, _ = model.DB().Table("login_logs").Where("uid", userID).Where("type", 2).Count()
	results.Success(c, "获取登录日志", map[string]interface{}{
		"page":     pageNo,
		"pageSize": pageSize,
		"total":    totalCount,
		"items":    list}, nil)
}

// 判断账号是否存在
func (api *Account) Isaccountexist(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	if parameter["id"] != nil {
		res1, err := model.DB().Table("admin_account").Where("id", "!=", parameter["id"]).Where("username", parameter["username"]).Value("id")
		if err != nil {
			results.Failed(c, "验证失败", err)
		} else if res1 != nil {
			results.Failed(c, "账号已存在", err)
		} else {
			results.Success(c, "验证通过", res1, nil)
		}
	} else {
		res2, err := model.DB().Table("admin_account").Where("username", parameter["username"]).Value("id")
		if err != nil {
			results.Failed(c, "验证失败", err)
		} else if res2 != nil {
			results.Failed(c, "账号已存在", err)
		} else {
			results.Success(c, "验证通过", res2, nil)
		}
	}
}
