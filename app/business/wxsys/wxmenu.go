package wxsys

import (
	"encoding/json"
	"fmt"
	"gofly/app/model"
	"gofly/route/middleware"
	"gofly/utils"
	"gofly/utils/results"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 用于自动注册路由
type Wxmenu struct {
}

func init() {
	fpath := Wxmenu{}
	utils.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取菜单-微信服务器上
func (api *Wxmenu) Get_menu(c *gin.Context) {
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	//获取公众号配置
	wxconfig, _ := model.DB().Table("business_wxsys_officonfig").Where("businessID", user.BusinessID).Fields("id,name,AppID,AppSecret,expires_access_token,access_token").First()
	//更新access_token
	AccessTokenHost := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", wxconfig["AppID"], wxconfig["AppSecret"])
	var (
		access_token  string
		wxAccessToken WxAccessToken
	)
	timestamp := time.Now().Unix()                                       //10位时间戳
	expires_access_token_int := wxconfig["expires_access_token"].(int64) //数据库的时间传戳
	//获取access_token，如果缓存中有，则直接取出数据使用；否则重新调用微信端接口获取
	client := &http.Client{}
	//判断access_token是否过期
	if wxconfig["access_token"] == "" || expires_access_token_int == 0 || (timestamp-expires_access_token_int) > 7000 { //重新请求access_token
		request, _ := http.NewRequest("GET", AccessTokenHost, nil)
		response, _ := client.Do(request)
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			results.Failed(c, "请求AccessToken失败1", err.Error())
			return
		}
		err = json.Unmarshal(body, &wxAccessToken)
		if err != nil {
			results.Failed(c, "解析AccessToken失败", err.Error())
			return
		}
		if wxAccessToken.Errcode == 0 {
			access_token = wxAccessToken.Access_token
		} else {
			results.Failed(c, "获取AccessToken失败", wxAccessToken.Errmsg)
			return
		}
		//添加access_tokens时间
		model.DB().Table("business_wxsys_officonfig").Where("id", wxconfig["id"]).Data(map[string]interface{}{"access_token": access_token, "expires_access_token": time.Now().Unix()}).Update()
	} else {
		//缓存中存在access_token，直接读取
		access_token = wxconfig["access_token"].(string)
	}
	//获取 菜单接口
	wxmenu_data, err := utils.Get_x(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=%s", access_token))
	if err != nil {
		results.Failed(c, "获取微信菜单失败1", err)
	} else {
		var data_parameter map[string]interface{}
		if err := json.Unmarshal([]byte(wxmenu_data), &data_parameter); err == nil {
			if _, ok := data_parameter["errcode"]; ok {
				results.Failed(c, "获取菜单接口失败", data_parameter)
			} else {
				results.Success(c, "获取微信菜单", map[string]interface{}{
					"name":   wxconfig["name"],
					"wxmenu": data_parameter,
				}, nil)
			}
		}
	}
}

// 保存微信菜单
func (api *Wxmenu) SaveMenuOnly(c *gin.Context) {
	//获取post传过来的data
	body, _ := ioutil.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	var f_id float64 = 0
	if parameter["id"] != nil {
		f_id = parameter["id"].(float64)
	}
	if parameter["menu"] != nil {
		parameter["menu"] = utils.JSONMarshalToString(parameter["menu"])
	}
	if f_id == 0 {
		delete(parameter, "id")
		parameter["accountID"] = user.Accountid
		parameter["businessID"] = user.BusinessID
		addId, err := model.DB().Table("business_wxsys_wxmenu").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			if addId != 0 {
				model.DB().Table("business_wxsys_wxmenu").
					Data(map[string]interface{}{"weigh": addId}).
					Where("id", addId).
					Update()
			}
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		res, err := model.DB().Table("business_wxsys_wxmenu").
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

// 获取菜单列表
func (api *Wxmenu) Get_menuList(c *gin.Context) {
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	list, err := model.DB().Table("business_wxsys_wxmenu").Where("businessID", user.BusinessID).Order("id desc").Get()
	if err != nil {
		results.Failed(c, "获取菜单失败", err)
	} else {
		results.Success(c, "获取菜单", list, nil)
	}
}

// 删除
func (api *Wxmenu) Del_menu(c *gin.Context) {
	//获取post传过来的data
	body, _ := ioutil.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	res2, err := model.DB().Table("business_wxsys_wxmenu").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		results.Success(c, "删除成功！", res2, nil)
	}
}

// 发布菜单
func (api *Wxmenu) SaveMenu(c *gin.Context) {
	//获取post传过来的data
	body, _ := ioutil.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	//获取公众号配置
	wxconfig, _ := model.DB().Table("business_wxsys_officonfig").Where("businessID", user.BusinessID).Fields("id,name,AppID,AppSecret,expires_access_token,access_token").First()
	//更新access_token
	AccessTokenHost := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", wxconfig["AppID"], wxconfig["AppSecret"])
	timestamp := strconv.FormatInt(time.Now().Unix(), 10) //10位时间戳
	var (
		access_token  string
		wxAccessToken WxAccessToken
	)
	//当前时间戳转int
	intNum, _ := strconv.Atoi(timestamp)
	timestampint := int64(intNum)
	expires_access_token_int := wxconfig["expires_access_token"].(int64) //数据库的时间传戳
	//获取access_token，如果缓存中有，则直接取出数据使用；否则重新调用微信端接口获取
	client := &http.Client{}
	//判断access_token是否过期
	if wxconfig["access_token"] == "" || expires_access_token_int == 0 || (timestampint-expires_access_token_int) > 7200 { //重新请求access_token
		request, _ := http.NewRequest("GET", AccessTokenHost, nil)
		response, _ := client.Do(request)
		defer response.Body.Close() //最后再执行
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			results.Failed(c, "请求AccessToken失败1", err.Error())
			return
		}
		err = json.Unmarshal(body, &wxAccessToken)
		if err != nil {
			results.Failed(c, "解析AccessToken失败", err.Error())
			return
		}
		if wxAccessToken.Errcode == 0 {
			access_token = wxAccessToken.Access_token
		} else {
			results.Failed(c, "获取AccessToken失败", wxAccessToken.Errmsg)
			return
		}
		//添加access_tokens时间
		model.DB().Table("client_system_wxconfig").Where("id", wxconfig["id"]).Data(map[string]interface{}{"access_token": access_token, "expires_access_token": timestamp}).Update()
	} else {
		//缓存中存在access_token，直接读取
		access_token = wxconfig["access_token"].(string)
	}
	//获取 菜单接口
	wxmenu_data, err := utils.Post_strdata(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s", access_token), parameter["data"].(string), "")
	if err != nil {
		results.Failed(c, "获取微信openid失败", err)
	} else {
		var data_parameter map[string]interface{}
		if err := json.Unmarshal([]byte(wxmenu_data), &data_parameter); err == nil {
			if data_parameter["errcode"].(float64) != 0 {
				results.Failed(c, "创建微信菜单失败", data_parameter)
			} else {
				results.Success(c, "创建微信菜单成功", data_parameter, parameter)
			}
		}
	}
}

// 类型
type WxAccessToken struct {
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
}
type WxJsApiTicket struct {
	Ticket     string `json:"ticket"`
	Expires_in int    `json:"expires_in"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}
type WxSignature struct {
	Noncestr  string `json:"noncestr"`
	Timestamp string `json:"timestamp"`
	Url       string `json:"url"`
	Signature string `json:"signature"`
	AppID     string `json:"appId"`
}

type WxSignRtn struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data WxSignature `json:"data"`
}
