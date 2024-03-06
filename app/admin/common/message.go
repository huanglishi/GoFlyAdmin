package common

/**
* 系统消息
 */
import (
	"encoding/json"
	"gofly/model"
	"gofly/route/middleware"
	"gofly/utils/gf"
	"gofly/utils/results"
	"io"
	"reflect"

	"github.com/gin-gonic/gin"
)

func init() {
	gf.Register(&Message{}, reflect.TypeOf(Message{}).PkgPath())
}

type Message struct {
}

// 获取消息列表
func (api *Message) Get_list(c *gin.Context) {
	//用户
	getuser, _ := c.Get("user")
	user := getuser.(*middleware.UserClaims)
	usertype := 1 //用户类型
	list, err := model.DB().Table("common_message").Fields("id,type,title,path,content,isread,createtime").
		WhereIn("usertype", []interface{}{0, usertype}).Where("touid", user.ID).
		// Limit(pageSize).Page(pageNo).
		Order("id desc").Get()
	if err != nil {
		results.Failed(c, "加载数据失败", err)
	} else {
		var totalCount int64
		totalCount, _ = model.DB().Table("common_message").WhereIn("usertype", []interface{}{0, usertype}).Where("touid", user.ID).Count()
		results.Success(c, "获取全部列表", map[string]interface{}{
			"total": totalCount,
			"items": list,
		}, nil)
	}

}

// 设置为已读
func (api *Message) Read(c *gin.Context) {
	//获取post传过来的data
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	b_ids, _ := json.Marshal(parameter["ids"])
	var ids_arr []interface{}
	json.Unmarshal([]byte(b_ids), &ids_arr)
	res2, err := model.DB().Table("common_message").WhereIn("id", ids_arr).Data(map[string]interface{}{"isread": 1}).Update()
	if err != nil {
		results.Failed(c, "更新失败！", err)
	} else {
		msg := "更新成功！"
		if res2 == 0 {
			msg = "暂无数据更新"
		}
		results.Success(c, msg, res2, nil)
	}
}
