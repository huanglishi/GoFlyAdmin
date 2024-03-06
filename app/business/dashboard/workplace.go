package dashboard

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
	"time"

	"gofly/utils/gform"

	"github.com/gin-gonic/gin"
)

// 用于自动注册路由
type Workplace struct {
}

// 初始化生成路由
func init() {
	fpath := Workplace{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 统计总数据
func (api *Workplace) Get_statistical(c *gin.Context) {
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	article_count, _ := model.DB().Table("business_website_article_content").Where("businessID", user.BusinessID).Count("*")
	leavemessage_count, err := model.DB().Table("business_website_leavemessage").Where("businessID", user.BusinessID).Count("*")
	visit_record, _ := model.DB().Table("business_website_visit_record").Where("businessID", user.BusinessID).Count("*")
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	yesterday := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.Local)
	yse_star_time := gf.StringTimestamp(yesterday.Format("2006-01-02")+" 00:00", "datetime")
	yse_end_time := gf.StringTimestamp(yesterday.Format("2006-01-02")+" 23:59", "datetime")
	visit_record_yesterday, _ := model.DB().Table("business_website_visit_record").Where("businessID", user.BusinessID).WhereBetween("createtime", []interface{}{yse_star_time, yse_end_time}).Count("*")
	today_star_time := gf.StringTimestamp(today.Format("2006-01-02")+" 00:00", "datetime")
	today_end_time := gf.StringTimestamp(today.Format("2006-01-02")+" 23:59", "datetime")
	visit_record_today, _ := model.DB().Table("business_website_visit_record").Where("businessID", user.BusinessID).WhereBetween("createtime", []interface{}{today_star_time, today_end_time}).Count("*")
	var visit_ratio float64 = 0
	if visit_record_yesterday > 0 {
		num1, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64((visit_record_today-visit_record_yesterday))/float64(visit_record_yesterday)), 64)
		visit_ratio = num1 * 100
		if visit_ratio < 0 {
			visit_ratio = 0
		}
	}
	if err != nil {
		results.Failed(c, "统计访问数据失败", err)
	} else {
		results.Success(c, "统计总数据", map[string]interface{}{
			"article_count":          article_count,
			"leavemessage_count":     leavemessage_count,
			"visit_record_yesterday": visit_record_yesterday,
			"visit_record_today":     visit_record_today,
			"visit_record":           gf.InterfaceToInt(visit_record),
			"visit_ratio":            visit_ratio}, nil)
	}
}

// 统计热门文章
func (api *Workplace) Get_popular(c *gin.Context) {
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	MDB := model.DB().Table("business_website_article_content").Where("status", 0).Where("businessID", user.BusinessID)
	list, err := MDB.Fields("id,title,visits,createtime").Order("visits desc").Limit(6).Get()
	if err != nil {
		results.Failed(c, "统计热门文章失败", err)
	} else {
		if list == nil {
			list = make([]gform.Data, 0) //赋空值
		}
		for _, val := range list {
			now := time.Now()
			today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
			yesterday := time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.Local)
			yse_star_time := gf.StringTimestamp(yesterday.Format("2006-01-02")+" 00:00", "datetime")
			yse_end_time := gf.StringTimestamp(yesterday.Format("2006-01-02")+" 23:59", "datetime")
			visit_record_yesterday, _ := model.DB().Table("business_website_visit_record").Where("article_id", val["id"]).WhereBetween("createtime", []interface{}{yse_star_time, yse_end_time}).Count("*")
			today_star_time := gf.StringTimestamp(today.Format("2006-01-02")+" 00:00", "datetime")
			today_end_time := gf.StringTimestamp(today.Format("2006-01-02")+" 23:59", "datetime")
			visit_record_today, _ := model.DB().Table("business_website_visit_record").Where("article_id", val["id"]).WhereBetween("createtime", []interface{}{today_star_time, today_end_time}).Count("*")
			var visit_ratio float64 = 0
			if visit_record_yesterday > 0 {
				num1, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64((visit_record_today-visit_record_yesterday))/float64(visit_record_yesterday)), 64)
				visit_ratio = num1 * 100
				if visit_ratio < 0 {
					visit_ratio = 0
				}
			}
			val["visit_ratio"] = visit_ratio
		}
		results.Success(c, "统计热门文档", list, nil)
	}
}

