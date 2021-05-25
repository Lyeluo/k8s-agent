package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Server struct {
	Port int `JSON:"port"`
}

type Config struct {
	Server    Server
	LogConfig LogConfig
}

type LogConfig struct {
	Level      string `json:"level"`
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
}

// 加载配置文件
func init() {
	var config Config
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read conf failed, err:%s \n", err))
	}

	// 配置flag
	flagInit()

	//绑定配置文件到结构体
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
	// 监控配置文件变化
	viper.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量config
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生变更...")
		if err := viper.Unmarshal(&config); err != nil {
			panic(fmt.Errorf("unmarshal config failed, err:%s \n", err))
		}
	})
	// 初始化log
	if err := InitLogger(&config.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

}

// 定义命令行传参数
func flagInit() {
	pflag.Int("server.port", 8080, "server port set ")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
}
