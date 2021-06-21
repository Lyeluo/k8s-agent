package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	namespace := server.Group("/api/v1/metrics")
	{
		namespace.POST("/node", service.MetricsNode)

		namespace.GET("/namespace/:namespace", service.MetricsNamespace)
	}
}
