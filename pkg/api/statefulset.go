package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	statefulset := server.Group("/api/v1/statefulset/namespace/:namespace/")
	{
		statefulset.POST("/list", service.StatefulsetList)

		statefulset.POST("/create", service.StatefulsetCreate)

		statefulset.DELETE("/:name", service.StatefulsetDelete)

		statefulset.GET("/:name/get", service.StatefulsetGet)

		statefulset.POST("/:name/update", service.StatefulsetUpdate)
	}
}
