package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	deployment := server.Group("/api/v1/event/namespace/:namespace/")
	{
		deployment.POST("/list", service.EventList)

	}
}
