package developer

import (
	"bufio"
	"fmt"
	"gofly/utils"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gohouse/gorose/v2"
)

// 判断文件是否存在不存在则创建
func CreatApicodeFile(root_path string, data gorose.Data) {
	url := data["url"].(string)
	url_arr := strings.Split(url, `/`)
	methods := url_arr[len(url_arr)-1]
	filename := url_arr[len(url_arr)-2]
	model_path := strings.Split(url, filename)
	folder_path := root_path + model_path[0]
	//1创建文件夹
	file_path := fmt.Sprintf("%s%s", "app/", folder_path)
	//如果没有filepath文件目录就创建一个
	if _, err := os.Stat(file_path); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(file_path, os.ModePerm)
		}
	}
	//2创建文件
	filego_path := fmt.Sprintf("%s%s%s.go", "app/", folder_path, filename)
	if _, err := os.Stat(filego_path); err != nil {
		if !os.IsExist(err) {
			os.Create(filego_path)
			//复制文件
			err := CopyFileContents("resource/staticfile/codetpl/apicode.gos", filego_path)
			if err != nil {
				panic(err)
			}
			//修复头部-数据表
			ChangPackage(filego_path, url_arr[1], filename, data["tablename"].(string))
			//添加控制路由
			CheckApiIsAddController(root_path, root_path+"/"+url_arr[1])
		}
	}
	//创建list
	if data["method"] == "get" {
		if data["getdata_type"] == "list" {
			CreatList(filego_path, methods, data["fields"].(string))
		} else if data["getdata_type"] == "detail" {
			CreatDetail(filego_path, methods, data["fields"].(string))
		}
	} else if data["method"] == "post" {
		CreatSave(filego_path, methods)
	} else if data["method"] == "delete" {
		CreatDel(filego_path, methods)
	}
}

