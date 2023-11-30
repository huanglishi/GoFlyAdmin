package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	DBconf DBconf `yaml:"dbconf"`
	App    App    `yaml:"app"`
	Jwt    Jwt    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Log    Log    `mapstructure:"log" json:"log" yaml:"log"`
}

// 读取Yaml配置文件，并转换成Config对象  struct结构
func (config *Config) InitializeConfig() *Config {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// fmt.Println("path=", path)
	vip := viper.New()
	vip.AddConfigPath(path + "/resource") //设置读取的文件路径
	vip.SetConfigName("config")           //设置读取的文件名
	vip.SetConfigType("yaml")             //设置文件的类型
	//尝试进行配置读取
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}
	// 监听配置文件
	vip.WatchConfig()
	vip.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := vip.Unmarshal(&config); err != nil {
			fmt.Println(err)
		}
	})

	err = vip.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return config
}
func listenSignal() {
	go func() {
		// 执行重启命令
		cmd := exec.Command("gofly", "run", "daemon", "restart")
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println(err)
		}
		defer stdout.Close()

		if err := cmd.Start(); err != nil {
			panic(err)
		}
		reader := bufio.NewReader(stdout)
		//实时循环读取输出流中的一行内容
		for {
			line, err2 := reader.ReadString('\n')
			if err2 != nil || io.EOF == err2 {
				break
			}
			fmt.Print(line)
		}

		if err := cmd.Wait(); err != nil {
			fmt.Println(err)
		}
		opBytes, _ := io.ReadAll(stdout)
		fmt.Print(string(opBytes))

	}()
}
