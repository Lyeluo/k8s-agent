package service

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/model"
	myutil "k8s.agent/pkg/util"
	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

//查询configmap
func ConfigMapList(c *gin.Context) {
	namespace := c.Param("namespace")

	listOptions, err := myutil.GetListOptions(c)
	if err != nil {
		zap.L().Sugar().Errorf("查询configmap失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	configMapList, err := config.GetK8sClient().CoreV1().ConfigMaps(namespace).List(listOptions)
	if err != nil {
		zap.L().Sugar().Errorf("查询configmap失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, configMapList, model.NoErr.Msg))

}

// 创建configmap
func ConfigMapCreate(c *gin.Context) {
	configmap := coreV1.ConfigMap{}

	err := c.BindJSON(&configmap)
	if err != nil {
		zap.L().Sugar().Errorf("创建configmap失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	configmapResult, err := config.GetK8sClient().CoreV1().ConfigMaps(configmap.Namespace).Create(&configmap)
	if err != nil {
		zap.L().Sugar().Errorf("创建configmap失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, configmapResult, model.NoErr.Msg))

}

// 根据名称删除
func ConfigMapDelete(c *gin.Context) {
	namespace := c.Param("namespace")
	configMapName := c.Param("name")

	err := config.GetK8sClient().CoreV1().ConfigMaps(namespace).Delete(configMapName, &v1.DeleteOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("删除configmap失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, nil, model.NoErr.Msg))

}

// 修改
func ConfigMapPatch(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")

	data := make(map[string]interface{})
	if err := c.BindJSON(&data); err != nil {
		zap.L().Sugar().Errorf("更新configmap失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}
	playLoadBytes, _ := json.Marshal(data)
	configmapResult, err := config.GetK8sClient().CoreV1().ConfigMaps(namespace).Patch(name, types.StrategicMergePatchType, playLoadBytes)
	if err != nil {
		zap.L().Sugar().Errorf("更新configmap失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, configmapResult, model.NoErr.Msg))

}
