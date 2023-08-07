package results

import (
	"gofly/route/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 请求成功的时候 使用该方法返回信息
func Success(ctx *gin.Context, msg string, data interface{}, exdata interface{}) {
	token := ctx.Request.Header.Get("Authorization")
	var newtoken interface{}
	if token != "" {
		tockenarr := middleware.Refresh(token)
		if tockenarr != nil {
			newtoken = tockenarr
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": msg,
		"data":    data,
		"exdata":  exdata,
		"token":   newtoken,
		"time":    time.Now().Unix(),
	})
}

// 请求成功的时候 使用该方法返回信息
func SuccessLogin(ctx *gin.Context, msg string, data interface{}, token, exdata interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": msg,
		"data":    data,
		"token":   token,
		"exdata":  exdata,
		"time":    time.Now().Unix(),
	})
}

// 请求失败的时候, 使用该方法返回信息
func Failed(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": msg,
		"data":    data,
		"time":    time.Now().Unix(),
	})
}
