package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	v1Service := server.Group("/api/v1/service/namespace/:namespace/")
	{
		v1Service.POST("/list", service.ServiceList)

		v1Service.POST("/create", service.ServiceCreate)

		v1Service.DELETE("/:name", service.ServiceDelete)

		v1Service.GET("/:name/get", service.ServiceGet)

		v1Service.PATCH("/:name", service.ServicePatch)

		v1Service.POST("/:name/update", service.ServiceUpdate)
	}
}
