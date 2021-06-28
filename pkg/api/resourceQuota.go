package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	namespace := server.Group("/api/v1/quota/:namespace")
	{

		namespace.GET("/list", service.QuotaList)

		namespace.GET("/:name", service.QuotaGet)

		namespace.POST("/create", service.QuotaCreate)

		namespace.POST("/update", service.QuotaUpdate)

		namespace.DELETE("/:name", service.QuotaDelete)
	}
}
