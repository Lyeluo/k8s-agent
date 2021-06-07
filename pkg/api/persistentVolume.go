package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	persistentVolume := server.Group("/api/v1/persistentVolume")
	{
		persistentVolume.POST("/list", service.PersistentVolumeList)

		persistentVolume.POST("/create", service.PersistentVolumeCreate)

		persistentVolume.PATCH("/:name/update", service.PersistentVolumePatch)

		persistentVolume.DELETE("/:name", service.PersistentVolumeDelete)
	}
}
