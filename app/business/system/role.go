package system

import (
	"encoding/json"
	"fmt"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils"
	"gofly/utils/results"
	"io"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose/v2"
)

// 用于自动注册路由
type Role struct {
}

// 初始化生成路由
func init() {
	fpath := Role{}
	utils.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取数据列表-子树结构
func (api *Role) Get_list(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	status := c.DefaultQuery("status", "")
	getuser, _ := c.Get("user") //当前用户
	user := getuser.(*middleware.UserClaims)
	role_id, _ := model.DB().Table("business_auth_role_access").Where("uid", user.ID).Pluck("role_id")
	role_ids := utils.GetAllChilIds("business_auth_role", role_id.([]interface{})) //批量获取子节点id
	all_role_id := utils.MergeArr(role_id.([]interface{}), role_ids)
	//查找条件
	MDB := model.DB().Table("business_auth_role")
	roleList, _ := MDB.Where(func() {
		MDB.WhereIn("id", all_role_id)
		if name != "" {
			MDB.Where("name", "like", "%"+name+"%")
		}
		if status != "" {
			MDB.Where("status", status)
		}
	}).OrWhere(func() {
		MDB.Where("businessID", user.BusinessID)
		if name != "" {
			MDB.Where("name", "like", "%"+name+"%")
		}
		if status != "" {
			MDB.Where("status", status)
		}
	}).Order("weigh asc").Get()
	roleList = utils.GetTreeArray(roleList, 0, "")
	if roleList == nil {
		roleList = make([]gorose.Data, 0)
	}
	results.Success(c, "获取拥有角色列表", roleList, all_role_id)
}

// 表单获取选择父级
func (api *Role) Get_parent(c *gin.Context) {
	getuser, _ := c.Get("user") //当前用户
	user := getuser.(*middleware.UserClaims)
	role_id, _ := model.DB().Table("business_auth_role_access").Where("uid", user.ID).Pluck("role_id")
	role_ids := utils.GetAllChilIds("business_auth_role", role_id.([]interface{})) //批量获取子节点id
	all_role_id := utils.MergeArr(role_id.([]interface{}), role_ids)
	menuList, _ := model.DB().Table("business_auth_role").WhereIn("id", all_role_id).OrWhere("businessID", user.BusinessID).Fields("id,pid,name").Order("weigh asc").Get()
	menuList = utils.GetMenuChildrenArray(menuList, 0, "pid")
	results.Success(c, "部门父级数据！", menuList, nil)
}

// 表单获取菜单
func (api *Role) Get_menuList(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	pid := c.DefaultQuery("pid", "0")
	MDB := model.DB().Table("business_auth_rule").Where("status", 0)
	if id == "0" || pid == "0" { //获取本账号所拥有的权限
		//账号信息
		getuser, _ := c.Get("user") //当前用户
		user := getuser.(*middleware.UserClaims)
		role_id, _ := model.DB().Table("business_auth_role_access").Where("uid", user.ID).Pluck("role_id")
		menu_id, _ := model.DB().Table("business_auth_role").WhereIn("id", role_id.([]interface{})).Pluck("rules")
		if !IsContain(menu_id.([]interface{}), "*") { //不是超级权限-过滤菜单权限
			getmenus := ArraymoreMerge(menu_id.([]interface{}))
			MDB = MDB.WhereIn("id", getmenus)
		}
	} else {
		//获取用户权限
		pid, _ := model.DB().Table("business_auth_role").Where("id", id).Value("pid") //获取父级权限
		menu_id_str, _ := model.DB().Table("business_auth_role").Where("id", pid).Value("rules")
		if !strings.Contains(menu_id_str.(string), "*") { //不是超级权限-过滤菜单权限
			getmenus := Axplode(menu_id_str)
			MDB = MDB.WhereIn("id", getmenus)
		}
	}
	menuList, _ := MDB.Fields("id,pid,title,locale").Order("orderNo asc").Get()
	for _, val := range menuList {
		if val["title"] == "" {
			val["title"] = val["locale"]
		}
		delete(val, "locale")
	}
	menuList = GetMenuChildrenArray(menuList, 0)
	results.Success(c, "获取菜单数据", menuList, nil)
}

// 保存编辑
func (api *Role) Save(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
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
	if parameter["menu"] != nil && parameter["menu"] != "*" {
		rules := utils.GetRulesID("business_auth_role", "pid", parameter["menu"]) //获取子菜单包含的父级ID
		rudata := rules.([]interface{})
		var rulesStr []string
		for _, v := range rudata {
			str := fmt.Sprintf("%v", v) //interface{}强转string
			rulesStr = append(rulesStr, str)
		}
		parameter["rules"] = strings.Join(rulesStr, ",")
		parameter["menu"] = utils.JSONMarshalToString(parameter["menu"])
	}
	parameter["createtime"] = time.Now().Unix()
	if f_id == 0 {
		delete(parameter, "id")
		parameter["businessID"] = user.BusinessID
		addId, err := model.DB().Table("business_auth_role").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			if addId != 0 {
				model.DB().Table("business_auth_role").
					Data(map[string]interface{}{"weigh": addId}).
					Where("id", addId).
					Update()
			}
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		delete(parameter, "children")
		delete(parameter, "spacer")
		res, err := model.DB().Table("business_auth_role").
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
func (api *Role) UpStatus(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	res2, err := model.DB().Table("business_auth_role").Where("id", parameter["id"]).Data(map[string]interface{}{"status": parameter["status"]}).Update()
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
func (api *Role) Del(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	res2, err := model.DB().Table("business_auth_role").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		results.Success(c, "删除成功！", res2, nil)
	}
}
