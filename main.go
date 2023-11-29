package main

import (
	"fmt"
	"gofly/bootstrap"
	"gofly/global"
	"runtime"
	"strconv"
)

func main() {
	// 初始化配置
	global.App.Config.InitializeConfig()
	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("项目启动成功")
	//加载配置
	cpu_num, _ := strconv.Atoi(global.App.Config.App.CPUnum)
	mycpu := runtime.NumCPU()
	if cpu_num > mycpu { //如果配置cpu核数大于当前计算机核数，则等当前计算机核数
		cpu_num = mycpu
	}
	if cpu_num > 0 {
		runtime.GOMAXPROCS(cpu_num)
		global.App.Log.Info(fmt.Sprintf("当前计算机核数: %v个,调用：%v个", mycpu, cpu_num))
	} else {
		runtime.GOMAXPROCS(mycpu)
		global.App.Log.Info(fmt.Sprintf("当前计算机核数: %v个,调用：%v个", mycpu, mycpu))
	}

	// 启动服务器
	bootstrap.RunServer()
}
