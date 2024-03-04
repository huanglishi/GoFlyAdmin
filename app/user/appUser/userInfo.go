package appUser

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	com "gofly/app/user/utils"
	"gofly/global"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils"
	"gofly/utils/results"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// 上传用户图片-头像、背景、照片墙
var uploadNumMax = "9"

// "1"头像 "2"背景 "3"联系方式二维码图片 "4"身份认证图片 "9"照片墙
const (
	userImageFace     = "1"
	userImageBg       = "2"
	userImageContact  = "3"
	userImageIdentity = "4"
	userImagePhotos   = "9"
)

// 用户上传图片时对应获取数据库存储位置接口
func getImagePosition(imgType string) (userInfo string, imageState string) {
	switch imgType {
	case userImageFace:
		userInfo = "FaceImg"
		imageState = "FaceImgState"
	case userImageBg:
		userInfo = "BgImg"
		imageState = "BgImgState"
	case userImageContact:
		userInfo = "ContactQrCode"
	case userImageIdentity:
		userInfo = "IdentityPic"
		imageState = "IdentityState"
	case userImagePhotos:
		userInfo = "Photos"
	}
	return userInfo, imageState
}

const (
	IMG_REVIEWING = 0 //图片审核中
	IMG_PASSED    = 1 //审核通过
	IMG_REJIECT   = 2 //审核失败
)

type UsrInfo struct {
	Uid               string `json:"uid"`                 //用户唯一ID
	DeviceId          string `json:"device_id"`           //设备唯一id
	DeviceType        uint   `json:"device_type"`         //设备端类型 设备类型1: IOS, 2: Android, 8: Android Pad
	LoginState        uint   `json:"login_state"`         // 登录状态
	Sexual            uint   `json:"sexual"`              //性别0女 1男 2 其他
	TelNum            string `json:"tel_num"`             //注册的电话号码
	UserName          string `json:"user_name"`           //用户名
	Password          string `json:"password"`            //账户密码
	PassSalt          string `json:"_"`                   //密码盐--
	Account           string `json:"account"`             // 账户余额
	Age               uint   `json:"age"`                 //年龄
	Signature         string `json:"signature"`           //签名
	SelfIntroduction  string `json:"self_introduction"`   //自我介绍
	FriendTag         uint   `json:"friend_tag"`          //0 找人聊天 1 找对象 2 找知己 4 随便看看
	City              string `json:"city"`                //城市
	Longitude         string `json:"longitude"`           //经度
	Latitude          string `json:"latitude"`            //纬度
	FaceImg           string `json:"face_img"`            //头像地址
	FaceImgState      uint   `json:"face_img_state"`      //头像审核状态 0 审核中 1审核通过 2 不通过
	BgImg             string `json:"bg_img"`              //背景图地址
	BgImgState        uint   `json:"bg_img_state"`        //背景图审核状态 0 审核中 1审核通过 2 不通过
	IdentityPic       string `json:"identity_pic"`        //身份认证图片地址
	IdentityState     uint   `json:"identity_state"`      //身份认证审核状态 0 审核中 1审核通过 2 不通过
	CarType           uint   `json:"car_type"`            //车型图片
	MeetGift          uint   `json:"meet_gift"`           //见面礼物类型
	CustomGift        string `json:"custom_gift"`         //自定礼物价格
	ContactUnlockGift uint   `json:"contact_unlock_gift"` //联系方式解锁
	MeetUnlockGift    uint   `json:"meet_unlock_gift"`    //见面礼物类型
	CreateTime        string `json:"create_time"`         //注册时间
	UpdateTime        string `json:"update_time"`         //更新时间
}

// 获取图片认证状态
func (api *UserOp) GetImgAuthState(c *gin.Context) {
	uid := c.Query("uid")
	res, err := model.DB().Table("app_user").Fields("FaceImgState", "BgImgState", "IdentityState").Where("Uid", uid).First()
	if res == nil || err != nil {
		results.Failed(c, "入参错误！", userNotExist)
		return
	}
	results.Success(c, "返回审核状态", apiSuccess, map[string]interface{}{
		"face_img_state": res["FaceImgState"].(int64),
		"bg_img_state":   res["BgImgState"].(int64),
		"identity_state": res["IdentityState"].(int64),
	})
}

