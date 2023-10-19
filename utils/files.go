package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 覆盖写入文件
// filePath文件路径
func WriteToFile(filePath string, content string) error {
	if _, err := os.Stat(filePath); err != nil {
		if !os.IsExist(err) {
			pathstr_arr := strings.Split(filePath, `/`)
			path_dirs := strings.Split(filePath, (pathstr_arr[len(pathstr_arr)-1]))
			os.MkdirAll(path_dirs[0], os.ModePerm)
			os.Create(filePath)
		}
	}
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content), n)
		defer f.Close()
	}
	return err
}

// 逐行读取文件
// filePath文件路径
func ReaderFileByline(filePath string) []interface{} {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	var list []interface{}
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		list = append(list, string(line))
	}
	return list
}

// 一次性读取全部文件
// filePath文件路径
func ReaderFileBystring(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}
	return string(bytes)
}
