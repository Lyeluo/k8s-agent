package config

import (
	"fmt"
	"net"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Server struct {
	Port         int    `JSON:"port"`
	SecretToken  string `JSON:"secretToken"`
	AgentAddress string `JSON:"address"`
	RegistryCron string `JSON:"registryCron"`
	AutoRegistry bool   `JSON:"autoRegistry"`
	AdminAddress string `JSON:"adminAddress"`
	ObjectId     string `JSON:"objectId"`
}

type Config struct {
	Server    Server
	LogConfig LogConfig
	Auth      Auth
}

type Auth struct {
	IsOpen      bool   `JSON:"open"`
	SecretToken string `JSON:"secretToken"`
}

type LogConfig struct {
	Level      string `JSON:"level"`
	Filename   string `JSON:"filename"`
	MaxSize    int    `JSON:"maxsize"`
	MaxAge     int    `JSON:"max_age"`
	MaxBackups int    `JSON:"max_backups"`
}

var config *Config

// 加载配置文件
func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	fmt.Println("开始读取配置文件。。。")

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
	// 服务端口
	pflag.Int("server.port", 8080, "server port set ")
	addrs, err := getClientIp()
	if err != nil {
		panic(err)
	}
	// 服务代理端地址
	pflag.String("server.address", addrs, "server address set ")
	// 与服务器通信的秘钥
	pflag.String("server.secretToken", "", "server secretToken set ")
	// server的服务端地址
	pflag.String("server.adminAddress", "", "admin address set ")
	// objectId
	pflag.String("server.objectId", "", "agent objectId set ")
	// 允许agent调用的secretToken
	pflag.String("auth.secretToken", "", "agent secretToken set ")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
}

func getClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", nil
}
