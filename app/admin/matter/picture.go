package matter

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
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

// 用于自动注册路由
type Picture struct{}

func init() {
	fpath := Picture{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 获取列表
func (api *Picture) Get_list(c *gin.Context) {
	title := c.DefaultQuery("title", "")
	status := c.DefaultQuery("status", "")
	createdTime := c.DefaultQuery("createdTime", "")
	page := c.DefaultQuery("page", "1")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageNo, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(_pageSize)
	MDB := model.DB().Table("common_picture")
	MDBC := model.DB().Table("common_picture")
	if title != "" {
		MDB.Where("title", "like", "%"+title+"%")
		MDBC.Where("title", "like", "%"+title+"%")
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
	list, err := MDB.Limit(pageSize).Page(pageNo).Order("id desc").Get()
	if err != nil {
		results.Failed(c, err.Error(), nil)
	} else {
		rooturl, _ := model.DB().Table("common_config").Where("keyname", "rooturl").Value("keyvalue")
		for _, val := range list {
			if _, ok := val["url"]; ok && val["url"] != "" && !strings.Contains(val["url"].(string), "http") && rooturl != nil {
				val["url"] = rooturl.(string) + val["url"].(string)
			}
			if _, ok := val["cid"]; ok {
				catename, _ := model.DB().Table("common_picture_cate").Where("id", val["cid"]).Value("name")
				val["catename"] = catename
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
func (api *Picture) Save(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	delete(parameter, "catename")
	parameter["createtime"] = time.Now().Unix()
	res, err := model.DB().Table("common_picture").
		Data(parameter).
		Where("id", parameter["id"]).
		Update()
	if err != nil {
		results.Failed(c, "更新失败", err)
	} else {
		results.Success(c, "更新成功！", res, nil)
	}
}

// 更新状态
func (api *Picture) UpStatus(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	res2, err := model.DB().Table("common_picture").Where("id", parameter["id"]).Data(map[string]interface{}{"status": parameter["status"]}).Update()
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
func (api *Picture) Del(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	ids := parameter["ids"]
	file_list, _ := model.DB().Table("common_picture").WhereIn("id", ids.([]interface{})).Pluck("url")
	res2, err := model.DB().Table("common_picture").WhereIn("id", ids.([]interface{})).Delete()
	if err != nil {
		results.Failed(c, "删除失败", err)
	} else {
		if file_list != nil {
			gf.Del_file(file_list.([]interface{}))
		}
		results.Success(c, "删除成功！", res2, nil)
	}
}

// 上传图片
func (api *Picture) UploadFile(context *gin.Context) {
	// 单个文件
	cid := context.DefaultPostForm("cid", "")
	typeId := context.DefaultPostForm("type", "")
	Id := context.DefaultPostForm("id", "0")
	file, err := context.FormFile("file")
	if err != nil {
		results.Failed(context, "获取数据失败，", err)
		return
	}
	nowTime := time.Now().Unix() //当前时间
	getuser, _ := context.Get("user")
	user := getuser.(*middleware.UserClaims)
	//判断文件是否已经传过
	fileContent, _ := file.Open()
	var byteContainer []byte
	byteContainer = make([]byte, 1000000)
	fileContent.Read(byteContainer)
	m_d5 := md5.New()
	m_d5.Write(byteContainer)
	sha1_str := hex.EncodeToString(m_d5.Sum(nil))
	rooturl, _ := model.DB().Table("common_config").Where("keyname", "rooturl").Value("keyvalue")
	//查找该用户是否传过
	attachment, _ := model.DB().Table("common_picture").Where("uid", user.ID).
		Where("sha1", sha1_str).Fields("id,name,title,url,filesize,mimetype,storage").First()
	if attachment != nil { //文件是否已经存在
		//更新到最前面
		var nid interface{}
		if Id != "0" {
			model.DB().Table("common_picture").Data(map[string]interface{}{"title": attachment["title"], "name": attachment["name"], "url": attachment["url"]}).Where("id", Id).Update()
			nid = Id
		} else {
			delete(attachment, "id")
			attachment["cid"] = cid
			attachment["type"] = typeId
			file_id, _ := model.DB().Table("common_picture").Data(attachment).InsertGetId()
			nid = file_id
			//更新排序
			model.DB().Table("common_picture").Data(map[string]interface{}{"weigh": file_id}).Where("id", file_id).Update()
		}
		results.Success(context, "文件已上传", map[string]interface{}{"id": nid, "title": attachment["title"], "url": attachment["url"]}, nil)
	} else {
		file_path := fmt.Sprintf("%s%s%s", "resource/uploads/", time.Now().Format("20060102"), "/")
		//如果没有filepath文件目录就创建一个
		if _, err := os.Stat(file_path); err != nil {
			if !os.IsExist(err) {
				os.MkdirAll(file_path, os.ModePerm)
			}
		}
		//上传到的路径
		filename_arr := strings.Split(file.Filename, ".")
		//重新名片-lunix系统不支持中文
		name_str := gf.Md5Str(fmt.Sprintf("%v%s", nowTime, filename_arr[0]))   //组装文件保存名字
		file_Filename := fmt.Sprintf("%s%s%s", name_str, ".", filename_arr[1]) //文件加.后缀
		path := file_path + file_Filename
		// 上传文件到指定的目录
		err = context.SaveUploadedFile(file, path)
		if err != nil { //上传失败
			context.JSON(200, gin.H{
				"uid":      sha1_str,
				"name":     file.Filename,
				"status":   "error",
				"response": "上传失败",
				"time":     nowTime,
			})
		} else { //上传成功
			//保存数据
			dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
			Insertdata := map[string]interface{}{
				"uid":        user.ID,
				"type":       typeId,
				"cid":        cid,
				"sha1":       sha1_str,
				"title":      filename_arr[0],
				"name":       file.Filename,
				"url":        path,
				"storage":    dir + strings.Replace(path, "/", "\\", -1),
				"createtime": nowTime,
				"filesize":   file.Size,
				"mimetype":   file.Header["Content-Type"][0],
			}
			//保存数据
			var nid interface{}
			if Id != "0" {
				model.DB().Table("common_picture").Data(Insertdata).Where("id", Id).Update()
				nid = Id
			} else {
				file_id, _ := model.DB().Table("common_picture").Data(Insertdata).InsertGetId()
				nid = file_id
				//更新排序
				model.DB().Table("common_picture").Data(map[string]interface{}{"weigh": file_id}).Where("id", file_id).Update()
			}
			//返回数据
			results.Success(context, "上传成功1", map[string]interface{}{"id": nid, "title": Insertdata["title"], "url": rooturl.(string) + Insertdata["url"].(string)}, nil)
		}
	}
	context.Abort()
}