// 获取进几天的访问记录
func (api *Workplace) Get_visitlist(c *gin.Context) {
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	qustr := fmt.Sprintf("%v", user.BusinessID)
	list, _ := model.DB().Query("select  FROM_UNIXTIME(createtime,'%Y-%m-%d')x,COUNT(*) as y from business_website_visit_record where businessID=" + qustr + " group by x order by x desc limit 10;")
	results.Success(c, "访问记录", list, nil)
}

// 获取网站常用统计数量
func (api *Workplace) Get_siteCount(c *gin.Context) {
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	site_count, _ := model.DB().Table("business_website_site").Where("businessID", user.BusinessID).Where("step", 2).Count("*")
	leavemessage_count, _ := model.DB().Table("business_website_leavemessage").Where("businessID", user.BusinessID).Where("status", 0).Count("*")
	results.Success(c, "获取网站常用统计数量数据", map[string]interface{}{
		"site_count":         site_count,
		"leavemessage_count": leavemessage_count}, nil)
}

// 获取公告信息
func (api *Workplace) Get_message(c *gin.Context) {
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	list, _ := model.DB().Table("common_message").WhereIn("usertype", []interface{}{0, 2}).Where("touid", 0).OrWhere("touid", user.ID).Limit(5).Fields("id,type,title,path,isread,createtime").Get()
	if list == nil {
		list = make([]gform.Data, 0)
	}
	countnum, _ := model.DB().Table("common_message").WhereIn("usertype", []interface{}{0, 2}).Where("touid", 0).OrWhere("touid", user.ID).Count("*")
	results.Success(c, "获取公告信息", map[string]interface{}{"list": list, "count": countnum}, nil)
}

// 获取公告详情
func (api *Workplace) Get_msmContent(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	data, _ := model.DB().Table("common_message").Where("id", id).Limit(5).Value("content")
	model.DB().Table("common_message").Where("id", id).Data(map[string]interface{}{"isread": 1}).Update()
	results.Success(c, "获取公告信息", data, nil)
}

// 1获取快捷操作
func (api *Workplace) Get_quick(c *gin.Context) {
	getuser, _ := c.Get("user") //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	list, err := model.DB().Table("business_home_quickop").Where("businessID", user.BusinessID).OrWhere("is_common", 1).Fields("id,uid,path_url,name,icon,type,is_common,weigh").Order("weigh asc,id asc").Get()
	if err != nil {
		results.Failed(c, "获取快捷操作失败", err)
	} else {
		results.Success(c, "获取快捷操作数据", list, nil)
	}
}

// 3保存快捷操作
func (api *Workplace) SaveQuick(c *gin.Context) {
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
		parameter["uid"] = user.ID
		parameter["businessID"] = user.BusinessID
		addId, err := model.DB().Table("business_home_quickop").Data(parameter).InsertGetId()
		if err != nil {
			results.Failed(c, "添加失败", err)
		} else {
			if addId != 0 {
				model.DB().Table("business_home_quickop").
					Data(map[string]interface{}{"weigh": addId}).
					Where("id", addId).
					Update()
			}
			results.Success(c, "添加成功！", addId, nil)
		}
	} else {
		res, err := model.DB().Table("business_home_quickop").
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

// 3删除快捷操作
func (api *Workplace) Del_quick(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	res2, err := model.DB().Table("business_home_quickop").Where("id", parameter["id"]).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		results.Success(c, "删除成功！", res2, nil)
	}
}
