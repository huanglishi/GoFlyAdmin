package datacenter

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils/gf"
	"gofly/utils/results"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// 用于自动注册路由
type Upfile struct{}

func init() {
	fpath := Upfile{}
	gf.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 上传图片
func (api *Upfile) UploadFile(c *gin.Context) {
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	//判断存储空间是否超出
	var usesize interface{}
	var fileSize interface{}
	usesize, _ = model.DB().Table("business_attachment").Where("businessID", user.BusinessID).Where("type", 0).Sum("filesize")
	if usesize == nil {
		usesize = 0
	}
	fileSize, _ = model.DB().Table("business_account").Where("id", user.BusinessID).Value("fileSize")
	if fileSize == nil {
		fileSize = 0
	}
	if gf.InterfaceToInt(usesize) >= gf.InterfaceToInt(fileSize) {
		results.Failed(c, "您的存储空间已满,请您先去购买存储空间！", nil)
		return
	}
	// 单个文件
	pid := c.DefaultPostForm("pid", "")
	filetype := c.DefaultPostForm("filetype", "image") //文件类型
	file, err := c.FormFile("file")
	if err != nil {
		results.Failed(c, "获取数据失败，", err)
		return
	}
	nowTime := time.Now().Unix() //当前时间
	rooturl, _ := model.DB().Table("common_config").Where("keyname", "rooturl").Value("keyvalue")
	//判断文件是否已经传过
	fileContent, _ := file.Open()
	var byteContainer []byte
	byteContainer = make([]byte, 1000000)
	fileContent.Read(byteContainer)
	m_d5 := md5.New()
	m_d5.Write(byteContainer)
	sha1_str := hex.EncodeToString(m_d5.Sum(nil))
	//查找该用户是否传过
	attachment, _ := model.DB().Table("business_attachment").Where("businessID", user.BusinessID).
		Where("sha1", sha1_str).Fields("id,pid,name,title,type,url,filesize,mimetype,cover_url").First()
	if attachment != nil { //文件是否已经存在
		//更新到最前面
		maxId, _ := model.DB().Table("business_attachment").Where("businessID", user.BusinessID).Order("weigh desc").Value("id")
		if maxId != nil {
			model.DB().Table("business_attachment").Data(map[string]interface{}{"weigh": maxId.(int64) + 1, "pid": pid}).Where("id", attachment["id"]).Update()
		}
		attachment["url"] = rooturl.(string) + attachment["url"].(string)
		results.Success(c, "文件已上传", attachment, nil)
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
		name_str := md5Str(fmt.Sprintf("%v%s", nowTime, filename_arr[0]))      //组装文件保存名字
		file_Filename := fmt.Sprintf("%s%s%s", name_str, ".", filename_arr[1]) //文件加.后缀
		path := file_path + file_Filename
		// 上传文件到指定的目录
		err = c.SaveUploadedFile(file, path)
		if err != nil { //上传失败
			c.JSON(200, gin.H{
				"uid":      sha1_str,
				"name":     file.Filename,
				"status":   "error",
				"response": "上传失败",
				"time":     nowTime,
			})
		} else { //上传成功
			//保存数据
			dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
			var ftype int64 = 0
			var cover_url string = ""
			if filetype == "video" {
				ftype = 2
				videopath := fmt.Sprintf("./%s", path)
				pathroot := strings.Split(path, ".")
				imgpath := fmt.Sprintf("./%s", pathroot[0])
				fname, err := GetSnapshot(videopath, imgpath, 1)
				if err == nil {
					cover_url = fname
				}
			}
			Insertdata := map[string]interface{}{
				"accountID":  user.Accountid,
				"businessID": user.BusinessID,
				"type":       ftype,
				"pid":        pid,
				"sha1":       sha1_str,
				"title":      filename_arr[0],
				"name":       file.Filename,
				"url":        path,
				"cover_url":  cover_url, //视频封面
				"storage":    dir + "/" + path,
				"createtime": nowTime,
				"filesize":   file.Size,
				"mimetype":   file.Header["Content-Type"][0],
			}
			//保存数据
			file_id, _ := model.DB().Table("business_attachment").Data(Insertdata).InsertGetId()
			//更新排序
			model.DB().Table("business_attachment").Data(map[string]interface{}{"weigh": file_id}).Where("id", file_id).Update()
			//返回数据
			getdata, _ := model.DB().Table("business_attachment").Where("id", file_id).Fields("id,pid,name,title,type,url,filesize,mimetype,cover_url").First()
			getdata["url"] = rooturl.(string) + getdata["url"].(string)
			results.Success(c, "上传成功", getdata, nil)
		}
	}
}

// 视频截取第一帧作为封面
func invokeFfmpeg(urlpath string, path string, ffmpegPath string) string {
	fmt.Println("urlpath:", urlpath)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(14400000)*time.Millisecond)
	cmd := exec.CommandContext(ctx, ffmpegPath,
		"-loglevel", "error",
		"-i", urlpath,
		"-ss", "1",
		"-f", "image2",
		"./"+path+".jpg")
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	defer cancel()
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	var outputerror string
	err := cmd.Run()
	if err != nil {
		outputerror += fmt.Sprintf("cmderr:%v;", err)
	}
	if stderr.Len() != 0 {
		outputerror += fmt.Sprintf("stderr:%v;", stderr.String())
	}
	if ctx.Err() != nil {
		outputerror += fmt.Sprintf("ctxerr:%v;", ctx.Err())
	}
	// fmt.Println("invokeFfmpeg err:", outputerror)
	return path + ".jpg"
}

// md5加密
func md5Str(origin string) string {
	m := md5.New()
	m.Write([]byte(origin))
	return hex.EncodeToString(m.Sum(nil))
}