// 更新位置信息
func (api *UserOp) UpdateLocation(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	if parameter["latitude"] == nil || parameter["longitude"] == nil ||
		parameter["uid"] == nil {
		results.Failed(c, "入参错误！", paramError)
		return
	}
	res, err := model.DB().Table("app_user").Fields().Where("Uid", parameter["uid"].(string)).First()
	if res == nil || err != nil {
		results.Failed(c, "用户不存在！", userNotExist)
		return
	}
	_, err = model.DB().Table("app_user").Where("Uid", parameter["uid"].(string)).
		Data(map[string]interface{}{
			"Latitude":   parameter["latitude"].(string),
			"Longitude":  parameter["longitude"].(string),
			"UpdateTime": com.DBTimeStamp(),
		}).Update()
	if err != nil {
		results.Failed(c, "账号不存在!", userNotExist)
		return
	}
	results.Success(c, "更新成功", apiSuccess, nil)
}

// 更新资料
func (api *UserOp) UpdateNormalInfo(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var userInfo UsrInfo
	_ = json.Unmarshal(body, &userInfo)

	res, err := model.DB().Table("app_user").Where("Uid", userInfo.Uid).First()
	if res == nil || err != nil {
		results.Failed(c, "入参错误！", userNotExist)
		return
	}
	_, err = model.DB().Table("app_user").Where("Uid", userInfo.Uid).Data(userInfo).Update()
	if err != nil {
		results.Failed(c, "更新失败!", serverInternalErr)
		return
	}
	results.Success(c, "更新成功", apiSuccess, nil)
}

// 更新账户余额
func (api *UserOp) UpdateAccountInfo(c *gin.Context) {
	var parameter map[string]interface{}
	body, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(body, &parameter)

	if parameter["uid"] == nil || parameter["amount"] == nil || parameter["type"] == nil {
		results.Failed(c, "入参错误！", paramError)
		return
	}
	uid := parameter["uid"].(string)

	res, err := model.DB().Table("app_user").Fields("Uid", "Account").
		Where("Uid", uid).First()
	if res == nil || err != nil {
		results.Failed(c, "入参错误！", userNotExist)
		return
	}

	amount := parameter["amount"].(string)
	from := int(parameter["type"].(float64))
	inAmount := com.Account2Float(amount)
	oldAccount := com.Account2Float(res["Account"].(string))
	_, err = model.DB().Table("app_user").Where("Uid", uid).
		Data(map[string]interface{}{
			"Account": com.Account2String(inAmount + oldAccount),
		}).Update()
	if err != nil {
		results.Failed(c, "更新失败!", serverInternalErr)
		return
	}
	_, err = model.DB().Table("app_user_recharge_record").Data(map[string]interface{}{
		"Uid":        uid,
		"Amount":     amount,
		"Type":       from,
		"UpdateTime": com.DBTimeStamp(),
	}).Insert()
	results.Success(c, "更新成功", apiSuccess, nil)
}

func init() {
	utils.Register(&UserOp{}, reflect.TypeOf(UserOp{}).PkgPath())
}

// 用于自动注册路由
type UserOp struct {
}

// 阿里云短信验证
func (api *UserOp) SmsVerify(c *gin.Context) {
	//参数一：连接的节点地址（有很多节点选择，这里我选择杭州）
	//参数二：AccessKey ID
	//参数三：AccessKey Secret
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI5t**********AWVNP", "qgzyEW4cA*******PBl7iyw")

	request := dysmsapi.CreateSendSmsRequest()       //创建请求
	request.Scheme = "https"                         //请求协议
	request.PhoneNumbers = "155******85"             //接收短信的手机号码
	request.SignName = "教程"                          //短信签名名称
	request.TemplateCode = "SMS_******236"           //短信模板ID
	par, err := json.Marshal(map[string]interface{}{ //定义短信模板参数（具体需要几个参数根据自己短信模板格式）
		"code": "123456",
	})
	request.TemplateParam = string(par) //将短信模板参数传入短信模板
	//切换设备--使用验证码登录
	/*if deviceId != res["DeviceId"].(string) {
		results.Failed(c, "切换设备，请使用验证码重新登录", pass)
		return
	}*/
	response, err := client.SendSms(request) //调用阿里云API发送信息
	if err != nil {                          //处理错误
		fmt.Print(err.Error())
		results.Failed(c, "短信验证失败", err)
		return
	}
	results.Success(c, "短信验证成功", response, nil)
	fmt.Printf("response is %#v\n", response) //控制台输出响应
}

func isNullParam(inStr string) bool {
	if inStr == "" {
		return true
	}
	return false
}

