package common

import (
	"bytes"
	"context"
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
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	gf.Register(&Uploadfile{}, reflect.TypeOf(Uploadfile{}).PkgPath())
}

type Uploadfile struct {
}

// 1.上传单文件
func (api *Uploadfile) Onefile(c *gin.Context) {
	cid := c.DefaultPostForm("cid", "1")
	// 单个文件
	file, err := c.FormFile("file")
	if err != nil {
		results.Failed(c, "获取数据失败，", err)
		return
	}
	nowTime := time.Now().Unix() //当前时间
	getuser, _ := c.Get("user")  //取值 实现了跨中间件取值
	user := getuser.(*middleware.UserClaims)
	//时间查询-获取当天时间
	day_time := time.Now().Format("2006-01-02")
	//文件唯一性
	file_uniname := fmt.Sprintf("%s%s%v", file.Filename, day_time, user.ID)
	sha1_str := gf.Md5(file_uniname)
	//开始
	day_star, _ := time.Parse("2006-01-02 15:04:05", day_time+" 00:00:00")
	day_star_times := day_star.Unix() //时间戳
	//结束
	day_end, _ := time.Parse("2006-01-02 15:04:05", day_time+" 23:59:59")
	day_end_times := day_end.Unix() //时间戳
	attachment, _ := model.DB().Table("attachment").Where("uid", user.ID).
		WhereBetween("uploadtime", []interface{}{day_star_times, day_end_times}).
		Where("sha1", sha1_str).Fields("id,title,url").First()
	if attachment != nil { //文件是否已经存在
		c.JSON(200, gin.H{
			"id":       attachment["id"],
			"uid":      sha1_str,
			"name":     attachment["name"],
			"status":   "done",
			"url":      attachment["url"],
			"response": "文件已上传",
			"time":     nowTime,
		})
		c.Abort()
		return
	}
	file_path := fmt.Sprintf("%s%s%s", "resource/uploads/", time.Now().Format("20060102"), "/")
	//如果没有filepath文件目录就创建一个
	if _, err := os.Stat(file_path); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(file_path, os.ModePerm)
		}
	}
	//上传到的路径
	filename_arr := strings.Split(file.Filename, ".")
	name_str := gf.Md5(fmt.Sprintf("%v%s", nowTime, filename_arr[0])) //组装文件保存名字
	//path := 'resource/uploads/20060102150405test.xlsx'
	file_Filename := fmt.Sprintf("%s%s%s", name_str, ".", filename_arr[1]) //文件加.后缀
	path := file_path + file_Filename
	// fmt.Println("path1:", path) //路径+文件名上传
	// 上传文件到指定的目录
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.JSON(200, gin.H{
			"uid":      sha1_str,
			"name":     file.Filename,
			"status":   "error",
			"response": "上传失败",
			"time":     nowTime,
		})
	} else {
		//保存数据
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

		//判断是否是视频-获取封面
		var cover_url string = ""
		if strings.Contains(file.Header["Content-Type"][0], "video/mp4") {
			//封面路径
			// fmt.Println("dir:", path)
			var ffmpegPath string = "ffmpeg"
			vurlpath := dir + strings.Replace("/"+path, "/", "\\", -1)
			cover_url = getLastFrame(vurlpath, file_path+name_str, ffmpegPath)
		}

		Insertdata := map[string]interface{}{
			"accountID":  user.Accountid,
			"cid":        cid,
			"uid":        user.ID,
			"sha1":       sha1_str,
			"title":      filename_arr[0],
			"name":       file.Filename,
			"url":        path,
			"cover_url":  cover_url,
			"storage":    dir + strings.Replace(path, "/", "\\", -1),
			"uploadtime": nowTime,
			"updatetime": nowTime,
			"filesize":   file.Size,
			"mimetype":   file.Header["Content-Type"][0],
		}
		file_id, _ := model.DB().Table("attachment").Data(Insertdata).InsertGetId()
		c.JSON(200, gin.H{
			"id":        file_id,
			"uid":       sha1_str,
			"name":      file.Filename,
			"status":    "done",
			"url":       path,
			"thumb":     path,
			"cover_url": cover_url,
			"response":  "上传成功",
			// "file":     file.Header,
			"time": nowTime,
		})
	}
}

// 2. 获取视频中最后一帧的图片 url=视频地址,path=图片地址
func getLastFrame(url string, path string, ffmpegPath string) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(50000)*time.Millisecond)
	cmd := exec.CommandContext(ctx, ffmpegPath,
		"-loglevel", "error",
		"-y",
		"-ss", "13",
		"-t", "1",
		"-i", url,
		"-vframes", "1",
		path+".jpg")
	defer cancel()
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	var outputerror string
	err := cmd.Run()
	if err != nil {
		outputerror += fmt.Sprintf("lastframecmd—err1:%v;", err)
	}
	if stderr.Len() != 0 {
		outputerror += fmt.Sprintf("lastframestd—err2:%v;", stderr.String())
	}
	if ctx.Err() != nil {
		outputerror += fmt.Sprintf("lastframectx—err3:%v;", ctx.Err())
	}
	return path + ".jpg"
}

// 3.显示图片
func (api *Uploadfile) Get_image(c *gin.Context) {
	imageName := c.Query("url")
	imgrul := strings.Split(imageName, "?")
	c.File(imgrul[0])
}

// 4.显示图片base64
func (api *Uploadfile) Get_imagebase(c *gin.Context) {
	imageName := c.Query("url")
	file, _ := os.ReadFile(imageName)
	c.Writer.WriteString(string(file))
}
