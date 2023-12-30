package utils

import (
	"io"
	"mime/multipart"
	"os"
)

func saveItem(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func SaveUserImage(file *multipart.FileHeader, uid string, preFix string) string {
	// 获取文件对象
	// 设置保存路径及文件名
	savePath := "resource/uploads/" + preFix + "/" + file.Filename
	if _, err := os.Stat(savePath); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(savePath, os.ModePerm)
		}
	}

	return savePath
}