// 获取Token
func getToken(appUser map[string]interface{}) interface{} {
	token := middleware.GenerateAppToken(&middleware.AppUserClaims{
		DeviceId:       appUser["device_id"].(string),
		MacAddr:        appUser["mac_addr"].(string),
		StandardClaims: jwt.StandardClaims{},
	})
	return token
}

// 用户获取token
func (api *UserOp) GetRefreshToken(c *gin.Context) {
	var appUser = make(map[string]interface{})
	if isNullParam(c.Query("device_id")) || isNullParam(c.Query("mac_addr")) {
		results.Failed(c, "token获取失败", "in param is null")
		return
	}
	appUser["device_id"] = c.Query("device_id")
	appUser["mac_addr"] = c.Query("mac_addr")
	token := getToken(appUser)

	results.SuccessLogin(c, "token", nil, token, nil)
}

// 注册用户
func (api *UserOp) RegisterUser(c *gin.Context) {
	var regInfo UsrInfo
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(body, &regInfo)
	//如果该用户手机号已注册 提示直接登录
	userdata, _ := model.DB().Table("app_user").Fields("id").Where("TelNum", regInfo.TelNum).First()
	if userdata != nil {
		results.Failed(c, "该手机号已注册,请登录！", telNumExist)
		return
	}
	//如果该用户名存在 提示重新修改
	userdata, _ = model.DB().Table("app_user").Fields("id").Where("UserName", regInfo.UserName).First()
	if userdata != nil {
		results.Failed(c, "该用户名已存在,请重新修改！", userNameExist)
		return
	}
	//创建用户
	regInfo.Uid = uuid.NewV4().String()
	regInfo.UpdateTime = com.DBTimeStamp()
	regInfo.CreateTime = regInfo.UpdateTime
	if regInfo.Password != "" {
		salt := time.Now().Unix()
		regInfo.PassSalt = strconv.Itoa(int(salt))
		regInfo.Password = utils.Md5(regInfo.Password + regInfo.PassSalt)
	}
	//设置默认背景图和头像图
	regInfo.FaceImg = "face.png"
	regInfo.BgImg = "bg.png"
	setUserDefaultFaceBgPath(regInfo.Uid, regInfo.Sexual)
	_, err := model.DB().Table("app_user").Data(regInfo).Insert()
	if err != nil {
		log.Fatal("Post_registerUser db insert", err)
		results.Failed(c, "添加账号失败", dbInsertError)
		return
	}
	//返回 uid--手机保存
	results.Success(c, "注册成功", apiSuccess, regInfo.Uid)
}

func setUserDefaultFaceBgPath(uid string, sexual uint) {
	preFix := global.App.Config.Userconf.ImagePrePath
	facePath := fmt.Sprintf("%s/%s%s1%s", preFix, uid, "/", "/")
	if _, err := os.Stat(facePath); err != nil {
		if !os.IsExist(err) {
			err := os.MkdirAll(facePath, os.ModePerm)
			fmt.Println(err)
		}
	}
	bgPath := fmt.Sprintf("%s/%s%s2%s", preFix, uid, "/", "/")
	if _, err := os.Stat(bgPath); err != nil {
		if !os.IsExist(err) {
			err := os.MkdirAll(bgPath, os.ModePerm)
			fmt.Println(err)
		}
	}
	var defaultFaceImg, defaultBgImg string
	if sexual == 0 {
		defaultFaceImg = fmt.Sprintf("%s/%s", preFix, "face_women.png")
		defaultBgImg = fmt.Sprintf("%s/%s", preFix, "bg_women.png")
	} else {
		defaultFaceImg = fmt.Sprintf("%s/%s", preFix, "face_man.png")
		defaultBgImg = fmt.Sprintf("%s/%s", preFix, "bg_man.png")
	}

	com.FileExchange(defaultFaceImg, facePath+"face.png")
	com.FileExchange(defaultBgImg, bgPath+"bg.png")
}

