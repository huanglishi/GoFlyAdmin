package datacenter

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

// 用于自动注册路由
type Attachment struct{}

func init() {
	fpath := Attachment{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取列表
func (api *Attachment) Get_list(c *gin.Context) {
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	createdTime := c.DefaultQuery("createdTime", "")
	name := c.DefaultQuery("name", "")
	status := c.DefaultQuery("status", "")
	page := c.DefaultQuery("page", "1")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(_pageSize)
	MDB := model.DB().Table("business_attachment").Where("businessID", user.BusinessID).Where("type", "!=", 1)
	MDBC := model.DB().Table("business_attachment").Where("businessID", user.BusinessID).Where("type", "!=", 1)
	if name != "" {
		MDB.Where("name", "like", "%"+name+"%")
		MDBC.Where("name", "like", "%"+name+"%")
	}
	if status != "" {
		MDB.Where("status", status)
		MDBC.Where("status", status)
	}
	if createdTime != "" {
		datetime_arr := strings.Split(createdTime, ",")
		star_time := gf.StringTimestamp(datetime_arr[0]+" 00:00", "datetime")
		end_time := gf.StringTimestamp(datetime_arr[1]+" 23:59", "datetime")
		MDB.WhereBetween("createtime", []interface{}{star_time, end_time})
		MDBC.WhereBetween("createtime", []interface{}{star_time, end_time})
	}
	list, err := MDB.Limit(pageSize).Page(pageNo).Fields("id,url,type,title,mimetype,cover_url,createtime,pid").Order("id desc").Get()
	if err != nil {
		results.Failed(c, err.Error(), nil)
	} else {
		rooturl, _ := model.DB().Table("common_config").Where("keyname", "rooturl").Value("keyvalue")
		for _, val := range list {
			if _, ok := val["url"]; ok && val["url"] != "" && !strings.Contains(val["url"].(string), "http") && rooturl != nil {
				val["url"] = rooturl.(string) + val["url"].(string)
			}
			if _, ok := val["cover_url"]; ok && val["cover_url"] != "" && !strings.Contains(val["cover_url"].(string), "http") && rooturl != nil {
				val["cover_url"] = rooturl.(string) + val["cover_url"].(string)
			}
		}

		//统计
		var allnumber int64
		var usesize interface{}
		var fileSize interface{}
		allnumber, _ = model.DB().Table("business_attachment").Where("businessID", user.BusinessID).Where("type", 0).Count()
		usesize, _ = model.DB().Table("business_attachment").Where("businessID", user.BusinessID).Where("type", 0).Sum("filesize")
		fileSize, _ = model.DB().Table("business_account").Where("id", user.BusinessID).Value("fileSize")
		datainfo := map[string]interface{}{"allnumber": allnumber, "usesize": usesize, "fileSize": fileSize}
		var totalCount int64
		totalCount, _ = MDBC.Count("*")
		results.Success(c, "获取全部列表", map[string]interface{}{
			"datainfo": datainfo,
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}, nil)
	}
}

// 获取分类列表
func (api *Attachment) Get_pictureCate(c *gin.Context) {
	list, err := model.DB().Table("common_picture_cate").Where("status", 0).Fields("id,name,type").Order("weigh desc,id desc").Get()
	if err != nil {
		results.Failed(c, err.Error(), nil)
	} else {
		results.Success(c, "获取选择列表", list, nil)
	}
}

// 获取图片库列表
func (api *Attachment) Get_picture(c *gin.Context) {
	createdTime := c.DefaultQuery("createdTime", "")
	cid := c.DefaultQuery("cid", "0")
	types := c.DefaultQuery("type", "0")
	title := c.DefaultQuery("searchword", "")
	page := c.DefaultQuery("page", "1")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(_pageSize)
	MDB := model.DB().Table("common_picture").
		Fields("id,cid,url,type,title,mimetype,cover_url,createtime").Where("status", 0).Where("type", types)
	MDBC := model.DB().Table("common_picture").Where("status", 0).Where("type", types)
	if cid != "0" {
		MDB.Where("cid", cid)
		MDBC.Where("cid", cid)
	}
	if title != "" {
		MDB.Where("title", "like", "%"+title+"%")
		MDBC.Where("title", "like", "%"+title+"%")
	}
	if createdTime != "" {
		datetime_arr := strings.Split(createdTime, ",")
		star_time := gf.StringTimestamp(datetime_arr[0]+" 00:00", "datetime")
		end_time := gf.StringTimestamp(datetime_arr[1]+" 23:59", "datetime")
		MDB.WhereBetween("createtime", []interface{}{star_time, end_time})
		MDBC.WhereBetween("createtime", []interface{}{star_time, end_time})
	}
	list, err := MDB.Limit(pageSize).Page(pageNo).Order("id desc").Get()
	if err != nil {
		results.Failed(c, err.Error(), nil)
	} else {
		rooturl, _ := model.DB().Table("common_config").Where("keyname", "rooturl").Value("keyvalue")
		for _, val := range list {
			if _, ok := val["image"]; ok && val["image"] != "" && !strings.Contains(val["image"].(string), "http") && rooturl != nil {
				val["image"] = rooturl.(string) + val["image"].(string)
			}
		}
		var totalCount int64
		totalCount, _ = MDBC.Count("*")
		results.Success(c, "获取全部列表", map[string]interface{}{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}, nil)
	}
}

// 保存
func (api *Attachment) Save(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	var f_id float64 = 0
	if parameter["id"] != nil {
		f_id = parameter["id"].(float64)
	}
	if f_id == 0 {
		delete(parameter, "id")
		parameter["createtime"] = time.Now().Unix()
		parameter["businessID"] = user.BusinessID
		getcount, _ := model.DB().Table("business_attachment").Where("businessID", user.BusinessID).Where("pid", parameter["pid"]).Where("title", "like", fmt.Sprintf("%s%v%s", "%", parameter["title"], "%")).Count()
		parameter["title"] = fmt.Sprintf("%s%v", parameter["title"], gf.InterfaceToInt(getcount)+1)
		addId, err := model.DB().Table("business_attachment").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			//更新排序
			model.DB().Table("business_attachment").Data(map[string]interface{}{"weigh": addId}).Where("id", addId).Update()
			getdata, _ := model.DB().Table("business_attachment").Where("id", addId).Fields("id,pid,name,title,type,url,filesize,mimetype,storage").Get()
			results.Success(c, "添加成功！", getdata, nil)
		}
	} else {
		delete(parameter, "catename")
		res, err := model.DB().Table("business_attachment").
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

// 删除
func (api *Attachment) Del(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	file_list, _ := model.DB().Table("business_attachment").WhereIn("id", ids.([]interface{})).Where("type", "!=", 1).Pluck("url")
	res2, err := model.DB().Table("business_attachment").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		if file_list != nil {
			gf.Del_file(file_list.([]interface{}))
		}
		results.Success(c, "删除成功！", res2, nil)
	}
}

// 获取我的附件
func (api *Attachment) Get_myFiles(c *gin.Context) {
	searchword := c.DefaultQuery("searchword", "")
	filetype := c.DefaultQuery("filetype", "image")
	pid := c.DefaultQuery("pid", "0")
	//当前用户
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	whereMap := model.DB().Table("business_attachment").Where("businessID", user.BusinessID).Where("pid", pid)
	if searchword != "" {
		whereMap.Where("title", "like", "%"+searchword+"%")
	}
	if filetype == "video" {
		whereMap.WhereIn("type", []interface{}{1, 2})
	} else { //默认图片
		whereMap.WhereIn("type", []interface{}{0, 1})
	}
	list, err := whereMap.
		Fields("id,pid,name,title,type,url,filesize,mimetype,storage,cover_url,is_common").Order("type desc,weigh desc,id desc").Get()
	if err != nil {
		results.Failed(c, "加载数据失败", err)
	} else {
		rooturl, _ := model.DB().Table("common_config").Where("keyname", "rooturl").Value("keyvalue")
		for _, val := range list {
			if _, ok := val["url"]; ok && val["url"] != "" && !strings.Contains(val["url"].(string), "http") && rooturl != nil {
				val["url"] = rooturl.(string) + val["url"].(string)
			}
			if _, ok := val["cover_url"]; ok && val["cover_url"] != "" && !strings.Contains(val["cover_url"].(string), "http") && rooturl != nil {
				val["cover_url"] = rooturl.(string) + val["cover_url"].(string)
			}
		}
		common_lisr, _ := model.DB().Table("business_attachment").Where("pid", pid).Where("is_common", 1).Fields("id,pid,name,title,type,url,filesize,mimetype,storage,cover_url,is_common").Order("type desc,weigh desc,id desc").Get()
		if list != nil {
			list = append(common_lisr, list...)
		} else {
			list = common_lisr
		}
		var totalCount int64
		//获取目录菜单
		allids := getAllParentIds(pid)
		allids = append(allids, pid)
		dirmenu, _ := model.DB().Table("business_attachment").WhereIn("id", allids).Fields("id,pid,title").Get()
		results.Success(c, "获取附件列表", map[string]interface{}{
			"total":   totalCount,
			"dirmenu": dirmenu,
			"allids":  allids,
			"items":   list,
		}, nil)
	}
}

// 工具
func getAllParentIds(id interface{}) []interface{} {
	var parent_ids []interface{}
	parent_id, _ := model.DB().Table("business_attachment").Where("id", id).Value("pid")
	if parent_id != nil {
		parent_ids = append(parent_ids, parent_id)
		parent_ids = append(parent_ids, getAllParentIds(parent_id)...)
	}
	return parent_ids
}

// 更新图片目录
func (api *Attachment) UpImgPid(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	res2, err := model.DB().Table("business_attachment").Where("id", parameter["imgid"]).Data(map[string]interface{}{"pid": parameter["pid"]}).Update()
	if err != nil {
		results.Failed(c, "更新失败！", err)
	} else {
		msg := "更新目录成功！"
		if res2 == 0 {
			msg = "暂无目录更新"
		}
		results.Success(c, msg, res2, nil)
	}
}
