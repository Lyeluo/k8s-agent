package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	deployment := server.Group("/api/v1/deployment/namespace/:namespace/")
	{
		deployment.GET("/list", service.DeploymentList)

		deployment.POST("/create", service.DeploymentCreate)

		deployment.DELETE("/:name", service.DeploymentDelete)

		deployment.PATCH("/:name", service.DeploymentPatch)

		deployment.POST("/:name/update", service.DeploymentUpdate)
	}
}
