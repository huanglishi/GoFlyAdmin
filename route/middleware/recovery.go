package middleware

import (
	"gofly/global"
	"gofly/utils/gf"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func CustomRecovery() gin.HandlerFunc {
	//加载配置
	conf := global.App.Config
	tiemstr := time.Now().Format("200601/02")
	return gin.RecoveryWithWriter(
		&lumberjack.Logger{
			Filename:   conf.Log.RootDir + "/" + tiemstr + "_err.log",
			MaxSize:    conf.Log.MaxSize,
			MaxBackups: conf.Log.MaxBackups,
			MaxAge:     conf.Log.MaxAge,
			Compress:   conf.Log.Compress,
		},
		gf.ServerError)
}
