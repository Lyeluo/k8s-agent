package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	namespace := server.Group("/api/v1/limitRange")
	{
		namespace.GET("/:namespace/get", service.NamespaceGet)

		namespace.POST("/create", service.NamespaceCreate)

		namespace.POST("/update", service.NamespaceUpdate)

		namespace.DELETE("/:name", service.NamespaceDelete)
	}
}
