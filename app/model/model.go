package model

import (
	"fmt"
	"gofly/global"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2" //数据库操作
	"go.uber.org/zap"
)

var err error
var engin *gorose.Engin

// 取得数据库连接实例
func MyInit(applog *zap.Logger) {
	global.App.Config.InitializeConfig()
	engin, err = gorose.Open(&gorose.Config{Driver: global.App.Config.DBconf.Driver, Dsn: global.App.Config.DBconf.Source})
	if err != nil {
		applog.Info(fmt.Sprintf("数据库连接实例错误: %v\n", err))
	} else {
		engin.GetExecuteDB().SetMaxIdleConns(10)                  //连接池最大空闲连接数,不设置, 默认无
		engin.GetExecuteDB().SetMaxOpenConns(50)                  // 连接池最大连接数,不设置, 默认无限
		engin.GetExecuteDB().SetConnMaxLifetime(59 * time.Second) //时间比超时时间短
	}
}

// controller层调用
func DB() gorose.IOrm {
	return engin.NewOrm()
}

// 取得总行数
func GetTotal(tablename string, wheres map[string]interface{}) int64 {
	total, _ := DB().Table(tablename).Where(wheres).Count()
	return total
}
