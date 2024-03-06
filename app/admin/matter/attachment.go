package matter

import (
	"encoding/json"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils/gf"
	"gofly/utils/results"
	"io"
	"reflect"
	"strconv"
	"strings"

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
	name := c.DefaultQuery("name", "")
	createdTime := c.DefaultQuery("createdTime", "")
	title := c.DefaultQuery("title", "")
	page := c.DefaultQuery("page", "1")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(_pageSize)
	MDB := model.DB().Table("attachment").
		Fields("id,url,imagewidth,imageheight,imagetype,filesize,mimetype,uploadtime,sha1,title,name,cover_url").Where("businessID", user.BusinessID)
	MDBC := model.DB().Table("attachment").Where("businessID", user.BusinessID)
	if name != "" {
		MDB.Where("name", "like", "%"+name+"%")
		MDBC.Where("name", "like", "%"+name+"%")
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
			if _, ok := val["url"]; ok && val["url"] != "" && !strings.Contains(val["url"].(string), "http") && rooturl != nil {
				val["url"] = rooturl.(string) + val["url"].(string)
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

// 删除
func (api *Attachment) Del(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	file_list, _ := model.DB().Table("attachment").WhereIn("id", ids.([]interface{})).Pluck("url")
	res2, err := model.DB().Table("attachment").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		if file_list != nil {
			gf.Del_file(file_list.([]interface{}))
		}
		results.Success(c, "删除成功！", res2, nil)
	}
}
