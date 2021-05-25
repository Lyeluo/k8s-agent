package config

import (
	"sync"

	"github.com/gin-gonic/gin"
	"k8s.agent/pkg/filter"
)

var once sync.Once
var instance *gin.Engine

func GetWebServer() *gin.Engine {
	once.Do(func() {
		if instance == nil {
			// 初始化web对象，此处可以设置全局配置
			instance = gin.New()
			instance.Use(GinLogger(), GinRecovery(true), filter.MiddlewareAuth)
		} else {
			instance = &gin.Engine{}
		}
	})
	return instance
}
