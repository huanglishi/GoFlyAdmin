package developer

import (
	"encoding/json"
	"gofly/model"
	"gofly/utils"
	"gofly/utils/results"
	"io"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

// 用于自动注册路由
type Apicode struct {
}

func init() {
	fpath := Apicode{}
	utils.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 生成api接口代码
func (api *Apicode) Installcode(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	data, err := model.DB().Table("common_apitext").Where("id", parameter["id"]).Fields("cid,url,getdata_type,tablename,apicode_type,is_install,fields,method").First()
	if err != nil {
		results.Failed(c, "生成api接口代码失败", err)
	} else {
		type_id, _ := model.DB().Table("common_apitext_group").Where("id", data["cid"]).Value("type_id")
		rooturl, _ := model.DB().Table("common_apitext_type").Where("id", type_id).Value("rooturl")
		root_path := "business"
		if rooturl != nil { //模块名称
			rooturl_arr := strings.Split(rooturl.(string), `/`)
			root_path = rooturl_arr[len(rooturl_arr)-1]
		}
		//创建文件
		CreatApicodeFile(root_path, data)
		model.DB().Table("common_apitext").
			Data(map[string]interface{}{"is_install": 1}).
			Where("id", parameter["id"]).
			Update()
		results.Success(c, "生成api接口代码成功！", data, nil)
	}
}

// 卸载api接口代码-改变方法
func (api *Apicode) Uninstallcode(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	data, err := model.DB().Table("common_apitext").Where("id", parameter["id"]).Fields("cid,url,getdata_type,tablename,apicode_type,is_install,fields,method").First()
	if err != nil {
		results.Failed(c, "卸载失败", err)
	} else {
		type_id, _ := model.DB().Table("common_apitext_group").Where("id", data["cid"]).Value("type_id")
		rooturl, _ := model.DB().Table("common_apitext_type").Where("id", type_id).Value("rooturl")
		root_path := "business"
		if rooturl != nil { //模块名称
			rooturl_arr := strings.Split(rooturl.(string), `/`)
			root_path = rooturl_arr[len(rooturl_arr)-1]
		}
		UnApicodeFile(root_path, data)
		model.DB().Table("common_apitext").Data(map[string]interface{}{"is_install": 2}).Where("id", parameter["id"]).Update()
		results.Success(c, "卸载成功！", data, nil)
	}
}

// 删除文件
func (api *Apicode) RemoveFile(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	data, err := model.DB().Table("common_apitext").Where("id", parameter["id"]).Fields("cid,url,getdata_type,tablename,apicode_type,is_install,fields,method").First()
	if err != nil {
		results.Failed(c, "删除文件失败", err)
	} else {
		type_id, _ := model.DB().Table("common_apitext_group").Where("id", data["cid"]).Value("type_id")
		rooturl, _ := model.DB().Table("common_apitext_type").Where("id", type_id).Value("rooturl")
		root_path := "business"
		if rooturl != nil { //模块名称
			rooturl_arr := strings.Split(rooturl.(string), `/`)
			root_path = rooturl_arr[len(rooturl_arr)-1]
		}
		//判断删除文件
		url := data["url"].(string)
		url_arr := strings.Split(url, `/`)
		filename := url_arr[len(url_arr)-1]
		model_path := strings.Split(url, filename)
		haselist, _ := model.DB().Table("common_apitext").Where("url", "like", model_path[0]+"%").Where("is_install", 1).Count("*")
		if haselist == 0 {
			RemoveModel(root_path, data) //删除文件-如果没人其他文件则移除文件夹及路由
		}
		model.DB().Table("common_apitext").Data(map[string]interface{}{"is_install": 0}).Where("id", parameter["id"]).Update()
		results.Success(c, "删除文件成功！", data, haselist)
	}
}
