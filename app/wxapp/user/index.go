package user

import (
	"encoding/json"
	"fmt"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils/gf"
	"gofly/utils/results"
	"io"
	"reflect"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

/**
*使用 Index 是省略路径中的index
*本路径为： /admin/user/login -省去了index
 */
type Index struct{}

func init() {
	fpath := Index{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

/**
*1.《登录》
 */
func (api *Index) Get_openid(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	businessID := c.GetHeader("Businessid")
	account, _ := model.DB().Table("business_wxsys_wxappconfig").Where("businessID", businessID).Fields("id,accountID,businessID,AppID,AppSecret").First()
	ref := gf.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code", account["AppID"], account["AppSecret"], code))
	var parameter map[string]interface{}
	if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
		rooturl, _ := model.DB().Table("common_config").Where("keyname", "rooturl").Value("keyvalue")
		user, _ := model.DB().Table("business_wxsys_user").Where("wxapp_openid", parameter["openid"]).Fields("id,accountID,businessID,name,avatar,nickname,openid,unionid").First()
		if user != nil {
			//token
			token := getToken(user)
			if !strings.Contains(user["avatar"].(string), "http") && rooturl != nil {
				user["avatar"] = rooturl.(string) + user["avatar"].(string)
			}
			results.SuccessLogin(c, "直接获取已有的用户数据", user, token, nil)
		} else { //不存在则添加一条
			parameter["accountID"] = account["accountID"]
			parameter["businessID"] = account["businessID"]
			parameter["wxapp_openid"] = parameter["openid"]
			parameter["createtime"] = time.Now().Unix()
			parameter["avatar"] = "resource/staticfile/avatar.png"
			delete(parameter, "session_key")
			delete(parameter, "openid")
			addId, err := model.DB().Table("business_wxsys_user").Data(parameter).InsertGetId()
			if err != nil {
				results.Failed(c, "添加账号失败", err)
			} else {
				model.DB().Table("business_wxsys_user").Data(map[string]interface{}{"name": fmt.Sprintf("%s%v", "hl_", addId)}).Where("id", addId).Update()
				user, _ := model.DB().Table("business_wxsys_user").Where("id", addId).Fields("id,accountID,businessID,name,avatar,nickname,openid,unionid").First()
				//token
				token := getToken(user)
				if !strings.Contains(user["avatar"].(string), "http") && rooturl != nil {
					user["avatar"] = rooturl.(string) + user["avatar"].(string)
				}
				results.SuccessLogin(c, "添加并获取token！", user, token, nil)
			}
		}
	} else {
		results.Failed(c, "获取openid失败", err)
	}
}

// 获取Token
func getToken(user map[string]interface{}) interface{} {
	token := middleware.GenerateToken(&middleware.UserClaims{
		ID:             user["id"].(int64),
		Accountid:      user["accountID"].(int64),
		BusinessID:     user["businessID"].(int64),
		Openid:         user["openid"].(string),
		StandardClaims: jwt.StandardClaims{},
	})
	return token
}

// 获取用户手机
func (api *Index) Get_phone(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	businessID := c.GetHeader("Businessid")
	account, _ := model.DB().Table("business_wxsys_wxappconfig").Where("businessID", businessID).Fields("id,accountID,businessID,AppID,AppSecret,access_token,secret,expires_in").First()
	if account == nil {
		results.Failed(c, "小程序账号不存在", nil)
		return
	}
	expires_in := account["expires_in"].(int64)
	nowtime := time.Now().Unix()
	if expires_in-nowtime <= 0 {
		url_str := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%v&secret=%v", account["expires_in"], account["secret"].(string))
		access_ref := gf.Get(url_str)
		var access_parameter map[string]interface{}
		if err := json.Unmarshal([]byte(access_ref), &access_parameter); err == nil {
			account["access_token"] = access_parameter["access_token"]
			model.DB().Table("business_wxsys_wxappconfig").Data(access_parameter).Where("id", account["id"]).Update()
		} else {
			results.Failed(c, "获取access_token失败", err)
			return
		}
	}
	ref := gf.Post("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token="+account["access_token"].(string),
		map[string]interface{}{"code": code}, "application/json")
	var parameter map[string]interface{}
	if err := json.Unmarshal([]byte(ref), &parameter); err == nil {
		if parameter["errmsg"] == "ok" {
			var phone_info map[string]interface{}
			m_phone_info, _ := json.Marshal(parameter["phone_info"])
			json.Unmarshal(m_phone_info, &phone_info)
			//当前用户
			token := c.Request.Header.Get("Authorization")
			user := middleware.ParseToken(token)
			_, err := model.DB().Table("business_wxsys_user").Data(map[string]interface{}{"mobile": phone_info["phoneNumber"]}).Where("id", user.ID).Update()
			if err == nil {
				results.Success(c, "获取用户手机", "ok", nil)
			} else {
				results.Failed(c, "获取手机失败1", err)
			}
		} else {
			results.Failed(c, "获取手机失败2", parameter)
		}
	} else {
		results.Failed(c, "获取手机失败", err)
	}
}

// 获取用户信息
func (api *Index) Get_userinfo(c *gin.Context) {
	//当前用户
	token := c.Request.Header.Get("Authorization")
	user := middleware.ParseToken(token)
	data, err := model.DB().Table("business_wxsys_user").Where("id", user.ID).First()
	if err != nil {
		results.Failed(c, "获取用户信息失败", err)
	} else {
		results.Success(c, "获取用户信息成功！", data, nil)
	}
}

// 保存
func (api *Index) UpuserInfo(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	token := c.Request.Header.Get("Authorization")
	user := middleware.ParseToken(token)
	_ = json.Unmarshal(body, &parameter)
	res, err := model.DB().Table("business_wxsys_user").
		Data(parameter).
		Where("id", user.ID).
		Update()
	if err != nil {
		results.Failed(c, "更新失败", err)
	} else {
		results.Success(c, "更新成功！", res, nil)
	}
}

// 获取token-APi
func (api *Index) Get_apitoken(c *gin.Context) {
	user_id := c.DefaultQuery("user_id", "")
	userdata, _ := model.DB().Table("business_wxsys_user").Where("id", user_id).Fields("id,accountID,businessID,name,wxapp_openid").First()
	if userdata == nil {
		results.Failed(c, "账号不存在", nil)
	} else {
		token := middleware.GenerateToken(&middleware.UserClaims{
			ID:             userdata["id"].(int64),
			Accountid:      userdata["accountID"].(int64),
			BusinessID:     userdata["businessID"].(int64),
			Openid:         userdata["wxapp_openid"].(string),
			StandardClaims: jwt.StandardClaims{},
		})
		results.Success(c, "获取测试Token", token, nil)
	}
}
