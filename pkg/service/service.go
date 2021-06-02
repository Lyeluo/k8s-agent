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

// 查询当前命名空间下的所有service
func ServiceList(c *gin.Context) {
	namespace := c.Param("namespace")

	listOptions, err := myutil.GetListOptions(c)
	if err != nil {
		zap.L().Sugar().Errorf("更新deployment失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	serviceList, err := config.GetK8sClient().CoreV1().Services(namespace).List(listOptions)
	if err != nil {
		zap.L().Sugar().Errorf("查询service失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, serviceList, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, serviceList, model.NoErr.Msg))

}

// 当前namespace 下创建service
func ServiceCreate(c *gin.Context) {
	service := v1.Service{}

	err := c.BindJSON(&service)
	if err != nil {
		zap.L().Sugar().Errorf("创建service失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	serviceResult, err := config.GetK8sClient().CoreV1().Services(service.Namespace).Create(&service)
	if err != nil {
		zap.L().Sugar().Errorf("创建service失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, serviceResult, model.NoErr.Msg))
}

// 删除service
func ServiceDelete(c *gin.Context) {
	namespace := c.Param("namespace")
	serviceName := c.Param("name")

	err := config.GetK8sClient().CoreV1().Services(namespace).Delete(serviceName, &metav1.DeleteOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("删除service失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, nil, model.NoErr.Msg))

}

// 更新service
func ServicePatch(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	// todo: 绑定这一步不确定是否必须
	data := make(map[string]interface{})
	if err := c.BindJSON(&data); err != nil {
		zap.L().Sugar().Errorf("更新service失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}
	playLoadBytes, _ := json.Marshal(data)

	deploymentResult, err := config.GetK8sClient().CoreV1().Services(namespace).Patch(name, types.JSONPatchType, playLoadBytes)
	if err != nil {
		zap.L().Sugar().Errorf("更新service失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, deploymentResult, model.NoErr.Msg))

}

// 更新service（重新创建）
func ServiceUpdate(c *gin.Context) {
	service := v1.Service{}

	err := c.BindJSON(&service)
	if err != nil {
		zap.L().Sugar().Errorf("更新service失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	serviceResult, err := config.GetK8sClient().CoreV1().Services(service.Namespace).Update(&service)
	if err != nil {
		zap.L().Sugar().Errorf("更新service失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, serviceResult, model.NoErr.Msg))

}
