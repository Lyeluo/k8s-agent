package service

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"k8s.agent/pkg/config"
	"k8s.agent/pkg/model"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// 根据namespace查询deploymen的列表
func DeploymentList(c *gin.Context) {
	namespace := c.Param("namespace")

	deploymentList, err := config.GetK8sConfig().AppsV1().Deployments(namespace).List(metav1.ListOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("查询namespace失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, deploymentList, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, deploymentList, model.NoErr.Msg))

}

// 根据name 与 namespace删除deployment
func DeploymentDelete(c *gin.Context) {
	namespace := c.Param("namespace")
	deploymentName := c.Param("name")

	err := config.GetK8sConfig().AppsV1().Deployments(namespace).Delete(deploymentName, &metav1.DeleteOptions{})
	if err != nil {
		zap.L().Sugar().Errorf("删除deployment失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, nil, model.NoErr.Msg))
}

// 创建deployment
func DeploymentCreate(c *gin.Context) {
	deployment := appsv1.Deployment{}

	err := c.BindJSON(&deployment)
	if err != nil {
		zap.L().Sugar().Errorf("创建deployment失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	deploymentResult, err := config.GetK8sConfig().AppsV1().Deployments(deployment.Namespace).Create(&deployment)
	if err != nil {
		zap.L().Sugar().Errorf("创建deployment失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, deploymentResult, model.NoErr.Msg))
}

// 更新deploymen（删除原有然后重新部署）
func DeploymentUpdate(c *gin.Context) {
	deployment := appsv1.Deployment{}

	err := c.BindJSON(&deployment)
	if err != nil {
		zap.L().Sugar().Errorf("更新deployment失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	deploymentResult, err := config.GetK8sConfig().AppsV1().Deployments(deployment.Namespace).Update(&deployment)
	if err != nil {
		zap.L().Sugar().Errorf("更新deployment失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, deploymentResult, model.NoErr.Msg))

}

// 更新deployment(patch)
func DeploymentPatch(c *gin.Context) {
	name := c.Param("name")
	namespace := c.Param("namespace")

	data := make(map[string]interface{})
	if err := c.BindJSON(&data); err != nil {
		zap.L().Sugar().Errorf("更新deployment失败，原因: %s", err.Error())
		c.JSON(http.StatusBadRequest, model.NewResponse(false, nil, err.Error()))
		return
	}

	playLoadBytes, _ := json.Marshal(data)

	deploymentResult, err := config.GetK8sConfig().AppsV1().Deployments(namespace).Patch(name, types.StrategicMergePatchType, playLoadBytes)
	if err != nil {
		zap.L().Sugar().Errorf("更新deployment失败，原因: %s", err.Error())
		c.JSON(http.StatusOK, model.NewResponse(false, nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, model.NewResponse(true, deploymentResult, model.NoErr.Msg))

}
