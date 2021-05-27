package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	deployment := server.Group("/api/v1/service/namespace/:namespace/")
	{
		deployment.GET("/list", service.ServiceList)

		deployment.POST("/create", service.ServiceCreate)

		deployment.DELETE("/:name", service.ServiceDelete)

		deployment.PATCH("/:name", service.ServicePatch)

		deployment.POST("/:name/update", service.ServiceUpdate)
	}
}
