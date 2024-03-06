package service

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"gofly/model"
	"gofly/utils/gf"
	"log"
	"reflect"
	"sort"
	"strings"
	"time"

	"gofly/utils/gform"

	"github.com/gin-gonic/gin"
)

/**
*使用 Index 是省略路径中的index
*本路径为： /admin/user/login -省去了index
 */
func init() {
	gf.Register(&Index{}, reflect.TypeOf(Index{}).PkgPath())
}

type Index struct{}

/**
* url接口配置
 */
func (api *Index) GetPost_api(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	Data, err := model.DB().Table("business_wxsys_officonfig").Where("businessID", id).Fields("Token,accountID,businessID").First()
	if err != nil {
		log.Println("获取账号信息失败!")
		return
	} else {
		if c.Request.Method == "GET" {
			//微信参数
			signature := c.DefaultQuery("signature", "")
			timestamp := c.DefaultQuery("timestamp", "")
			nonce := c.DefaultQuery("nonce", "")
			echostr := c.DefaultQuery("echostr", "")
			ok := CheckSignature(signature, timestamp, nonce, Data["Token"].(string))
			if !ok {
				log.Println("微信公众号接入校验失败!")
				return
			}
			log.Println("微信公众号接入校验成功!")
			_, _ = c.Writer.WriteString(echostr)
		} else { //post请求
			log.Println("post请求!")
			openid := c.DefaultQuery("openid", "")
			log.Printf("Openid: %s\n", openid)
			postdata, err := c.GetRawData()
			if err != nil {
				log.Fatalln(err)
			}
			msgText, err := ReceiveCommonMsg(postdata)
			if err != nil {
				log.Fatalln(err)
			}
			if msgText.Event == "subscribe" { //关注事件

				Onsubscribe(msgText, openid, Data)
			} else if msgText.Event == "unsubscribe" { //取消订阅
				model.DB().Table("business_wxsys_user").Where("openid", openid).Data(map[string]interface{}{"subscribe": 0}).Update()
			}
			log.Printf("[消息接收] - 收到消息, 消息类型为: %s, FromUserName: %s\n", msgText.Event, msgText.FromUserName)
		}
	}
}

// 处理关注事件
func Onsubscribe(msgText WxReceiveCommonMsg, openid string, Data gform.Data) {
	//判断账号是否存在
	log.Println("判断账号是否存在!")
	user, _ := model.DB().Table("business_wxsys_user").Where("openid", openid).Fields("id").First()
	if user == nil {
		log.Println("新增账号!")
		userid, err := model.DB().Table("business_wxsys_user").Data(map[string]interface{}{
			"openid":     openid,
			"accountID":  Data["accountID"],
			"businessID": Data["businessID"],
			"subscribe":  1,
			"avatar":     "resource/staticfile/avatar.png",
			"createtime": time.Now().Unix(),
		}).InsertGetId()
		model.DB().Table("business_wxsys_user").Data(map[string]interface{}{
			"nickname": fmt.Sprintf("U_%v", userid),
		}).Where("id", userid).Update()
		log.Printf("添加失败: %s\n", err)
	} else {
		model.DB().Table("business_wxsys_user").Where("id", user["id"]).Data(map[string]interface{}{"subscribe": 1}).Update()
	}
}

// WxReceiveCommonMsg 接收普通消息
type WxReceiveCommonMsg struct {
	ToUserName   string //接收者 开发者 微信号
	FromUserName string //发送者 发送方帐号（一个OpenID）
	Content      string //文本内容
	CreateTime   int64  //创建时间
	MsgType      string //消息类型
	MsgId        int64  //消息id
	PicUrl       string //图片url
	MediaId      string //媒体id
	Event        string //事件类型，VIEW
	EventKey     string //事件KEY值，设置的跳转URL
	MenuId       string
	Format       string
	Recognition  string
	ThumbMediaId string //缩略图媒体ID
}

// WxReceiveFunc (接收到消息之后，会将消息交于这个函数处理)
var WxReceiveFunc func(msg WxReceiveCommonMsg) error

// 处理接口事件
func ReceiveCommonMsg(msgData []byte) (WxReceiveCommonMsg, error) {
	fmt.Printf("received weixin msgData:\n%s\n", msgData)
	msg := WxReceiveCommonMsg{}
	err := xml.Unmarshal(msgData, &msg)
	if WxReceiveFunc == nil {
		return msg, err
	}
	err = WxReceiveFunc(msg)
	return msg, err
}

// CheckSignature 微信公众号签名检查
func CheckSignature(signature, timestamp, nonce, token string) bool {
	arr := []string{timestamp, nonce, token}
	// 字典序排序
	sort.Strings(arr)

	n := len(timestamp) + len(nonce) + len(token)
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(arr); i++ {
		b.WriteString(arr[i])
	}

	return Sha1(b.String()) == signature
}

// 进行Sha1编码
func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
