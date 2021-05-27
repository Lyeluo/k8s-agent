package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	deployment := server.Group("/api/v1/pod/namespace/:namespace/")
	{
		deployment.GET("/deployment/:deployment/list", service.PodListByDeployment)

		deployment.PATCH("/:name", service.PodPatch)
	}
}
