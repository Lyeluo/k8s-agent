package api

import (
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/service"
)

func init() {

	server := config.GetWebServer()
	persistentVolumeClaim := server.Group("/api/v1/persistentVolumeClaim/namespace/:namespace")
	{
		persistentVolumeClaim.POST("/list", service.PersistentVolumeClaimList)

		persistentVolumeClaim.POST("/create", service.PersistentVolumeClaimCreate)
		// 因为大部分字段在创建后不允许修改，所以不提供修改api
		// persistentVolumeClaim.PATCH("/:name/update", service.PersistentVolumeClaimPatch)

		persistentVolumeClaim.DELETE("/:name", service.PersistentVolumeClaimDelete)
	}
}
