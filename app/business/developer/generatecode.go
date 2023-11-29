package developer

import (
	"encoding/json"
	"gofly/global"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils"
	"gofly/utils/results"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/**
* 代码安装
 */
// 用于自动注册路由
type Generatecode struct{}

func init() {
	fpath := Generatecode{}
	utils.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取列表
func (api *Generatecode) Get_list(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	page := c.DefaultQuery("page", "1")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(_pageSize)
	MDB := model.DB().Table("common_generatecode").Where("status", 0)
	MDBC := model.DB().Table("common_generatecode").Where("status", 0)
	if name != "" {
		MDB.Where("tablename", "like", "%"+name+"%")
		MDBC.Where("tablename", "like", "%"+name+"%")
	}
	list, err := MDB.Limit(pageSize).Page(pageNo).Order("id desc").Get()
	if err != nil {
		results.Failed(c, err.Error(), nil)
	} else {
		var totalCount int64
		totalCount, _ = MDBC.Count("*")
		results.Success(c, "获取全部列表", map[string]interface{}{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}, nil)
	}
}

// 获取数据表字段
func (api *Generatecode) Get_dbfield(c *gin.Context) {
	tablename := c.DefaultQuery("tablename", "")
	if tablename == "" {
		results.Failed(c, "请传数据表名称", nil)
	} else {
		tablename_arr := strings.Split(tablename, ",")
		//获取数据库名
		var dielddata_list []map[string]interface{}
		for _, Val := range tablename_arr {
			dielddata, _ := model.DB().Query("select COLUMN_NAME,COLUMN_COMMENT,DATA_TYPE from information_schema.columns where TABLE_SCHEMA='" + global.App.Config.DBconf.Database + "' AND TABLE_NAME='" + Val + "'")
			for _, data := range dielddata {
				if data["COLUMN_COMMENT"] == "" && data["COLUMN_NAME"] == "id" {
					data["COLUMN_COMMENT"] = "ID"
				}
				dielddata_list = append(dielddata_list, map[string]interface{}{"value": data["COLUMN_NAME"], "label": data["COLUMN_COMMENT"], "type": data["DATA_TYPE"]})
			}
		}
		results.Success(c, "获取数据表字段", dielddata_list, nil)
	}
}

// 获取数据库列表
func (api *Generatecode) Get_tablelist(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	dbtalbelist, _ := model.DB().Table("common_generatecode").Where("status", 0).Where("id", "!=", id).Fields("tablename as value,comment as label").Order("id desc").Get()
	results.Success(c, "获取数据库列表", dbtalbelist, nil)
}

// 保存-安装
func (api *Generatecode) Save(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	var f_id float64 = 0
	if parameter["id"] != nil {
		f_id = parameter["id"].(float64)
	}
	if f_id == 0 {
		//一般不走这里
		delete(parameter, "id")
		//获取数据库名
		tablenames, _ := model.DB().Table("common_generatecode").Pluck("tablename")
		dbtalbelist, _ := model.DB().Query("SELECT TABLE_NAME,TABLE_COMMENT,ENGINE,TABLE_ROWS,TABLE_COLLATION,AUTO_INCREMENT FROM information_schema.TABLES WHERE table_schema='" + global.App.Config.DBconf.Database + "'")
		save_arr := []map[string]interface{}{}
		for _, val := range dbtalbelist {
			webb, _ := json.Marshal(val)
			var webjson map[string]interface{}
			_ = json.Unmarshal(webb, &webjson)
			if val["TABLE_COMMENT"] == nil {
				val["TABLE_COMMENT"] = ""
			}
			if val["AUTO_INCREMENT"] == nil {
				val["AUTO_INCREMENT"] = ""
			}
			midata := map[string]interface{}{"tablename": val["TABLE_NAME"], "comment": val["TABLE_COMMENT"],
				"engine":     val["ENGINE"],
				"table_rows": val["TABLE_ROWS"],
				"createtime": utils.NowTimestamp(), "updatetime": utils.NowTimestamp(),
				"collation": val["TABLE_COLLATION"], "auto_increment": val["AUTO_INCREMENT"]}
			if utils.IsContain(tablenames.([]interface{}), val["TABLE_NAME"].(string)) {
				delete(midata, "createtime")
				model.DB().Table("common_generatecode").Data(midata).Where("tablename", val["TABLE_NAME"]).Update()
			} else {
				save_arr = append(save_arr, midata)
			}
		}
		if save_arr != nil && len(save_arr) > 0 {
			_, err := model.DB().Table("common_generatecode").Data(save_arr).Insert()
			if err != nil {
				results.Failed(c, "更新失败", nil)
			} else {
				results.Success(c, "更新成功！", save_arr, nil)
			}
		} else {
			results.Success(c, "已更新全部", dbtalbelist, nil)
		}
	} else { //更新并生成
		//1生成菜单
		getuser, _ := c.Get("user")
		user := getuser.(*middleware.UserClaims)
		findrule, _ := model.DB().Table("business_auth_rule").Where("routePath", parameter["routePath"]).OrWhere("routeName", parameter["routeName"]).Fields("id").First()
		var isok = false
		if findrule == nil {
			save_arr := map[string]interface{}{"createtime": time.Now().Unix(),
				"title": parameter["comment"], "type": 1, "uid": user.ID,
				"icon": parameter["icon"], "routePath": parameter["routePath"], "routeName": parameter["routeName"],
				"pid": parameter["pid"], "component": parameter["component"],
			}
			getId, err := model.DB().Table("business_auth_rule").Data(save_arr).InsertGetId()
			if err != nil {
				results.Failed(c, "添加菜单失败", err)
			} else { //更新排序
				model.DB().Table("business_auth_rule").
					Data(map[string]interface{}{"orderNo": getId}).
					Where("id", getId).
					Update()
				parameter["rule_id"] = getId
				isok = true
			}
		} else {
			isok = true
			parameter["rule_id"] = findrule["id"]
		}
		//菜单添加好后添加代码
		if isok {
			/***************************后端**************************/
			file_path := filepath.Join("app/", utils.InterfaceTostring(parameter["api_path"]))
			//1. 如果没有filepath文件目录就创建一个
			if _, err := os.Stat(file_path); err != nil {
				if !os.IsExist(err) {
					os.MkdirAll(file_path, os.ModePerm)
				}
			}
			//2. 替换文件内容
			filename_arr := strings.Split(parameter["api_filename"].(string), `.`) //文件名称
			packgename_arr := strings.Split(parameter["api_path"].(string), `/`)
			//2.1 模块名称
			modelname := "business"
			if len(packgename_arr) > 0 {
				modelname = packgename_arr[0]
			}
			//2.2 文件名称
			filename := "index"
			if len(filename_arr) > 0 {
				filename = filename_arr[0]
			}
			//2.3 包名
			packageName := ""
			if len(packgename_arr) > 0 {
				packageName = packgename_arr[len(packgename_arr)-1]
			}
			// //创建后端代码
			go MarkeGoCode(file_path, filename, packageName, parameter)
			// // //3. 查看是否添加文件到控制器文件
			go CheckIsAddController(modelname, utils.InterfaceTostring(parameter["api_path"]))
			/******************************前端******************************/
			component_arr := strings.Split(parameter["component"].(string), `/`)
			componentpah_arr := strings.Split(parameter["component"].(string), (component_arr[len(component_arr)-1]))
			vue_path := filepath.Join(global.App.Config.App.Vueobjroot, "/src/views/", componentpah_arr[0]) //前端文件路径
			//1. 如果没有filepath文件目录就创建一个
			if _, err := os.Stat(vue_path); err != nil {
				if !os.IsExist(err) {
					os.MkdirAll(vue_path, os.ModePerm)
				}
			}
			//2. 复制前端模板到新创建文件夹下
			CopyAllDir(filepath.Join("resource/staticfile/codetpl/vue/", utils.InterfaceTostring(parameter["tpl_type"])), vue_path)
			//3. 修改模板文件内容
			if parameter["tpl_type"] == "contentcatelist" { //如果是关联分类则更新分类api.ts
				ApitsReplay(filepath.Join(vue_path, "cate/api.ts"), packageName, filename+"cate")
			}
			//修改api/index.ts文件
			ApitsReplay(filepath.Join(vue_path, "api/index.ts"), packageName, filename)
			//替换data.ts
			UpFieldData(filepath.Join(vue_path, "data.ts"), parameter["tablefieldname"]) //更新data.ts
			//替换AddForm.vue表单-根据读取（parameter["tablefieldname"]）修vue
			UpFieldAddForm(filepath.Join(vue_path, "AddForm.vue"), utils.InterfaceTostring(parameter["fields"]), parameter["tablefieldname"]) //更新表单
			parameter["is_install"] = 1
			delete(parameter, "tablefieldname")
			res, err := model.DB().Table("common_generatecode").
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
}

// 更新状态
func (api *Generatecode) UpStatus(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	res2, err := model.DB().Table("common_generatecode").Where("id", parameter["id"]).Data(map[string]interface{}{"status": parameter["status"]}).Update()
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

// 删除/卸载
func (api *Generatecode) Del(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	if parameter["is_install"] != nil && utils.InterfaceToInt(parameter["is_install"]) == 1 { //卸载
		isok, err := common_uninstall(parameter["id"])
		if isok {
			model.DB().Table("common_generatecode").Where("id", parameter["id"]).Data(map[string]interface{}{"is_install": 2}).Update()
			results.Success(c, "卸载成功！", nil, nil)
		} else {
			results.Failed(c, "卸载失败", err)
		}
	} else { //删除
		res2, err := model.DB().Table("common_generatecode").Where("id", parameter["id"]).Delete()
		if err != nil {
			results.Failed(c, "删除失败", err)
		} else {
			results.Success(c, "删除成功！", res2, nil)
		}
	}
}

// 卸载
func (api *Generatecode) Uninstallcode(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	isok, err := common_uninstall(parameter["id"])
	if isok {
		results.Success(c, "卸载成功！", 0, nil)
	} else {
		results.Failed(c, "卸载失败", err)
	}
}

// 卸载通用方法
func common_uninstall(id interface{}) (bool, error) {
	data, err := model.DB().Table("common_generatecode").Where("id", id).Fields("rule_id,api_path,api_filename,tpl_type,component").First()
	if err != nil {
		return false, err
	} else {
		file_path := filepath.Join("app/", utils.InterfaceTostring(data["api_path"]))
		//判断后端代码是否存在删除后端代码
		filego_path := filepath.Join(file_path, utils.InterfaceTostring(data["api_filename"]))
		if _, err := os.Stat(filego_path); err == nil {
			//删除菜单
			model.DB().Table("business_auth_rule").Where("id", data["rule_id"]).Delete()
			model.DB().Table("common_generatecode").Data(map[string]interface{}{"is_install": 0}).Where("id", id).Update()
			go UnInstallCodeFile(data)
		}
		return true, nil
	}
}

// 获取内容
func (api *Generatecode) Get_content(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		results.Failed(c, "请传参数id", nil)
	} else {
		res2, err := model.DB().Table("common_generatecode").Where("id", id).First()
		if err != nil {
			results.Failed(c, "获取内容失败", err)
		} else {
			results.Success(c, "获取内容成功！", res2, nil)
		}
	}

}
