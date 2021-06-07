package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	node := server.Group("/api/v1/node")
	{
		node.POST("/list", service.NodeList)

		node.PATCH("/:name/update", service.NodeUpdate)
	}
}
