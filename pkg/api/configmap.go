package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	configmap := server.Group("/api/v1/configmap/namespace/:namespace/")
	{
		configmap.POST("/list", service.ConfigMapList)

		configmap.POST("/create", service.ConfigMapCreate)

		configmap.DELETE("/:name", service.ConfigMapDelete)

		configmap.GET("/:name/get", service.ConfigMapGet)

		configmap.PATCH("/:name", service.ConfigMapPatch)

		// configmap.POST("/:name/update", service.ConfigMapUpdate)
	}
}
