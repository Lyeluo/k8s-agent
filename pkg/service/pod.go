package service

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

//查询当前deployment下所有pod（根据label匹配）
func PodListByDeployment(c *gin.Context) {
	namespace := c.Param("namespace")

	deploymentName := c.Param("deployment")

	deployment, err := config.GetK8sConfig().AppsV1().Deployments(namespace).Get(deploymentName, metav1.GetOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("查询pod失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, deployment, err.Error()))
		return
	}

	listSelector := model.GetK8sSelectorByMap(deployment.Spec.Selector.MatchLabels)
	option := metav1.ListOptions{
		LabelSelector: listSelector.String(),
	}
	podList, err := config.GetK8sConfig().CoreV1().Pods(namespace).List(option)
	if err != nil {
		zap.L().Sugar().Errorf("查询pod失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, podList, err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.NewResponse(true, podList, model.NoErr.Msg))
}

// 更改pod（patch）
func PodPatch(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	data := make(map[string]interface{})
	if err := c.BindJSON(&data); err != nil {
		zap.L().Sugar().Errorf("更新deployment失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	playLoadBytes, _ := json.Marshal(data)

	podResult, err := config.GetK8sConfig().CoreV1().Pods(namespace).Patch(name, types.StrategicMergePatchType, playLoadBytes)
	if err != nil {
		zap.L().Sugar().Errorf("更新pod失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, podResult, model.NoErr.Msg))
}
