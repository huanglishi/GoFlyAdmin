package utils

import (
	"fmt"
	"github.com/gohouse/gorose/v2"
	"github.com/shopspring/decimal"
	"gofly/global"
	"gofly/utils"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"os"
	"sort"
	"strings"
	"time"
)

var timeTemplate1 = "2006-01-02 15:04:05"
var LOC, _ = time.LoadLocation("Asia/Shanghai")

func GetTodayStart() string {
	year, month, day := time.Now().Date()
	location, _ := time.LoadLocation("Asia/Shanghai")
	unix := time.Date(year, month, day, 0, 0, 0, 0, location).Format(timeTemplate1)
	return unix
}

func GetAmountPayCoefficient(payPrice float64) float64 {
	for _, v := range global.App.Config.Userconf.UserIncomeRateMap {
		if payPrice > v.Min && payPrice <= v.Max {
			return v.TackOffRate
		}
	}
	return 0
}

func GetGiftConfig(giftType uint) float64 {
	for _, v := range global.App.Config.Userconf.UserGiftMap {
		if v.GiftType == giftType {
			return v.GiftValue
		}
	}
	return 0
}

func DBTimeStamp() string {
	return time.Now().Format(timeTemplate1)
}
func U64TimeStamp(timeStamp string) time.Time {
	time, _ := time.ParseInLocation(timeStamp, timeTemplate1, LOC)
	return time
}
func Account2String(fAccount float64) string {
	return decimal.NewFromFloat(fAccount).String()
}

func Account2Float(sAccount string) float64 {
	fAccount, _ := decimal.NewFromString(sAccount)
	return fAccount.InexactFloat64()
}

func FloatDifference(sAccount, s1Account float64) float64 {
	rst := decimal.NewFromFloat(sAccount).Add(decimal.NewFromFloat(s1Account))
	return rst.InexactFloat64()
}
func FileExchange(fileRead string, fileWrite string) {
	fileRd, err := os.OpenFile(fileRead, os.O_RDONLY, 0777)
	fmt.Println(err)
	readBytes, err := ioutil.ReadAll(fileRd)
	fmt.Println(err)
	fileRd.Close()

	file1, err := os.OpenFile(fileWrite, os.O_RDWR|os.O_CREATE, 0777)
	fmt.Println(err)
	_, err = file1.Write(readBytes)
	fmt.Println(err)
	err = file1.Close()
	fmt.Println(err)
}

func GetFileFullPath(fileHandle *multipart.FileHeader, filePath string, pos string) (fullPath string) {
	var filenameArr = strings.Split(fileHandle.Filename, ".")
	nameStr := utils.Md5(fmt.Sprintf("%v%s", time.Now().Unix(), filenameArr[0]))
	var newName string
	if pos == "" {
		newName = nameStr + "." + filenameArr[1]
	} else {
		newName = nameStr + "_" + pos + "." + filenameArr[1]
	}
	// 上传文件到指定的目录
	return filePath + newName
}
func GenerateOrderID() string {
	timestamp := time.Now().Unix()
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(1000) // 生成一个0到999之间的随机数
	orderID := fmt.Sprintf("%d%d", timestamp, randomNumber)
	return orderID
}

func GetUserImgUrl(uid string, imgType uint, image string) string {
	switch imgType {
	case 1:
		return fmt.Sprintf("%s%s%s/%s/%d/%s", global.App.Config.Userconf.ServerAddress, "/", global.App.Config.Userconf.ImagePrePath, uid, imgType, image)
	case 2:
		return fmt.Sprintf("%s%s%s/%s/%d/%s", global.App.Config.Userconf.ServerAddress, "/", global.App.Config.Userconf.ImagePrePath, uid, imgType, image)
	case 9:
		list := strings.Split(image, ",")
		var rstList []string
		for _, v := range list {
			rstList = append(rstList, fmt.Sprintf("%s%s%s/%s/%d/%s",
				global.App.Config.Userconf.ServerAddress, "/", global.App.Config.Userconf.ImagePrePath, uid, imgType, v))
		}
		return strings.Join(rstList, ",")
	case 10:
		list := strings.Split(image, ",")
		var rstList []string
		for _, v := range list {
			rstList = append(rstList, fmt.Sprintf("%s%s%s/%s/%d/%s",
				global.App.Config.Userconf.ServerAddress, "/", global.App.Config.Userconf.ImagePrePath, uid, imgType, v))
		}
		return strings.Join(rstList, ",")
	}

	return ""
}