// 1.创建list方法
func CreatList(filePath, methods, fields string) {
	//替换文件内容
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), "get_list") {
			datestr := strings.ReplaceAll(string(a), "get_list", utils.FirstUpper(methods))
			result += datestr + "\n"
		} else if strings.Contains(string(a), "{fields}") {
			datestr := strings.ReplaceAll(string(a), "{fields}", fields)
			result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 2.创建Detail方法
func CreatDetail(filePath, methods, fields string) {
	//替换文件内容
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), "get_detail") {
			datestr := strings.ReplaceAll(string(a), "get_detail", utils.FirstUpper(methods))
			result += datestr + "\n"
		} else if strings.Contains(string(a), "{fields}") {
			datestr := strings.ReplaceAll(string(a), "{fields}", fields)
			result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 3.创建Save方法
func CreatSave(filePath, methods string) {
	//替换文件内容
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), "save(") {
			datestr := strings.ReplaceAll(string(a), "save(", utils.FirstUpper(methods)+"(")
			result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 4.创建Del方法
func CreatDel(filePath, methods string) {
	//替换文件内容
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), "del(") {
			datestr := strings.ReplaceAll(string(a), "del(", utils.FirstUpper(methods)+"(")
			result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 创建修改go文件package
func ChangPackage(filePath, packageName, filename, tablename string) {
	//替换文件内容
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), "packageName") {
			datestr := strings.ReplaceAll(string(a), "packageName", packageName)
			result += datestr + "\n"
		} else if strings.Contains(string(a), "Replace") {
			datestr := strings.ReplaceAll(string(a), "Replace", utils.FirstUpper(filename))
			result += datestr + "\n"
		} else if strings.Contains(string(a), "{tablename}") {
			datestr := strings.ReplaceAll(string(a), "{tablename}", tablename)
			result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 检查该类是否添加到控制器
func CheckApiIsAddController(modelname, path string) {
	filePath := "app/" + modelname + "/controller.go"
	//1判断文件没有则添加
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("不存在！")
			os.Create(filePath)
			//复制文件
			err := CopyFileContents("resource/staticfile/codetpl/controller.gos", filePath)
			if err != nil {
				panic(err)
			}
		}
	}
	//2修改文件
	con_path := "gofly/app/" + path
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	ishase := false
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), con_path) {
			ishase = true
		}
		result += string(a) + "\n"
	}
	if ishase == false {
		addstr := "	_ \"" + con_path + "\"\n"
		datestr := strings.ReplaceAll(result, ")", addstr)
		result = datestr + ")\n"
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 检查该类控制器四否存在路由移除
func CheckApiRemoveController(modelname, path string) {
	filePath := "app/" + modelname + "/controller.go"
	con_path := "gofly/app/" + path
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), con_path) { //存在路由则移除
			// datestr := strings.ReplaceAll(string(a), utils.FirstUpper(methods), "get_list")
			// result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}

	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 卸载接口
func UnApicodeFile(root_path string, data gorose.Data) {
	url := data["url"].(string)
	url_arr := strings.Split(url, `/`)
	methods := url_arr[len(url_arr)-1]
	filename := url_arr[len(url_arr)-2]
	model_path := strings.Split(url, filename)
	folder_path := root_path + model_path[0]
	//2判断文件是否存在
	filego_path := fmt.Sprintf("%s%s%s.go", "app/", folder_path, filename)
	if _, err := os.Stat(filego_path); err == nil {
		if data["method"] == "get" {
			if data["getdata_type"] == "list" {
				UnList(filego_path, methods, data["fields"].(string))
			} else if data["getdata_type"] == "detail" {
				UnDetail(filego_path, methods, data["fields"].(string))
			}
		} else if data["method"] == "post" {
			UnSave(filego_path, methods)
		} else if data["method"] == "delete" {
			UnDel(filego_path, methods)
		}
	}
}

// 1.卸载list方法
func UnList(filePath, methods, fields string) {
	//替换文件内容
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), utils.FirstUpper(methods)) {
			datestr := strings.ReplaceAll(string(a), utils.FirstUpper(methods), "get_list")
			result += datestr + "\n"
		} else if strings.Contains(string(a), fields) {
			datestr := strings.ReplaceAll(string(a), fields, "{fields}")
			result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 2.卸载Detail方法
func UnDetail(filePath, methods, fields string) {
	//替换文件内容
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), utils.FirstUpper(methods)) {
			datestr := strings.ReplaceAll(string(a), utils.FirstUpper(methods), "get_detail")
			result += datestr + "\n"
		} else if strings.Contains(string(a), fields) {
			datestr := strings.ReplaceAll(string(a), fields, "{fields}")
			result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 3.卸载Save方法
func UnSave(filePath, methods string) {
	//替换文件内容
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), utils.FirstUpper(methods)+"(") {
			datestr := strings.ReplaceAll(string(a), utils.FirstUpper(methods)+"(", "save(")
			result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 4.自饿着Del方法
func UnDel(filePath, methods string) {
	//替换文件内容
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var result = ""
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(a), utils.FirstUpper(methods)+"(") {
			datestr := strings.ReplaceAll(string(a), utils.FirstUpper(methods)+"(", "del(")
			result += datestr + "\n"
		} else {
			result += string(a) + "\n"
		}
	}
	fw, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666) //os.O_TRUNC清空文件重新写入，否则原文件内容可能残留
	w := bufio.NewWriter(fw)
	w.WriteString(result)
	if err != nil {
		panic(err)
	}
	w.Flush()
}

// 卸载时候删除文件
func RemoveModel(root_path string, data gorose.Data) {
	url := data["url"].(string)
	url_arr := strings.Split(url, `/`)
	filename := url_arr[len(url_arr)-2]
	model_path := strings.Split(url, filename)
	folder_path := root_path + model_path[0]
	//2创建文件
	filego_path := fmt.Sprintf("%s%s%s.go", "app/", folder_path, filename)
	if _, err := os.Stat(filego_path); err == nil {
		//1.文件存在删除文件
		os.Remove(filego_path)
		//2.删除文件夹
		file_path := fmt.Sprintf("%s%s", "app/", folder_path)
		dir, _ := ioutil.ReadDir(file_path)
		if len(dir) == 0 {
			os.RemoveAll(file_path)
			//3.移除路由
			CheckApiRemoveController(root_path, root_path+"/"+url_arr[1])
		}
	}
}
