package developer

import (
	"encoding/json"
	"fmt"
	"gofly/global"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils/gf"
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
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
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

// 更新生成代码的数据表
func (api *Generatecode) UpCodeTable(c *gin.Context) {
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
			"rule_name":  val["TABLE_COMMENT"],
			"engine":     val["ENGINE"],
			"table_rows": val["TABLE_ROWS"],
			"createtime": gf.NowTimestamp(), "updatetime": gf.NowTimestamp(),
			"collation": val["TABLE_COLLATION"], "auto_increment": val["AUTO_INCREMENT"]}
		if gf.IsContain(tablenames.([]interface{}), val["TABLE_NAME"].(string)) {
			delete(midata, "createtime")
			delete(midata, "rule_name")
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
}

// 保存-生成代码
func (api *Generatecode) Save(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	codedata := parameter["codedata"].(map[string]interface{})
	field_list := parameter["field_list"].([]interface{})
	list := parameter["list"].([]interface{})
	search_list := parameter["search_list"].([]interface{})
	//更新字段列表数据
	for _, fval := range field_list {
		item := fval.(map[string]interface{})
		model.DB().Table("common_generatecode_field").Data(item).Where("id", item["id"]).Update()
	}
	for _, lval := range list {
		item := lval.(map[string]interface{})
		model.DB().Table("common_generatecode_field").Data(item).Where("id", item["id"]).Update()
	}
	for _, sval := range search_list {
		item := sval.(map[string]interface{})
		model.DB().Table("common_generatecode_field").Data(item).Where("id", item["id"]).Update()
	}
	//1生成菜单
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	findrule, _ := model.DB().Table("business_auth_rule").Where("routePath", codedata["routePath"]).OrWhere("routeName", codedata["routeName"]).Fields("id").First()
	var isok = false
	if findrule == nil {
		save_arr := map[string]interface{}{"createtime": time.Now().Unix(),
			"title": codedata["rule_name"], "type": 1, "uid": user.ID,
			"icon": codedata["icon"], "routePath": codedata["routePath"], "routeName": codedata["routeName"],
			"pid": codedata["pid"], "component": codedata["component"],
		}
		getId, err := model.DB().Table("business_auth_rule").Data(save_arr).InsertGetId()
		if err != nil {
			results.Failed(c, "添加菜单失败", err)
		} else { //更新排序
			model.DB().Table("business_auth_rule").
				Data(map[string]interface{}{"orderNo": getId}).
				Where("id", getId).
				Update()
			codedata["rule_id"] = getId
			isok = true
		}
	} else {
		isok = true
		codedata["rule_id"] = findrule["id"]
	}
	//菜单添加好后添加代码
	if isok {
		/***************************后端**************************/
		file_path := filepath.Join("app/", gf.InterfaceTostring(codedata["api_path"]))
		//1. 如果没有filepath文件目录就创建一个
		if _, err := os.Stat(file_path); err != nil {
			if !os.IsExist(err) {
				os.MkdirAll(file_path, os.ModePerm)
			}
		}
		//2. 替换文件内容
		filename_arr := strings.Split(codedata["api_filename"].(string), `.`) //文件名称
		packgename_arr := strings.Split(codedata["api_path"].(string), `/`)
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
		//创建后端代码
		fields_inter, _ := model.DB().Table("common_generatecode_field").Where("generatecode_id", codedata["id"]).Where("islist", 1).
			Order("list_weigh asc,id asc").Pluck("field")
		if fields_inter != nil {
			fields_arr := fields_inter.([]interface{})
			var str_arr = make([]string, len(fields_arr))
			for k, v := range fields_arr {
				str_arr[k] = fmt.Sprintf("%v", v)
			}
			codedata["fields"] = strings.Join(str_arr, ",")
		} else {
			codedata["fields"] = ""
		}
		go MarkeGoCode(file_path, filename, packageName, codedata)
		//3. 查看是否添加文件到控制器文件
		go CheckIsAddController(modelname, gf.InterfaceTostring(codedata["api_path"]))
		/******************************前端******************************/
		component_arr := strings.Split(codedata["component"].(string), `/`)
		componentpah_arr := strings.Split(codedata["component"].(string), (component_arr[len(component_arr)-1]))
		vue_path := filepath.Join(global.App.Config.App.Vueobjroot, "/src/views/", componentpah_arr[0]) //前端文件路径
		//1. 如果没有filepath文件目录就创建一个
		if _, err := os.Stat(vue_path); err != nil {
			if !os.IsExist(err) {
				os.MkdirAll(vue_path, os.ModePerm)
			}
		}
		//2. 复制前端模板到新创建文件夹下
		CopyAllDir(filepath.Join("resource/developer/codetpl/vue/", gf.InterfaceTostring(codedata["tpl_type"])), vue_path)
		//3. 修改模板文件内容
		if codedata["tpl_type"] == "contentcatelist" { //如果是关联分类则更新分类api.ts
			ApitsReplay(filepath.Join(vue_path, "cate/api.ts"), packageName, filename+"cate")
		}
		//修改api/index.ts文件
		ApitsReplay(filepath.Join(vue_path, "api/index.ts"), packageName, filename)
		//替换data.ts
		listfield, _ := model.DB().Table("common_generatecode_field").Where("generatecode_id", codedata["id"]).Where("islist", 1).
			Fields("id,name,field,align,width").Order("list_weigh asc,id asc").Get()
		UpFieldData(filepath.Join(vue_path, "data.ts"), listfield) //更新data.ts
		//替换AddForm.vue表单
		formfield, _ := model.DB().Table("common_generatecode_field").Where("generatecode_id", codedata["id"]).Where("isform", 1).
			Fields("id,name,field,required,formtype,datatable,datatablename").Order("field_weigh asc,id asc").Get()
		UpFieldAddForm(filepath.Join(vue_path, "AddForm.vue"), codedata["fields"], formfield) //更新表单
		/*************最后更新代码生成表数据***************************/
		codedata["is_install"] = 1
		res, err := model.DB().Table("common_generatecode").Data(codedata).Where("id", codedata["id"]).Update()
		if err != nil {
			results.Failed(c, "更新失败", err)
		} else {
			results.Success(c, "更新成功！", res, nil)
		}
	} else {
		results.Failed(c, "添加菜单失败", nil)
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
	if parameter["is_install"] != nil && gf.InterfaceToInt(parameter["is_install"]) == 1 { //卸载
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
		file_path := filepath.Join("app/", gf.InterfaceTostring(data["api_path"]))
		//判断后端代码是否存在删除后端代码
		filego_path := filepath.Join(file_path, gf.InterfaceTostring(data["api_filename"]))
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
func (api *Generatecode) GetContent(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		results.Failed(c, "请传参数id", nil)
	} else {
		data, err := model.DB().Table("common_generatecode").Fields("id,tablename,comment,pid,rule_id,rule_name,icon,is_install,routePath,routeName,component,api_path,api_filename,cate_tablename,tpl_type").Where("id", id).First()
		if err != nil {
			results.Failed(c, "获取内容失败", err)
		} else {
			if data == nil {
				results.Failed(c, "生成数据表不存在", err)
			} else {

				var dielddata_list []map[string]interface{}
				var haseids []interface{}
				dielddata, _ := model.DB().Query("select COLUMN_NAME,COLUMN_COMMENT,DATA_TYPE,CHARACTER_MAXIMUM_LENGTH from information_schema.columns where TABLE_SCHEMA='" + global.App.Config.DBconf.Database + "' AND TABLE_NAME='" + data["tablename"].(string) + "'")
				for _, data := range dielddata {
					if data["COLUMN_COMMENT"] == "" && data["COLUMN_NAME"] == "id" {
						data["COLUMN_COMMENT"] = "ID"
					}
					formtype := "text"
					searchtype := "text"
					isorder := 0
					if data["COLUMN_NAME"] == "id" {
						isorder = 1
					}
					def_value := "[]"
					if data["DATA_TYPE"] == "int" {
						formtype = "number"
					} else if data["DATA_TYPE"] == "varchar" && data["CHARACTER_MAXIMUM_LENGTH"] == 225 {
						formtype = "textarea"
					} else if data["DATA_TYPE"] == "text" {
						formtype = "editor"
					} else if data["DATA_TYPE"] == "enum" {
						formtype = "select"
						searchtype = "select"
					} else if strings.HasSuffix(gf.InterfaceTostring(data["COLUMN_NAME"]), "date") {
						formtype = "date"
					} else if strings.HasSuffix(gf.InterfaceTostring(data["COLUMN_NAME"]), "datetime") {
						formtype = "datetime"
					} else if strings.HasSuffix(gf.InterfaceTostring(data["COLUMN_NAME"]), "time") {
						formtype = "time"
					} else if strings.HasSuffix(gf.InterfaceTostring(data["COLUMN_NAME"]), "image") {
						formtype = "image"
					} else if strings.HasSuffix(gf.InterfaceTostring(data["COLUMN_NAME"]), "images") {
						formtype = "images"
					} else if strings.HasSuffix(gf.InterfaceTostring(data["COLUMN_NAME"]), "file") {
						formtype = "file"
					} else if strings.HasSuffix(gf.InterfaceTostring(data["COLUMN_NAME"]), "files") {
						formtype = "files"
					}
					if fieldval, _ := model.DB().Table("common_generatecode_field").Where("generatecode_id", id).Where("field", data["COLUMN_NAME"]).Value("id"); fieldval != nil {
						haseids = append(haseids, fieldval)
					} else {
						dielddata_list = append(dielddata_list, map[string]interface{}{"generatecode_id": id, "name": data["COLUMN_COMMENT"], "field": data["COLUMN_NAME"], "formtype": formtype, "def_value": def_value, "searchtype": searchtype, "isorder": isorder})
					}
				}
				if haseids != nil {
					model.DB().Table("common_generatecode_field").Where("generatecode_id", id).WhereNotIn("id", haseids).Delete()
				}
				if dielddata_list != nil {
					model.DB().Table("common_generatecode_field").Data(dielddata_list).Insert()
				}
				field_list, _ := model.DB().Table("common_generatecode_field").Where("generatecode_id", id).
					Fields("id,isform,name,field,required,formtype,datatable,datatablename,field_weigh").Order("field_weigh asc,id asc").Get()
				for _, fval := range field_list {
					if gf.InterfaceToInt(fval["isform"]) == 1 {
						fval["isform"] = true
					} else {
						fval["isform"] = false
					}
					if gf.InterfaceToInt(fval["required"]) == 1 {
						fval["required"] = true
					} else {
						fval["required"] = false
					}
				}
				list, _ := model.DB().Table("common_generatecode_field").Where("generatecode_id", id).
					Fields("id,islist,name,field,isorder,align,width,list_weigh").Order("list_weigh asc,id asc").Get()
				for _, lval := range list {
					if gf.InterfaceToInt(lval["islist"]) == 1 {
						lval["islist"] = true
					} else {
						lval["islist"] = false
					}
					if gf.InterfaceToInt(lval["isorder"]) == 1 {
						lval["isorder"] = true
					} else {
						lval["isorder"] = false
					}
				}
				search_list, _ := model.DB().Table("common_generatecode_field").Where("generatecode_id", id).
					Fields("id,issearch,name,searchway,searchtype,search_weigh").Order("search_weigh asc,id asc").Get()
				for _, sval := range search_list {
					if gf.InterfaceToInt(sval["issearch"]) == 1 {
						sval["issearch"] = true
					} else {
						sval["issearch"] = false
					}
				}
				results.Success(c, "获取生成表单信息成功！", gf.Map{"data": data, "field_list": field_list, "list": list, "search_list": search_list}, nil)
			}
		}
	}

}