type fileBy []string

func (arr fileBy) Len() int {
	return len(arr)
}
func (arr fileBy) Less(i, j int) bool {
	if arr[i] != "" && arr[j] != "" {
		step1 := strings.Split(strings.Split(arr[i], "_")[1], ".")[0]
		step2 := strings.Split(strings.Split(arr[j], "_")[1], ".")[0]
		return step1 < step2
	}
	return false
}
func (arr fileBy) Swap(i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func ReadDirFileNames(filePath string) (rst []string) {
	rd, _ := ioutil.ReadDir(filePath)
	for _, fi := range rd {
		if !fi.IsDir() {
			if fi.Name() != "" {
				rst = append(rst, fi.Name())
			}
		}
	}
	var byDesc fileBy = rst
	sort.Sort(byDesc)
	return rst
}
func ModelDBUsrInfo(dbUsr gorose.Data) map[string]interface{} {
	Sexual := dbUsr["Sexual"].(int64)
	resp := map[string]interface{}{
		"user_name":      dbUsr["UserName"].(string),
		"face_img":       GetUserImgUrl(dbUsr["Uid"].(string), 1, dbUsr["FaceImg"].(string)),
		"identity_state": dbUsr["IdentityState"].(int64),
	}
	if dbUsr["FaceImgState"] != nil {
		resp["face_img_state"] = dbUsr["FaceImgState"].(int64)
	}
	if dbUsr["Age"] != nil {
		resp["age"] = dbUsr["Age"].(int64)
	}
	if dbUsr["signature"] != nil {
		resp["signature"] = dbUsr["Signature"].(string)
	}
	if dbUsr["self_introduction"] != nil {
		resp["self_introduction"] = dbUsr["SelfIntroduction"].(string)
	}
	if dbUsr["friend_tag"] != nil {
		resp["friend_tag"] = dbUsr["FriendTag"].(int64)
	}
	if dbUsr["BgImg"] != nil {
		resp["bg_img"] = GetUserImgUrl(dbUsr["Uid"].(string), 2, dbUsr["BgImg"].(string))
	}
	if dbUsr["BgImgState"] != nil {
		resp["bg_img_state"] = dbUsr["BgImgState"].(int64)
	}
	if dbUsr["distance"] != nil {
		resp["distance"] = dbUsr["distance"].(float64)
	}
	if dbUsr["TelNum"] != nil {
		resp["tel_num"] = dbUsr["TelNum"].(string)
	}
	if dbUsr["IdentityPic"] != nil {
		resp["identity_pic"] = GetUserImgUrl(dbUsr["Uid"].(string), 4, dbUsr["IdentityPic"].(string))
	}
	if Sexual == 0 {
		if dbUsr["MeetUnlockGift"] != nil {
			resp["meet_unlock_gift"] = dbUsr["MeetUnlockGift"].(int64)
		}
	}
	if Sexual == 1 {
		if dbUsr["CarType"] != nil {
			resp["car_type"] = dbUsr["CarType"].(int64)
		}
		if dbUsr["CustomGift"] != nil {
			resp["custom_gift"] = dbUsr["CustomGift"].(string)
		}
		if dbUsr["MeetGift"] != nil {
			resp["meet_gift"] = dbUsr["MeetGift"].(int64)
		}
	}
	return resp
}
