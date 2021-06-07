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

// 查询所有的pv
func PersistentVolumeList(c *gin.Context) {

	listOptions, err := myutil.GetListOptions(c)
	if err != nil {
		zap.L().Sugar().Errorf("查询PersistentVolume信息失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	persistentVolumeList, err := config.GetK8sClient().CoreV1().PersistentVolumes().List(listOptions)
	if err != nil {
		zap.L().Sugar().Errorf("查询PersistentVolume信息失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, persistentVolumeList, model.NoErr.Msg))

}

// 创建pv
func PersistentVolumeCreate(c *gin.Context) {
	persistentVolume := v1.PersistentVolume{}

	err := c.BindJSON(&persistentVolume)
	if err != nil {
		zap.L().Sugar().Errorf("创建PersistentVolume失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	persistentVolumeResult, err := config.GetK8sClient().CoreV1().PersistentVolumes().Create(&persistentVolume)
	if err != nil {
		zap.L().Sugar().Errorf("创建PersistentVolume失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, persistentVolumeResult, model.NoErr.Msg))

}

//更新pv（patch）
func PersistentVolumePatch(c *gin.Context) {
	name := c.Param("name")

	data := make(map[string]interface{})
	if err := c.BindJSON(&data); err != nil {
		zap.L().Sugar().Errorf("更新PersistentVolume失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}
	playLoadBytes, _ := json.Marshal(data)
	persistentResult, err := config.GetK8sClient().CoreV1().PersistentVolumes().Patch(name, types.StrategicMergePatchType, playLoadBytes)
	if err != nil {
		zap.L().Sugar().Errorf("更新PersistentVolume失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, persistentResult, model.NoErr.Msg))

}

// 删除pv
func PersistentVolumeDelete(c *gin.Context) {
	name := c.Param("name")

	err := config.GetK8sClient().CoreV1().PersistentVolumes().Delete(name, &metav1.DeleteOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("删除PersistentVolume失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, nil, model.NoErr.Msg))

}
