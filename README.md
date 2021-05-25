## 通用web框架开发模板
此项目为golang通用web项目开发基础框架，主要使用架构有：
- Gin: golang最主流web框架
- Zap: Uber开源的日志收集框架
- Viper: 读取配置文件的框架
## 添加配置
1. 添加配置文件，需将配置信息添加到`config.toml`文件。该文件中的配置会在变更后自动刷新。
2. 在加载配置文件`pkg.config.appConfig.go`文件中添加配置对应的结构体，并将该结构体设置为Config结构体的一个属性，如：
```go
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
```
注意：LogConfig与Config中的属性值都必须设置为大写，否则无法对该属性进行赋值
3. 获取配置使用viper的api，如：
```go
serverPort := viper.GetInt("server.port")
```
## 命令行参数
1. 命令行参数建议使用组件库中的`spf13/pflag`。添加位置在`pkg.config.appConfig.go`文件中
2. 使用pflag设置的命令行参数可以替换配置文件中的内容，如：
```go
func flagInit() {
	pflag.Int("server.port", 8080, "server port set ")
	pflag.Parse()
	// 将viper与命令行参数进行绑定
	viper.BindPFlags(pflag.CommandLine)
}
```
## 日志记录
1. 配置文件中有日志的相关配置，如：
```toml
[logConfig]
	level      = ""
    # 日志文件的位置
	filename   = "./web.log"
    # 在进行切割之前，日志文件的最大大小（以MB为单位）
	maxSize    = 200
    # 保留旧文件的最大天数
	maxAge     = 10
     # 保留旧文件的最大个数
	maxBackups = 3  
```
2. 使用日志时，只需要使用zap的工具包即可`zap.L().Xxx()`,如：
```go
zap.L().Error("我是一个错误日志!")
```
3. 日志信息会默认在控制台和文件中同时输出
## 路由开发
1. 所有的路由在pkg.api目录下的api包中开发，在init方法中写入路由，系统会自动加载。如：
```go
package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()

	v1 := server.Group("/api/v1")
    // 配置登陆拦截器
	v1.Use(MiddlewareAuth)
	{
		v1.GET("/:namespace/list", service.Helloworld)

	}
}
```
2. 不限定文件名称与`init()`方法数量，只需要在此包下的init方法中即可
## 工具类
所有工具类维护在pkg.util目录下,目前有：
- string.go: 操作字符串的工具类
- struct.go：操作结构体的工具类
- time.go： 格式化时间的工具类