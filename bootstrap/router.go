package bootstrap

import (
	"context"
	"fmt"
	"gofly/global"
	"gofly/model"
	"gofly/route"
	"gofly/utils/gf"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// 优雅重启/停止服务器
func RunServer() {
	//加载路由
	r := route.InitRouter()
	routes := ""
	for _, route := range r.Routes() {
		if !strings.Contains(route.Path, "/admin/") && route.Path != "/" && !strings.Contains(route.Path, "/*filepath") {
			routes = routes + fmt.Sprintf("%v\n", route.Path)
		}
	}
	filePath := "runtime/app/routers.txt"
	gf.WriteToFile(filePath, routes)
	model.MyInit(1) //初始化数据
	if global.App.Config.App.Env == "dev" {
		fmt.Printf("\n %c[1;40;32m%s%c[0m\n", 0x1B, "在浏览器访问：​​http://127.0.0.1:"+global.App.Config.App.Port+"/common/install/index ​进行​安装​", 0x1B)
		r.Run(":" + global.App.Config.App.Port)
	} else { //优雅-生成环境使用
		//换一种启动方式
		srv := &http.Server{
			Addr:    ":" + global.App.Config.App.Port,
			Handler: r,
		}
		global.App.Log.Info("启动端口：" + global.App.Config.App.Port)
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				str := fmt.Sprintf("listen: %s\n", err) //拼接字符串
				global.App.Log.Error(str)
			}
		}()

		// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		global.App.Log.Info("关闭服务器 ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			str := fmt.Sprintf("服务器关闭： %s\n", err) //拼接字符串
			global.App.Log.Error(str)
		}
		global.App.Log.Info("服务器正在退出 ...")
	}
}
