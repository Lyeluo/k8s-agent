package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	deployment := server.Group("/api/v1/deployment/namespace/:namespace/")
	{
		deployment.POST("/list", service.DeploymentList)

		deployment.POST("/create", service.DeploymentCreate)

		deployment.GET("/:name/get", service.DeploymentGet)

		deployment.DELETE("/:name", service.DeploymentDelete)

		deployment.PATCH("/:name", service.DeploymentPatch)

		deployment.POST("/:name/update", service.DeploymentUpdate)
	}
}
