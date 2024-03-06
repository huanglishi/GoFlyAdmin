package developer

import (
	"encoding/json"
	"fmt"
	"gofly/global"
	"gofly/model"
	"gofly/utils/gf"
	"gofly/utils/results"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"gofly/utils/gform"

	"github.com/gin-gonic/gin"
)

/**
* 接口文档
 */
// 用于自动注册路由
type Devapi struct {
}

func init() {
	fpath := Devapi{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取部门列表
func (api *Devapi) Get_list(c *gin.Context) {
	cid := c.DefaultQuery("cid", "0")
	title := c.DefaultQuery("title", "")
	url := c.DefaultQuery("url", "")
	page := c.DefaultQuery("page", "1")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(_pageSize)
	MDB := model.DB().Table("common_apidoc")
	CDB := model.DB().Table("common_apidoc")
	if cid != "0" {
		MDB.Where("cid", cid)
		CDB.Where("cid", cid)
	}
	if title != "" {
		MDB.Where("title", "like", "%"+title+"%")
		CDB.Where("title", "like", "%"+title+"%")
	}
	if url != "" {
		MDB.Where("url", "like", "%"+url+"%")
		CDB.Where("url", "like", "%"+url+"%")
	}
	list, err := MDB.Limit(pageSize).Page(pageNo).Order("id desc").Get()
	if err != nil {
		results.Failed(c, err.Error(), nil)
	} else {
		for _, val := range list {
			groupdata, _ := model.DB().Table("common_apidoc_group a").
				Join("common_apidoc_type t", "t.id", "=", "a.type_id").
				Where("a.id", val["cid"]).Fields("a.name,a.type_id,t.model_name").First()
			val["groupname"] = groupdata["name"]
			val["type_id"] = groupdata["type_id"]
			if gf.InterfaceToInt(val["apicode_type"]) == 2 {
				val["url"] = fmt.Sprintf("/%v%v", groupdata["model_name"], val["url"])
			}
		}
		var totalCount int64
		totalCount, _ = CDB.Count()
		results.Success(c, "获取数据列表", map[string]interface{}{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}, nil)
	}
}

// 获取分组
func (api *Devapi) Get_group(c *gin.Context) {
	list, _ := model.DB().Table("common_apidoc_group").Fields("id,pid,name").Order("id asc").Get()
	list = gf.GetMenuChildrenArray(list, 0, "pid")
	if list == nil {
		list = make([]gform.Data, 0)
	}
	results.Success(c, "获取分组列表", list, nil)
}

// 保存
func (api *Devapi) Save(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	//当前用户
	var f_id float64 = 0
	if parameter["id"] != nil {
		f_id = parameter["id"].(float64)
	}
	parameter["createtime"] = time.Now().Unix()
	if f_id == 0 {
		delete(parameter, "id")
		addId, err := model.DB().Table("common_apidoc").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		delete(parameter, "groupname")
		delete(parameter, "type_id")
		res, err := model.DB().Table("common_apidoc").
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
func (api *Devapi) UpStatus(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	res2, err := model.DB().Table("common_apidoc").Where("id", parameter["id"]).Data(map[string]interface{}{"status": parameter["status"]}).Update()
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
func (api *Devapi) Del(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	res2, err := model.DB().Table("common_apidoc").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		results.Success(c, "删除成功！", res2, nil)
	}
}

// 获取数据库字段
func (api *Devapi) Get_DBField(c *gin.Context) {
	tablename := c.DefaultQuery("tablename", "")
	if tablename == "" {
		results.Failed(c, "请传数据表名称", nil)
	} else {
		tablename_arr := strings.Split(tablename, ",")
		//获取数据库名
		var dielddata_list []map[string]interface{}
		for _, Val := range tablename_arr {
			dielddata, _ := model.DB().Query("select COLUMN_NAME,COLUMN_COMMENT,COLUMN_TYPE,DATA_TYPE,CHARACTER_MAXIMUM_LENGTH,IS_NULLABLE,COLUMN_DEFAULT,NUMERIC_PRECISION from information_schema.columns where TABLE_SCHEMA='" + global.App.Config.DBconf.Database + "' AND TABLE_NAME='" + Val + "'")
			for _, data := range dielddata {
				data["tablename"] = Val
				dielddata_list = append(dielddata_list, data)
			}
		}
		results.Success(c, "获取数据库字段", dielddata_list, tablename)
	}
}

// 获取锁数据表
func (api *Devapi) Get_tables(c *gin.Context) {
	tablelist, _ := model.DB().Query("select TABLE_NAME,TABLE_COMMENT from information_schema.tables where table_schema = '" + global.App.Config.DBconf.Database + "'")
	var talbe_list []interface{}
	for _, Val := range tablelist {
		if !strings.Contains(fmt.Sprintf("%v", Val["TABLE_NAME"]), "admin_") && !strings.Contains(fmt.Sprintf("%v", Val["TABLE_NAME"]), "login_") && Val["TABLE_NAME"] != "attachment" {
			talbe_list = append(talbe_list, map[string]interface{}{"name": Val["TABLE_NAME"], "title": Val["TABLE_COMMENT"]})
		}
	}
	results.Success(c, "获取锁数据表", talbe_list, nil)
}

// 获取所有路由列表
func (api *Devapi) Get_routes(c *gin.Context) {
	filePath := "runtime/app/routers.txt"
	list := gf.ReaderFileByline(filePath)
	results.Success(c, "获取所有路由列表", list, nil)
}

// 获取模块列表
func (api *Devapi) GetModel(c *gin.Context) {
	var list []string
	files, err := os.ReadDir("./app")
	if err != nil {
		results.Failed(c, "获取目录错误", err)
		return
	}
	for _, file := range files {
		if file.IsDir() {
			list = append(list, file.Name())
		}
	}
	results.Success(c, "获取模块目录列表", list, nil)
}

// 获取指定数据库表的数据
func (api *Devapi) Get_tablelist(c *gin.Context) {
	tablename := c.DefaultQuery("tablename", "")
	pageSize := c.DefaultQuery("pageSize", "10")
	if tablename == "" {
		results.Failed(c, "请传数据库名称", nil)
	} else {
		seachword := c.DefaultQuery("seachword", "")
		MDB := model.DB().Table(tablename)
		if seachword != "" {
			MDB.Where("name", "like", "%"+seachword+"%")
		}
		list, err := MDB.Limit(gf.InterfaceToInt(pageSize)).Order("id desc").Get()
		if err != nil {
			results.Failed(c, err.Error(), nil)
		} else {
			results.Success(c, "获取指定数据库表的数据", list, nil)
		}
	}
}
