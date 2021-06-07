package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	event := server.Group("/api/v1/event/namespace/:namespace/")
	{
		event.POST("/list", service.EventList)

	}
}
