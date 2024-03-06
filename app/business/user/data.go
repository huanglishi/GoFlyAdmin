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

// 用于自动注册路由
type Data struct {
}

// 初始化生成路由
func init() {
	fpath := Data{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取账号信息
func (api *Data) Get_user(c *gin.Context) {
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	data, err := model.DB().Table("business_account").Where("id", user.ID).Fields("id,status,nickname,name,mobile,email,company,remark,city,area,address,createtime").First()
	if err != nil {
		results.Failed(c, "获取账号信息失败", err)
	} else {
		results.Success(c, "获取账号信息成功！", data, nil)
	}
}

// 更新账号信息
func (api *Data) SaveInfo(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	delete(parameter, "id")
	res, err := model.DB().Table("business_account").
		Data(parameter).
		Where("id", user.ID).
		Update()
	if err != nil {
		results.Failed(c, "更新失败", err)
	} else {
		results.Success(c, "更新成功！", res, nil)
	}
}

// 校验密码
func (api *Data) CheckPassword(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	data, err := model.DB().Table("business_account").Where("id", user.ID).Fields("password,salt").First()
	if err != nil {
		results.Success(c, "账号不存在", false, nil)
	} else {
		if data == nil {
			results.Success(c, "账号不存在", false, nil)
		} else {
			pass := gf.Md5(parameter["password"].(string) + data["salt"].(string))
			if pass != data["password"] {
				results.Success(c, "您输入的密码不正确！", false, nil)
			} else {
				results.Success(c, "密码验证成功", true, nil)
			}
		}
	}
}

// 更新密码
func (api *Data) ChangePassword(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	data, err := model.DB().Table("business_account").Where("id", user.ID).Fields("password,salt").First()
	if err != nil {
		results.Failed(c, "账号不存在", nil)
	} else {
		if data == nil {
			results.Failed(c, "账号不存在", nil)
		} else {
			pass := gf.Md5(parameter["oldpassword"].(string) + data["salt"].(string))
			if pass != data["password"] {
				results.Success(c, "您输入的原来密码不正确！", false, nil)
			} else {
				newpass := gf.Md5(parameter["password"].(string) + data["salt"].(string))
				model.DB().Table("business_account").
					Data(map[string]interface{}{"password": newpass}).
					Where("id", user.ID).
					Update()
				results.Success(c, "密码修改成功!", true, nil)
			}
		}
	}
}
