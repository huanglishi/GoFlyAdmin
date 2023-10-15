package common

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"fmt"
	"gofly/global"
	"gofly/model"
	"gofly/utils"
	"gofly/utils/results"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

/**
* 项目安装
 */
type Install struct {
}

func init() {
	fpath := Install{}
	utils.Register(&fpath, reflect.TypeOf(fpath).PkgPath())
}

// 安装页面
func (api *Install) Index(context *gin.Context) {
	_path, err := os.Getwd()
	if err != nil {
		results.Failed(context, "项目路径获取失败", nil)
		return
	}
	filePath := path.Join(_path, "/resource/staticfile/template/install.lock")

	if _, err := os.Stat(filePath); err == nil {
		context.HTML(http.StatusOK, "isinstall.html", gin.H{
			"title": "已经安装页面",
		})
	} else {
		context.HTML(http.StatusOK, "install.html", gin.H{
			"title": "安装页面",
			"path":  path.Join(_path, "../", "vue"), // 当前项目的绝对路径的上一级目录
		})
	}

}

// 安装
func (api *Install) Save(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	_path, err := os.Getwd() //获取当前路径
	if err != nil {
		results.Failed(c, "项目路径获取失败", nil)
		return
	}
	model.CreateDataBase(parameter["username"], parameter["password"], parameter["hostname"], parameter["hostport"], parameter["database"])
	//2.修改数据库配置
	cferr := upConfFieldData(_path, parameter)
	if cferr != nil {
		results.Failed(c, "修改数据库配置失败", nil)
		return
	}
	model.MyInit(2) //初始化数据
	//创建数据库
	//导入书库配置
	SqlPath := path.Join(_path, "/resource/staticfile/template/gofly_single.sql")
	sqls, sqlerr := os.ReadFile(SqlPath)
	if sqlerr != nil {
		results.Failed(c, "数据库文件不存在："+SqlPath, nil)
		return
	}
	sqlArr := strings.Split(string(sqls), ";")
	for _, sql := range sqlArr {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		model.ExecSql(sql)
	}
	//3.修改后台账号
	salt := time.Now().Unix()
	adminpass := fmt.Sprintf("%v%v", utils.Md5(parameter["adminPassword"].(string)), salt)
	businesspass := fmt.Sprintf("%v%v", utils.Md5(parameter["businessPassword"].(string)), salt)
	model.DB().Table("admin_user").Data(map[string]interface{}{"username": parameter["adminUsername"], "password": utils.Md5(adminpass), "salt": salt}).Where("id", 1).Update()
	model.DB().Table("business_user").Data(map[string]interface{}{"username": parameter["businessUsername"], "password": utils.Md5(businesspass), "salt": salt}).Where("id", 1).Update()
	//4.创建安装锁文件
	filePath := path.Join(_path, "/resource/staticfile/template/install.lock")
	os.Create(filePath)
	//5.安装前端页面
	if _, ok := parameter["vuepath"]; ok && parameter["vuepath"] != "" {
		parameter["vueobjroot"] = fmt.Sprintf("%v/gofly_business", parameter["vuepath"]) //更新前端路径
		//1 如果没有filepath文件目录就创建一个
		file_path := fmt.Sprintf("%v", parameter["vuepath"])
		if _, err := os.Stat(file_path); err != nil {
			if !os.IsExist(err) {
				os.MkdirAll(file_path, os.ModePerm)
			}
		}
		//2 复制前端文件到指定位置
		vuesoure_path := path.Join(_path, "/resource/staticfile/template/vuecode/")
		CopyDir(vuesoure_path, file_path)
		//3 解压文件
		business_vue_path := file_path + "/gofly_business.zip"
		admin_vue_path := file_path + "/gofly_admin.zip"
		Unzip(business_vue_path, file_path)
		Unzip(admin_vue_path, file_path)
		//删除zip文件
		os.RemoveAll(business_vue_path)
		os.RemoveAll(admin_vue_path)
	}

	results.Success(c, "安装成功,去前端刷新试试！", parameter, nil)
}

// 更新配置文件
func upConfFieldData(_path string, parameter map[string]interface{}) error {
	file_path := path.Join(_path, "/config/settings.yml")
	f, err := os.Open(file_path)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	var is_hose = false
	for {
		is_hose = false
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		for keys, Val := range parameter {
			if strings.Contains(string(a), keys) {
				is_hose = true
				datestr := strings.ReplaceAll(string(a), string(a), fmt.Sprintf("     %v: %v\n", keys, Val))
				result += datestr
			}
		}
		if !is_hose {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
	return nil
}

// DeCompress 解压文件 返回解压的目录
// zipFile 完整文件路径，dest文件目录
func DeCompress(zipFile, dest string) (string, error) {
	// 打开zip文件
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return "", err
	}
	defer func() {
		err := reader.Close()
		if err != nil {
			global.App.Log.Info(fmt.Sprintf("解压文件关闭失败: %v\n", err.Error()))
		}
	}()
	var (
		first string // 记录第一次的解压的名字
		order int    = 0
	)
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return "", err
		}
		filename := filepath.Join(dest, file.Name)
		//记录第一次的名字
		if order == 0 {
			first = filename
		}
		order += 1
		if file.FileInfo().IsDir() {
			err = os.MkdirAll(filename, 0755)
			if err != nil {
				return "", err
			}
		} else {
			w, err := os.Create(filename)
			if err != nil {
				return "", err
			}
			//defer w.Close()
			_, err = io.Copy(w, rc)
			if err != nil {
				return "", err
			}
			iErr := w.Close()
			if iErr != nil {
				global.App.Log.Info(fmt.Sprintf("[unzip]: close io %s\n", iErr.Error()))
			}
			fErr := rc.Close()
			if fErr != nil {
				global.App.Log.Info(fmt.Sprintf("[unzip]: close io %s\n", fErr.Error()))
			}
		}
	}
	return first, nil
}

// Unzip decompresses a zip file to specified directory.
// Note that the destination directory don't need to specify the trailing path separator.
// If the destination directory doesn't exist, it will be created automatically.
func Unzip(zipath, dir string) error {
	// Open zip file.
	reader, err := zip.OpenReader(zipath)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if err := unzipFile(file, dir); err != nil {
			return err
		}
	}
	return nil
}

func unzipFile(file *zip.File, dir string) error {
	// Prevent path traversal vulnerability.
	// Such as if the file name is "../../../path/to/file.txt" which will be cleaned to "path/to/file.txt".
	name := strings.TrimPrefix(filepath.Join(string(filepath.Separator), file.Name), string(filepath.Separator))
	filePath := path.Join(dir, name)

	// Create the directory of file.
	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	// Open the file.
	r, err := file.Open()
	if err != nil {
		return err
	}
	defer r.Close()

	// Create the file.
	w, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer w.Close()

	// Save the decompressed file content.
	_, err = io.Copy(w, r)
	return err
}

// 2复制整个文件夹下文件到另个文件夹 targetPath文件夹，destPath复制的文件
func CopyDir(targetPath string, destPath string) error {
	err := filepath.Walk(targetPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		destPath := filepath.Join(destPath, path[len(targetPath):])
		//如果是个文件夹则创建这个文件夹
		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}
		//如果是文件则生成这个文件
		return copyFile(path, destPath)

	})
	return err
}

// 复制单个文件
func copyFile(srcFile, destFile string) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer src.Close()
	//创建复制的文件
	dest, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer dest.Close()
	//复制内容到文件
	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}
	//让复制的文件将内容存到硬盘而不是缓存
	err = dest.Sync()
	if err != nil {
		return err
	}

	return nil
}
