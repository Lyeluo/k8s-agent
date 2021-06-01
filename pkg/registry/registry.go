package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	_ "k8s.agent/pkg/config"
)

var (
	c *cron.Cron
)

func init() {

	autoRegistry := viper.GetBool("server.autoRegistry")

	if !autoRegistry {
		return
	}

	if err := registry(); err != nil {
		zap.L().Sugar().Errorf("注册地址到中心失败，原因：%s", err)
		panic(err)
	}
	// 开启定时任务
	registryJob()

	//监听退出
	exitRegistry()
}

func exitRegistry() {
	//创建监听退出chan
	chann := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(chann, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range chann {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("程序已关闭，正在退出。。。。。", s)
				deregistry()

			default:
				fmt.Println("other", s)
			}
		}
	}()

}

// 定时上报信息到中心
func registryJob() {

	c = cron.New(cron.WithSeconds())

	spec := viper.GetString("server.registryCron")

	c.AddFunc(spec, func() {
		err := registry()
		if err != nil {
			zap.L().Sugar().Errorf("定时上报地址失败，原因：%s", err)
		}
	})

	c.Start()
}

type RegistryVo struct {
	Port     int    `JSON:"port"`
	Address  string `JSON:"address"`
	ObjectId string `JSON:"objectId"`
}

type RegistryClient struct {
	RegistryReq   *http.Request
	DeRegistryReq *http.Request
}

var once sync.Once
var instance *RegistryClient

func GetRegistryClient() *RegistryClient {
	once.Do(func() {
		if instance == nil {
			// 初始化web对象，此处可以设置全局配置
			instance = getRegistryClient()
		} else {
			instance = &RegistryClient{}
		}
	})
	return instance
}

//初始化registry的client
func getRegistryClient() (registryClient *RegistryClient) {

	adminAddress := viper.GetString("server.adminAddress")

	registryURL := strings.Join([]string{adminAddress, "/registry/up"}, "")
	registryVo := &RegistryVo{
		Port:     viper.GetInt("server.port"),
		Address:  viper.GetString("server.address"),
		ObjectId: viper.GetString("server.objectId"),
	}

	body, err := json.Marshal(registryVo)
	if err != nil {
		zap.L().Sugar().Errorf("注册地址到中心失败，原因：%s", err)
		return
	}
	//  注册的httpclient
	registryReq, err := http.NewRequest("POST", registryURL, bytes.NewBuffer(body))
	if err != nil {
		zap.L().Sugar().Errorf("注册地址到中心失败，原因：%s", err)
	}
	registryReq.Header.Set("secretToken", viper.GetString("server.secretToken"))

	// 注销的httpclient
	deRegistryURL := strings.Join([]string{adminAddress, "/registry/down"}, "")
	deRegistryReq, err := http.NewRequest("GET", deRegistryURL, nil)
	if err != nil {
		zap.L().Sugar().Errorf("注册地址到中心失败，原因：%s", err)
	}
	deRegistryReq.Header.Set("secretToken", viper.GetString("server.secretToken"))

	registryClient = &RegistryClient{
		RegistryReq:   registryReq,
		DeRegistryReq: deRegistryReq,
	}
	return
}

func registry() (err error) {
	client := &http.Client{}
	resp, err := client.Do(getRegistryClient().RegistryReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func deregistry() {
	// 停止定时任务
	c.Stop()

	// 1.注销
	client := &http.Client{}
	_, err := client.Do(getRegistryClient().DeRegistryReq)
	if err != nil {
		zap.L().Sugar().Errorf("下线agent失败，报错内容： %s", err)
	}
	// 退出程序
	os.Exit(0)
}