// 用户登录-用户名密码
func (api *UserOp) Login(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	if parameter["user_name"] == nil || parameter["password"] == nil ||
		parameter["device_id"] == nil || parameter["uid"] == nil {
		results.Failed(c, "请提交用户账号或密码！", paramError)
		return
	}
	username := parameter["user_name"].(string)
	password := parameter["password"].(string)
	uid := parameter["uid"]

	res, err := model.DB().Table("app_user").Fields("Uid,DeviceId,UserName,PassSalt,Password").Where("Uid", uid).First()
	if res == nil || err != nil {
		results.Failed(c, "账号不存在!", userNotExist)
		return
	}
	if username != res["UserName"] {
		results.Failed(c, "账号不存在!", userNotExist)
		return
	}
	pass := utils.Md5(password + res["PassSalt"].(string))
	if pass != res["Password"] {
		results.Failed(c, "您输入的密码不正确!", passwdIncorrect)
		return
	}
	if parameter["device_id"].(string) != res["DeviceId"].(string) {
		results.Failed(c, "切换设备,请使用验证码重新登录", deviceIdExist)
		return
	}
	model.DB().Table("app_user").Where("Uid", res["Uid"]).
		Data(map[string]interface{}{
			"LoginState": 1,
			"UpdateTime": com.DBTimeStamp()}).Update()
	//登录日志
	model.DB().Table("app_user_login_logs").
		Data(map[string]interface{}{
			"Uid": res["Uid"], "InOrOut": "in",
			"LoginTime": com.DBTimeStamp(),
			"LoginIp":   utils.GetIp(c)}).Insert()
	results.Success(c, "登录成功返回token！", apiSuccess, nil)
}

// 上传单张图片
func (api *UserOp) UploadOne(c *gin.Context) {
	imgType := c.Query("photo_type") //图片类型
	Uid := c.Query("uid")
	//除照片墙可以多张图,其余均为一张
	if imgType < userImageFace || imgType > userImagePhotos {
		results.Failed(c, "参数错误！", paramError)
		return
	}
	if Uid == "" {
		results.Failed(c, "用户Uid未填写", paramError)
		return
	}
	dbFiled, imageState := getImagePosition(imgType)

	res, err := model.DB().Table("app_user").Fields("Uid", dbFiled).
		Where("Uid", Uid).First()
	if res == nil || err != nil {
		results.Failed(c, "账号不存在!", userNotExist)
		return
	}

	preFix := global.App.Config.Userconf.ImagePrePath
	filePath := fmt.Sprintf("%s%s%s%s%s", preFix, Uid, "/", imgType, "/")
	//如果没有filepath文件目录就创建一个
	if _, err := os.Stat(filePath); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(filePath, os.ModePerm)
		}
	}
	fileHandle, err := c.FormFile("file")
	if err != nil {
		results.Failed(c, err.Error(), imageUploadErr)
		return
	}
	err = c.SaveUploadedFile(fileHandle, com.GetFileFullPath(fileHandle, filePath, ""))
	if err != nil {
		results.Failed(c, err.Error(), imageUploadErr)
		return
	}
	if res[dbFiled] != nil {
		oldName := res[dbFiled].(string)
		os.Remove(filePath + oldName)
	}
	dbPath := com.ReadDirFileNames(filePath)
	var updateInfo = make(map[string]interface{})
	if len(dbPath) > 0 {
		updateInfo[dbFiled] = dbPath[0]
	}
	if imageState != "" {
		updateInfo[imageState] = 0
	}
	_, err = model.DB().Table("app_user").Where("Uid", Uid).Data(updateInfo).Update()
	if err != nil {
		results.Failed(c, err.Error(), dbInsertError)
		return
	}
	var updateState = make(map[string]interface{})
	switch imgType {
	case "1": //头像
		updateState["FaceImgState"] = 0
	case "2": //背景
		updateState["BgImgState"] = 0
	case "4": //身份
		updateState["IdentityState"] = 0
	}
	model.DB().Table("app_user").Where("Uid", Uid).Data(updateState).Update()
	//上传成功 存在旧图则删除
	results.Success(c, "上传成功,审核中...", apiSuccess, nil)
}

// 修改密码
func (api *UserOp) ChangePwd(c *gin.Context) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	if parameter["uid"] == nil || parameter["passwordOld"] == nil || parameter["passwordNew"] == nil {
		results.Failed(c, "参数错误！", paramError)
		return
	}
	uid := parameter["uid"].(string)
	userdata, err := model.DB().Table("app_user").Where("Uid", uid).Fields("Password,PassSalt").First()
	if err != nil {
		results.Failed(c, "查找账号失败！", paramError)
		return
	} else {
		pass := utils.Md5(parameter["passwordOld"].(string) + userdata["PassSalt"].(string))
		if userdata["Password"] != pass {
			results.Failed(c, "原来密码输入错误！", paramError)
			return
		} else {
			newPass := utils.Md5(parameter["passwordNew"].(string) + userdata["PassSalt"].(string))
			res, err := model.DB().Table("app_user").Where("Uid", uid).
				Data(map[string]interface{}{"Password": newPass}).Update()
			if err != nil {
				results.Failed(c, "修改密码失败", serverInternalErr)
				return
			} else {
				results.Success(c, "修改密码成功！", res, nil)
				return
			}
		}
	}
}

