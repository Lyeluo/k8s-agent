package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	pod := server.Group("/api/v1/pod/namespace/:namespace/")
	{
		pod.GET("/deployment/:deployment/list", service.PodListByDeployment)

		pod.PATCH("/:name", service.PodPatch)

		pod.POST("/:name/logs", service.PodLogs)

		pod.POST("/:name/exec", service.PodExec)
	}
}
