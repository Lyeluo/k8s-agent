package service

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/model"
	myutil "k8s.agent/pkg/util"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// 查询所有的pvc
func PersistentVolumeClaimList(c *gin.Context) {
	namespace := c.Param("namespace")

	listOptions, err := myutil.GetListOptions(c)
	if err != nil {
		zap.L().Sugar().Errorf("查询PersistentVolumeClaim信息失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	PersistentVolumeClaimList, err := config.GetK8sClient().CoreV1().PersistentVolumeClaims(namespace).List(listOptions)
	if err != nil {
		zap.L().Sugar().Errorf("查询PersistentVolumeClaim信息失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, PersistentVolumeClaimList, model.NoErr.Msg))

}

// 创建pvc
func PersistentVolumeClaimCreate(c *gin.Context) {
	namespace := c.Param("namespace")
	PersistentVolumeClaim := v1.PersistentVolumeClaim{}

	err := c.BindJSON(&PersistentVolumeClaim)
	if err != nil {
		zap.L().Sugar().Errorf("创建PersistentVolumeClaim失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	PersistentVolumeClaimResult, err := config.GetK8sClient().CoreV1().PersistentVolumeClaims(namespace).Create(&PersistentVolumeClaim)
	if err != nil {
		zap.L().Sugar().Errorf("创建PersistentVolumeClaim失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, PersistentVolumeClaimResult, model.NoErr.Msg))

}

//更新pvc（patch）
func PersistentVolumeClaimPatch(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")

	data := make(map[string]interface{})
	if err := c.BindJSON(&data); err != nil {
		zap.L().Sugar().Errorf("更新PersistentVolumeClaim失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}
	playLoadBytes, _ := json.Marshal(data)
	persistentResult, err := config.GetK8sClient().CoreV1().PersistentVolumeClaims(namespace).Patch(name, types.StrategicMergePatchType, playLoadBytes)
	if err != nil {
		zap.L().Sugar().Errorf("更新PersistentVolumeClaim失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, persistentResult, model.NoErr.Msg))

}

// 删除pvc
func PersistentVolumeClaimDelete(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")

	err := config.GetK8sClient().CoreV1().PersistentVolumeClaims(namespace).Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("删除PersistentVolumeClaim失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, nil, model.NoErr.Msg))

}