// 退出登录
func (api *UserOp) Logout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	uid := c.Query("uid")
	if token != "" {
		middleware.Refresh(token)
		_, err := model.DB().Table("app_user").Where("Uid", uid).
			Data(map[string]interface{}{"LoginState": 0}).Update()
		if err != nil {
			results.Failed(c, "退出登录失败", paramError)
		}
	}
	results.Success(c, "退出登录成功", true, nil)
}

func userSocialListQueryStr(longitude, latitude interface{}, sexual, limit, offset uint) (queryStr string) {
	/*
		SELECT s.Uid,s.UserName,s.Age,s.Signature,s.SelfIntroduction,s.FriendTag,s.FaceImg,s.IdentityState,
		(st_distance (point (Longitude, Latitude),point(108.979654,34.358591)) / 0.0111) AS distance
		 FROM app_user s WHERE Sexual=0 AND FaceImgState=1 AND BgImgState=1 ORDER BY distance LIMIT 10 OFFSET 0
	*/
	//未获取经纬度时
	if longitude == nil || latitude == nil {
		if sexual == 0 { //查询女性用户
			queryStr = fmt.Sprintf(
				"SELECT Uid,UserName,Age,Signature,SelfIntroduction,FriendTag,FaceImg,BgImg,IdentityState,ContactUnlockGift,MeetUnlockGift\n"+
					"FROM app_user WHERE Sexual=%d AND FaceImgState=1 AND BgImgState=1 ORDER BY UpdateTime desc LIMIT %d OFFSET %d",
				sexual, limit, offset)
		} else { //查询男性用户
			queryStr = fmt.Sprintf(
				"SELECT Uid,UserName,Age,Signature,SelfIntroduction,FriendTag,FaceImg,BgImg,IdentityState,CarType,MeetGift,CustomGift\n"+
					"FROM app_user WHERE Sexual=%d AND FaceImgState=1 AND BgImgState=1 ORDER BY UpdateTime desc LIMIT %d OFFSET %d",
				sexual, limit, offset)
		}
	} else {
		if sexual == 0 {
			queryStr = fmt.Sprintf(
				"SELECT Uid,UserName,Age,Signature,SelfIntroduction,FriendTag,FaceImg,BgImg,IdentityState,ContactUnlockGift,MeetUnlockGift,\n"+
					"(st_distance (point (Longitude, Latitude),point(%s,%s) ) / 0.0111) AS distance FROM app_user\n"+
					"WHERE Sexual=%d AND FaceImgState=1 AND BgImgState=1 ORDER BY distance LIMIT %d OFFSET %d",
				longitude.(string), latitude.(string), sexual, limit, offset)
		} else if sexual == 1 {
			queryStr = fmt.Sprintf(
				"SELECT Uid,UserName,Age,Signature,SelfIntroduction,FriendTag,FaceImg,BgImg,IdentityState,CarType,MeetGift,CustomGift,\n"+
					"(st_distance (point (Longitude, Latitude),point(%s,%s) ) / 0.0111) AS distance FROM app_user\n"+
					"WHERE Sexual=%d AND FaceImgState=1 AND BgImgState=1 ORDER BY distance LIMIT %d OFFSET %d",
				longitude.(string), latitude.(string), sexual, limit, offset)
		}
	}
	fmt.Println(queryStr)
	return queryStr
}

// 获取用户交友卡片列表
func (api *UserOp) GetUserSocialList(c *gin.Context) {
	//获取post传过来的data
	var parameter map[string]interface{}
	body, _ := io.ReadAll(c.Request.Body)
	_ = json.Unmarshal(body, &parameter)
	if parameter["limit"] == nil || parameter["offset"] == nil || parameter["sexual"] == nil {
		results.Failed(c, "参数错误", paramError)
		return
	}
	limit := uint(parameter["limit"].(float64))
	offset := uint(parameter["offset"].(float64))
	sexual := uint(parameter["sexual"].(float64))

	queryStr := userSocialListQueryStr(parameter["longitude"], parameter["latitude"], sexual, limit, offset)
	res, err := model.DB().Query(queryStr)
	if res == nil || err != nil {
		results.Failed(c, "获取数据失败", serverInternalErr)
		return
	}

	var resp []map[string]interface{}
	for _, v := range res {
		resp = append(resp, com.ModelDBUsrInfo(v))
	}
	results.Success(c, "查询成功", apiSuccess, res)
}

