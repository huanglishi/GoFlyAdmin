package appUser

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	com "gofly/app/user/utils"
	"gofly/global"
	"gofly/model"
	"gofly/utils"
	"gofly/utils/results"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func init() {
	utils.Register(&UserStatement{}, reflect.TypeOf(UserStatement{}).PkgPath())
}

type Comment struct {
	MomentId    string   `json:"moment_id"`
	Moment      Moment   `gorm:"foreignKey:MomentId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Uid         string   `json:"uid"`
	User        UsrInfo  `gorm:"foreignKey:Uid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ParentId    *int     `json:"parent_id"`
	Parent      *Comment `gorm:"foreignKey:ParentId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Content     string   `json:"content"`
	CreatedTime string   `json:"created_time"`
	UpdatedTime string   `json:"updated_time"`
}
type Moment struct {
	Uid          string `json:"uid"`
	MomentId     string `json:"moment_id"`
	Content      string `json:"content"`
	LikeCount    uint   `json:"like_count"`
	MomentState  uint8  `json:"moment_state"`
	CommentCount uint   `json:"comment_count"`
	Photos       string `json:"photos"`
	CreateTime   string `json:"create_time"`
	UpdateTime   string `json:"update_time"`
}

type CommentTree struct {
	CommentId   string                 `json:"commentId"`
	Content     string                 `json:"content"`
	Author      map[string]interface{} `json:"author"`
	CreatedTime string                 `json:"created_time"`
	Children    []*CommentTree         `json:"children"`
}

// 用于自动注册路由
type UserStatement struct {
}

// 发表说说+上传多张图片
func (api *UserStatement) AddMoment(c *gin.Context) {
	Uid := c.Query("uid")
	formHandle, _ := c.MultipartForm()
	images := formHandle.File["file"]
	content := formHandle.Value["content"]
	//userContent := c.Request.MultipartForm.Value["content"]
	if Uid == "" {
		results.Failed(c, "参数错误！", paramError)
		return
	}
	res, err := model.DB().Table("app_user").Fields("MomentCount").Where("Uid", Uid).First()
	if res == nil || err != nil {
		results.Failed(c, "账号不存在!", userNotExist)
		return
	}

	preFix := global.App.Config.Userconf.ImagePrePath
	articleImgPrefix := com.GenerateOrderID()
	filePath := fmt.Sprintf("%s%s%s%s%d%s%s%s", preFix, "/", Uid, "/", 9, "/", articleImgPrefix, "/")
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
	momentId := com.GenerateOrderID()
	_, err = model.DB().Table("app_user_moment").Data(map[string]interface{}{
		"MomentId":    momentId,
		"Uid":         Uid,
		"Content":     strings.Join(content, ","),
		"Photos":      strings.Join(newImageFiles, ","),
		"MomentState": 0, //审核中
		"CreateTime":  com.DBTimeStamp(),
		"UpdateTime":  com.DBTimeStamp(),
	}).Insert()
	if err != nil {
		results.Failed(c, err.Error(), dbInsertError)
		return
	}
	//更新用户moment数量
	currentMomentCount := res["MomentCount"].(int64)
	_, err = model.DB().Table("app_user").Data(map[string]interface{}{
		"MomentCount": currentMomentCount + 1,
		"UpdateTime":  com.DBTimeStamp(),
	}).Insert()
	results.Success(c, "发表成功,审核中...", apiSuccess, map[string]interface{}{
		"moment_id": momentId,
		"photos":    com.GetUserImgUrl(Uid, 9, strings.Join(newImageFiles, ",")),
	})
}

// 发表评论
func (api *UserStatement) AddComment(c *gin.Context) {
	var parameter map[string]interface{}
	bytes, _ := io.ReadAll(c.Request.Body)
	json.Unmarshal(bytes, &parameter)
	if parameter["uid"] == nil || parameter["moment_id"] == nil || parameter["content"] == nil || parameter["parent_id"] == nil {
		results.Failed(c, "参数错误", paramError)
		return
	}
	momentId := parameter["moment_id"].(string)
	commentId := com.GenerateOrderID()
	_, err := model.DB().Table("app_user_comment").Data(map[string]interface{}{
		"MomentId":   parameter["moment_id"].(string),
		"CommentId":  commentId,
		"Uid":        parameter["uid"].(string),
		"ParentId":   parameter["parent_id"].(string),
		"Content":    parameter["content"].(string),
		"CreateTime": com.DBTimeStamp(),
		"UpdateTime": com.DBTimeStamp(),
	}).Insert()
	if err != nil {
		results.Failed(c, err.Error(), dbInsertError)
		return
	}
	//更新 moment
	momentInfo, err := model.DB().Table("app_user_moment").Where("MomentId", momentId).Fields("CommentCount").First()
	_, err = model.DB().Table("app_user_moment").Where("MomentId", momentId).Data(map[string]interface{}{
		"CommentCount": momentInfo["CommentCount"].(int64) + 1,
		"UpdateTime":   com.DBTimeStamp(),
	}).Update()

	results.Success(c, "评论成功.", apiSuccess, map[string]interface{}{"comment_id": commentId})
}

// 获取评论
func (api *UserStatement) GetComments(c *gin.Context) {
	var commentTrees []CommentTree
	momentId := c.Query("moment_id")
	if momentId == "" {
		results.Failed(c, "参数错误", paramError)
		return
	}
	commentTrees = GetMomentComment(momentId)
	results.Success(c, "查询成功", apiSuccess, commentTrees)
}
func GetMomentComment(momentId string) []CommentTree {
	var commentTrees []CommentTree
	comments, _ := model.DB().Table("app_user_comment").Where("MomentId", momentId).Where("ParentId", "=", "").Get()
	for _, comment := range comments {
		commentId := comment["CommentId"].(string)
		uid := comment["Uid"].(string)
		userInfo, _ := model.DB().Table("app_user").Fields("UserName", "FaceImg").Where("Uid", uid).First()
		commentTree := CommentTree{
			CommentId: commentId,
			Content:   comment["Content"].(string),
			Author: gin.H{
				"uid":       uid,
				"user_name": userInfo["UserName"].(string),
				"face_img":  com.GetUserImgUrl(uid, 1, userInfo["FaceImg"].(string))},
			CreatedTime: comment["CreateTime"].(time.Time).String(),
			Children:    []*CommentTree{},
		}
		GetMomentCommentChild(commentId, &commentTree)
		commentTrees = append(commentTrees, commentTree)
	}
	return commentTrees
}

// 查询二级及以下的多级评论
func GetMomentCommentChild(parentId string, commentTree *CommentTree) {
	comments, _ := model.DB().Table("app_user_comment").Where("ParentId", parentId).Get()
	for _, v := range comments {
		cid := v["CommentId"].(string)
		uid := v["Uid"].(string)
		if cid == uid {
			continue
		}
		userInfo, _ := model.DB().Table("app_user").Fields("UserName", "FaceImg").Where("Uid", uid).First()
		child := CommentTree{
			CommentId: cid,
			Content:   v["Content"].(string),
			Author: gin.H{
				"uid":       uid,
				"user_name": userInfo["UserName"].(string),
				"face_img":  com.GetUserImgUrl(uid, 1, userInfo["FaceImg"].(string)),
			},
			CreatedTime: v["CreateTime"].(time.Time).String(),
			Children:    []*CommentTree{},
		}
		commentTree.Children = append(commentTree.Children, &child)
		GetMomentCommentChild(cid, &child)
	}
}

// 用户删除动态+背景图
func (api *UserStatement) DeleteStatement(c *gin.Context) {
	momentId := c.Query("moment_id") //图片类型
	if momentId == "" {
		results.Failed(c, "参数错误", paramError)
		return
	}
	res, err := model.DB().Table("app_user_moment").Fields("Photos", "Uid").
		Where("MomentId", momentId).First()
	if res == nil || err != nil {
		results.Failed(c, "动态不存在!", userNotExist)
		return
	}
	//删除动态图片
	photos := res["Photos"].(string)
	if len(photos) > 0 {
		uid := res["Uid"].(string)
		files := strings.Split(photos, ",")
		preFix := global.App.Config.Userconf.ImagePrePath
		filePath := fmt.Sprintf("%s%s%s%d%s", preFix, uid, "/", 9, "/")
		prePath := strings.Split(files[0], "/")
		os.RemoveAll(filePath + prePath[0])
	}
	//删除动态
	_, err = model.DB().Table("app_user_moment").Where("MomentId", momentId).Delete()
	if err != nil {
		//TODO add log
	}
	results.Success(c, "删除成功", apiSuccess, nil)
}

// 获取单个用户的全部动态
func (api *UserStatement) GetUserMoments(c *gin.Context) {
	uid := c.Query("uid") //图片类型
	if uid == "" {
		results.Failed(c, "参数错误", paramError)
		return
	}
	res, err := model.DB().Table("app_user_moment").Fields("MomentId", "Content", "LikeCount", "CommentCount", "Photos", "CreateTime").
		Where("Uid", uid).Order("CreateTime desc").Get()
	if res == nil || err != nil || len(res) < 1 {
		results.Failed(c, "动态不存在!", userNotExist)
		return
	}
	var resp []map[string]interface{}
	//删除动态图片
	for _, v := range res {
		var rst = map[string]interface{}{
			"MomentId":     v["MomentId"].(string),
			"CommentCount": v["CommentCount"].(int64),
			"LikeCount":    v["LikeCount"].(int64),
			"Photos":       v["Photos"].(string),
			"CreateTime":   v["CreateTime"].(time.Time).String(),
		}
		resp = append(resp, rst)
	}
	//删除动态
	results.Success(c, "查询成功", apiSuccess, resp)
}

func userMomentListQueryStr(longitude, latitude interface{}, sexual, limit, offset uint) (queryStr string) {
	/*
		SELECT s.Uid,s.UserName,s.Age,s.Signature,s.SelfIntroduction,s.FriendTag,s.FaceImg,s.IdentityState,
		(st_distance (point (Longitude, Latitude),point(108.979654,34.358591)) / 0.0111) AS distance
		 FROM app_user s WHERE Sexual=0 AND FaceImgState=1 AND BgImgState=1 ORDER BY distance LIMIT 10 OFFSET 0
	*/
	//未获取经纬度时
	if longitude == nil || latitude == nil {
		if sexual == 0 { //查询女性用户
			queryStr = fmt.Sprintf(
				"SELECT DISTINCT s.Uid,s.UserName,s.Age,s.Sexual,s.FriendTag,s.FaceImg,s.IdentityState,s.MeetUnlockGift,MomentState\n"+
					"FROM app_user s cross join app_user_moment b ON s.Uid=b.Uid\n"+
					"WHERE Sexual=0 AND FaceImgState=1 AND BgImgState=1 AND MomentCount>0 And MomentState=1\n"+
					"ORDER BY b.UpdateTime desc LIMIT %d OFFSET %d",
				limit, offset)
		} else if sexual == 1 { //查询男性用户
			queryStr = fmt.Sprintf(
				"SELECT DISTINCT s.Uid,s.UserName,s.Age,s.Sexual,s.FriendTag,s.FaceImg,s.IdentityState,s.CarType,s.MeetGift,CustomGift\n"+
					"FROM app_user s cross join app_user_moment b ON s.Uid=b.Uid\n"+
					"WHERE Sexual=1 AND FaceImgState=1 AND BgImgState=1 AND MomentCount>0 And MomentState=1\n"+
					"ORDER BY b.UpdateTime desc LIMIT %d OFFSET %d",
				limit, offset)
		}
	} else {
		if sexual == 0 {
			queryStr = fmt.Sprintf(
				"SELECT DISTINCT s.Uid,s.UserName,s.Age,s.Sexual,s.FriendTag,s.FaceImg,s.IdentityState,s.MeetUnlockGift,MomentState,\n"+
					"(st_distance (point (Longitude, Latitude),point(%s,%s) ) / 0.0111) AS distance "+
					"FROM app_user s cross join app_user_moment b ON s.Uid=b.Uid\n"+
					"WHERE Sexual=0 AND FaceImgState=1 AND BgImgState=1 AND MomentCount>0 And MomentState=1\n"+
					"ORDER BY distance LIMIT %d OFFSET %d",
				longitude.(string), latitude.(string), limit, offset)
		} else if sexual == 1 {
			queryStr = fmt.Sprintf(
				"SELECT DISTINCT s.Uid,s.UserName,s.Age,s.Sexual,s.FriendTag,s.FaceImg,s.IdentityState,s.CarType,s.MeetGift,s.CustomGift,MomentState,\n"+
					"(st_distance (point (Longitude, Latitude),point(%s,%s) ) / 0.0111) AS distance\n"+
					"FROM app_user s cross join app_user_moment b ON s.Uid=b.Uid\n"+
					"WHERE Sexual=1 AND FaceImgState=1 AND BgImgState=1 AND MomentCount>0 And MomentState=1\n"+
					"ORDER BY distance LIMIT %d OFFSET %d",
				longitude.(string), latitude.(string), limit, offset)
		}
	}
	fmt.Println(queryStr)
	return queryStr
}

// 获取动态列表
func (api *UserStatement) GetMomentList(c *gin.Context) {
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
	queryStr := userMomentListQueryStr(parameter["longitude"], parameter["latitude"], sexual, limit, offset)
	res, err := model.DB().Query(queryStr)
	if res == nil || err != nil {
		return
	}
	var resp []map[string]interface{}
	for _, v := range res {
		data := com.ModelDBUsrInfo(v)
		uid := v["Uid"].(string)
		momentInfo, _ := model.DB().Table("app_user_moment").Fields("MomentId", "Content", "LikeCount", "CommentCount", "Photos", "UpdateTime").
			Where("Uid", uid).Order("UpdateTime desc").First()
		momentId := momentInfo["MomentId"].(string)
		data["moment_id"] = momentId
		data["content"] = momentInfo["Content"].(string)
		data["comment_count"] = momentInfo["CommentCount"].(int64)
		data["like_count"] = momentInfo["LikeCount"].(int64)
		data["photos"] = com.GetUserImgUrl(uid, 9, momentInfo["Photos"].(string)) //
		data["update_time"] = momentInfo["UpdateTime"].(time.Time).String()
		resp = append(resp, data)
	}
	//删除动态
	results.Success(c, "查询成功", apiSuccess, resp)
}

// 动态点赞数量
func (api *UserStatement) PostMomentLikes(c *gin.Context) {
	momentId := c.Query("moment_id")

	momentInfo, _ := model.DB().Table("app_user_moment").Fields("LikeCount").
		Where("MomentId", momentId).First()
	currentLikeCount := momentInfo["LikeCount"].(int64)

	model.DB().Table("app_user_moment").Where("MomentId", momentId).Data(
		map[string]interface{}{"LikeCount": currentLikeCount + 1}).Update()
	results.Success(c, "点赞成功", apiSuccess, map[string]interface{}{
		"like_count": currentLikeCount + 1})
}

// 图片位置交换
func (api *UserOp) PhotoPosChange(c *gin.Context) {
	imgType := c.Query("photo_type") //图片类型
	pos := c.Query("pos")            //图片数量
	Uid := c.Query("uid")
	if imgType != userImagePhotos || Uid == "" || len(pos) != 3 {
		results.Failed(c, "参数错误", paramError)
		return
	}
	dbFiled, _ := getImagePosition(imgType)
	res, err := model.DB().Table("app_user").Fields("Uid", dbFiled).
		Where("Uid", Uid).First()
	if res == nil || err != nil {
		results.Failed(c, "账号不存在!", userNotExist)
		return
	}
	if res[dbFiled] == nil || res[dbFiled].(string) == "" {
		results.Failed(c, "无相关图片,请重新上传!", imageXchangeErr)
		return
	}
	preFix := global.App.Config.Userconf.ImagePrePath
	filePath := fmt.Sprintf("%s%s%s%d%s", preFix, Uid, "/", 9, "/")
	fieldsStr := res[dbFiled].(string)
	fieldsArr := strings.Split(fieldsStr, ",")

	exchanges := strings.Split(pos, "_")
	pos0, _ := strconv.Atoi(exchanges[0])
	pos1, _ := strconv.Atoi(exchanges[1])

	tmpName := uuid.NewV4().String()
	endFix := strings.Split(fieldsArr[pos0], ".")
	tmpName = tmpName + "." + endFix[1]
	com.FileExchange(filePath+fieldsArr[pos0], filePath+tmpName)
	com.FileExchange(filePath+fieldsArr[pos1], filePath+fieldsArr[pos0])
	com.FileExchange(filePath+tmpName, filePath+fieldsArr[pos1])

	os.Remove(filePath + tmpName)

	results.Success(c, "交换成功", apiSuccess, nil)
}
