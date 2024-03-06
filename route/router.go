package route

import (
	//一定要导入这个Controller包，用来注册需要访问的方法
	//这里路由-由构架是添加-开发者仅在指定工程目录下controller.go文件添加宝即可

	// _ "gofly/app/admin"
	// _ "gofly/app/business"
	// _ "gofly/app/wxoffi"
	// _ "gofly/app/common"
	// _ "gofly/app/wxapp"
	// _ "gofly/app/home"
	_ "gofly/app"
	"net/http"

	"gofly/global"
	"gofly/route/middleware"
	"strings"
	"time"

	//工具
	"gofly/utils/gf"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 路由初始化
func InitRouter() *gin.Engine {
	//初始化路由
	R := gin.Default()
	R.SetTrustedProxies([]string{"127.0.0.1"})
	/**静态资源处理*/
	// a.1.前端项目静态资源
	// R.StaticFile("/favicon.ico", "./resource/webadmin/favicon.ico")
	//a.2.附件访问
	R.Static("/resource", "./resource")
	//a.3.业务后台
	R.Static("/webadmin", "./resource/webadmin")
	R.Static("/webbusiness", "./resource/webbusiness")
	R.LoadHTMLFiles("./resource/developer/template/install.html", "./resource/developer/template/isinstall.html")
	//访问域名根目录重定向
	R.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, global.App.Config.App.Rootview)
	})
	//控制台日志级别
	gin.SetMode(global.App.Config.App.RunlogType) //ReleaseMode 为方便调试，Gin 框架在运行的时候默认是debug模式，在控制台默认会打印出很多调试日志，上线的时候我们需要关闭debug模式，改为release模式。
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	R.MaxMultipartMemory = 8 << 20 // 8 MiB
	//0.跨域访问-注意跨域要放在gin.Default下
	var str_arr []string
	if global.App.Config.App.Allowurl != "" {
		str_arr = strings.Split(global.App.Config.App.Allowurl, `,`)
	} else {
		str_arr = []string{"http://localhost:8080"}
	}

	R.Use(cors.New(cors.Config{
		AllowOrigins: str_arr,
		// AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Content-Type", "Authorization", "Businessid", "verify-encrypt", "ignoreCancelToken", "verify-time"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//1.错误日志
	R.Use(gin.Logger(), middleware.CustomRecovery())
	//2.限流rate-limit 中间件
	R.Use(middleware.LimitHandler())
	//3.判断接口是否合法
	R.Use(middleware.ValidityAPi())
	//4.验证token
	R.Use(middleware.JwtVerify)
	//5.找不到路由
	R.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		c.JSON(404, gin.H{"code": 404, "message": "您" + method + "请求地址：" + path + "不存在！"})
	})
	//绑定基本路由，访问路径：/User/List
	gf.Bind(R)
	return R
}
