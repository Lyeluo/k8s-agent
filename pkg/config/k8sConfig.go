package config

import (
	"os"
	"path/filepath"
	"sync"

	_ "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var k8sOnce sync.Once
var clientset *kubernetes.Clientset

// 初始化k8s的配置
func init() {
	_ = GetK8sConfig()
}

func GetK8sConfig() *kubernetes.Clientset {
	k8sOnce.Do(func() {
		if instance == nil {
			clientset = getK8sClient()
		} else {
			clientset = &kubernetes.Clientset{}
		}

	})
	return clientset
}

/*
 * 获取k8s的client对象
 */
func getK8sClient() (clientset *kubernetes.Clientset) {

	//在 kubeconfig 中使用当前上下文环境，config 获取支持 url 和 path 方式
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homeDir(), ".kube", "config"))
	if err != nil {
		panic(err.Error())
	}

	// config, err := rest.InClusterConfig()
	// if err != nil {
	// 	panic(err.Error())
	// }

	// 根据指定的 config 创建一个新的 clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