// 获取单个用户信息
func (api *UserOp) GetUserInfo(c *gin.Context) {
	//获取post传过来的data
	uid := c.Query("uid")
	if uid == "" {
		results.Failed(c, "参数错误", paramError)
		return
	}
	res, err := model.DB().Table("app_user").Where("Uid", uid).
		Fields("Uid", "UserName", "Age", "Signature", "SelfIntroduction", "FriendTag", "FaceImg", "BgImg",
			"IdentityState", "CarType", "MeetGift", "CustomGift", "ContactUnlockGift", "MeetUnlockGift").First()
	if res == nil || err != nil {
		results.Failed(c, "获取数据失败", serverInternalErr)
		return
	}
	results.Success(c, "查询成功", apiSuccess, com.ModelDBUsrInfo(res))
}

// 用户反馈信息
func (api *UserOp) UserFeedBack(c *gin.Context) {
	Uid := c.Query("uid")
	formHandle, _ := c.MultipartForm()
	images := formHandle.File["file"]
	content := formHandle.Value["content"]
	if Uid == "" {
		results.Failed(c, "参数错误！", paramError)
		return
	}
	feedbackInfo, err := model.DB().Table("app_user_feedback").Where("Uid", Uid).First()
	if feedbackInfo != nil {
		results.Failed(c, "您的反馈处理中..", feedbackExist)
		return
	}
	preFix := global.App.Config.Userconf.ImagePrePath
	articleImgPrefix := com.GenerateOrderID()
	filePath := fmt.Sprintf("%s%s%s%s%d%s%s%s", preFix, "/", Uid, "/", 10, "/", articleImgPrefix, "/")
	//如果没有filepath文件目录就创建一个
	if _, err := os.Stat(filePath); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(filePath, os.ModePerm)
		}
	}

	for pos, fileHandle := range images {
		err = c.SaveUploadedFile(fileHandle, com.GetFileFullPath(fileHandle, filePath, strconv.Itoa(pos)))
		if err != nil {
			results.Failed(c, err.Error(), imageUploadErr)
			return
		}
	}
	imageFiles := com.ReadDirFileNames(filePath)
	var newImageFiles []string
	for _, v := range imageFiles {
		newImageFiles = append(newImageFiles, fmt.Sprintf("%s/%s", articleImgPrefix, v))
	}
	feedbackId := com.GenerateOrderID()
	timeNow := com.DBTimeStamp()
	_, err = model.DB().Table("app_user_feedback").Data(map[string]interface{}{
		"Uid":        Uid,
		"FeedbackId": feedbackId,
		"Content":    strings.Join(content, ","),
		"Photos":     strings.Join(newImageFiles, ","),
		"CreateTime": timeNow,
		"UpdateTime": timeNow,
	}).Insert()
	if err != nil {
		results.Failed(c, err.Error(), dbInsertError)
		return
	}
	results.Success(c, "提交成功,处理中...", apiSuccess, map[string]interface{}{
		"feedback_id": feedbackId,
		"content":     strings.Join(content, ","),
		"photos":      com.GetUserImgUrl(Uid, 10, strings.Join(newImageFiles, ",")),
	})
}

func (api *UserOp) GetUserFeedBack(c *gin.Context) {
	//获取post传过来的data
	uid := c.Query("uid")
	if uid == "" {
		results.Failed(c, "参数错误", paramError)
		return
	}
	res, err := model.DB().Table("app_user_feedback").Where("Uid", uid).
		Fields("Uid", "FeedbackId", "FeedbackState", "Content", "Photos", "CreateTime").Order("CreateTime asc").First()
	if res == nil || err != nil {
		results.Failed(c, "获取数据失败", serverInternalErr)
		return
	}

	results.Success(c, "查询成功", apiSuccess, map[string]interface{}{
		"uid":            uid,
		"feedback_id":    res["FeedbackId"].(string),
		"feedback_state": res["FeedbackState"].(int64),
		"content":        res["Content"].(string),
		"photos":         com.GetUserImgUrl(uid, 10, res["Photos"].(string)),
		"create_time":    res["CreateTime"].(time.Time).String(),
	})
}
