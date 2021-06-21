package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	namespace := server.Group("/api/v1/limitRange/:namespace")
	{

		namespace.GET("/list", service.LimitRangeList)

		namespace.GET("/:name", service.LimitRangeGet)

		namespace.POST("/create", service.LimitRangeCreate)

		namespace.POST("/update", service.LimitRangeUpdate)

		namespace.DELETE("/:name", service.LimitRangeDelete)
	}
}
