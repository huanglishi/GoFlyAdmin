package user

import (
	"encoding/json"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils/gf"
	"gofly/utils/results"
	"io"
	"reflect"

	"github.com/gin-gonic/gin"
)

func init() {
	gf.Register(&Account{}, reflect.TypeOf(Account{}).PkgPath())
}

// 用于自动注册路由
type Account struct {
}

// 系统设置-个人资料
func (api *Account) Get_userdata(c *gin.Context) {
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	userdata, err := model.DB().Table("admin_account").Fields("id,accountID,username,nickname,avatar,status,password,salt,tel,email,city,remark").Where("id", user.ID).First()
	if err != nil {
		results.Failed(c, " 获取用户数据失败", err)
	} else {
		results.Success(c, " 获取编辑用户数据", userdata, nil)
	}
}

// 系统设置-获取菜单
func (api *Account) Get_menu(c *gin.Context) {
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	//获取用户权限菜单
	role_id, acerr := model.DB().Table("admin_auth_role_access").Where("uid", user.ID).Pluck("role_id")
	if acerr != nil {
		results.Failed(c, "获取role_access失败", acerr)
	}
	if role_id == nil {
		results.Failed(c, "您没有使用权限", nil)
	}
	menu_ids, rerr := model.DB().Table("admin_auth_role").WhereIn("id", role_id.([]interface{})).Pluck("rules")
	if rerr != nil {
		results.Failed(c, "查找auth_role败！", rerr)
	}
	RMDB := model.DB().Table("admin_auth_rule")
	var roles []interface{}
	if !gf.IsContain(menu_ids.([]interface{}), "*") { //不是超级权限-过滤菜单权限
		getmenus := gf.ArrayMerge(menu_ids.([]interface{}))
		RMDB = RMDB.WhereIn("id", getmenus)
		roles = getmenus
	} else {
		roles = make([]interface{}, 0)
	}
	nemu_list, ruleerr := RMDB.Where("status", 0).WhereIn("type", []interface{}{0, 1}).Order("orderNo asc").Get()
	if ruleerr != nil {
		results.Failed(c, "获取菜单错误", ruleerr)
	}
	rulemenu := GetMenuArray(nemu_list, 0, roles)
	results.Success(c, " 获取菜单", rulemenu, nil)
}

// 保存数据
func (api *Account) Upuserinfo(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	res, err := model.DB().Table("admin_account").
		Data(parameter).
		Where("id", user.ID).
		Update()
	if err != nil {
		results.Failed(c, "更新失败", err)
	} else {
		results.Success(c, " 更新用户数据成功", res, nil)
	}
}

// 更新头像
func (api *Account) Upavatar(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	res, err := model.DB().Table("admin_account").Where("id", user.ID).Data(map[string]interface{}{"avatar": parameter["url"]}).Update()
	if err != nil {
		results.Failed(c, "更新头像失败！", err)
	} else {
		results.Success(c, " 更新头像成功", res, nil)
	}
}

// 修改密码
func (api *Account) Changepwd(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	userdata, err := model.DB().Table("admin_account").Where("id", user.ID).Fields("password,salt").First()
	if err != nil {
		results.Failed(c, "查找账号失败！", err)
	} else {
		pass := gf.Md5(parameter["passwordOld"].(string) + userdata["salt"].(string))
		if userdata["password"] != pass {
			results.Failed(c, "原来密码输入错误！", err)
		} else {
			newpass := gf.Md5(parameter["passwordNew"].(string) + userdata["salt"].(string))
			res, err := model.DB().Table("admin_account").
				Data(map[string]interface{}{"password": newpass}).
				Where("id", user.ID).
				Update()
			if err != nil {
				results.Failed(c, "修改密码失败", err)
			} else {
				results.Success(c, "修改密码成功！", res, nil)
			}
		}
	}
}
